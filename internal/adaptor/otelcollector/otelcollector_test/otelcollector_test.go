package oterlcollector_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/otelcollector"
)

var options = otelcollector.Options{
	Config: otelcollector.Config{Endpoint: "localhost:4317"},
	AppInfo: otelcollector.AppInfo{
		Host:    "localhost",
		Port:    "8000",
		Name:    "your-app-name",
		Version: "0.1.0",
		Env:     "development",
	},
}

func TestIncrementInt64Counter(t *testing.T) {
	client := otelcollector.New(options)
	if err := client.Configure(); err != nil {
		t.Fatalf("failed to initialize collector client: %v", err)
	}

	for i := 0; i < 12; i++ {
		err := client.IncrementInt64Counter(
			context.Background(),
			"increment_int_64_counter",
			100,
			"this is description",
			map[string]interface{}{"fo": "ba", "fo1": "ba"},
		)
		if err != nil {
			t.Fatalf("err increment 64: %v", err)
		}

	}

	if err := client.Shutdown(context.Background()); err != nil {
		t.Fatalf("failed to shutdown metric client: %v", err)
	}

	fmt.Println("Metrics recorded successfully!")
}
