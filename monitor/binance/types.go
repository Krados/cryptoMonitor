package binance

import "github.com/shopspring/decimal"

type KlineResp [][]interface{}

type LatestPriceResp struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
}
