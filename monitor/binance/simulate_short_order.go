package binance

import (
	"cryptoMonitor/config"
	"fmt"
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
	lossTick := config.Get().DataSource.ProfitStrategy.LossTick
	shortTick := config.Get().DataSource.ProfitStrategy.ShortTick
	pNow := decimal.New(0, 0)
	shortR := config.Get().DataSource.ProfitStrategy.ShortR
	pMax := s.EnterPrice.Add(lossTick)
	tmpMsg := fmt.Sprintf("signal:short start uuid:%s symbol:%s pE:%s", orderUUID, s.Symbol, s.EnterPrice)
	log.Info(tmpMsg)
	for {
		tmp, ok := GetPriceMap().Load(s.Symbol)
		if !ok {
			continue
		}
		pNow = tmp.(decimal.Decimal)
		if pNow.LessThan(pMin) {
			pMin = pNow
		}
		if pNow.LessThanOrEqual(s.EnterPrice.Sub(shortTick)) {
			tmpI := s.EnterPrice.Sub(pMin)
			tmpI = tmpI.Mul(shortR)
			if tmpI.GreaterThanOrEqual(s.EnterPrice.Sub(pNow)) {
				tmpMsg := fmt.Sprintf("signal:short win uuid:%s symbol:%s pE:%s pMax:%s pNow:%s",
					orderUUID, s.Symbol, s.EnterPrice, pMax, pNow)
				log.Info(tmpMsg)
				break
			}
		}
		if pNow.GreaterThanOrEqual(pMax) {
			tmpMsg := fmt.Sprintf("signal:short lose uuid:%s symbol:%s pE:%s pMin:%s pNow:%s",
				orderUUID, s.Symbol, s.EnterPrice, pMin, pNow)
			log.Infof(tmpMsg)
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}
