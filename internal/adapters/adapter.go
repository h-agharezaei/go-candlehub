package adapters

import (
    "time"

    "candlehub/internal/model"
)

type MarketAdapter interface {
    Asset() string
    FetchMinuteCandles(from time.Time) ([]model.Candle, error)
}