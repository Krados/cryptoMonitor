package binance

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

type ShouldAttempt struct {
	PlaceOrderDirection int `json:"place_order_direction"`
	HoldDirection       int `json:"hold_direction"`
}

type KlineResp [][]interface{}

type LatestPriceResp struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}
