package binance

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Monitor struct {
	WatchList      []config.WatchList
	Interval       time.Duration
	BaseURL        string
	KlineURI       string
	LatestPriceURI string
	Dispatcher     lib.Dispatcher
}

func NewMonitor() Monitor {
	var tmpMonitor Monitor
	tmpMonitor.Interval = config.GetConfig().DataSource.Interval
	tmpMonitor.BaseURL = config.GetConfig().DataSource.APISetting.Base
	tmpMonitor.KlineURI = config.GetConfig().DataSource.APISetting.KlineURI
	tmpMonitor.LatestPriceURI = config.GetConfig().DataSource.APISetting.LatestPriceURI

	tmpMonitor.Dispatcher = lib.NewDispatcher(30)
	tmpMonitor.Dispatcher.Start()
	for i := 0; i < len(config.GetConfig().DataSource.WatchList); i++ {
		tmpMonitor.WatchList = append(tmpMonitor.WatchList,
			config.GetConfig().DataSource.WatchList[i])
	}
	return tmpMonitor
}

func (m Monitor) Run() {
	for i := 0; i < len(m.WatchList); i++ {
		go m.GetKlineDataInterval(m.WatchList[i].Symbol, m.WatchList[i].Interval, m.WatchList[i].Limit)
	}
}

func (m Monitor) GetKlineDataInterval(symbol string, interval string, limit int) {
	ticker := time.NewTicker(m.Interval)
	for range ticker.C {
		kResp, err := m.GetKlineData(symbol, interval, limit)
		if err != nil {
			log.Warningln(err)
			continue
		}
		m.Dispatcher.Dispatch(CollectJob{
			KResp:  kResp,
			Symbol: symbol,
		})
	}
}

func (m Monitor) GetKlineData(symbol string, interval string, limit int) (kResp KlineResp, err error) {
	reqUrl := fmt.Sprintf("%s%s?symbol=%s&interval=%s&limit=%d",
		m.BaseURL, m.KlineURI, symbol, interval, limit)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dataByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var tmpResp KlineResp
	err = json.Unmarshal(dataByte, &tmpResp)
	if err != nil {
		return
	}
	kResp = tmpResp

	return
}

func (m Monitor) GetLatestPrice(symbol string) (priceResp LatestPriceResp, err error) {
	reqUrl := fmt.Sprintf("%s%s?symbol=%s",
		m.BaseURL, m.LatestPriceURI, symbol)
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dataByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	var tmpResp LatestPriceResp
	err = json.Unmarshal(dataByte, &tmpResp)
	if err != nil {
		return
	}
	priceResp = tmpResp

	return
}
