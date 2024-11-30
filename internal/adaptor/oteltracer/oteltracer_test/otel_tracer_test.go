package oteltracer_test

import (
	"context"
	"errors"
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/oteltracer/codes"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/oteltracer"
)

func TestTrace(t *testing.T) {
	t.Parallel()

	cfg := oteltracer.Config{
		Endpoint: "localhost:4318", // Or localhost:4318
		AppName:  "TEST!",
		AppEnv:   "development",
	}

	ctx := context.Background()

	tracerClient := oteltracer.New(cfg)

	if err := tracerClient.Configure(); err != nil {
		t.Fatalf("OpenTelemetry: %v", err)
	}

	Handler(tracerClient)

	if err := tracerClient.Shutdown(ctx); err != nil {
		t.Logf("error shutdown: %v", err)
	}
}

func Handler(wrapper *oteltracer.OpenTelemetry) {
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// defer cancel()
	ctx := context.Background()

	spanCtx, span := wrapper.Span(ctx, "test")
	defer span.End(true)

	// IMPL ME
	Usecase(spanCtx, wrapper)
}

func Usecase(ctx context.Context, wrapper *oteltracer.OpenTelemetry) {
	attrs := map[string]interface{}{
		"user": map[string]string{
			"first_name": "John",
			"last_name":  "Doe",
		},
	}

	spanCtx, span := wrapper.Span(
		ctx,
		"test",
	)
	span.SetAttributes(attrs)
	span.SetName("Change.usecase")
	span.AddEvent("New.event", map[string]interface{}{"foo": "bar"})

	defer span.End()

	// IMPL ME
	Repository(spanCtx, wrapper)
}

func Repository(ctx context.Context, wrapper *oteltracer.OpenTelemetry) {
	attrs := map[string]interface{}{
		"user_id": 123,
	}
	_, span := wrapper.Span(
		ctx,
		"get.user",
	)
	span.SetAttributes(attrs)
	span.SetStatus(codes.Error.Uint32(), "Description.for.status")
	span.RecordError(errors.New("record.error"))

	defer span.End()
	// IMPL ME
}
