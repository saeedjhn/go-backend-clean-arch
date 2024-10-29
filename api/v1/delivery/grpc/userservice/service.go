package userservice

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/mapper/usermapper"
	"google.golang.org/grpc"

	pb "github.com/saeedjhn/go-backend-clean-arch/api/v1/proto/user/gen"
)

type Interactor interface {
	Register(ctx context.Context, req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	Login(ctx context.Context, req userdto.LoginRequest) (userdto.LoginResponse, error)
	Profile(ctx context.Context, req userdto.ProfileRequest) (userdto.ProfileResponse, error)
	Tasks(ctx context.Context, req userdto.TasksRequest) (userdto.TasksResponse, error)
	CreateTask(ctx context.Context, req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error)
	RefreshToken(ctx context.Context, req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error)
}

type Presenter interface {
	ProfileResponse(resp userdto.ProfileResponse) *pb.ProfileResponse
	Error(err error) error
}

var _ pb.UserServiceServer = (*Service)(nil)

type Service struct {
	pb.UserServiceServer
	present        Presenter
	userInteractor Interactor
}

func New(present Presenter, itr Interactor) *Service {
	return &Service{present: present, userInteractor: itr}
}

func (u Service) Create(ctx context.Context, req *pb.CreateRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u Service) Get(ctx context.Context, req *pb.GetRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u Service) Profile(ctx context.Context, req *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	resp, err := u.userInteractor.Profile(ctx, usermapper.MapProfileRequestFromProtobuf(req))

	if err != nil {
		return &pb.ProfileResponse{}, u.present.Error(err)
	}

	return u.present.ProfileResponse(resp), nil
}

func (u Service) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u Service) Delete(ctx context.Context, req *pb.DeleteRequest) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u Service) List(empty *empty.Empty, g grpc.ServerStreamingServer[pb.User]) error {
	//TODO implement me
	panic("implement me")
}

func (u Service) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
