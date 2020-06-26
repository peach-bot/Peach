package main

import (
	"flag"
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

	sharded := flag.Bool("sharded", false, "determines weather bot runs in shards or not")
	TOKEN := flag.String("token", "", "token override instead of secrets")
	flag.Parse()

	c, err := CreateClient(log, *sharded)
	if err != nil {
		log.Fatal(err, "\nUnable to create new client, exiting...")
	}

	if *sharded == false {
		c.GatewayURL = "wss://gateway.discord.gg/"
	}

	// Set discord epoch and sequence
	snowflake.Epoch = 1420070400000

	// Settings
	c.Sharded = *sharded
	c.Compress = true
	c.LargeThreshold = 250
	c.GuildSubscriptions = true
	if *TOKEN == "" {
		c.TOKEN = os.Getenv("BOTTOKEN")
	} else {
		c.TOKEN = *TOKEN
	}
	c.MissingHeartbeatAcks = 5
	c.GatewayURL = c.GatewayURL + "?v=" + APIVersion + "&encoding=json"

	done := make(chan bool)
	err = c.Run()
	if err != nil {
		c.Log.Fatal(err)
	}

	<-done
}
