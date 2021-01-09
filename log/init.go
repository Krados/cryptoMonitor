package log

import (
	"cryptoMonitor/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

func Init() (err error) {
	logPath := "./logs"
	if _, tmpErr := os.Stat(logPath); os.IsNotExist(tmpErr) {
		err = os.Mkdir(logPath, 0755)
		if err != nil {
			return
		}
	}
	fileName := fmt.Sprintf("./logs/%s", config.Get().DataSource.LogFileName)
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	return
}
