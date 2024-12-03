package usermapper

import (
	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"
	"strconv"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MapProfileRequestFromHTTP() {

}

func MapProfileResponseToHTTP() {}

func MapProfileRequestFromProtobuf(req *pb.ProfileRequest) userdto.ProfileRequest {
	id, _ := strconv.ParseUint(req.GetId(), 10, 64)

	return userdto.ProfileRequest{ID: id}
}

func MapProfileResponseToProtobuf(resp userdto.ProfileResponse) *pb.ProfileResponse {
	return &pb.ProfileResponse{User: &pb.User{
		Id:        strconv.FormatUint(resp.User.ID, 10),
		Name:      resp.User.Name,
		Mobile:    resp.User.Mobile,
		Email:     resp.User.Email,
		CreatedAt: timestamppb.New(resp.User.CreatedAt),
		UpdatedAt: timestamppb.New(resp.User.UpdatedAt),
	}}
}
