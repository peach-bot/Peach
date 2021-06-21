package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

func createlog() *logrus.Logger {
	// Set log format, output and level
	l := logrus.New()
	l.SetReportCaller(true)
	l.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		QuoteEmptyFields: true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return " Coord ", ""
		},
	})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	return l
}

func main() {
	var err error
	l := createlog()
	l.Info("shard coordinator starting...")

	c := new(Coordinator)
	c.loadJson()
	c.log = l
	c.createdb(c.Config.Coordinator.DBCredentials)

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
	api.HandleFunc("/guilds/{guildID}", c.pathGetGuildSettings).Methods(http.MethodGet)
	api.HandleFunc("/users/{userID}", c.pathGetUserSettings).Methods(http.MethodGet)

	s := &http.Server{
		Addr:    ":" + c.Config.Coordinator.Port,
		Handler: r,
	}

	// run
	done := make(chan bool)

	switch c.Config.Coordinator.CertType {
	case "build":
		go s.ListenAndServeTLS(c.Config.Coordinator.Domain+".cert.pem", c.Config.Coordinator.Domain+".key.pem")
	case "letsencrypt":
		go s.ListenAndServeTLS("/etc/letsencrypt/live/"+c.Config.Coordinator.Domain+"/fullchain.pem", "/etc/letsencrypt/live/"+c.Config.Coordinator.Domain+"/privkey.pem")
	case "none":
		go http.ListenAndServe(":"+c.Config.Coordinator.Port, r)
	}

	// log ready
	l.Info("shard coordinator online")
	<-done
}
