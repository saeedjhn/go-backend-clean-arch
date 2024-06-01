package userusecase

import (
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/domain"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/security/bcrypt"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
)

func (u *UserInteractor) Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error) {
	const op = message.OpUserUsecaseRegister

	user := domain.User{
		Name:   req.Name,
		Mobile: req.Mobile,
	}

	encryptPass, err := bcrypt.Generate(req.Password, bcrypt.DefaultCost)
	user.Password = encryptPass

	createdUser, err := u.repository.Register(user)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	return userdto.RegisterResponse{
		User: userdto.UserInfo{
			ID:     createdUser.ID,
			Mobile: createdUser.Mobile,
			Name:   createdUser.Name,
		},
	}, nil
}
