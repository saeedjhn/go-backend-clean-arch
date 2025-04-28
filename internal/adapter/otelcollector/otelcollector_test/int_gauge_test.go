package oterlcollector_test

import (
	"context"
	"testing"
)

func TestIntGauge_MemoryUsage(t *testing.T) {
	go simulateComputation(1000)

	for i := range 500 {
		if err := _otelClient.IntGauge(
			context.Background(),
			"int_gauge_memory-usage",
			getMemoryUsage(),
			"This is a description for memory usage",
		); err != nil {
			t.Errorf("Error incrementing  counter at iteration %d: %v", i, err)
		}

		// time.Sleep(time.Second)
	}
	// time.Sleep(5 * time.Minute)
}
