package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

var clustersecret string = os.Getenv("CLUSTERSECRET")
var db database

func createlog() *logrus.Logger {
	// Set log format, output and level
	l := logrus.New()
	l.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		QuoteEmptyFields: true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
	})
	l.SetOutput(os.Stdout)
	l.SetLevel(log.DebugLevel)
	return l
}

func main() {
	var err error
	l := createlog()
	l.Info("shard coordinator starting...")

	dbc := strings.Split(os.Getenv("DATABASE"), ", ")
	port, err := strconv.Atoi(dbc[4])
	if err != nil {
		log.Fatal("Passed invalid database port.")
	}
	db = database{nil}
	db.dbconn, err = pgx.Connect(pgx.ConnConfig{Database: dbc[0], User: dbc[1], Password: dbc[2], Host: dbc[3], Port: uint16(port)})
	if err != nil {
		l.Fatal(err)
	}
	defer db.dbconn.Close()

	c := new(clientCoordinator)
	c.log = l
	err = c.create()
	if err != nil {
		c.log.Fatal(err)
	}

	// setup paths
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/login", c.pathLogin).Methods(http.MethodGet)
	api.HandleFunc("/ready", c.pathReady).Methods(http.MethodGet)
	api.HandleFunc("/heartbeat", c.pathHeartbeat).Methods(http.MethodGet)

	// run
	done := make(chan bool)
	go http.ListenAndServe(":8080", r)

	// log ready
	l.Info("shard coordinator online")
	<-done
}
