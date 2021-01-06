package binance

const (
	InLong = iota
	InShort
	Consolidation
)

const (
	HoldLong = iota
	HoldShort
)

type ShouldAttempt struct {
	PlaceOrderDirection int `json:"place_order_direction"`
	HoldDirection       int `json:"hold_direction"`
}
