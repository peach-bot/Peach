package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

func keepAlive() {
	for {
		time.Sleep(time.Hour)
	}
}

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
			return " Launch", ""
		},
	})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.InfoLevel)
	return l
}

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var l Launcher
	l.Log = createLog()
	l.ID = fmt.Sprint(r.Int())

	err := l.loadJson()
	if err != nil {
		l.Log.Fatal(err)
	}

	l.Log.Print(l.Config)

	l.Stop = make(chan interface{})
	l.SetupCloseHandler()

	go keepAlive()

	// for i := 0; i < l.Config.Clients.Shards; i++ {
	go l.runClient()
	time.Sleep(1 * time.Second)
	// }

	select {
	case <-l.Stop:
		break
	}
}
