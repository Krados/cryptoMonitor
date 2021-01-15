package strategy

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
)

type SmaCross struct {
}

func (s SmaCross) Calculate(data []lib.KlineData) (prediction lib.DirectionPrediction, err error) {
	prediction.Name = "SmaCross"
	n1 := config.Get().DataSource.Strategy.SmaCross.N1K
	n2 := config.Get().DataSource.Strategy.SmaCross.N2K
	n1Ma, err := NkMa(n1, 1, data)
	n2Ma, err := NkMa(n2, 1, data)
	n1Pma, err := NkMa(n1, 2, data)
	n2Pma, err := NkMa(n2, 2, data)
	if err != nil {
		prediction.PlaceOrderDirection = lib.InUnknown
		prediction.HoldDirection = lib.UnknownHold
	}

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
