package grpc

import (
	"fmt"
	"net"

	"github.com/saeedjhn/go-backend-clean-arch/api/v1/delivery/grpc/userservice"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	App *bootstrap.Application
}

func New(app *bootstrap.Application) *Server {
	return &Server{App: app}
}

func (s Server) Run() error {

	addr := fmt.Sprintf(":%s", s.App.Config.GRPCServer.Port)

	listen, err := net.Listen(s.App.Config.GRPCServer.Network, addr)
	if err != nil {
		return fmt.Errorf(
			"failed to start listening on %s:%s due to: %w",
			s.App.Config.GRPCServer.Network,
			addr,
			err,
		)
	}

	s.App.Logger.Set().Named("GRPC.Server").Info("Start.Server", zap.Any("Server.Config", s.App.Config.GRPCServer))

	gs := grpc.NewServer()

	// Register xxxservices
	userservice.Register(gs)

	reflection.Register(gs)

	if err = gs.Serve(listen); err != nil {
		return fmt.Errorf("gRPC server failed to start serving: %w", err)
	}

	return nil
}
