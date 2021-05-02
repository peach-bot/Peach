package main

import (
	"fmt"
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
	ID          string
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

func (l *Launcher) runClient() {
	for {
		cmd := &exec.Cmd{
			Path: "./discordclient-" + VERSION + ".exe",
			Args: []string{
				"./discordclient-" + VERSION + ".exe",
				fmt.Sprintf("--log=%s", l.Config.Launcher.LogLevel),
				fmt.Sprintf("--sharded=%t", true),
				fmt.Sprintf("--token=%s", ""),
				fmt.Sprintf("--ccurl=%s", shellescape.Quote(l.Config.Launcher.CoordinatorURL)),
				fmt.Sprintf("--secret=%s", l.Config.Secret),
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

func (c *Client) stop() {
	c.Process.Signal(syscall.SIGTERM)
}

func (l *Launcher) SetupCloseHandler() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}
