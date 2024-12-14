package oteltracer

import (
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var _ contract.Span = (*Span)(nil)

// Span wraps a trace.Span and provides additional utility methods
// to manage and enrich the span with attributes, events, and errors.
type Span struct {
	span trace.Span // The underlying OpenTelemetry span being wrapped.
}

// NewSpan creates a new Span instance by wrapping the provided trace.Span.
func NewSpan(span trace.Span) *Span {
	return &Span{span: span}
}

// SetAttributes sets multiple attributes on the span.
// Attributes are passed as a map of key-value pairs and are converted
// into OpenTelemetry-compatible attributes.
func (s *Span) SetAttributes(attrs map[string]interface{}) {
	var otelAttrs []attribute.KeyValue

	for k, v := range attrs {
		otelAttrs = append(otelAttrs, attribute.String(k, fmt.Sprintf("%#v", v)))
	}

	s.span.SetAttributes(otelAttrs...)
}

// AddEvent adds an event to the span, optionally with attributes.
// Attributes can be provided as one or more maps of key-value pairs.
// If no attributes are provided, only the event name is added.
func (s *Span) AddEvent(name string, attrs ...map[string]interface{}) {
	if len(attrs) != 0 {
		otelAttrs := s.convertToOtelAttributes(attrs)

		s.span.AddEvent(name, trace.WithAttributes(otelAttrs...))

		return
	}

	s.span.AddEvent(name)
}

// SetName updates the name of the span.
// This is useful when the span name needs to be dynamically changed
// based on runtime information.
func (s *Span) SetName(name string) {
	s.span.SetName(name)
}

// SetStatus sets the status of the span.
// The status is represented by a code (e.g., OK, ERROR) and a description
// providing additional context.
func (s *Span) SetStatus(code uint32, description string) {
	s.span.SetStatus(codes.Code(code), description)
}

// RecordError records an error on the span, optionally with attributes.
// Attributes can be provided as one or more maps of key-value pairs.
// This is useful for adding context to the error being recorded.
func (s *Span) RecordError(err error, attrs ...map[string]interface{}) {
	if len(attrs) != 0 {
		otelAttrs := s.convertToOtelAttributes(attrs)

		s.span.RecordError(err, trace.WithAttributes(otelAttrs...))

		return
	}

	s.span.RecordError(err)
}

// End ends the span, optionally enabling stack trace capture.
// If no argument is provided, stack trace capture is disabled by default.
func (s *Span) End(enableStackTrace ...bool) {
	stackTraceEnabled := false // Default value

	if len(enableStackTrace) > 0 {
		stackTraceEnabled = enableStackTrace[0]
	}

	s.span.End(trace.WithStackTrace(stackTraceEnabled))
}

// convertToOtelAttributes converts a slice of attribute maps into a slice of OpenTelemetry KeyValue attributes.
func (s *Span) convertToOtelAttributes(attrs []map[string]interface{}) []attribute.KeyValue {
	otelAttrs := make([]attribute.KeyValue, 0, len(attrs))

	for _, attr := range attrs {
		for k, v := range attr {
			otelAttrs = append(otelAttrs, attribute.String(k, fmt.Sprintf("%#v", v)))
		}
	}

	return otelAttrs
}
