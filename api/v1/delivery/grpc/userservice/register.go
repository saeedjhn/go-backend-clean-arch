package userservice

import (
	pb "github.com/saeedjhn/go-backend-clean-arch/api/v1/proto/user/gen"
	"github.com/saeedjhn/go-backend-clean-arch/internal/bootstrap"
	"github.com/saeedjhn/go-backend-clean-arch/internal/presenter/grpcpresenter/usergrpcpresenter"
	"google.golang.org/grpc"
)

func Register(app *bootstrap.Application, gs grpc.ServiceRegistrar) {
	p := usergrpcpresenter.New()

	us := New(p, app.Usecase.UserIntr)

	pb.RegisterUserServiceServer(gs, us)
}
