package user

import (
	"strconv"

	"github.com/saeedjhn/go-backend-clean-arch/internal/types"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapProfileRequestFromProtobuf(req *pb.ProfileRequest) user.ProfileRequest {
	id, _ := strconv.ParseUint(req.GetId(), 10, 64)

	return user.ProfileRequest{ID: types.ID(id)}
}

func MapProfileResponseToProtobuf(resp user.ProfileResponse) *pb.ProfileResponse {
	return &pb.ProfileResponse{User: &pb.User{
		Id:        strconv.FormatUint(resp.UserInfo.ID.Uint64(), 10),
		Name:      resp.UserInfo.Name,
		Mobile:    resp.UserInfo.Mobile,
		Email:     resp.UserInfo.Email,
		CreatedAt: timestamppb.New(resp.UserInfo.CreatedAt),
		UpdatedAt: timestamppb.New(resp.UserInfo.UpdatedAt),
	}}
}
