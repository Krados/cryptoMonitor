package monitor

import "cryptoMonitor/monitor/binance"

func Start() {
	binance.InitRunner()
	binance.NewMonitor().Run()
}
