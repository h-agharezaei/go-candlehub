package main

import (
    "log"

    "candlehub/internal/adapters/yahoo"
    "candlehub/internal/scheduler"
)

func main() {
    log.Println("🕯️ CandleHub Collector started")

    goldAdapter := yahoo.NewGoldAdapter()
    sched := scheduler.NewScheduler(goldAdapter)

    sched.Start()
}
