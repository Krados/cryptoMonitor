package monitor

import "cryptoMonitor/monitor/binance"

func Start() {
	binance.InitRunner()
	binance.NewPriceMonitor().Run()
	binance.NewKlineMonitor().Run()
}
