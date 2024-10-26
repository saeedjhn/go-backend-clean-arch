package userservice

import (
	"context"
	"time"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/saeedjhn/go-backend-clean-arch/api/v1/proto/users/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Interactor interface {
	Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error)
	Login(req userdto.LoginRequest) (userdto.LoginResponse, error)
	Profile(req userdto.ProfileRequest) (userdto.ProfileResponse, error)
	Tasks(req userdto.TasksRequest) (userdto.TasksResponse, error)
	CreateTask(req userdto.CreateTaskRequest) (userdto.CreateTaskResponse, error)
	RefreshToken(req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error)
}

var _ pb.UserServiceServer = (*UserService)(nil)

type UserService struct {
	pb.UserServiceServer
	userInteractor Interactor
}

func (u UserService) CreateUser(ctx context.Context, request *pb.CreateUserRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.User, error) {
	// Fake User
	return &pb.User{
		Id:        1,
		Name:      "John",
		Mobile:    "Doe",
		Email:     "johndoe@gmail.com",
		CreatedAt: timestamppb.New(time.Now()),
		UpdatedAt: timestamppb.New(time.Now()),
	}, nil
}

func (u UserService) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*empty.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserService) ListUsers(e *empty.Empty, server pb.UserService_ListUsersServer) error {
	//TODO implement me
	panic("implement me")
}
