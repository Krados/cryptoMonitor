package binance

import (
	"cryptoMonitor/config"
	"fmt"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var _Price sync.Map

type PriceMonitor struct {
	Interval       time.Duration
	BaseURL        string
	LatestPriceURI string
}

func NewPriceMonitor() PriceMonitor {
	return PriceMonitor{
		Interval:       config.Get().DataSource.Interval,
		BaseURL:        config.Get().DataSource.APISetting.Base,
		LatestPriceURI: config.Get().DataSource.APISetting.LatestPriceURI,
	}
}

func (m PriceMonitor) Run() {
	ticker := time.NewTicker(m.Interval)
	reqUrl := fmt.Sprintf("%s%s", m.BaseURL, m.LatestPriceURI)
	go func() {
		for range ticker.C {
			err := m.GetPriceAndSet(reqUrl)
			if err != nil {
				log.Warningln(err)
				continue
			}
		}
	}()

}

func (m PriceMonitor) GetPriceAndSet(reqUrl string) (err error) {
	pResp, err := GetLatestPrice(reqUrl)
	if err != nil {
		return
	}
	for _, val := range pResp {
		_Price.LoadOrStore(val.Symbol, val.Price)
	}

	return
}

func GetPriceMap() *sync.Map {
	return &_Price
}
