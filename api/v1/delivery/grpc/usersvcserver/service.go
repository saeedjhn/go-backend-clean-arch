package usersvcserver

import (
	"context"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"google.golang.org/grpc"

	pb "github.com/saeedjhn/go-backend-clean-arch/api/v1/proto/user/gen"
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

var _ pb.UserServiceServer = (*Service)(nil)

type Service struct {
	pb.UserServiceServer
	userInteractor Interactor
}

func New(itr Interactor) *Service {
	return &Service{userInteractor: itr}
}

func (u Service) Create(ctx context.Context, request *pb.CreateRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u Service) Get(ctx context.Context, request *pb.GetRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u Service) Profile(ctx context.Context, request *pb.ProfileRequest) (*pb.ProfileResponse, error) {
	// Fake User
	id, _ := strconv.ParseUint(request.GetId(), 10, 64)
	pReqDTO := userdto.ProfileRequest{ID: id}

	pRespDTO, err := u.userInteractor.Profile(pReqDTO)
	if err != nil {
		return nil, err
	}

	return &pb.ProfileResponse{
		User: &pb.User{
			Id:        strconv.FormatUint(pRespDTO.User.ID, 10),
			Name:      pRespDTO.User.Name,
			Mobile:    pRespDTO.User.Mobile,
			Email:     pRespDTO.User.Email,
			CreatedAt: timestamppb.New(pRespDTO.User.CreatedAt),
			UpdatedAt: timestamppb.New(pRespDTO.User.UpdatedAt),
		},
	}, nil
}

func (u Service) Update(ctx context.Context, request *pb.UpdateRequest) (*pb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u Service) Delete(ctx context.Context, request *pb.DeleteRequest) (*empty.Empty, error) {
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
