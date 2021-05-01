package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var redactSensitive bool

func createLog() *logrus.Logger {
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
			return " Client", ""
		},
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
	ccURL := flag.String("ccurl", "", "url of the coordinator")
	secret := flag.String("secret", "", "secret for communicating with the coordinator")
	spotifyid := flag.String("spotifyid", "", "Spotify client id for spotify extension")
	spotifysecret := flag.String("spotifysecret", "", "Spotify client secret for spotify extension")
	redactSensitiveFlag := flag.Bool("redactsensitive", true, "Set to true to sensitive tokens and secrets from logs")
	flag.Parse()
	log.Infof("Sharded: %t, LogLevel: %s, coordinator URL: %s", *sharded, *loglevel, *ccURL)

	redactSensitive = *redactSensitiveFlag

	switch *loglevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)

	c, err := CreateClient(log, *sharded, *ccURL, *secret)
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
	if *TOKEN != "" {
		c.TOKEN = *TOKEN
	}
	c.MissingHeartbeatAcks = 5
	c.GatewayURL = c.GatewayURL + "?v=" + APIVersion + "&encoding=json"
	c.UserAgent = fmt.Sprintf("DiscordBot (https://github.com/peach-bot/Peach, %s)", VERSION)
	c.httpRetries = 5

	for {

		c.wsConn = nil
		c.httpClient = &http.Client{}
		c.GuildCache = cache.New(120*time.Minute, 5*time.Minute)
		c.ChannelCache = cache.New(120*time.Minute, 5*time.Minute)
		c.Extensions.setup(c, *spotifyid, *spotifysecret)

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
		case <-stop:
			c.Quit <- nil
			c.Log.Info("Quitting...")
			return
		}
	}
}
