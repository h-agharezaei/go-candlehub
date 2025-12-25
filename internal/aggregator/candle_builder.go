package aggregator

import (
    "sort"
    "time"

    "candlehub/internal/model"
)

type Builder struct {
    symbol    string
    timeframe time.Duration
}

func NewBuilder(symbol string, timeframe time.Duration) *Builder {
    return &Builder{
        symbol:    symbol,
        timeframe: timeframe,
    }
}

func (b *Builder) Build(minuteCandles []model.Candle) []model.Candle {
    buckets := make(map[int64][]model.Candle)

    for _, c := range minuteCandles {
        key := c.Time.Truncate(b.timeframe).Unix()
        buckets[key] = append(buckets[key], c)
    }

    result := make([]model.Candle, 0)

    for _, group := range buckets {
        sort.Slice(group, func(i, j int) bool {
            return group[i].Time.Before(group[j].Time)
        })

        open := group[0].Open
        close := group[len(group)-1].Close
        high := open
        low := open
        volume := 0.0

        for _, c := range group {
            if c.High > high {
                high = c.High
            }
            if c.Low < low {
                low = c.Low
            }
            volume += c.Volume
        }

        result = append(result, model.Candle{
            Symbol:    b.symbol,
            Timeframe: b.timeframe.String(),
            Time:      group[0].Time.Truncate(b.timeframe),
            Open:      open,
            High:      high,
            Low:       low,
            Close:     close,
            Volume:    volume,
        })
    }

    return result
}
