package user

import (
	"strconv"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapProfileRequestFromHTTP() {

}

func MapProfileResponseToHTTP() {}

func MapProfileRequestFromProtobuf(req *pb.ProfileRequest) user.ProfileRequest {
	id, _ := strconv.ParseUint(req.GetId(), 10, 64)

	return user.ProfileRequest{ID: id}
}

func MapProfileResponseToProtobuf(resp user.ProfileResponse) *pb.ProfileResponse {
	return &pb.ProfileResponse{User: &pb.User{
		Id:        strconv.FormatUint(resp.Data.ID, 10),
		Name:      resp.Data.Name,
		Mobile:    resp.Data.Mobile,
		Email:     resp.Data.Email,
		CreatedAt: timestamppb.New(resp.Data.CreatedAt),
		UpdatedAt: timestamppb.New(resp.Data.UpdatedAt),
	}}
}
