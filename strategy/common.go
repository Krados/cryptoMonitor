package strategy

import (
	"cryptoMonitor/lib"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
)

func NkMa(n, offset int, data []lib.KlineData) (nkMa decimal.Decimal, err error) {
	i1 := len(data) - n - offset
	i2 := len(data) - offset
	if i1 < 0 || i2 < 0 {
		err = errors.New(fmt.Sprintf("params wrong len:%d,n:%d,offset:%d", len(data), n, offset))
		return
	}
	newData := data[i1:i2]
	var tmpTotal decimal.Decimal
	for _, val := range newData {
		tmpTotal = tmpTotal.Add(val.ClosePrice)
	}
	if len(newData) == 0 {
		return
	}
	nkMa = tmpTotal.Div(decimal.NewFromInt(int64(len(newData))))

	return
}
