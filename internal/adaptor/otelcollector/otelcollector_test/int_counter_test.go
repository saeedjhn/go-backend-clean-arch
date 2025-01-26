package oterlcollector_test

import (
	"context"
	"testing"
)

func TestIntCounter_Increment_RecordsPositiveValues(t *testing.T) {
	for i := range 12 {
		err := _otelClient.IntCounter(
			context.Background(),
			"int_counter",
			100,
			"This is a description for increment operation",
		)
		if err != nil {
			t.Errorf("Error incrementing  counter at iteration %d: %v", i, err)
		}
	}
}

func TestIntUpDownCounter(t *testing.T) {
	t.Run("Increment_RecordsPositiveValues", func(t *testing.T) {
		for i := range 12 {
			err := _otelClient.IntUpDownCounter(
				context.Background(),
				"int_up_down_counter",
				10,
				"This is a description for increment operation",
			)
			if err != nil {
				t.Errorf("Error incrementing  counter at iteration %d: %v", i, err)
			}
		}
	})

	t.Run("Decrement_RecordsNegativeValues", func(t *testing.T) {
		for i := range 10 {
			err := _otelClient.IntUpDownCounter(
				context.Background(),
				"int_up_down_counter",
				-10,
				"This is a description for decrement operation",
			)
			if err != nil {
				t.Errorf("Error decrementing  counter at iteration %d: %v", i, err)
			}
		}
	})
}

func TestAsyncIntUpDownCounter(t *testing.T) {
	t.Run("Increment_RecordsPositiveValues", func(t *testing.T) {
		name := "async_int_up_down_counter"
		description := "This is a description for async int up/down counter"
		count := 100

		for i := range 5 {
			err := _otelClient.AsyncIntUpDownCounter(
				name,
				count,
				description,
			)
			if err != nil {
				t.Errorf("Error incrementing  intUpDown counter at iteration %d: %v", i, err)
				// t.Errorf("Error incrementing  intUpDown counter at iteration %v", err)
			}
		}
	})

	t.Run("Decrement_RecordsNegativeValues", func(t *testing.T) {
		name := "async_int_up_down_counter"
		description := "This is a description for async int up/down counter"
		count := -90

		err := _otelClient.AsyncIntUpDownCounter(
			name,
			count,
			description,
		)
		if err != nil {
			t.Errorf("Error creating AsyncIntUpDownCounter: %v", err)
		}
	})
}

func TestAsyncIntCounter_Increment_RecordsPositiveValues(t *testing.T) {
	name := "async_int_counter"
	description := "This is a description for async int counter"
	count := 100

	if err := _otelClient.AsyncIntCounter(
		name,
		count,
		description,
	); err != nil {
		t.Errorf("Error creating AsyncIntCounter: %v", err)
	}
}
