package binance

import (
	"errors"
	"sync"
)

var signalMap sync.Map

func GetSignal(symbol string) (int, error) {
	val, ok := signalMap.Load(symbol)
	if !ok {
		return -1, errors.New("!ok")
	}
	return val.(int), nil
}

func SetSignal(symbol string, i int) {
	signalMap.Store(symbol, i)
}
