package binance

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetKlineData(reqUrl string) (kResp KlineResp, err error) {
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dataByte, err := ioutil.ReadAll(resp.Body)
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
	req, err := http.NewRequest("GET", reqUrl, nil)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dataByte, err := ioutil.ReadAll(resp.Body)
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
