package user

import (
	"context"

	userusecase "github.com/saeedjhn/go-backend-clean-arch/internal/usecase/user"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/proto/user/gen"

	"github.com/golang/protobuf/ptypes/empty" //nolint:gomodguard // nothing
	"google.golang.org/grpc"
)

// type Interactor interface {
// 	Register(ctx context.Context, req user.RegisterRequest) (user.RegisterResponse, error)
// 	Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error)
// 	Profile(ctx context.Context, req user.ProfileRequest) (user.ProfileResponse, error)
// 	RefreshToken(ctx context.Context, req user.RefreshTokenRequest) (user.RefreshTokenResponse, error)
// }

var _ gen.UserServiceServer = (*Service)(nil)

type Service struct {
	gen.UserServiceServer
	userIntr userusecase.Interactor
}

func New(userInteractor userusecase.Interactor) Service {
	return Service{userIntr: userInteractor}
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
	r := MapProfileRequestFromProtobuf(req)

	resp, err := u.userIntr.Profile(ctx, r)
	if err != nil {
		return &gen.ProfileResponse{}, err
	}

	return MapProfileResponseToProtobuf(resp), nil
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
