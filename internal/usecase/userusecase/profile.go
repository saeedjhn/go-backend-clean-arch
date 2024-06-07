package userusecase

import (
	"go-backend-clean-arch/internal/dto/userdto"
)

func (u *UserInteractor) Profile(req userdto.ProfileRequest) (userdto.ProfileResponse, error) {
	user, err := u.repository.GetByID(req.ID)
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
	}}, nil
}
