package bootstrap

import (
	"context"
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/internal/buildinfo"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/contract"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/otelcollector"
)

func NewCollector(config *configs.Config, buildinf buildinfo.Info) (contract.Collector, error) {
	client := otelcollector.New(otelcollector.Options{
		Config: otelcollector.Config{Endpoint: config.Collector.Endpoint},
		AppInfo: otelcollector.AppInfo{
			Host:    config.HTTPServer.Host,
			Port:    config.HTTPServer.Port,
			Name:    config.Application.Name,
			Version: buildinf.Version,
			Env:     config.Application.Env.String(),
		},
	})

	if err := client.Configure(); err != nil {
		return nil, fmt.Errorf("failed to initialize collector: %w", err)
	}

	return client, nil
}

func ShutdownCollector(ctx context.Context, collector contract.Collector) error {
	return collector.Shutdown(ctx)
}
