package strategy

import (
	"cryptoMonitor/config"
	"cryptoMonitor/lib"
	"github.com/shopspring/decimal"
	"sync"
)

type Executioner struct {
	klineList   []lib.KlineData
	strategies  []string
	predictions []lib.DirectionPrediction
	sync.Mutex
}

type FinalSuggestion struct {
	PlaceOrderDirection int `json:"place_order_direction"`
	HoldDirection       int `json:"hold_direction"`
}

func NewStrategyExecutioner() *Executioner {
	return &Executioner{}
}

func (e *Executioner) SetKline(in []lib.KlineData) *Executioner {
	e.klineList = in
	return e
}

func (e *Executioner) GetKline() []lib.KlineData {
	return e.klineList
}

func (e *Executioner) SetStrategy(in []string) *Executioner {
	e.strategies = in
	return e
}

func (e *Executioner) Exec() FinalSuggestion {
	var wg sync.WaitGroup
	for _, val := range e.strategies {
		instance := NewStrategy(val)
		if instance == nil {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			prediction, _ := instance.Calculate(e.GetKline())
			e.AddPrediction(prediction)
		}()
	}
	wg.Wait()
	return e.WrapUp()
}

func (e *Executioner) AddPrediction(prediction lib.DirectionPrediction) {
	e.Lock()
	defer e.Unlock()
	e.predictions = append(e.predictions, prediction)
}

func (e *Executioner) WrapUp() FinalSuggestion {
	var suggestion FinalSuggestion
	totalPrediction := decimal.NewFromInt(int64(len(e.predictions)))
	weight := config.Get().DataSource.Strategy.Weight
	inLongCount := 0
	inShortCount := 0
	holdLongCount := 0
	holdShortCount := 0
	for _, val := range e.predictions {
		if val.PlaceOrderDirection == lib.InLong {
			inLongCount += 1
		} else if val.PlaceOrderDirection == lib.InShort {
			inShortCount += 1
		}
		if val.HoldDirection == lib.HoldLong {
			holdLongCount += 1
		} else if val.HoldDirection == lib.HoldShort {
			holdShortCount += 1
		}
	}
	if inLongCount == 0 && inShortCount == 0 {
		suggestion.PlaceOrderDirection = lib.InUnknown
	} else {
		if inLongCount > 0 {
			d := decimal.NewFromInt(int64(inLongCount))
			if d.Div(totalPrediction).GreaterThanOrEqual(weight) {
				suggestion.PlaceOrderDirection = lib.InLong
			} else {
				suggestion.PlaceOrderDirection = lib.InUnknown
			}
		}
		if inShortCount > 0 {
			d := decimal.NewFromInt(int64(inShortCount))
			if d.Div(totalPrediction).GreaterThanOrEqual(weight) {
				suggestion.PlaceOrderDirection = lib.InShort
			} else {
				suggestion.PlaceOrderDirection = lib.InUnknown
			}
		}
	}
	if holdLongCount == 0 && holdShortCount == 0 {
		suggestion.HoldDirection = lib.UnknownHold
	} else {
		if holdLongCount > 0 {
			d := decimal.NewFromInt(int64(holdLongCount))
			if d.Div(totalPrediction).GreaterThanOrEqual(weight) {
				suggestion.PlaceOrderDirection = lib.HoldLong
			} else {
				suggestion.PlaceOrderDirection = lib.UnknownHold
			}
		}
		if holdShortCount > 0 {
			d := decimal.NewFromInt(int64(holdShortCount))
			if d.Div(totalPrediction).GreaterThanOrEqual(weight) {
				suggestion.PlaceOrderDirection = lib.HoldShort
			} else {
				suggestion.PlaceOrderDirection = lib.UnknownHold
			}
		}
	}

	return suggestion
}
