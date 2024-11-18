package userusecase

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
)

func (i *Interactor) Profile(ctx context.Context, req userdto.ProfileRequest) (userdto.ProfileResponse, error) {
	user, err := i.repository.GetByID(ctx, req.ID)
	if err != nil {
		return userdto.ProfileResponse{}, err
	}

	return userdto.ProfileResponse{User: userdto.UserInfo{
		ID:        user.ID,
		Name:      user.Name,
		Mobile:    user.Mobile,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}}, nil // Or
	// return userdto.ProfileResponse{
	//	User: user.ToUserInfoDTO(),
	// }, nil
}
