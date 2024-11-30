package tracercontract

import (
	"context"
)

type Span interface {
	End(enableStackTrace ...bool)
	SetAttributes(attrs map[string]interface{})
	AddEvent(name string, attrs ...map[string]interface{})
	SetName(name string)
	SetStatus(code uint32, description string)
	RecordError(err error, attrs ...map[string]interface{})
}

// Tracer defines the methods required for tracing operations.
type Tracer interface {
	// Configure initializes the tracer with the necessary settings.
	Configure() error

	// Span creates and starts a new span with the given name and attributes.
	Span(ctx context.Context, name string) (context.Context, Span)

	// Shutdown gracefully shuts down the tracer provider and flushes spans.
	Shutdown(ctx context.Context) error

	// Optional methods

	// Inject injects the tracing context into a carrier (e.g., HTTP headers).
	// Inject(ctx context.Context, carrier map[string]string)

	// Extract extracts the tracing context from a carrier (e.g., HTTP headers).
	// Extract(ctx context.Context, carrier map[string]string) context.Context
}
