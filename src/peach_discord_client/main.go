package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

// VERSION of Peach
const VERSION = "v0.1.0"

var bottoken string = os.Getenv("BOTTOKEN")

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
	l := createLog()

	l.Info("shard node starting...")

	c, err := CreateClient()
	if err != nil {
		l.Fatal("Unable to create new client, exiting...")
	}
	c.log = l
	done := make(chan bool)
	err = c.Run()

	<-done
}
