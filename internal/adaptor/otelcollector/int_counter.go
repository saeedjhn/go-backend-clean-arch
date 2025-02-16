package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) IntCounter(
	ctx context.Context,
	name string,
	count int,
	description string,
	attrs ...map[string]interface{},
) error {
	count64 := int64(count)
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		val.(metric.Int64Counter).Add(ctx, count64, metric.WithAttributes(otelAttrs...)) //nolint:errcheck // nothing

		return nil
	}

	counter, err := o.meter.Int64Counter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create IntCounter for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Add(ctx, count64, metric.WithAttributes(otelAttrs...))

	return nil
}

func (o *OpenTelemetry) IntUpDownCounter(
	ctx context.Context,
	name string,
	count int,
	description string,
	attrs ...map[string]interface{},
) error {
	count64 := int64(count)
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		val.(metric.Int64UpDownCounter).Add(ctx, count64, metric.WithAttributes(otelAttrs...)) //nolint:errcheck // nothing

		return nil
	}

	counter, err := o.meter.Int64UpDownCounter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create IntUpDownCounter for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Add(ctx, count64, metric.WithAttributes(otelAttrs...))

	return nil
}
