package strategy

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
)

type SmaCross struct {
}

func (s SmaCross) Calculate(data []lib.KlineData) (prediction lib.DirectionPrediction, err error) {
	n1 := config.Get().DataSource.Strategy.SmaCross.N1K
	n2 := config.Get().DataSource.Strategy.SmaCross.N2K
	n1Ma := NkMa(n1, -2, data)
	n2Ma := NkMa(n2, -2, data)
	n1Pma := NkMa(n1, -3, data)
	n2Pma := NkMa(n2, -3, data)

	if n1Ma.GreaterThan(n2Ma) && n1Pma.LessThan(n2Pma) {
		prediction.PlaceOrderDirection = lib.InLong
		prediction.HoldDirection = lib.UnknownHold

	} else if n1Ma.LessThan(n2Ma) && n1Pma.GreaterThan(n2Pma) {
		prediction.PlaceOrderDirection = lib.InShort
		prediction.HoldDirection = lib.UnknownHold
	} else {
		prediction.PlaceOrderDirection = lib.InUnknown
		prediction.HoldDirection = lib.UnknownHold
	}

	return
}
