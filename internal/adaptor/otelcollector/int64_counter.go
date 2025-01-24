package otelcollector

import (
	"context"
	"fmt"
	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) Int64Counter(
	ctx context.Context,
	name string,
	count int64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		val.(metric.Int64Counter).Add(ctx, count, metric.WithAttributes(otelAttrs...))

		return nil
	}

	counter, err := o.meter.Int64Counter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create Int64Counter for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Add(ctx, count, metric.WithAttributes(otelAttrs...))

	return nil
}

func (o *OpenTelemetry) Int64UpDownCounter(
	ctx context.Context,
	name string,
	count int64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		val.(metric.Int64UpDownCounter).Add(ctx, count, metric.WithAttributes(otelAttrs...))

		return nil
	}

	counter, err := o.meter.Int64UpDownCounter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create Int64UpDownCounter for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Add(ctx, count, metric.WithAttributes(otelAttrs...))

	return nil
}
