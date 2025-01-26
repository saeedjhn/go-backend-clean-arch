package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) FloatGauge(
	ctx context.Context,
	name string,
	count float64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		val.(metric.Float64Gauge).Record(ctx, count, metric.WithAttributes(otelAttrs...))

		return nil
	}

	counter, err := o.meter.Float64Gauge(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create FloatGauge for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Record(ctx, count, metric.WithAttributes(otelAttrs...))

	return nil
}

func (o *OpenTelemetry) AsyncFloatGauge(
	name string,
	count float64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		counter := val.(metric.Float64ObservableGauge)
		_, err := o.meter.RegisterCallback(
			func(_ context.Context, o metric.Observer) error {
				o.ObserveFloat64(counter, count, metric.WithAttributes(otelAttrs...))
				return nil
			}, counter,
		)
		return err
	}

	counter, err := o.meter.Float64ObservableGauge(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create AsyncFloatGauge for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	_, err = o.meter.RegisterCallback(
		func(_ context.Context, o metric.Observer) error {
			o.ObserveFloat64(counter, count, metric.WithAttributes(otelAttrs...))
			return nil
		}, counter,
	)
	if err != nil {
		return err
	}

	return nil
}
