package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

var clustersecret string = os.Getenv("CLUSTERSECRET")
var db database

func createlog() *logrus.Logger {
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
	var err error
	l := createlog()
	l.Info("shard coordinator starting...")

	createdb(l)

	c := new(clientCoordinator)
	c.log = l

	c.heartbeatInterval = "10000ms"

	err = c.create()
	if err != nil {
		c.log.Fatal(err)
	}

	// setup paths
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", c.pathLogin).Methods(http.MethodGet)
	api.HandleFunc("/ready", c.pathReady).Methods(http.MethodGet)
	api.HandleFunc("/shards", c.pathGetShards).Methods(http.MethodGet)
	api.HandleFunc("/heartbeat", c.pathHeartbeat).Methods(http.MethodGet)
	api.HandleFunc("/guilds/{serverID}", c.pathGetServerSettings).Methods(http.MethodGet)

	// run
	done := make(chan bool)

	go http.ListenAndServe(":8080", r)

	// log ready
	l.Info("shard coordinator online")
	<-done
}
