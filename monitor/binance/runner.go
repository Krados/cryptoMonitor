package binance

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
)

var _Runner *lib.WorkRunner

func InitRunner() {
	_Runner = lib.NewWorkRunner(config.Get().DataSource.TradeRunnerNum)
}

func GetRunner() *lib.WorkRunner {
	return _Runner
}
