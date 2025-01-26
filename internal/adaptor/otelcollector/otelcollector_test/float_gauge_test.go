package oterlcollector_test

import (
	"context"
	"testing"
)

func TestFloatGauge_MemoryUsage(t *testing.T) {
	go simulateComputation(1000)

	for i := range 500 {
		if err := _otelClient.FloatGauge(
			context.Background(),
			"float_gauge_memory-usage",
			float64(getMemoryUsage()),
			"This is a description for memory usage",
		); err != nil {
			t.Errorf("Error incrementing  counter at iteration %d: %v", i, err)
		}

		// time.Sleep(2 * time.Second)
	}
	// time.Sleep(5 * time.Minute)
}
