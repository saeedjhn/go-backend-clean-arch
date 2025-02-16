package interceptor

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const _count = 1

type Manager struct {
	cfg       *configs.Config
	logger    contract.Logger
	collector contract.Collector
}

func New(
	cfg *configs.Config,
	logger contract.Logger,
	collector contract.Collector,
) *Manager {
	return &Manager{
		logger:    logger,
		cfg:       cfg,
		collector: collector,
	}
}

func (im *Manager) Logger(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()

	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)

	im.logger.Infof("Method: %s, Time: %v, Metadata: %v, Err: %v", info.FullMethod, time.Since(start), md, err)

	return reply, err
}

// TODO: api.grpc.delivery.interceptor.metrics-checkErrorAnyCollectorMethod

func (im *Manager) Metrics(
	ctx context.Context,
	req interface{}, // gRPC request
	info *grpc.UnaryServerInfo, // gRPC method
	handler grpc.UnaryHandler, // handler to process gRPC request
) (interface{}, error) {
	start := time.Now()

	resp, err := handler(ctx, req)

	status := http.StatusOK
	attrs := map[string]interface{}{
		"grpc_status": status,
		"grpc_method": info.FullMethod,
		"grpc_server": info.Server,
	}

	if err != nil {
		richErr := richerror.Analysis(err)
		status = httpstatus.MapkindToHTTPStatusCode(richErr.Kind())

		attrs["grpc_status"] = status

		err = im.collector.IntCounter(
			ctx,
			"grpc_errors_total",
			_count,
			"Total number of gRPC errors, categorized by method and status code",
			attrs,
		)
		im.logErr(
			"grpc_errors_total",
			"Total number of gRPC errors, categorized by method and status code",
			err,
		)
	}

	err = im.collector.FloatHistogram(
		ctx,
		"grpc_request_duration_seconds",
		time.Since(start).Seconds(),
		"Duration of gRPC requests in seconds, categorized by method and status",
		attrs,
	)
	im.logErr(
		"grpc_request_duration_seconds",
		"Duration of gRPC requests in seconds, categorized by method and status",
		err,
	)

	err = im.collector.IntCounter(
		ctx,
		"grpc_requests_total",
		_count,
		"Total number of gRPC requests, categorized by method and status code",
		attrs,
	)
	im.logErr(
		"grpc_requests_total",
		"Total number of gRPC requests, categorized by method and status code",
		err,
	)

	nameIntCounter := fmt.Sprintf("grpc_method_%s_requests_total", info.FullMethod)
	descIntCounter := fmt.Sprintf("Total number of gRPC requests for the method %s", info.FullMethod)
	err = im.collector.IntCounter(
		ctx,
		nameIntCounter,
		_count,
		descIntCounter,
		attrs,
	)
	im.logErr(
		nameIntCounter,
		descIntCounter,
		err,
	)

	err = im.collector.IntGauge(
		ctx,
		"grpc_active_connections",
		_count,
		"number of active grpc connections",
		attrs,
	)
	im.logErr(
		"grpc_active_connections",
		"number of active grpc connections",
		err,
	)

	return resp, err
}

func (im *Manager) logErr(metricName string, description string, err error) {
	if err != nil {
		im.logger.Errorf(
			"GRPC.Interceptor.Collector.[Name: %s].[Description: %s].[Error: %v]",
			metricName,
			description,
			err,
		)
	}
}
