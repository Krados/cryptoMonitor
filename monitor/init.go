package monitor

import (
	"cryptoMonitor/monitor/binance"
	"github.com/shopspring/decimal"
)

func Start() {
	binance.FinalBalance = &binance.BalanceTmp{
		Value: decimal.New(0, 0),
	}
	binance.InitRunner()
	binance.NewPriceMonitor().Run()
	binance.NewKlineMonitor().Run()
}
