package binance

import "github.com/shopspring/decimal"

type KlineResp [][]interface{}

type LatestPriceResp struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
}

type OrderResp struct {
	ClientOrderID string          `json:"clientOrderId"`
	CumQty        decimal.Decimal `json:"cumQty"`
	CumQuote      decimal.Decimal `json:"cumQuote"`
	ExecutedQty   decimal.Decimal `json:"executedQty"`
	OrderID       int             `json:"orderId"`
	AvgPrice      decimal.Decimal `json:"avgPrice"`
	OrigQty       decimal.Decimal `json:"origQty"`
	Price         decimal.Decimal `json:"price"`
	ReduceOnly    bool            `json:"reduceOnly"`
	Side          string          `json:"side"`
	PositionSide  string          `json:"positionSide"`
	Status        string          `json:"status"`
	StopPrice     decimal.Decimal `json:"stopPrice"`
	ClosePosition bool            `json:"closePosition"`
	Symbol        string          `json:"symbol"`
	TimeInForce   string          `json:"timeInForce"`
	Type          string          `json:"type"`
	OrigType      string          `json:"origType"`
	ActivatePrice decimal.Decimal `json:"activatePrice"`
	PriceRate     decimal.Decimal `json:"priceRate"`
	UpdateTime    int64           `json:"updateTime"`
	WorkingType   string          `json:"workingType"`
	PriceProtect  bool            `json:"priceProtect"`
}

type OpenOrderResp struct {
	AvgPrice      decimal.Decimal `json:"avgPrice"`
	ClientOrderID string          `json:"clientOrderId"`
	CumQuote      decimal.Decimal `json:"cumQuote"`
	ExecutedQty   decimal.Decimal `json:"executedQty"`
	OrderID       int             `json:"orderId"`
	OrigQty       decimal.Decimal `json:"origQty"`
	OrigType      string          `json:"origType"`
	Price         decimal.Decimal `json:"price"`
	ReduceOnly    bool            `json:"reduceOnly"`
	Side          string          `json:"side"`
	PositionSide  string          `json:"positionSide"`
	Status        string          `json:"status"`
	StopPrice     decimal.Decimal `json:"stopPrice"`
	ClosePosition bool            `json:"closePosition"`
	Symbol        string          `json:"symbol"`
	Time          int64           `json:"time"`
	TimeInForce   string          `json:"timeInForce"`
	Type          string          `json:"type"`
	ActivatePrice decimal.Decimal `json:"activatePrice"`
	PriceRate     decimal.Decimal `json:"priceRate"`
	UpdateTime    int64           `json:"updateTime"`
	WorkingType   string          `json:"workingType"`
	PriceProtect  bool            `json:"priceProtect"`
}
