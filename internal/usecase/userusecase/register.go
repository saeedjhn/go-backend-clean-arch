package userusecase

import (
	"errors"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/domain"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/kind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/security/bcrypt"
	"go-backend-clean-arch-according-to-go-standards-project-layout/pkg/message"
)

func (u *UserInteractor) Register(req userdto.RegisterRequest) (userdto.RegisterResponse, error) {
	const op = message.OpUserUsecaseRegister

	isUnique, err := u.repository.IsMobileUnique(req.Mobile)
	if err != nil {
		return userdto.RegisterResponse{}, err
	}

	if !isUnique {
		return userdto.RegisterResponse{},
			richerror.New(op).
				WithErr(errors.New(message.ErrorMsgMobileIsNotUnique)).
				WithMessage(message.ErrorMsgInvalidInput).
				WithKind(kind.KindStatusBadRequest)
	}

	createUser := domain.User{
		Name:   req.Name,
		Mobile: req.Mobile,
	}

	encryptPass, err := bcrypt.Generate(req.Password, bcrypt.DefaultCost)
	createUser.Password = encryptPass

	createdUser, err := u.repository.Register(createUser)
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
