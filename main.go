package main

import (
	"cryptoMonitor/api"
	"cryptoMonitor/cache"
	"cryptoMonitor/config"
	logger "cryptoMonitor/log"
	"cryptoMonitor/monitor"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init config
	if err := config.Init(); err != nil {
		log.Fatalln(err)
		return
	}

	// init logger
	if err := logger.Init(); err != nil {
		log.Fatalln(err)
		return
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
