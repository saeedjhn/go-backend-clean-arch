package oterlcollector_test

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/otelcollector"
	"log"
	"testing"
	"time"
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
