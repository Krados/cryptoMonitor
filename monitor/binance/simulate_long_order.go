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
	pMax := decimal.New(0, 0)
	tick := config.Get().DataSource.ProfitStrategy.Tick
	pMin := s.EnterPrice.Sub(tick)
	pNow := decimal.New(0, 0)
	longR := config.Get().DataSource.ProfitStrategy.LongR
	log.Infof("signal:long start uuid:%s symbol:%s pE:%s", orderUUID, s.Symbol, s.EnterPrice)

	for {
		tmp, ok := GetPriceMap().Load(s.Symbol)
		if !ok {
			continue
		}
		pNow = tmp.(decimal.Decimal)
		if pNow.GreaterThan(pMax) {
			pMax = pNow
		}
		if !pMax.IsZero() {
			tmp := pMax.Sub(s.EnterPrice)
			tmp = tmp.Mul(longR)
			tmp = tmp.Add(s.EnterPrice)
			if pNow.LessThanOrEqual(tmp) && pNow.GreaterThan(s.EnterPrice.Add(tick.Div(decimal.New(2, 0)))) {
				word := ""
				if pNow.GreaterThan(s.EnterPrice) {
					word = "enough"
				} else {
					word = "insufficient"
				}
				log.Infof("signal:long win %s uuid:%s symbol:%s pE:%s pMax:%s pNow:%s",
					word, orderUUID, s.Symbol, s.EnterPrice, pMax, pNow)
				FinalBalance.Lock()
				FinalBalance.Value = FinalBalance.Value.Add(pNow.Sub(s.EnterPrice))
				FinalBalance.Unlock()
				break
			}
		}
		if pNow.LessThanOrEqual(pMin) {
			log.Infof("signal:long lose uuid:%s symbol:%s pE:%s pMin:%s pNow:%s",
				orderUUID, s.Symbol, s.EnterPrice, pMin, pNow)
			FinalBalance.Lock()
			FinalBalance.Value = FinalBalance.Value.Add(pNow.Sub(s.EnterPrice))
			FinalBalance.Unlock()
			break
		}

		time.Sleep(time.Second)
	}
	log.Infof("final balance:%s", FinalBalance.Value)
}
