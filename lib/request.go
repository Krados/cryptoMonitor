package lib

import (
	"io"
	"io/ioutil"
	"net/http"
)

func SendRequest(method, reqUrl string, body io.Reader, headers map[string]string) (dataByte []byte, err error) {
	req, err := http.NewRequest(method, reqUrl, body)
	if err != nil {
		return
	}
	if len(headers) != 0 {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
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
