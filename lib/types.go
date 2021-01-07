package lib

import (
	"github.com/shopspring/decimal"
)

const (
	InLong = iota
	InShort
	Consolidation
	InUnknown
)

const (
	HoldLong = iota
	HoldShort
	UnknownHold
)

type KlineData struct {
	OpenPrice                decimal.Decimal
	HighestPrice             decimal.Decimal
	LowestPrice              decimal.Decimal
	ClosePrice               decimal.Decimal
	Volume                   decimal.Decimal
	QuoteAssetVolume         decimal.Decimal
	NumberOfTrades           decimal.Decimal
	TakerBuyBaseAssetVolume  decimal.Decimal
	TakerBuyQuoteAssetVolume decimal.Decimal
}

type DirectionPrediction struct {
	PlaceOrderDirection int `json:"place_order_direction"`
	HoldDirection       int `json:"hold_direction"`
}
