package oterlcollector_test

import (
	"context"
	"log"
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adapter/otelcollector"
)

var (
	_otelClient *otelcollector.OpenTelemetry //nolint:gochecknoglobals // nothing
	_options    = otelcollector.Options{     //nolint:gochecknoglobals // nothing
		Config: otelcollector.Config{Endpoint: "localhost:4317", Timeout: 5 * time.Minute},
		AppInfo: otelcollector.AppInfo{
			Host:    "localhost",
			Port:    "8000",
			Name:    "your-app-name",
			Version: "0.1.0",
			Env:     "development",
		},
	}
)

//go:generate go test -v -race -count=1

func TestMain(m *testing.M) {
	_otelClient = otelcollector.New(_options)

	if err := _otelClient.Configure(); err != nil {
		log.Panicf("Failed to initialize collector client: %v", err)
	}

	defer func() {
		if err := _otelClient.Shutdown(context.Background()); err != nil {
			log.Printf("Failed to shutdown metric client: %v", err)
		}
	}()

	m.Run()
}

func getMemoryUsage() int {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return int(memStats.Alloc) // Currently allocated memory in bytes
}

func randomNumber() int {
	rand.NewSource(time.Now().UnixNano())

	return rand.Intn(101)
}

func randomAPIResponse() float64 {
	return float64(time.Duration(rand.Intn(100)))
}

func simulateComputation(n int) {
	slice := make([]int, n)

	for k, v := range slice {
		_ = k * (2 + v)
		time.Sleep(500 * time.Millisecond)
	}
}
