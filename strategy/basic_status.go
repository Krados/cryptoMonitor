package strategy

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"github.com/shopspring/decimal"
)

type BasicStatus struct {
}

func (s BasicStatus) Calculate(data []lib.KlineData) (prediction lib.DirectionPrediction, err error) {
	n := config.Get().DataSource.Strategy.BasicStatus.NK
	n1Ma, _ := s.NkMa(n, -1, data)
	n2Ma, _ := s.NkMa(n, -2, data)

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

func (s BasicStatus) NkMa(n, offset int, data []lib.KlineData) (nkMa decimal.Decimal, err error) {
	newData := data[len(data)-n-offset : len(data)-offset]
	var tmpTotal decimal.Decimal
	for _, val := range newData {
		tmpTotal = tmpTotal.Add(val.ClosePrice)
	}
	nkMa = tmpTotal.Div(decimal.NewFromInt(int64(n)))

	return
}
