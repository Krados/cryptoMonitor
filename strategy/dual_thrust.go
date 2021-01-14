package strategy

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"github.com/shopspring/decimal"
)

type DualThrust struct {
}

func (s DualThrust) Calculate(data []lib.KlineData) (prediction lib.DirectionPrediction, err error) {
	n := config.Get().DataSource.Strategy.DualThrust.N1K
	kUp := config.Get().DataSource.Strategy.DualThrust.KUp
	kDown := config.Get().DataSource.Strategy.DualThrust.KDown
	maxH := s.MaxHighest(n, data)
	minL := s.MinLowest(n, data)
	maxC := s.MaxClose(n, data)
	minC := s.MinClose(n, data)
	latestK := data[len(data)-1]
	latestK_1 := data[len(data)-2]
	channelRange := s.ChannelRange(maxH, minL, maxC, minC)
	channelUp := s.ChannelUp(latestK.OpenPrice, channelRange, kUp)
	channelDown := s.ChannelDown(latestK.OpenPrice, channelRange, kDown)

	if latestK.ClosePrice.GreaterThan(channelUp) && latestK_1.ClosePrice.LessThan(channelUp) {
		prediction.PlaceOrderDirection = lib.InLong
		prediction.HoldDirection = lib.UnknownHold
	} else if latestK.ClosePrice.LessThan(channelDown) && latestK_1.ClosePrice.GreaterThan(channelDown) {
		prediction.PlaceOrderDirection = lib.InShort
		prediction.HoldDirection = lib.UnknownHold
	} else {
		prediction.PlaceOrderDirection = lib.InUnknown
		prediction.HoldDirection = lib.UnknownHold
	}

	return
}

func (s DualThrust) MaxHighest(n int, data []lib.KlineData) decimal.Decimal {
	all := make([]decimal.Decimal, 0)
	tmpData := data[len(data)-2-n : len(data)-2]
	for _, val := range tmpData {
		all = append(all, val.HighestPrice)
	}
	maxH := decimal.Max(all[0], all[1:]...)

	return maxH
}

func (s DualThrust) MinLowest(n int, data []lib.KlineData) decimal.Decimal {
	all := make([]decimal.Decimal, 0)
	tmpData := data[len(data)-2-n : len(data)-2]
	for _, val := range tmpData {
		all = append(all, val.LowestPrice)
	}
	minL := decimal.Min(all[0], all[1:]...)

	return minL
}

func (s DualThrust) MaxClose(n int, data []lib.KlineData) decimal.Decimal {
	all := make([]decimal.Decimal, 0)
	tmpData := data[len(data)-2-n : len(data)-2]
	for _, val := range tmpData {
		all = append(all, val.ClosePrice)
	}
	maxC := decimal.Max(all[0], all[1:]...)

	return maxC
}

func (s DualThrust) MinClose(n int, data []lib.KlineData) decimal.Decimal {
	all := make([]decimal.Decimal, 0)
	tmpData := data[len(data)-2-n : len(data)-2]
	for _, val := range tmpData {
		all = append(all, val.ClosePrice)
	}
	minC := decimal.Min(all[0], all[1:]...)

	return minC
}

func (s DualThrust) ChannelRange(maxH, minL, maxC, minC decimal.Decimal) decimal.Decimal {
	tmpD1 := maxH.Sub(minC)
	tmpD2 := maxC.Sub(minL)

	maxRange := decimal.Max(tmpD1, tmpD2)

	return maxRange
}

func (s DualThrust) ChannelUp(latestOpen, channelRange, k decimal.Decimal) decimal.Decimal {
	return channelRange.Mul(k).Add(latestOpen)
}

func (s DualThrust) ChannelDown(latestOpen, channelRange, k decimal.Decimal) decimal.Decimal {
	return channelRange.Mul(k).Sub(latestOpen)
}
