package binance

import (
	"crypto/hmac"
	"crypto/sha256"
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func GetKlineData(reqUrl string) (kResp KlineResp, err error) {
	dataByte, err := lib.SendRequest(http.MethodGet, reqUrl, nil, nil)
	if err != nil {
		return
	}
	var tmpResp KlineResp
	err = json.Unmarshal(dataByte, &tmpResp)
	if err != nil {
		return
	}
	kResp = tmpResp

	return
}

func GetLatestPrice(reqUrl string) (priceResp []LatestPriceResp, err error) {
	dataByte, err := lib.SendRequest(http.MethodGet, reqUrl, nil, nil)
	if err != nil {
		return
	}
	var tmpResp []LatestPriceResp
	err = json.Unmarshal(dataByte, &tmpResp)
	if err != nil {
		return
	}
	priceResp = tmpResp

	return
}

func SendOrder(
	symbol string,
	side string,
	oType string,
	quantity string,
	price string,
	timeInForce string,
	recvWindow string) (orderResp OrderResp, err error) {

	// prepare url values
	urlValues := url.Values{}
	urlValues.Set("symbol", symbol)
	urlValues.Set("side", side)
	urlValues.Set("type", oType)
	urlValues.Set("quantity", quantity)
	urlValues.Set("price", price)
	urlValues.Set("timeInForce", timeInForce)
	urlValues.Set("recvWindow", recvWindow)
	urlValues.Set("timestamp", fmt.Sprintf("%d", NowInMilliSecond()))
	key := config.Get().DataSource.APIKey
	secret := config.Get().DataSource.APISecret

	// add key to header
	m := make(map[string]string)
	m["X-MBX-APIKEY"] = key

	// hmac sha256 payload
	payload := urlValues.Encode()
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(payload))
	sha := hex.EncodeToString(h.Sum(nil))

	// send order
	reqUrl := fmt.Sprintf("%s/fapi/v1/order/test?%s&signature=%s",
		config.Get().DataSource.APISetting.Base, payload, sha)
	dataByte, err := lib.SendRequest(http.MethodPost, reqUrl, nil, m)
	if err != nil {
		return
	}

	// unmarshal
	var tmpOrderResp OrderResp
	err = json.Unmarshal(dataByte, &tmpOrderResp)
	if err != nil {
		return
	}
	orderResp = tmpOrderResp

	return
}

func NowInMilliSecond() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
