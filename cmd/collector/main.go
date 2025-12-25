package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"candlehub/internal/adapters/yahoo"
	"candlehub/internal/scheduler"
)

func main() {
	log.Println("🕯️ CandleHub Collector started")

	// Create context with cancellation for graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	goldAdapter := yahoo.NewGoldAdapter()
	sched := scheduler.NewScheduler(goldAdapter)

	// Start scheduler in goroutine
	go func() {
		sched.Start(ctx)
	}()

	// Wait for shutdown signal
	sig := <-sigChan
	log.Printf("📵 Received signal: %s, shutting down gracefully...", sig)

	// Cancel context to stop scheduler
	cancel()

	// Create a timeout context for graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// Wait for graceful shutdown or timeout
	done := make(chan struct{})
	go func() {
		// Wait a bit for scheduler to clean up
		time.Sleep(100 * time.Millisecond)
		close(done)
	}()

	select {
	case <-done:
		log.Println("🛑 CandleHub Collector stopped gracefully")
	case <-shutdownCtx.Done():
		log.Println("⚠️ Shutdown timeout exceeded, forcing exit")
	}
}
