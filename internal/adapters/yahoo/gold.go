package yahoo

import (
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"

	"candlehub/internal/model"
)

type GoldAdapter struct{}

func NewGoldAdapter() *GoldAdapter {
	return &GoldAdapter{}
}

func (a *GoldAdapter) Asset() string {
	return "XAUUSD"
}

func (a *GoldAdapter) FetchMinuteCandles(from time.Time) ([]model.Candle, error) {
	now := time.Now()
	params := &chart.Params{
		Symbol:   "XAUUSD=X",
		Interval: datetime.OneMin,
		Start:    datetime.New(&from),
		End:      datetime.New(&now),
	}

	iter := chart.Get(params)

	candles := make([]model.Candle, 0)
	for iter.Next() {
		bar := iter.Bar()
		candles = append(candles, model.Candle{
			Symbol: a.Asset(),
			Time:   time.Unix(int64(bar.Timestamp), 0),
			Open:   func() float64 { f, _ := bar.Open.Float64(); return f }(),
			High:   func() float64 { f, _ := bar.High.Float64(); return f }(),
			Low:    func() float64 { f, _ := bar.Low.Float64(); return f }(),
			Close:  func() float64 { f, _ := bar.Close.Float64(); return f }(),
			Volume: float64(bar.Volume),
		})
	}

	return candles, iter.Err()
}
