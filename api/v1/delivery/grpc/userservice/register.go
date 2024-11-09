package userservice

import (
	pb "github.com/saeedjhn/go-backend-clean-arch/api/v1/proto/user/gen"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"google.golang.org/grpc"
)

func Register(app *bootstrap.Application, gs grpc.ServiceRegistrar) {
	us := New(app.Usecase.UserIntr)

	pb.RegisterUserServiceServer(gs, us)
}
