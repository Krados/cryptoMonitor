package binance

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	//urlValues.Set("price", price)
	//urlValues.Set("timeInForce", timeInForce)
	urlValues.Set("recvWindow", recvWindow)
	urlValues.Set("timestamp", fmt.Sprintf("%d", lib.NowInMilliSecond()))
	key := config.Get().DataSource.APIKey
	secret := config.Get().DataSource.APISecret

	// add key to header
	m := make(map[string]string)
	m["X-MBX-APIKEY"] = key

	// hmac sha256 payload
	payload := urlValues.Encode()
	sha := lib.HmacSha256(secret, payload)

	// send order
	reqUrl := fmt.Sprintf("%s/fapi/v1/order?%s&signature=%s",
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

func OpenOrder(symbol string) (openOrderResp OpenOrderResp, exists bool, err error) {
	// prepare url values
	urlValues := url.Values{}
	urlValues.Set("symbol", symbol)
	urlValues.Set("timestamp", fmt.Sprintf("%d", lib.NowInMilliSecond()))

	// add key to header
	m := make(map[string]string)
	m["X-MBX-APIKEY"] = config.Get().DataSource.APIKey

	// hmac sha256 payload
	payload := urlValues.Encode()
	sha := lib.HmacSha256(config.Get().DataSource.APISecret, payload)

	// get open order
	reqUrl := fmt.Sprintf("%s/fapi/v1/openOrder?%s&signature=%s",
		config.Get().DataSource.APISetting.Base, payload, sha)
	dataByte, err := lib.SendRequest(http.MethodGet, reqUrl, nil, m)
	if err != nil {
		return
	}

	// try unmarshal to ErrorResp
	var tmpErrorResp ErrorResp
	tryErr := json.Unmarshal(dataByte, &tmpErrorResp)
	if tryErr == nil {
		return
	}

	// unmarshal
	var tmpOpenOrderResp OpenOrderResp
	err = json.Unmarshal(dataByte, &tmpOpenOrderResp)
	if err != nil {
		return
	}
	openOrderResp = tmpOpenOrderResp
	exists = true

	return
}

func Balance() (balanceResp []BalanceResp, err error) {
	// prepare url values
	urlValues := url.Values{}
	urlValues.Set("timestamp", fmt.Sprintf("%d", lib.NowInMilliSecond()))

	// add key to header
	m := make(map[string]string)
	m["X-MBX-APIKEY"] = config.Get().DataSource.APIKey

	// hmac sha256 payload
	payload := urlValues.Encode()
	sha := lib.HmacSha256(config.Get().DataSource.APISecret, payload)

	// get balance
	reqUrl := fmt.Sprintf("%s/fapi/v2/balance?%s&signature=%s",
		config.Get().DataSource.APISetting.Base, payload, sha)
	dataByte, err := lib.SendRequest(http.MethodGet, reqUrl, nil, m)
	if err != nil {
		return
	}

	// unmarshal
	var tmpBalanceResp []BalanceResp
	err = json.Unmarshal(dataByte, &tmpBalanceResp)
	if err != nil {
		return
	}
	balanceResp = tmpBalanceResp

	return
}

func GetOrder(symbol string, orderId int) (orderResp GetOrderResp, err error) {
	// prepare url values
	urlValues := url.Values{}
	urlValues.Set("symbol", symbol)
	urlValues.Set("orderId", fmt.Sprintf("%d", orderId))
	urlValues.Set("timestamp", fmt.Sprintf("%d", lib.NowInMilliSecond()))

	// add key to header
	m := make(map[string]string)
	m["X-MBX-APIKEY"] = config.Get().DataSource.APIKey

	// hmac sha256 payload
	payload := urlValues.Encode()
	sha := lib.HmacSha256(config.Get().DataSource.APISecret, payload)

	// get order
	reqUrl := fmt.Sprintf("%s/fapi/v1/order?%s&signature=%s",
		config.Get().DataSource.APISetting.Base, payload, sha)
	dataByte, err := lib.SendRequest(http.MethodGet, reqUrl, nil, m)
	if err != nil {
		return
	}

	// unmarshal
	var tmpOrderResp GetOrderResp
	err = json.Unmarshal(dataByte, &tmpOrderResp)
	if err != nil {
		return
	}
	orderResp = tmpOrderResp

	return
}
