package binance

import (
	"github.com/shopspring/decimal"
	"sync"
)

var FinalBalance *BalanceTmp

type BalanceTmp struct {
	Value decimal.Decimal
	sync.Mutex
}
