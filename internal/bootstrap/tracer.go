package bootstrap

import (
	"context"
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/buildinfo"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-backend-clean-arch/configs"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adapter/oteltracer"
)

func NewTracer(config *configs.Config, buildinf buildinfo.Info) (contract.Tracer, error) {
	client := oteltracer.New(oteltracer.Options{
		Config: oteltracer.Config{Endpoint: config.Tracer.Endpoint},
		AppInfo: oteltracer.AppInfo{
			Host:    config.HTTPServer.Host,
			Port:    config.HTTPServer.Port,
			Name:    config.Application.Name,
			Version: buildinf.Version,
			Env:     config.Application.Env.String(),
		},
	})

	if err := client.Configure(); err != nil {
		return nil, fmt.Errorf("failed to initialize tracer: %w", err)
	}

	return client, nil
}

func ShutdownTracer(ctx context.Context, trc contract.Tracer) error {
	return trc.Shutdown(ctx)
}
