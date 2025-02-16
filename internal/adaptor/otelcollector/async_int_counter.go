package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) AsyncIntCounter( //nolint:dupl // 10-48 lines are duplicate
	name string,
	count int,
	description string,
	attrs ...map[string]interface{},
) error {
	count64 := int64(count)
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		counter := val.(metric.Int64ObservableCounter) //nolint:errcheck // nothing
		_, err := o.meter.RegisterCallback(
			func(_ context.Context, o metric.Observer) error {
				o.ObserveInt64(counter, count64, metric.WithAttributes(otelAttrs...))
				return nil
			}, counter,
		)
		return err
	}

	counter, err := o.meter.Int64ObservableCounter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create Int64ObservableCounter for %s: %w", name, err)
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

func (o *OpenTelemetry) AsyncIntUpDownCounter( //nolint:dupl // 50-89 lines are duplicate
	name string,
	count int,
	description string,
	attrs ...map[string]interface{},
) error {
	count64 := int64(count)
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		counter := val.(metric.Int64ObservableUpDownCounter) //nolint:errcheck // nothing
		_, err := o.meter.RegisterCallback(
			func(_ context.Context, o metric.Observer) error {
				o.ObserveInt64(counter, count64, metric.WithAttributes(otelAttrs...))
				return nil
			}, counter,
		)
		return err
	}

	counter, err := o.meter.Int64ObservableUpDownCounter(name, metric.WithDescription(description))
	if err != nil {
		return fmt.Errorf("failed to create Int64ObservableUpDownCounter for %s: %w", name, err)
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
