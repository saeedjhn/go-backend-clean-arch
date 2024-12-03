package usergrpcpresenter

import (
	"fmt"
	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"
	"strconv"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Presenter struct {
}

func New() *Presenter {
	return &Presenter{}
}

func (p Presenter) ProfileResponse(resp userdto.ProfileResponse) *pb.ProfileResponse {
	return &pb.ProfileResponse{User: &pb.User{
		Id:        strconv.FormatUint(resp.User.ID, 10),
		Name:      resp.User.Name,
		Mobile:    resp.User.Mobile,
		Email:     resp.User.Email,
		CreatedAt: timestamppb.New(resp.User.CreatedAt),
		UpdatedAt: timestamppb.New(resp.User.UpdatedAt),
	}}
}

func (p Presenter) Error(err error) error {
	return fmt.Errorf("gRPC error: %w", err)
}
