package otelcollector

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/metric"
)

func (o *OpenTelemetry) IntHistogram(
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
		val.(metric.Int64Histogram).Record(ctx, count64, metric.WithAttributes(otelAttrs...)) //nolint:errcheck // nothing

		return nil
	}

	counter, err := o.meter.Int64Histogram(
		name,
		metric.WithExplicitBucketBoundaries(o.bucketBoundaries...),
		metric.WithDescription(description),
	)
	if err != nil {
		return fmt.Errorf("failed to create IntHistogram for %s: %w", name, err)
	}

	o.counterCache.Store(name, counter)
	counter.Record(ctx, count64, metric.WithAttributes(otelAttrs...))

	return nil
}
