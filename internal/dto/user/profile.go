package user

import (
	"strconv"

	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ProfileRequest struct {
	ID uint64 `json:"id"`
}

type ProfileResponse struct {
	Data UserInfo `json:"data"`
}

func (p ProfileResponse) ToProtobuf() *pb.ProfileResponse {
	return &pb.ProfileResponse{User: &pb.User{
		Id:        strconv.FormatUint(p.Data.ID, 10),
		Name:      p.Data.Name,
		Mobile:    p.Data.Mobile,
		Email:     p.Data.Email,
		CreatedAt: timestamppb.New(p.Data.CreatedAt),
		UpdatedAt: timestamppb.New(p.Data.UpdatedAt),
	}}
}
