package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	// Set log format, output and level
	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		QuoteEmptyFields: true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.Info("shard node starting...")
	for {

	}
}
