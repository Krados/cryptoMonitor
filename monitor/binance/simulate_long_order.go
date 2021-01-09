package binance

import (
	"cryptoMonitor/config"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

type SimulateLongOrder struct {
	EnterPrice decimal.Decimal
	Symbol     string
}

func (s SimulateLongOrder) Action() {
	orderUUID := uuid.NewV4()
	var triggerPrice decimal.Decimal
	var tmpPriceForTrigger decimal.Decimal
	for {
		val, ok := GetPriceMap().Load(s.Symbol)
		if !ok {
			continue
		}
		tmpPriceForTrigger = val.(decimal.Decimal)
		if tmpPriceForTrigger.GreaterThan(s.EnterPrice) {
			triggerPrice = tmpPriceForTrigger
			break
		}
		time.Sleep(time.Second)
	}

	priceMax := decimal.New(0, 0)
	priceMin := decimal.New(0, 0)
	var tmpPrice decimal.Decimal
	var tmpForMaxSubTrigger decimal.Decimal
	var tmpMaxMulR decimal.Decimal
	var tmpMinMulR decimal.Decimal
	for {
		val, ok := GetPriceMap().Load(s.Symbol)
		if !ok {
			continue
		}
		tmpPrice = val.(decimal.Decimal)
		if tmpPrice.GreaterThan(priceMax) {
			priceMax = tmpPrice
		}
		if priceMin.LessThan(tmpPrice) {
			priceMin = tmpPrice
		}

		// out the order cuz long enough
		tmpForMaxSubTrigger = priceMax.Sub(triggerPrice)
		tmpMaxMulR = priceMax.Mul(config.Get().DataSource.ProfitStrategy.LongR)
		if tmpForMaxSubTrigger.LessThan(tmpMaxMulR) {
			log.Infof("uuid:%s symbol:%s enterPrice:%s triggerPrice:%s priceMax:%s longR:%s",
				orderUUID, s.Symbol, s.EnterPrice, triggerPrice, priceMax, config.Get().DataSource.ProfitStrategy.LongR)
			break
		}

		// out the order cuz short enough
		tmpMinMulR = priceMin.Mul(config.Get().DataSource.ProfitStrategy.ShortR)
		if tmpMinMulR.GreaterThan(triggerPrice) {
			log.Infof("uuid:%s symbol:%s enterPrice:%s triggerPrice:%s priceMin:%s shortR:%s",
				orderUUID, s.Symbol, s.EnterPrice, triggerPrice, priceMin, config.Get().DataSource.ProfitStrategy.ShortR)
			break
		}

		time.Sleep(time.Second)
	}
}
