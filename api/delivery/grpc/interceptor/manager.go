package interceptor

import (
	"context"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/contract"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Manager struct {
	cfg    *configs.Config
	logger contract.Logger
	// metr   metrics.Metrics
}

func New(
	cfg *configs.Config,
	logger contract.Logger,
	// metr metrics.Metrics,
) *Manager {
	return &Manager{
		logger: logger,
		cfg:    cfg,
		// metr:   metr,
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

// func (im *Manager) Metrics(
// 	ctx context.Context,
// 	req interface{},
// 	info *grpc.UnaryServerInfo,
// 	handler grpc.UnaryHandler,
// ) (interface{}, error) {
// 	start := time.Now()
//
// 	resp, err := handler(ctx, req)
//
// 	var status = http.StatusOK
// 	if err != nil {
// 		status = grpc_errors.MapGRPCErrCodeToHttpStatus(grpc_errors.ParseGRPCErrStatusCode(err))
// 	}
// 	im.metr.ObserveResponseTime(status, info.FullMethod, info.FullMethod, time.Since(start).Seconds())
// 	im.metr.IncHits(status, info.FullMethod, info.FullMethod)
//
// 	return resp, err
// }
