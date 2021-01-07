package strategy

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

type SmaCross struct {
}

func (s SmaCross) Calculate(data []lib.KlineData) (prediction lib.DirectionPrediction, err error) {
	n1 := config.Get().DataSource.Strategy.SmaCross.N1K
	n2 := config.Get().DataSource.Strategy.SmaCross.N2K
	n1Ma, err := s.NkMa(n1, data)
	if err != nil {
		log.Warningln(err)
		return
	}
	n2Ma, err := s.NkMa(n2, data)
	if err != nil {
		log.Warningln(err)
		return
	}
	n1Pma, err := s.NkPMa(n1, data)
	if err != nil {
		log.Warningln(err)
		return
	}
	n2Pma, err := s.NkPMa(n2, data)
	if err != nil {
		log.Warningln(err)
		return
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

func (s SmaCross) NkMa(n int, data []lib.KlineData) (nkMa decimal.Decimal, err error) {
	newData := data[len(data)-n-2 : len(data)-2]
	var tmpTotal decimal.Decimal
	for _, val := range newData {
		tmpTotal = tmpTotal.Add(val.ClosePrice)
	}
	nkMa = tmpTotal.Div(decimal.NewFromInt(int64(n)))

	return
}

func (s SmaCross) NkPMa(n int, data []lib.KlineData) (nkPMa decimal.Decimal, err error) {
	newData := data[len(data)-n-3 : len(data)-3]
	var tmpTotal decimal.Decimal
	for _, val := range newData {
		tmpTotal = tmpTotal.Add(val.ClosePrice)
	}
	nkPMa = tmpTotal.Div(decimal.NewFromInt(int64(n)))

	return
}
