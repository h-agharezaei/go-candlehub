package yahoo

import (
    "time"

    "github.com/piquette/finance-go/chart"

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
    params := &chart.Params{
        Symbol:   "XAUUSD=X",
        Interval: chart.Interval1Min,
        Period1:  from,
        Period2:  time.Now(),
    }

    iter := chart.Get(params)

    candles := make([]model.Candle, 0)
    for iter.Next() {
        bar := iter.Bar()
        candles = append(candles, model.Candle{
            Symbol: a.Asset(),
            Time:   time.Unix(bar.Timestamp, 0),
            Open:   bar.Open.InexactFloat64(),
            High:   bar.High.InexactFloat64(),
            Low:    bar.Low.InexactFloat64(),
            Close:  bar.Close.InexactFloat64(),
            Volume: bar.Volume.InexactFloat64(),
        })
    }

    return candles, iter.Err()
}
