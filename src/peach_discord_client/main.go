package main

import (
	"flag"
	"os"

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
	log.Info("Shard node starting...")

	// command line flags
	sharded := flag.Bool("sharded", false, "determines weather bot runs in shards or not")
	TOKEN := flag.String("token", "", "token override instead of secrets")
	loglevel := flag.String("log", "info", "declares how verbose the logging should be ('debug', 'info', 'error')")
	flag.Parse()

	switch *loglevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	}

	for {
		c, err := CreateClient(log, *sharded)
		if err != nil {
			log.Fatal(err, "\nUnable to create new client, exiting...")
		}

		if *sharded == false {
			c.GatewayURL = "wss://gateway.discord.gg/"
		}

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

		err = c.Run()
		if err != nil {
			c.Log.Fatal(err)
		}
		select {
		case <-c.Reconnect:
			c.Log.Info("Reconnecting...")
		case <-c.Quit:
			c.Log.Info("Quitting...")
			return
		}
	}
}
