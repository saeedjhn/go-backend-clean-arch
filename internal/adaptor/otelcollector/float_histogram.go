package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) FloatHistogram(
	ctx context.Context,
	name string,
	count float64,
	description string,
	attrs ...map[string]interface{},
) error {
	otelAttrs := o.setAttrs(attrs)

	val, ok := o.counterCache.Load(name)
	if ok {
		val.(metric.Float64Histogram).Record(ctx, count, metric.WithAttributes(otelAttrs...)) //nolint:errcheck // nothing

		return nil
	}

	counter, err := o.meter.Float64Histogram(
		name,
		metric.WithExplicitBucketBoundaries(o.bucketBoundaries...),
		metric.WithDescription(description),
	)
	if err != nil {
		return fmt.Errorf("failed to create FloatHistogram for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Record(ctx, count, metric.WithAttributes(otelAttrs...))

	return nil
}
