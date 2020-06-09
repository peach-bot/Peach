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
	l.SetLevel(logrus.DebugLevel)
	return l
}

func main() {
	log := createLog()

	log.Info("Shard node starting...")

	c, err := CreateClient(log)
	if err != nil {
		log.Fatal(err, "\nUnable to create new client, exiting...")
	}

	// Set discord epoch and sequence
	snowflake.Epoch = 1420070400000

	// Settings
	c.Compress = true
	c.LargeThreshold = 250
	c.GuildSubscriptions = true
	c.TOKEN = os.Getenv("BOTTOKEN")
	c.MissingHeartbeatAcks = 5
	c.GatewayURL = c.GatewayURL + "?v=" + APIVersion + "&encoding=json"

	done := make(chan bool)
	err = c.Run()
	if err != nil {
		c.Log.Fatal(err)
	}

	<-done
}
