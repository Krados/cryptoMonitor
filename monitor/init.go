package monitor

import "cryptoMonitor/monitor/binance"

func Start() {
	binance.NewMonitor().Run()
}
