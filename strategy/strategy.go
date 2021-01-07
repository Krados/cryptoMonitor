package strategy

import "cryptoMonitor/lib"

type Strategy interface {
	Calculate(data []lib.KlineData) (lib.DirectionPrediction, error)
}

func NewStrategy(key string) Strategy {
	switch key {
	case "sma_cross":
		return SmaCross{}
	}
	return nil
}
