package bootstrap

import (
	"context"
	"fmt"
	"github.com/saeedjhn/go-backend-clean-arch/configs"

	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/oteltracer"
)

func NewTracer(
	config oteltracer.Config,
	appInfo configs.Application,
	serverInfo configs.HTTPServer,
) (contract.Tracer, error) {
	tracerClient := oteltracer.New(oteltracer.Options{
		Config: oteltracer.Config{Endpoint: config.Endpoint},
		AppInfo: oteltracer.AppInfo{
			Host:    serverInfo.Host,
			Port:    serverInfo.Port,
			Name:    appInfo.Name,
			Version: appInfo.Version,
			Env:     appInfo.Env.String(),
		},
	})

	if err := tracerClient.Configure(); err != nil {
		return nil, fmt.Errorf("failed to initialize tracing: %w", err)
	}

	return tracerClient, nil
}

func ShutdownTracer(ctx context.Context, trc contract.Tracer) error {
	return trc.Shutdown(ctx)
}
