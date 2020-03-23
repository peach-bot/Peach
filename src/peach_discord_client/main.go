package main

import (
	"os"

	"github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
)

// VERSION of Peach
const VERSION = "v0.1.0"

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
	log := createLog()

	log.Info("shard node starting...")

	c, err := CreateClient()
	if err != nil {
		log.Fatal("Unable to create new client, exiting...")
	}
	c.Log = log

	// Set discord epoch and sequence
	snowflake.Epoch = 1420070400000

	// Settings
	c.Compress = true
	c.LargeThreshold = 250
	c.GuildSubscriptions = true
	c.TOKEN = os.Getenv("BOTTOKEN")
	c.MissingHeartbeatAcks = 5

	done := make(chan bool)
	err = c.Run()

	<-done
}
