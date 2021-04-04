package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/alessio/shellescape.v1"
)

type Launcher struct {
	Log         *logrus.Logger
	Stop        chan interface{}
	Config      Config
	Clients     []Client
	Coordinator Coordinator
}

type Client struct {
	Process *os.Process
	Pos     int
}

type Coordinator struct {
	Process *os.Process
}

type Config struct {
	Clients struct {
		Sharded        bool   `json:"sharded"`
		Shards         int    `json:"shards"`
		Token          string `json:"token"`
		LogLevel       string `json:"loglevel"`
		CoordinatorURL string `json:"coordinator"`
	} `json:"clients"`
	Clientcoordinator struct {
		Launch        bool   `json:"launch"`
		Port          string `json:"port"`
		DBCredentials string `json:"dbc"`
		CertType      string `json:"certtype"`
		Domain        string `json:"domain"`
	} `json:"clientcoordinator"`
	Secret string `json:"secret"`
}

func (l *Launcher) runClient() {
	cmd := &exec.Cmd{
		Path: "./discordclient.exe",
		Args: []string{
			"./discordclient.exe",
			fmt.Sprintf("--log=%s", l.Config.Clients.LogLevel),
			fmt.Sprintf("--sharded=%t", l.Config.Clients.Sharded),
			fmt.Sprintf("--token=%s", l.Config.Clients.Token),
			fmt.Sprintf("--ccurl=%s", shellescape.Quote(l.Config.Clients.CoordinatorURL)),
			fmt.Sprintf("--secret=%s", l.Config.Secret),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	var c Client
	c.Process = cmd.Process
	c.Pos = len(l.Clients)
	l.Clients = append(l.Clients, c)
	err := cmd.Run()
	if err != nil {
		l.Log.Error(err)
	}
	l.Clients = append(l.Clients[:c.Pos], l.Clients[c.Pos+1:]...)
}

func (l *Launcher) runCoordinator() {
	cmd := &exec.Cmd{
		Path: "./coordinator.exe",
		Args: []string{
			"./coordinator.exe",
			fmt.Sprintf("--secret=%s", l.Config.Secret),
			fmt.Sprintf("--dbc=%s", l.Config.Clientcoordinator.DBCredentials),
			fmt.Sprintf("--port=%s", l.Config.Clientcoordinator.Port),
			fmt.Sprintf("--certtype=%s", l.Config.Clientcoordinator.CertType),
			fmt.Sprintf("--domain=%s", l.Config.Clientcoordinator.Domain),
		},
		Stdout: os.Stdout,
	}
	var c Coordinator
	c.Process = cmd.Process
	l.Coordinator = c
	err := cmd.Run()
	if err != nil {
		l.Log.Fatal(err)
	}
}

func (c *Client) stop() {
	c.Process.Signal(syscall.SIGTERM)
}

func (c *Coordinator) stop() {
	c.Process.Signal(syscall.SIGTERM)
}

func (l *Launcher) loadJson() error {
	f, err := os.Open("launchcfg.json")
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	json.Unmarshal(b, &l.Config)

	return nil
}

func keepAlive() {
	for {
		time.Sleep(time.Hour)
	}
}

func (l *Launcher) SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		// for _, c := range l.Clients {
		// 	c.stop()
		// }
		// l.Coordinator.stop()

		os.Exit(0)
	}()
}

func createLog() *logrus.Logger {
	// Set log format, output and level
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		QuoteEmptyFields: true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
	})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.InfoLevel)
	return l
}

func main() {
	var l Launcher
	l.Log = createLog()

	err := l.loadJson()
	if err != nil {
		l.Log.Fatal(err)
	}

	l.Log.Print(l.Config)

	l.Stop = make(chan interface{})
	l.SetupCloseHandler()

	go keepAlive()

	if l.Config.Clientcoordinator.Launch {
		go l.runCoordinator()
	}

	for i := 0; i < l.Config.Clients.Shards; i++ {
		time.Sleep(5 * time.Second)
		go l.runClient()
	}

	select {
	case <-l.Stop:
		break
	}
}
