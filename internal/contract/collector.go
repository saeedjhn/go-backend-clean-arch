package contract

import "context"

type Collector interface {
	Configure() error // Configure initializes the tracer with the necessary settings.
	// IncrementCounter(name string, value int, labels map[string]string)
	// RecordGauge(name string, value float64, labels map[string]string)
	// ObserveHistogram(name string, value float64, labels map[string]string)

	Int64Counter(
		ctx context.Context,
		name string,
		count int64,
		description string,
		attrs ...map[string]interface{},
	) error
	Int64UpDownCounter(
		ctx context.Context,
		name string,
		count int64,
		description string,
		attrs ...map[string]interface{},
	) error
	Float64Counter(
		ctx context.Context,
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error
	Float64UpDownCounter(
		ctx context.Context,
		name string,
		count float64,
		description string,
		attrs ...map[string]interface{},
	) error
	AsyncInt64Counter(
		name string,
		count int64,
		description string,
		attrs ...map[string]interface{},
	) error
	AsyncInt64UpDownCounter(
		name string,
		count int64,
		description string,
		attrs ...map[string]interface{},
	) error

	Shutdown(ctx context.Context) error // Shutdown gracefully shuts down the tracer provider and flushes spans.
}

// type AsyncOperation interface {
// 	AsyncInt64Counter(
// 		name string,
// 		count int64,
// 		description string,
// 		attrs ...map[string]interface{},
// 	) error
// 	AsyncInt64UpDownCounter(
// 		name string,
// 		count int64,
// 		description string,
// 		attrs ...map[string]interface{},
// 	) error
// }
