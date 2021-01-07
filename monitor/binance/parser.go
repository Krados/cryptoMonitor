package binance

import (
	"cryptoMonitor/lib"
	"fmt"
	"github.com/shopspring/decimal"
)

func ParseKlineData(in KlineResp) (list []lib.KlineData, err error) {
	var tmpList []lib.KlineData
	var tmpD decimal.Decimal
	var tmpErr error
	for _, val := range in {
		var tmp lib.KlineData
		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[1]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.OpenPrice = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[2]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.HighestPrice = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[3]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.LowestPrice = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[4]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.ClosePrice = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[5]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.Volume = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[7]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.QuoteAssetVolume = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[7]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.QuoteAssetVolume = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[8]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.NumberOfTrades = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[9]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.TakerBuyBaseAssetVolume = tmpD

		tmpD, tmpErr = decimal.NewFromString(fmt.Sprintf("%v", val[10]))
		if tmpErr != nil {
			err = tmpErr
			return
		}
		tmp.TakerBuyQuoteAssetVolume = tmpD

		tmpList = append(tmpList, tmp)
	}
	list = tmpList

	return
}
