package binance

import (
	"cryptoMonitor/config"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

type SimulateShortOrder struct {
	EnterPrice decimal.Decimal
	Symbol     string
}

func (s SimulateShortOrder) Action() {
	orderUUID := uuid.NewV4()
	pMin := decimal.New(0, 0)
	tick := config.Get().DataSource.ProfitStrategy.Tick
	pMax := s.EnterPrice.Add(tick)
	pNow := decimal.New(0, 0)
	shortR := config.Get().DataSource.ProfitStrategy.ShortR
	log.Infof("signal:short start uuid:%s symbol:%s pE:%s", orderUUID, s.Symbol, s.EnterPrice)
	for {
		tmp, ok := GetPriceMap().Load(s.Symbol)
		if !ok {
			continue
		}
		pNow = tmp.(decimal.Decimal)
		if pNow.LessThan(pMin) {
			pMin = pNow
		}
		if !pMin.IsZero() {
			tmp := s.EnterPrice.Sub(pMin)
			tmp = tmp.Mul(shortR)
			tmp = s.EnterPrice.Sub(tmp)
			if pNow.GreaterThanOrEqual(tmp) && pNow.LessThan(s.EnterPrice.Sub(tick)) {
				word := ""
				if pNow.LessThan(s.EnterPrice) {
					word = "enough"
				} else {
					word = "insufficient"
				}
				log.Infof("signal:short win %s uuid:%s symbol:%s pE:%s pMin:%s pNow:%s",
					word, orderUUID, s.Symbol, s.EnterPrice, pMin, pNow)
				FinalBalance.Lock()
				FinalBalance.Value = FinalBalance.Value.Add(s.EnterPrice.Sub(pNow))
				FinalBalance.Unlock()
				break
			}
		}
		if pNow.GreaterThanOrEqual(pMax) {
			log.Infof("signal:short lose uuid:%s symbol:%s pE:%s pMax:%s pNow:%s",
				orderUUID, s.Symbol, s.EnterPrice, pMax, pNow)
			FinalBalance.Lock()
			FinalBalance.Value = FinalBalance.Value.Add(s.EnterPrice.Sub(pNow))
			FinalBalance.Unlock()
			break
		}

		time.Sleep(time.Millisecond * 500)
	}
	log.Infof("final balance:%s", FinalBalance.Value)
}
