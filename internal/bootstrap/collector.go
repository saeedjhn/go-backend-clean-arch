package bootstrap

import (
	"context"
	"fmt"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/otelcollector"
)

func NewCollector(
	config otelcollector.Config,
	appInfo configs.Application,
	serverInfo configs.HTTPServer,
) (contract.Collector, error) {
	client := otelcollector.New(otelcollector.Options{
		Config: otelcollector.Config{Endpoint: config.Endpoint},
		AppInfo: otelcollector.AppInfo{
			Host:    serverInfo.Host,
			Port:    serverInfo.Port,
			Name:    appInfo.Name,
			Version: appInfo.Version,
			Env:     appInfo.Env.String(),
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
