package contract

import "context"

type Collector interface {
	Configure() error // Configure initializes the tracer with the necessary settings.

	WithBucketBoundaries(bounds []float64) Collector

	IntCounter(
		ctx context.Context,
		name string,
		count int,
		description string,
		attrs ...map[string]interface{},
	) error
	IntUpDownCounter(
		ctx context.Context,
		name string,
		count int,
		description string,
		attrs ...map[string]interface{},
	) error
	IntGauge(
		ctx context.Context,
		name string,
		count int,
		description string,
		attrs ...map[string]interface{},
	) error
	IntHistogram(
		ctx context.Context,
		name string,
		count int,
		description string,
		attrs ...map[string]interface{},
	) error

	FloatCounter(
		ctx context.Context,
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error
	FloatUpDownCounter(
		ctx context.Context,
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error
	FloatGauge(
		ctx context.Context,
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error
	FloatHistogram(
		ctx context.Context,
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error

	AsyncIntCounter(
		name string,
		count int,
		description string,
		attrs ...map[string]interface{},
	) error
	AsyncIntUpDownCounter(
		name string,
		count int,
		description string,
		attrs ...map[string]interface{},
	) error
	AsyncIntGauge(
		name string,
		count int,
		description string,
		attrs ...map[string]interface{},
	) error
	AsyncFloatCounter(
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error
	AsyncFloatUpDownCounter(
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error
	AsyncFloatGauge(
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error

	Shutdown(ctx context.Context) error // Shutdown gracefully shuts down the tracer provider and flushes spans.
}
