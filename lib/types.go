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

const (
	InLongStr        = "long"
	InShortStr       = "short"
	ConsolidationStr = "consolidation"
	InUnknownStr     = "unknown"
	HoldLongStr      = "hold_long"
	HoldShortStr     = "hold_short"
	UnknownHoldStr   = "hold_unknown"
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
	Name                string `json:"name"`
	PlaceOrderDirection int    `json:"place_order_direction"`
	HoldDirection       int    `json:"hold_direction"`
}

func PlaceDirectionStr(direction int) string {
	switch direction {
	case InLong:
		return InLongStr
	case InShort:
		return InShortStr
	case Consolidation:
		return ConsolidationStr
	}

	return InUnknownStr
}

func HoldDirectionStr(direction int) string {
	switch direction {
	case HoldLong:
		return HoldLongStr
	case HoldShort:
		return HoldShortStr
	}

	return UnknownHoldStr
}
