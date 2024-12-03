package grpc

import (
	"fmt"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/grpc/userservice"
	"net"

	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"google.golang.org/grpc/reflection"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	app *bootstrap.Application
}

func New(app *bootstrap.Application) *Server {
	return &Server{app: app}
}

func (s Server) Run() error {

	addr := fmt.Sprintf(":%s", s.app.Config.GRPCServer.Port)

	listen, err := net.Listen(s.app.Config.GRPCServer.Network, addr)
	if err != nil {
		return fmt.Errorf(
			"failed to start listening on %s:%s due to: %w",
			s.app.Config.GRPCServer.Network,
			addr,
			err,
		)
	}

	s.app.Logger.Set().Named("GRPC.Server").Info("Start.Server", zap.Any("Server.Config", s.app.Config.GRPCServer))

	gs := grpc.NewServer()

	// Register xxxServiceServer
	userservice.Register(s.app, gs)

	if s.app.Config.Application.Env != configs.Production {
		reflection.Register(gs)
	}

	if err = gs.Serve(listen); err != nil {
		return fmt.Errorf("gRPC server failed to start serving: %w", err)
	}

	return nil
}
