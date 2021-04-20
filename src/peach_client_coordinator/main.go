package main

import (
	"flag"
	"net/http"
	"os"
	"runtime"

	"github.com/gorilla/mux"

	"github.com/sirupsen/logrus"
)

var clustersecret string
var db database

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
			return " CCoord", ""
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

	secret := flag.String("secret", "", "secret")
	dbc := flag.String("dbc", "", "data base credentials string")
	port := flag.String("port", "5000", "port the client coordinator should run on")
	certType := flag.String("certtype", "none", "build if cert is located in build folder, letsencrypt if cert is located under /etc/letsencrypt/live/")
	domain := flag.String("domain", "none", "domain")
	flag.Parse()
	clustersecret = *secret

	createdb(l, *dbc)

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

	s := &http.Server{
		Addr:    ":" + *port,
		Handler: r,
	}

	// run
	done := make(chan bool)

	switch *certType {
	case "build":
		go s.ListenAndServeTLS(*domain+".cert.pem", *domain+".key.pem")
	case "letsencrypt":
		go s.ListenAndServeTLS("/etc/letsencrypt/live/"+*domain+"/fullchain.pem", "/etc/letsencrypt/live/"+*domain+"/privkey.pem")
	case "none":
		go http.ListenAndServe(":"+*port, r)
	}

	// log ready
	l.Info("shard coordinator online")
	<-done
}
