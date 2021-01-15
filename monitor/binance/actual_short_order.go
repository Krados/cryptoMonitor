package binance

import (
	"cryptoMonitor/config"
	"cryptoMonitor/service"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"time"
)

type ActualShortOrder struct {
	EnterPrice decimal.Decimal
	Symbol     string
	Strategies []string
}

func (s ActualShortOrder) Action() {
	orderUUID := uuid.NewV4()
	pMin := decimal.New(0, 0)
	lossTick := config.Get().DataSource.ProfitStrategy.LossTick
	shortTick := config.Get().DataSource.ProfitStrategy.ShortTick
	pNow := decimal.New(0, 0)
	shortR := config.Get().DataSource.ProfitStrategy.ShortR

	// make sure there is no open order exists
	_, exists, err := OpenOrder(s.Symbol)
	if err == nil {
		if exists {
			log.Warnf("signal:short open order exists")
			return
		}
	}

	// get balance
	balanceResp, err := Balance()
	if err != nil {
		log.Warnf("signal:short get balance failed, err:%s", err)
		return
	}
	balance := decimal.New(0, 0)
	for _, val := range balanceResp {
		if val.Asset == "USDT" {
			balance = val.Balance
		}
	}
	if !balance.IsPositive() {
		log.Warnf("signal:short !balance.IsPositive(), err:%s")
		return
	}

	// send order
	oriOrder, err := SendOrder(s.Symbol, "SELL", "MARKET", "1", "", "", "5000")
	if err != nil {
		log.Warnf("signal:short send order failed, err:%s", err)
		return
	}

	// get order
	currentOrder, err := GetOrder(s.Symbol, oriOrder.OrderID)
	if err != nil {
		log.Warnf("signal:short get order failed, err:%s", err)
		return
	}

	s.EnterPrice = currentOrder.AvgPrice
	pMax := s.EnterPrice.Add(lossTick)
	tmpMsg := fmt.Sprintf("signal:short start uuid:%s symbol:%s pE:%s strategies:%s", orderUUID, s.Symbol, s.EnterPrice, s.Strategies)
	service.GetTelegramBot().SendMessage(tmpMsg)
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
				outOrder, err := SendOrder(s.Symbol, "BUY", "MARKET", "1", "", "", "5000")
				if err != nil {
					log.Warnf("signal:short send order failed, err:%s", err)
					return
				}
				log.Infof("%v", outOrder)
				tmpMsg := fmt.Sprintf("signal:short win uuid:%s symbol:%s pE:%s pMax:%s pNow:%s",
					orderUUID, s.Symbol, s.EnterPrice, pMax, pNow)
				service.GetTelegramBot().SendMessage(tmpMsg)
				log.Info(tmpMsg)
				break
			}
		}
		if pNow.GreaterThanOrEqual(pMax) {
			outOrder, err := SendOrder(s.Symbol, "BUY", "MARKET", "1", "", "", "5000")
			if err != nil {
				log.Warnf("signal:short send order failed, err:%s", err)
				return
			}
			log.Infof("%v", outOrder)
			tmpMsg := fmt.Sprintf("signal:short lose uuid:%s symbol:%s pE:%s pMin:%s pNow:%s",
				orderUUID, s.Symbol, s.EnterPrice, pMin, pNow)
			service.GetTelegramBot().SendMessage(tmpMsg)
			log.Infof(tmpMsg)
			break
		}
		time.Sleep(time.Millisecond * 500)
	}
}
