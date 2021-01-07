package binance

type KlineResp [][]interface{}

type LatestPriceResp struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
