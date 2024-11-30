package bootstrap

import (
	"context"
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/oteltracer"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract/tracercontract"
)

func NewTracer(c oteltracer.Config) (tracercontract.Tracer, error) {
	tracerClient := oteltracer.New(c)

	if err := tracerClient.Configure(); err != nil {
		return nil, fmt.Errorf("failed to initialize tracing: %w", err)
	}

	return tracerClient, nil
}

func ShutdownTracer(ctx context.Context, trc tracercontract.Tracer) error {
	return trc.Shutdown(ctx)
}
