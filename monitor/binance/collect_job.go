package binance

import (
	"cryptoMonitor/cache"
	"encoding/json"
)

type CollectJob struct {
	KResp  KlineResp
	Symbol string
}

func (c CollectJob) Exec() {
	//todo: use strategy to decide the data
	if len(c.KResp)%2 == 0 {
		keyByte := []byte(c.Symbol)
		data := ShouldAttempt{
			PlaceOrderDirection: InLong,
			HoldDirection:       HoldLong,
		}
		dataByte, err := json.Marshal(data)
		if err != nil {
			return
		}
		_ = cache.Get().Set(keyByte, dataByte, -1)
	}
}
