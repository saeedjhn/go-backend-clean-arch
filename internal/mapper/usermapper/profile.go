package usermapper

import (
	"strconv"

	pb "github.com/saeedjhn/go-backend-clean-arch/api/proto/user/gen"

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
		Id:        strconv.FormatUint(resp.Data.ID, 10),
		Name:      resp.Data.Name,
		Mobile:    resp.Data.Mobile,
		Email:     resp.Data.Email,
		CreatedAt: timestamppb.New(resp.Data.CreatedAt),
		UpdatedAt: timestamppb.New(resp.Data.UpdatedAt),
	}}
}
