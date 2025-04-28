package oterlcollector_test

import (
	"context"
	"testing"
)

func TestFloatCounter_Increment_RecordsPositiveValues(t *testing.T) {
	for i := range 12 {
		err := _otelClient.FloatCounter(
			context.Background(),
			"float_counter",
			100.5,
			"This is a description for increment operation with FloatCounter",
		)
		if err != nil {
			t.Errorf("Error incrementing  float counter at iteration %d: %v", i, err)
		}
	}
}

func TestFloatUpDownCounter(t *testing.T) {
	t.Run("Increment_RecordsPositiveValues", func(t *testing.T) {
		for i := range 12 {
			err := _otelClient.FloatUpDownCounter(
				context.Background(),
				"float_up_down_counter",
				10.5,
				"This is a description for increment operation with FloatUpDownCounter",
			)
			if err != nil {
				t.Errorf("Error incrementing  float counter at iteration %d: %v", i, err)
			}
		}
	})

	t.Run("Decrement_RecordsNegativeValues", func(t *testing.T) {
		for i := range 10 {
			err := _otelClient.FloatUpDownCounter(
				context.Background(),
				"float_up_down_counter",
				-10.5,
				"This is a description for decrement operation with FloatUpDownCounter",
			)
			if err != nil {
				t.Errorf("Error decrementing float counter at iteration %d: %v", i, err)
			}
		}
	})
}

func TestAsyncFloatCounter_Increment_RecordsPositiveValues(t *testing.T) {
	name := "async_float_counter"
	description := "This is a description for async float counter"
	count := 100.0

	if err := _otelClient.AsyncFloatCounter(
		name,
		count,
		description,
	); err != nil {
		t.Errorf("Error creating AsyncFloatCounter: %v", err)
	}
}

func TestAsyncFloatUpDownCounter(t *testing.T) {
	t.Run("Increment_RecordsPositiveValues", func(t *testing.T) {
		name := "async_float_up_down_counter"
		description := "This is a description for async float up/down counter"
		count := 100.0

		for i := range 5 {
			err := _otelClient.AsyncFloatUpDownCounter(
				name,
				count,
				description,
			)
			if err != nil {
				t.Errorf("Error incrementing  floatUpDown counter at iteration %d: %v", i, err)
			}
		}
	})

	t.Run("Decrement_RecordsNegativeValues", func(t *testing.T) {
		name := "async_float_up_down_counter"
		description := "This is a description for async float up/down counter"
		count := -90.0

		err := _otelClient.AsyncFloatUpDownCounter(
			name,
			count,
			description,
		)
		if err != nil {
			t.Errorf("Error creating AsyncFloatUpDownCounter: %v", err)
		}
	})
}
