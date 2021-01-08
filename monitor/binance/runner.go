package binance

import "cryptoMonitor/lib"

var _Runner *lib.WorkRunner

func InitRunner() {
	_Runner = lib.NewWorkRunner(100)
}

func GetRunner() *lib.WorkRunner {
	return _Runner
}
