package binance

import (
	"cryptoMonitor/config"
	"cryptoMonitor/service"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

type ActualLongOrder struct {
	EnterPrice decimal.Decimal
	Symbol     string
	Strategies []string
}

func (s ActualLongOrder) Action() {
	orderUUID := uuid.NewV4()
	pMax := decimal.New(0, 0)
	lossTick := config.Get().DataSource.ProfitStrategy.LossTick
	longTick := config.Get().DataSource.ProfitStrategy.LongTick
	pNow := decimal.New(0, 0)
	longR := config.Get().DataSource.ProfitStrategy.LongR

	// make sure there is no open order exists
	_, exists, err := OpenOrder(s.Symbol)
	if err == nil {
		if exists {
			log.Warnf("signal:long open order exists")
			return
		}
	}

	// get balance
	balanceResp, err := Balance()
	if err != nil {
		log.Warnf("signal:long get balance failed, err:%s", err)
		return
	}
	balance := decimal.New(0, 0)
	for _, val := range balanceResp {
		if val.Asset == "USDT" {
			balance = val.Balance
		}
	}
	if !balance.IsPositive() {
		log.Warnf("signal:long !balance.IsPositive(), err:%s")
		return
	}

	// send order
	oriOrder, err := SendOrder(s.Symbol, "BUY", "MARKET", "1", "", "", "5000")
	if err != nil {
		log.Warnf("signal:long send order failed, err:%s", err)
		return
	}

	// get order
	currentOrder, err := GetOrder(s.Symbol, oriOrder.OrderID)
	if err != nil {
		log.Warnf("signal:long get order failed, err:%s", err)
		return
	}

	s.EnterPrice = currentOrder.AvgPrice
	pMin := s.EnterPrice.Sub(lossTick)
	tmpMsg := fmt.Sprintf("signal:long start uuid:%s symbol:%s pE:%s strategies:%s", orderUUID, s.Symbol, s.EnterPrice, s.Strategies)
	service.GetTelegramBot().SendMessage(tmpMsg)
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
				outOrder, err := SendOrder(s.Symbol, "SELL", "MARKET", "1", "", "", "5000")
				if err != nil {
					log.Warnf("signal:long send order failed, err:%s", err)
					return
				}
				log.Infof("%v", outOrder)
				tmpMsg := fmt.Sprintf("signal:long win uuid:%s symbol:%s pE:%s pMax:%s pNow:%s",
					orderUUID, s.Symbol, s.EnterPrice, pMax, pNow)
				service.GetTelegramBot().SendMessage(tmpMsg)
				log.Info(tmpMsg)
				break
			}
		}
		if pNow.LessThanOrEqual(pMin) {
			outOrder, err := SendOrder(s.Symbol, "SELL", "MARKET", "1", "", "", "5000")
			if err != nil {
				log.Warnf("signal:long send order failed, err:%s", err)
				return
			}
			log.Infof("%v", outOrder)
			tmpMsg := fmt.Sprintf("signal:long lose uuid:%s symbol:%s pE:%s pMin:%s pNow:%s",
				orderUUID, s.Symbol, s.EnterPrice, pMin, pNow)
			service.GetTelegramBot().SendMessage(tmpMsg)
			log.Info(tmpMsg)
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}
