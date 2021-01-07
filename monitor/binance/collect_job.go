package binance

import (
	"cryptoMonitor/cache"
	"cryptoMonitor/strategy"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type CollectJob struct {
	KResp      KlineResp
	PResp      LatestPriceResp
	Symbol     string
	Strategies []string
}

func (c CollectJob) Exec() {
	kLines, err := ParseKlineData(c.KResp)
	if err != nil {
		log.Warningln(err)
		return
	}
	exec := strategy.NewStrategyExecutioner().
		SetKline(kLines).SetStrategy(c.Strategies)
	suggestion := exec.Exec()
	keyByte := []byte(c.Symbol)
	dataByte, err := json.Marshal(suggestion)
	if err != nil {
		log.Warningln(err)
		return
	}
	err = cache.Get().Set(keyByte, dataByte, -1)
	if err != nil {
		log.Warningln(err)
		return
	}
}
