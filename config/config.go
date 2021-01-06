package config

import (
	"encoding/json"
	"io/ioutil"
)

var _config *Config

func Init() (err error) {
	jsonByte, err  := ioutil.ReadFile("./config.json")
	if err != nil {
		return
	}

	var tmpConfig Config
	err = json.Unmarshal(jsonByte, &tmpConfig)
	if err != nil {
		return
	}
	_config = &tmpConfig

	return
}

func GetConfig() *Config {
	return _config
}
