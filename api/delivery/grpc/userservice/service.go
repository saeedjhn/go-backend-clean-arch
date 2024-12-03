package userservice

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"

	"github.com/golang/protobuf/ptypes/empty" //nolint:gomodguard // nothing
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/mapper/usermapper"
	"google.golang.org/grpc"
)

type Interactor interface {
	Register(ctx context.Context, req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	Login(ctx context.Context, req userdto.LoginRequest) (userdto.LoginResponse, error)
	Profile(ctx context.Context, req userdto.ProfileRequest) (userdto.ProfileResponse, error)
	Tasks(ctx context.Context, req userdto.TasksRequest) (userdto.TasksResponse, error)
	CreateTask(ctx context.Context, req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error)
	RefreshToken(ctx context.Context, req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error)
}

var _ gen.UserServiceServer = (*Service)(nil)

type Service struct {
	gen.UserServiceServer
	userIntr Interactor
}

func New(itr Interactor) *Service {
	return &Service{userIntr: itr}
}

func (u Service) Create(_ context.Context, _ *gen.CreateRequest) (*gen.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u Service) Get(_ context.Context, _ *gen.GetRequest) (*gen.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u Service) Profile(ctx context.Context, req *gen.ProfileRequest) (*gen.ProfileResponse, error) {
	r := usermapper.MapProfileRequestFromProtobuf(req)

	resp, err := u.userIntr.Profile(ctx, r)
	if err != nil {
		return &gen.ProfileResponse{}, err
	}

	return usermapper.MapProfileResponseToProtobuf(resp), nil
}

func (u Service) Update(_ context.Context, _ *gen.UpdateRequest) (*gen.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u Service) Delete(_ context.Context, _ *gen.DeleteRequest) (*empty.Empty, error) {
	// TODO implement me
	panic("implement me")
}

func (u Service) List(_ *empty.Empty, _ grpc.ServerStreamingServer[gen.User]) error {
	// TODO implement me
	panic("implement me")
}

// func (u Service) mustEmbedUnimplementedUserServiceServer() {
// TODO implement me
// 	panic("implement me")
// }
