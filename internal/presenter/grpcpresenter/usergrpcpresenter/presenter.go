package usergrpcpresenter

import (
	"fmt"
	"strconv"

	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"

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
		Id:        strconv.FormatUint(resp.Data.ID, 10),
		Name:      resp.Data.Name,
		Mobile:    resp.Data.Mobile,
		Email:     resp.Data.Email,
		CreatedAt: timestamppb.New(resp.Data.CreatedAt),
		UpdatedAt: timestamppb.New(resp.Data.UpdatedAt),
	}}
}

func (p Presenter) Error(err error) error {
	return fmt.Errorf("gRPC error: %w", err)
}
