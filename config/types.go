package config

import "time"

type Config struct {
	DataSource DataSource `json:"data_source"`
}

type APISetting struct {
	Base           string `json:"base"`
	KlineURI       string `json:"kline_uri"`
	LatestPriceURI string `json:"latest_price_uri"`
}

type WatchList struct {
	Symbol   string `json:"symbol"`
	Interval string `json:"interval"`
	Limit    int    `json:"limit"`
}

type DataSource struct {
	Name       string        `json:"name"`
	Interval   time.Duration `json:"interval"`
	APISetting APISetting    `json:"api_setting"`
	WatchList  []WatchList   `json:"watch_list"`
	Strategy   Strategy      `json:"strategy"`
}

type Strategy struct {
	N1K int `json:"n_1_k"`
	N2K int `json:"n_2_k"`
}
