package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) FloatCounter(
	ctx context.Context,
	name string,
	count float64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		val.(metric.Float64Counter).Add(ctx, count, metric.WithAttributes(otelAttrs...)) //nolint:errcheck // nothing

		return nil
	}

	counter, err := o.meter.Float64Counter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create FloatCounter for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Add(ctx, count, metric.WithAttributes(otelAttrs...))

	return nil
}

func (o *OpenTelemetry) FloatUpDownCounter(
	ctx context.Context,
	name string,
	count float64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		val.(metric.Float64UpDownCounter).Add(ctx, count, metric.WithAttributes(otelAttrs...)) //nolint:errcheck // nothing

		return nil
	}

	counter, err := o.meter.Float64Counter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create FloatUpDownCounter for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Add(ctx, count, metric.WithAttributes(otelAttrs...))

	return nil
}
