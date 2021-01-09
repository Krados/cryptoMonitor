package binance

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
)

type KlineMonitor struct {
	WatchList  []config.WatchList
	Interval   time.Duration
	BaseURL    string
	KlineURI   string
	Dispatcher lib.Dispatcher
}

func NewKlineMonitor() KlineMonitor {
	var tmpMonitor KlineMonitor
	tmpMonitor.Interval = config.Get().DataSource.Interval
	tmpMonitor.BaseURL = config.Get().DataSource.APISetting.Base
	tmpMonitor.KlineURI = config.Get().DataSource.APISetting.KlineURI

	tmpMonitor.Dispatcher = lib.NewDispatcher(30)
	tmpMonitor.Dispatcher.Start()
	for i := 0; i < len(config.Get().DataSource.WatchList); i++ {
		tmpMonitor.WatchList = append(tmpMonitor.WatchList,
			config.Get().DataSource.WatchList[i])
	}
	return tmpMonitor
}

func (m KlineMonitor) Run() {
	for i := 0; i < len(m.WatchList); i++ {
		go m.GetKlineDataInterval(m.WatchList[i].Symbol, m.WatchList[i].Interval, m.WatchList[i].Limit, m.WatchList[i].Strategies)
	}
}

func (m KlineMonitor) GetKlineDataInterval(symbol string, interval string, limit int, strategies []string) {
	ticker := time.NewTicker(m.Interval)
	for range ticker.C {
		reqUrl := fmt.Sprintf("%s%s?symbol=%s&interval=%s&limit=%d",
			m.BaseURL, m.KlineURI, symbol, interval, limit)
		kResp, err := GetKlineData(reqUrl)
		if err != nil {
			log.Warningln(err)
			continue
		}
		m.Dispatcher.Dispatch(CollectJob{
			KResp:      kResp,
			Symbol:     symbol,
			Strategies: strategies,
		})
	}
}
