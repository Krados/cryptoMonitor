package lib

import (
	"encoding/json"
	"errors"
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

	// make sure resp no error
	var errorResp ErrorResp
	err = json.Unmarshal(tmpDataByte, &errorResp)
	if err == nil {
		err = errors.New(string(tmpDataByte))
		return
	}

	dataByte = tmpDataByte

	return
}
