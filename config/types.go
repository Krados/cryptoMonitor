package config

import (
	"github.com/shopspring/decimal"
	"time"
)

type Config struct {
	DataSource DataSource `json:"data_source"`
}

type APISetting struct {
	Base           string `json:"base"`
	KlineURI       string `json:"kline_uri"`
	LatestPriceURI string `json:"latest_price_uri"`
}

type WatchList struct {
	Symbol     string   `json:"symbol"`
	Interval   string   `json:"interval"`
	Limit      int      `json:"limit"`
	Strategies []string `json:"strategies"`
}

type DataSource struct {
	Name       string        `json:"name"`
	Interval   time.Duration `json:"interval"`
	APISetting APISetting    `json:"api_setting"`
	WatchList  []WatchList   `json:"watch_list"`
	Strategy   Strategy      `json:"strategy"`
}

type SmaCross struct {
	N1K int `json:"n_1_k"`
	N2K int `json:"n_2_k"`
}

type Strategy struct {
	Weight     decimal.Decimal `json:"weight"`
	SmaCross   SmaCross        `json:"sma_cross"`
	DualThrust DualThrust      `json:"dual_thrust"`
}

type DualThrust struct {
	N1K   int             `json:"n_1_k"`
	KUp   decimal.Decimal `json:"k_up"`
	KDown decimal.Decimal `json:"k_down"`
}
