package oterlcollector_test

import (
	"context"
	"testing"
)

//go:generate go test -v -race -count=1

func TestInt64Counter_Increment_RecordsPositiveValues(t *testing.T) {
	for i := range 12 {
		err := _otelClient.Int64Counter(
			context.Background(),
			"int_64_counter",
			100,
			"This is a description for increment operation",
		)
		if err != nil {
			t.Errorf("Error incrementing 64-bit counter at iteration %d: %v", i, err)
		}
	}

	t.Log("Metrics for increment recorded successfully!")
}

func TestInt64UpDownCounter(t *testing.T) {
	t.Run("Increment_RecordsPositiveValues", func(t *testing.T) {
		for i := range 12 {
			err := _otelClient.Int64UpDownCounter(
				context.Background(),
				"int_64_up_down_counter",
				10,
				"This is a description for increment operation",
			)
			if err != nil {
				t.Errorf("Error incrementing 64-bit counter at iteration %d: %v", i, err)
			}
		}
		t.Log("Metrics for increment recorded successfully!")
	})

	t.Run("Decrement_RecordsNegativeValues", func(t *testing.T) {
		for i := range 10 {
			err := _otelClient.Int64UpDownCounter(
				context.Background(),
				"int_64_up_down_counter",
				-10,
				"This is a description for decrement operation",
			)
			if err != nil {
				t.Errorf("Error decrementing 64-bit counter at iteration %d: %v", i, err)
			}
		}
		t.Log("Metrics for decrement recorded successfully!")
	})
}

func TestAsyncInt64Counter(t *testing.T) {
	t.Run("AsyncInt64Counter_RecordsValues", func(t *testing.T) {
		name := "async_int_64_counter"
		description := "This is a description for async int64 counter"
		count := int64(100)

		if err := _otelClient.AsyncInt64Counter(
			name,
			count,
			description,
		); err != nil {
			t.Errorf("Error creating AsyncInt64Counter: %v", err)
		}

		t.Log("AsyncInt64Counter created and observed successfully!")
	})
}

func TestAsyncInt64UpDownCounter(t *testing.T) {
	t.Run("Increment_RecordsPositiveValues", func(t *testing.T) {
		name := "async_int_64_up_down_counter"
		description := "This is a description for async int64 up/down counter"
		count := int64(100)

		for i := range 5 {
			err := _otelClient.AsyncInt64UpDownCounter(
				name,
				count,
				description,
			)
			if err != nil {
				t.Errorf("Error incrementing 64-bit intUpDown counter at iteration %d: %v", i, err)
				// t.Errorf("Error incrementing 64-bit intUpDown counter at iteration %v", err)
			}
		}

		t.Log("AsyncInt64UpDownCounter created and observed successfully!")
	})

	t.Run("Decrement_RecordsNegativeValues", func(t *testing.T) {
		name := "async_int_64_up_down_counter"
		description := "This is a description for async int64 up/down counter"
		count := int64(-90)

		err := _otelClient.AsyncInt64UpDownCounter(
			name,
			count,
			description,
		)
		if err != nil {
			t.Errorf("Error creating AsyncInt64UpDownCounter: %v", err)
		}

		t.Log("AsyncInt64UpDownCounter created and observed successfully!")
	})
}

func TestFloat64Counter_Increment_RecordsPositiveValues(t *testing.T) {
	for i := range 12 {
		err := _otelClient.Float64Counter(
			context.Background(),
			"float_64_counter",
			100.5,
			"This is a description for increment operation with Float64Counter",
		)
		if err != nil {
			t.Errorf("Error incrementing 64-bit float counter at iteration %d: %v", i, err)
		}
	}

	t.Log("Metrics for Float64Counter increment recorded successfully!")
}

func TestFloat64UpDownCounter(t *testing.T) {
	t.Run("Increment_RecordsPositiveValues", func(t *testing.T) {
		for i := range 12 {
			err := _otelClient.Float64UpDownCounter(
				context.Background(),
				"float_64_up_down_counter",
				10.5,
				"This is a description for increment operation with Float64UpDownCounter",
			)
			if err != nil {
				t.Errorf("Error incrementing 64-bit float counter at iteration %d: %v", i, err)
			}
		}
		t.Log("Metrics for increment recorded successfully!")
	})

	t.Run("Decrement_RecordsNegativeValues", func(t *testing.T) {
		for i := range 10 {
			err := _otelClient.Float64UpDownCounter(
				context.Background(),
				"float_64_up_down_counter",
				-10.5,
				"This is a description for decrement operation with Float64UpDownCounter",
			)
			if err != nil {
				t.Errorf("Error decrementing 64-bit float counter at iteration %d: %v", i, err)
			}
		}
		t.Log("Metrics for decrement recorded successfully!")
	})
}
