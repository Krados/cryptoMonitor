package strategy

import (
	"cryptoMonitor/lib"
	"github.com/shopspring/decimal"
)

func NkMa(n, offset int, data []lib.KlineData) (nkMa decimal.Decimal) {
	newData := data[len(data)-n-offset : len(data)-offset]
	var tmpTotal decimal.Decimal
	for _, val := range newData {
		tmpTotal = tmpTotal.Add(val.ClosePrice)
	}
	nkMa = tmpTotal.Div(decimal.NewFromInt(int64(n)))

	return
}
