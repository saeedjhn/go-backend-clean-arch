package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) IntGauge(
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
		val.(metric.Int64Gauge).Record(ctx, count64, metric.WithAttributes(otelAttrs...))

		return nil
	}

	counter, err := o.meter.Int64Gauge(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create IntGauge for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Record(ctx, count64, metric.WithAttributes(otelAttrs...))

	return nil
}

func (o *OpenTelemetry) AsyncIntGauge(
	name string,
	count int,
	description string,
	attrs ...map[string]interface{},
) error {
	count64 := int64(count)
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		counter := val.(metric.Int64ObservableGauge)
		_, err := o.meter.RegisterCallback(
			func(_ context.Context, o metric.Observer) error {
				o.ObserveInt64(counter, count64, metric.WithAttributes(otelAttrs...))
				return nil
			}, counter,
		)
		return err
	}

	counter, err := o.meter.Int64ObservableGauge(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create AsyncIntGauge for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	_, err = o.meter.RegisterCallback(
		func(_ context.Context, o metric.Observer) error {
			o.ObserveInt64(counter, count64, metric.WithAttributes(otelAttrs...))
			return nil
		}, counter,
	)
	if err != nil {
		return err
	}

	return nil
}
