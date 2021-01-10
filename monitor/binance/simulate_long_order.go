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

	// waiting for trigger price
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
	triggerAt := time.Now()

	// waiting for (priceMax - priceNow) < priceMax * LongR
	priceMax := decimal.New(0, 0)
	var priceNow decimal.Decimal
	var tmpForMaxSubNow decimal.Decimal
	var tmpMaxMulR decimal.Decimal
	shortR := config.Get().DataSource.ProfitStrategy.ShortR
	longR := config.Get().DataSource.ProfitStrategy.LongR
	triggerMulShortR := triggerPrice.Mul(shortR)
	for {
		val, ok := GetPriceMap().Load(s.Symbol)
		if !ok {
			continue
		}
		priceNow = val.(decimal.Decimal)
		if priceNow.GreaterThan(priceMax) {
			priceMax = priceNow
		}

		// out the order cuz long enough
		tmpForMaxSubNow = priceMax.Sub(priceNow)
		tmpMaxMulR = priceMax.Mul(longR)
		if !tmpForMaxSubNow.IsZero() && tmpForMaxSubNow.LessThan(tmpMaxMulR) {
			log.Infof("triggerAt:%s uuid:%s symbol:%s enterPrice:%s priceNow:%s "+
				"triggerPrice:%s priceMax:%s tmpForMaxSubNow:%s tmpMaxMulR:%s longR:%s",
				triggerAt, orderUUID, s.Symbol, s.EnterPrice, priceNow,
				triggerPrice, priceMax, tmpForMaxSubNow, tmpMaxMulR, longR)
			break
		}

		// out the order cuz short enough
		if priceNow.LessThan(triggerMulShortR) {
			log.Infof("triggerAt:%s uuid:%s symbol:%s enterPrice:%s priceNow:%s "+
				"triggerPrice:%s triggerMulShortR:%s shortR:%s",
				triggerAt, orderUUID, s.Symbol, s.EnterPrice, priceNow,
				triggerPrice, triggerMulShortR, shortR)
			break
		}

		time.Sleep(time.Second)
	}
}
