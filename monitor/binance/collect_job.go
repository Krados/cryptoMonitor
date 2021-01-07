package binance

import (
	"cryptoMonitor/cache"
	"cryptoMonitor/config"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"

	log "github.com/sirupsen/logrus"
)

type CollectJob struct {
	KResp  KlineResp
	PResp  LatestPriceResp
	Symbol string
}

func (c CollectJob) Exec() {
	n1 := config.GetConfig().DataSource.Strategy.N1K
	n2 := config.GetConfig().DataSource.Strategy.N2K
	n1Ma, err := c.NkMa(n1, c.KResp)
	if err != nil {
		log.Warningln(err)
		return
	}
	n2Ma, err := c.NkMa(n2, c.KResp)
	if err != nil {
		log.Warningln(err)
		return
	}
	n1Pma, err := c.NkPMa(n1, c.KResp)
	if err != nil {
		log.Warningln(err)
		return
	}
	n2Pma, err := c.NkPMa(n2, c.KResp)
	if err != nil {
		log.Warningln(err)
		return
	}
	keyByte := []byte(c.Symbol)
	var dataByte []byte
	var attempt ShouldAttempt
	if n1Ma.GreaterThan(n2Ma) && n1Pma.LessThan(n2Pma) {
		log.Infoln("long_signal", c.Symbol)
		attempt.PlaceOrderDirection = InLong
		attempt.HoldDirection = UnknownHold
		dataByte, err = json.Marshal(attempt)
		if err != nil {
			log.Warningln(err)
			return
		}
		err = cache.Get().Set(keyByte, dataByte, -1)
		if err != nil {
			log.Warningln(err)
			return
		}

	} else if n1Ma.LessThan(n2Ma) && n1Pma.GreaterThan(n2Pma) {
		log.Infoln("short_signal", c.Symbol)
		attempt.PlaceOrderDirection = InShort
		attempt.HoldDirection = UnknownHold
		dataByte, err = json.Marshal(attempt)
		if err != nil {
			log.Warningln(err)
			return
		}
		err = cache.Get().Set(keyByte, dataByte, -1)
		if err != nil {
			log.Warningln(err)
			return
		}
	} else {
		log.Infoln("unknown_signal", c.Symbol)
		attempt.PlaceOrderDirection = InUnknown
		attempt.HoldDirection = UnknownHold
		dataByte, err = json.Marshal(attempt)
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
}

func (c CollectJob) NkMa(n int, k KlineResp) (nkMa decimal.Decimal, err error) {
	tmpK20 := c.KResp[len(k)-n-2 : len(k)-2]
	var tmpTotal decimal.Decimal
	for _, val := range tmpK20 {
		tmpD, tmpErr := decimal.NewFromString(fmt.Sprintf("%v", val[4]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmpTotal = tmpTotal.Add(tmpD)
	}
	nkMa = tmpTotal.Div(decimal.NewFromInt(int64(n)))

	return
}

func (c CollectJob) NkPMa(n int, k KlineResp) (nkPMa decimal.Decimal, err error) {
	tmpK20 := c.KResp[len(k)-n-3 : len(k)-3]
	var tmpTotal decimal.Decimal
	for _, val := range tmpK20 {
		tmpD, tmpErr := decimal.NewFromString(fmt.Sprintf("%v", val[4]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmpTotal = tmpTotal.Add(tmpD)
	}
	nkPMa = tmpTotal.Div(decimal.NewFromInt(int64(n)))

	return
}
