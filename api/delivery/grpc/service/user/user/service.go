package user

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/golang/protobuf/ptypes/empty" //nolint:gomodguard // nothing
	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"
	"google.golang.org/grpc"
)

type Interactor interface {
	Register(ctx context.Context, req user.RegisterRequest) (user.RegisterResponse, error)
	Login(ctx context.Context, req user.LoginRequest) (user.LoginResponse, error)
	Profile(ctx context.Context, req user.ProfileRequest) (user.ProfileResponse, error)
	RefreshToken(ctx context.Context, req user.RefreshTokenRequest) (user.RefreshTokenResponse, error)
}

var _ pb.UserServiceServer = (*Service)(nil)

type Service struct {
	pb.UserServiceServer
	userIntr Interactor
}

func New(userInteractor Interactor) *Service {
	return &Service{userIntr: userInteractor}
}

func (u Service) Create(_ context.Context, _ *pb.CreateRequest) (*pb.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u Service) Get(_ context.Context, _ *pb.GetRequest) (*pb.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u Service) Profile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	r := MapProfileRequestFromProtobuf(req)

	resp, err := u.userIntr.Profile(ctx, r)
	if err != nil {
		return &pb.ProfileResponse{}, err
	}

	return MapProfileResponseToProtobuf(resp), nil
}

func (u Service) Update(_ context.Context, _ *pb.UpdateRequest) (*pb.User, error) {
	// TODO implement me
	panic("implement me")
}

func (u Service) Delete(_ context.Context, _ *pb.DeleteRequest) (*empty.Empty, error) {
	// TODO implement me
	panic("implement me")
}

func (u Service) List(_ *empty.Empty, _ grpc.ServerStreamingServer[pb.User]) error {
	// TODO implement me
	panic("implement me")
}

// func (u Service) mustEmbedUnimplementedUserServiceServer() {
// TODO implement me
// 	panic("implement me")
// }
