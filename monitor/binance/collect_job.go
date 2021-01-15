package binance

import (
	"cryptoMonitor/cache"
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"cryptoMonitor/strategy"
	"encoding/json"
	log "github.com/sirupsen/logrus"
)

type CollectJob struct {
	KResp     KlineResp
	WatchList config.WatchList
}

func (c CollectJob) Exec() {
	kLines, err := ParseKlineData(c.KResp)
	if err != nil {
		log.Warningln(err)
		return
	}
	exec := strategy.NewStrategyExecutioner().
		SetKline(kLines).SetStrategy(c.WatchList.Strategies)
	suggestion := exec.Exec()
	keyByte := []byte(c.WatchList.Symbol)
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

	log.Debugf("symbol:%s, pd:%s, hd:%s, price:%v long_strategy:%s short_strategy:%s",
		c.WatchList.Symbol, lib.PlaceDirectionStr(suggestion.PlaceOrderDirection),
		lib.HoldDirectionStr(suggestion.HoldDirection), kLines[len(kLines)-1].ClosePrice,
		suggestion.InLongStrategies, suggestion.InShortStrategies)

	if suggestion.PlaceOrderDirection == lib.InUnknown {
		SetSignal(c.WatchList.Symbol, lib.InUnknown)
		return
	}

	if suggestion.PlaceOrderDirection == lib.InLong {
		val, err := GetSignal(c.WatchList.Symbol)
		if err != nil {
			log.Warnf("%s", err)
			return
		}
		if val != lib.InUnknown {
			return
		}
		SetSignal(c.WatchList.Symbol, lib.InLong)
		err = GetRunner().Receive(ActualLongOrder{
			WatchList: c.WatchList,
		})
		if err != nil {
			log.Debugf("watch long order failed , symbol:%s err:%s", c.WatchList.Symbol, err)
		}
	} else if suggestion.PlaceOrderDirection == lib.InShort {
		val, err := GetSignal(c.WatchList.Symbol)
		if err != nil {
			log.Warnf("%s", err)
			return
		}
		if val != lib.InUnknown {
			return
		}
		SetSignal(c.WatchList.Symbol, lib.InShort)
		err = GetRunner().Receive(ActualShortOrder{
			WatchList: c.WatchList,
		})
		if err != nil {
			log.Debugf("watch short order failed , symbol:%s err:%s", c.WatchList.Symbol, err)
		}
	}
}
