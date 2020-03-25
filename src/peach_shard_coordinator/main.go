package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	log "github.com/sirupsen/logrus"
)

var shards []shard
var bottoken string = os.Getenv("BOTTOKEN")

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
	log.Info("shard coordinator starting...")

	// setup paths
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/getshard", getShard).Methods(http.MethodGet)
	api.HandleFunc("/reserveshard", reserveShard).Methods(http.MethodPost)
	api.HandleFunc("/updateshard", updateShard).Methods(http.MethodPost)
	api.HandleFunc("/scale", scale).Methods(http.MethodGet)

	// initial creation of shards list
	resetShardCount(0)

	// run
	done := make(chan bool)
	go http.ListenAndServe(":8080", r)

	// log ready
	log.Info("shard coordinator online")

	// wait for goroutine to finish
	<-done
}
