package binance

import (
	"cryptoMonitor/cache"
	"cryptoMonitor/lib"
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

	log.Debugf("symbol:%s, pd:%s, hd:%s, price:%v",
		c.Symbol, lib.PlaceDirectionStr(suggestion.PlaceOrderDirection),
		lib.HoldDirectionStr(suggestion.HoldDirection), kLines[len(kLines)-1].ClosePrice)

	if suggestion.PlaceOrderDirection == lib.InUnknown {
		return
	}

	if suggestion.PlaceOrderDirection == lib.InLong {
		err = GetRunner().Receive(SimulateLongOrder{
			EnterPrice: kLines[len(kLines)-1].ClosePrice,
			Symbol:     c.Symbol,
		})
		if err != nil {
			log.Warnf("simulate long order failed , symbol:%s err:%s", c.Symbol, err)
		}
	} else if suggestion.PlaceOrderDirection == lib.InShort {
		err = GetRunner().Receive(SimulateShortOrder{})
		if err != nil {
			log.Warnf("simulate short order failed , symbol:%s err:%s", c.Symbol, err)
		}
	}
}
