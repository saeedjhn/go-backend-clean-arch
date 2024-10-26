package userservice

import (
	pb "github.com/saeedjhn/go-backend-clean-arch/api/v1/proto/users/gen"
	"google.golang.org/grpc"
)

func Register(gs grpc.ServiceRegistrar) {
	pb.RegisterUserServiceServer(gs, &UserService{})
}
