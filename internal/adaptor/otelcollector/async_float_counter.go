package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) AsyncFloatCounter( //nolint:dupl // 10-17 lines are duplicate
	name string,
	count float64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		counter := val.(metric.Float64ObservableCounter) //nolint:errcheck // nothing
		_, err := o.meter.RegisterCallback(
			func(_ context.Context, o metric.Observer) error {
				o.ObserveFloat64(counter, count, metric.WithAttributes(otelAttrs...))
				return nil
			}, counter,
		)
		return err
	}

	counter, err := o.meter.Float64ObservableCounter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create Float64ObservableCounter for %s: %w", name, err)
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

func (o *OpenTelemetry) AsyncFloatUpDownCounter( //nolint:dupl // 49-54 lines are duplicate
	name string,
	count float64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		counter := val.(metric.Float64ObservableUpDownCounter) //nolint:errcheck // nothing
		_, err := o.meter.RegisterCallback(
			func(_ context.Context, o metric.Observer) error {
				o.ObserveFloat64(counter, count, metric.WithAttributes(otelAttrs...))
				return nil
			}, counter,
		)
		return err
	}

	counter, err := o.meter.Float64ObservableUpDownCounter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create Float64ObservableUpDownCounter for %s: %w", name, err)
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
