package strategy

import "cryptoMonitor/lib"

type Strategy interface {
	Calculate(data []lib.KlineData) (lib.DirectionPrediction, error)
}

func NewStrategy(key string) Strategy {
	switch key {
	case "sma_cross":
		return SmaCross{}
	case "dual_thrust":
		return DualThrust{}
	case "basic_status":
		return BasicStatus{}
	}

	return nil
}
