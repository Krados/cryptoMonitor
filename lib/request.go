package lib

import (
	"io"
	"io/ioutil"
	"net/http"
)

func SendRequest(method string, reqUrl string, body io.Reader) (dataByte []byte, err error) {
	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	tmpDataByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	dataByte = tmpDataByte

	return
}
