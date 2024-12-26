package grpc

import (
	"fmt"
	"net"

	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/grpc/interceptor"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/saeedjhn/go-backend-clean-arch/api/delivery/grpc/user"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"google.golang.org/grpc"
)

type Server struct {
	app *bootstrap.Application
}

func New(app *bootstrap.Application) *Server {
	return &Server{app: app}
}

func (s Server) Run() error {
	intercep := interceptor.New(s.app.Config, s.app.Logger)

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

	s.app.Logger.Infow("Server.GRPC.Start", "config", s.app.Config.GRPCServer)

	gs := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     s.app.Config.GRPCServer.MaxConnectionIdle,
		Timeout:               s.app.Config.GRPCServer.Timeout,
		MaxConnectionAge:      s.app.Config.GRPCServer.MaxConnectionAge,
		Time:                  s.app.Config.GRPCServer.Timeout,
		MaxConnectionAgeGrace: s.app.Config.GRPCServer.MaxConnectionAgeGrace,
	}),
		grpc.UnaryInterceptor(intercep.Logger),
		grpc.ChainUnaryInterceptor(
			grpcctxtags.UnaryServerInterceptor(),
			grpcprometheus.UnaryServerInterceptor,
			grpcrecovery.UnaryServerInterceptor(),
		),
	)

	// Register xxxServiceServer
	user.Register(s.app, gs)

	grpcprometheus.Register(gs)

	if s.app.Config.Application.Env != configs.Production {
		reflection.Register(gs)
	}

	if err = gs.Serve(listen); err != nil {
		return fmt.Errorf("gRPC server failed to start serving: %w", err)
	}

	return nil
}
