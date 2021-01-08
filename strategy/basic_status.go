package strategy

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
)

type BasicStatus struct {
}

func (s BasicStatus) Calculate(data []lib.KlineData) (prediction lib.DirectionPrediction, err error) {
	n := config.Get().DataSource.Strategy.BasicStatus.NK
	n1Ma := NkMa(n, -2, data)
	n2Ma := NkMa(n, -3, data)

	diff := n1Ma.Sub(n2Ma)

	if diff.IsPositive() {
		prediction.PlaceOrderDirection = lib.InLong
		prediction.HoldDirection = lib.UnknownHold
	} else if diff.IsNegative() && !diff.IsZero() {
		prediction.PlaceOrderDirection = lib.InShort
		prediction.HoldDirection = lib.UnknownHold
	} else {
		prediction.PlaceOrderDirection = lib.Consolidation
		prediction.HoldDirection = lib.UnknownHold
	}

	return
}
