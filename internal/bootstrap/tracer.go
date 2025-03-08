package bootstrap

import (
	"context"
	"fmt"

	"github.com/saeedjhn/go-domain-driven-design/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-domain-driven-design/configs"

	"github.com/saeedjhn/go-domain-driven-design/internal/adaptor/oteltracer"
)

func NewTracer(
	config oteltracer.Config,
	appInfo configs.Application,
	serverInfo configs.HTTPServer,
) (contract.Tracer, error) {
	client := oteltracer.New(oteltracer.Options{
		Config: oteltracer.Config{Endpoint: config.Endpoint},
		AppInfo: oteltracer.AppInfo{
			Host:    serverInfo.Host,
			Port:    serverInfo.Port,
			Name:    appInfo.Name,
			Version: appInfo.Version,
			Env:     appInfo.Env.String(),
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
