package binance

import (
	"cryptoMonitor/config"
	"fmt"
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
	lossTick := config.Get().DataSource.ProfitStrategy.LossTick
	longTick := config.Get().DataSource.ProfitStrategy.LongTick
	pNow := decimal.New(0, 0)
	longR := config.Get().DataSource.ProfitStrategy.LongR
	pMin := s.EnterPrice.Sub(lossTick)
	tmpMsg := fmt.Sprintf("signal:long start uuid:%s symbol:%s pE:%s", orderUUID, s.Symbol, s.EnterPrice)
	log.Info(tmpMsg)
	for {
		tmp, ok := GetPriceMap().Load(s.Symbol)
		if !ok {
			continue
		}
		pNow = tmp.(decimal.Decimal)
		if pNow.GreaterThan(pMax) {
			pMax = pNow
		}
		if pNow.GreaterThanOrEqual(s.EnterPrice.Add(longTick)) {
			tmpI := pMax.Sub(s.EnterPrice)
			tmpI = tmpI.Mul(longR)
			if tmpI.GreaterThanOrEqual(pNow.Sub(s.EnterPrice)) {
				tmpMsg := fmt.Sprintf("signal:long win uuid:%s symbol:%s pE:%s pMax:%s pNow:%s",
					orderUUID, s.Symbol, s.EnterPrice, pMax, pNow)
				log.Info(tmpMsg)
				break
			}
		}
		if pNow.LessThanOrEqual(pMin) {
			tmpMsg := fmt.Sprintf("signal:long lose uuid:%s symbol:%s pE:%s pMin:%s pNow:%s",
				orderUUID, s.Symbol, s.EnterPrice, pMin, pNow)
			log.Info(tmpMsg)
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}
