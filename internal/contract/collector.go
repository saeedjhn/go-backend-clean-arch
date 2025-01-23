package contract

import "context"

type Collector interface {
	Configure() error // Configure initializes the tracer with the necessary settings.
	// IncrementCounter(name string, value int, labels map[string]string)
	// RecordGauge(name string, value float64, labels map[string]string)
	// ObserveHistogram(name string, value float64, labels map[string]string)
	IncrementInt64Counter(
		ctx context.Context,
		name string,
		incr int64,
		description string,
		attrs ...map[string]interface{},
	) error
	Shutdown(ctx context.Context) error // Shutdown gracefully shuts down the tracer provider and flushes spans.
}
