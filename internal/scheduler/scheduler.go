package scheduler

import (
    "log"
    "time"

    "candlehub/internal/adapters"
    "candlehub/internal/aggregator"
)

type Scheduler struct {
    adapter adapters.MarketAdapter
}

func NewScheduler(adapter adapters.MarketAdapter) *Scheduler {
    return &Scheduler{adapter: adapter}
}

func (s *Scheduler) Start() {
    ticker := time.NewTicker(1 * time.Minute)
    defer ticker.Stop()

    lastFetch := time.Now().Add(-30 * time.Minute)

    builder15m := aggregator.NewBuilder(s.adapter.Asset(), 15*time.Minute)
    builder1h := aggregator.NewBuilder(s.adapter.Asset(), 1*time.Hour)

    for {
        <-ticker.C

        candles, err := s.adapter.FetchMinuteCandles(lastFetch)
        if err != nil {
            log.Println("fetch error:", err)
            continue
        }

        if len(candles) == 0 {
            continue
        }

        lastFetch = candles[len(candles)-1].Time

        c15 := builder15m.Build(candles)
        c1h := builder1h.Build(candles)

        log.Printf("[%s] 15m candles: %d | 1h candles: %d\n",
            s.adapter.Asset(), len(c15), len(c1h))
    }
}
