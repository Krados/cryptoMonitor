package main

import (
	"cryptoMonitor/api"
	"cryptoMonitor/cache"
	"cryptoMonitor/config"
	"cryptoMonitor/monitor"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

func main() {
	// init config
	if err := config.Init(); err != nil {
		log.Fatalln(err)
	}

	// init cache
	cache.Init()

	// init monitor
	monitor.Start()

	// init router
	api.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
