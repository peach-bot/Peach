package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/alessio/shellescape"
	"github.com/sirupsen/logrus"
)

type Launcher struct {
	sync.Mutex
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
		Sharded             bool   `json:"sharded"`
		Shards              int    `json:"shards"`
		Token               string `json:"token"`
		LogLevel            string `json:"log_level"`
		CoordinatorURL      string `json:"coordinator"`
		SpotifyClientID     string `json:"spotify_client_id"`
		SpotifyClientSecret string `json:"spotify_client_secret"`
	} `json:"clients"`
	Clientcoordinator struct {
		Launch        bool   `json:"launch"`
		Port          string `json:"port"`
		DBCredentials string `json:"dbc"`
		CertType      string `json:"cert_type"`
		Domain        string `json:"domain"`
	} `json:"clientcoordinator"`
	Secret          string `json:"secret"`
	RedactSensitive bool   `json:"redact_sensitive"`
}

func (l *Launcher) runClient() {
	for {
		cmd := &exec.Cmd{
			Path: "./discordclient-" + VERSION + ".exe",
			Args: []string{
				"./discordclient-" + VERSION + ".exe",
				fmt.Sprintf("--log=%s", l.Config.Clients.LogLevel),
				fmt.Sprintf("--sharded=%t", l.Config.Clients.Sharded),
				fmt.Sprintf("--token=%s", l.Config.Clients.Token),
				fmt.Sprintf("--ccurl=%s", shellescape.Quote(l.Config.Clients.CoordinatorURL)),
				fmt.Sprintf("--secret=%s", l.Config.Secret),
				fmt.Sprintf("--spotifyid=%s", l.Config.Clients.SpotifyClientID),
				fmt.Sprintf("--spotifysecret=%s", l.Config.Clients.SpotifyClientSecret),
				fmt.Sprintf("--redactsensitive=%t", l.Config.RedactSensitive),
			},
			Stdout: os.Stdout,
			Stderr: os.Stderr,
		}
		var c Client
		c.Process = cmd.Process
		c.Pos = len(l.Clients)
		l.Lock()
		l.Clients = append(l.Clients, c)
		l.Unlock()
		err := cmd.Run()
		if err != nil {
			l.Log.Error(err)
		}
		l.Lock()
		l.Clients = append(l.Clients[:c.Pos], l.Clients[c.Pos+1:]...)
		l.Unlock()

		// Delay before trying to restart
		time.Sleep(5 * time.Second)
	}
}

func (l *Launcher) runCoordinator() {
	cmd := &exec.Cmd{
		Path: "./coordinator-" + VERSION + ".exe",
		Args: []string{
			"./coordinator-" + VERSION + ".exe",
			fmt.Sprintf("--secret=%s", l.Config.Secret),
			fmt.Sprintf("--dbc=%s", l.Config.Clientcoordinator.DBCredentials),
			fmt.Sprintf("--port=%s", l.Config.Clientcoordinator.Port),
			fmt.Sprintf("--certtype=%s", l.Config.Clientcoordinator.CertType),
			fmt.Sprintf("--domain=%s", l.Config.Clientcoordinator.Domain),
		},
		Stdout: os.Stdout,
		Stderr: os.Stderr,
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
