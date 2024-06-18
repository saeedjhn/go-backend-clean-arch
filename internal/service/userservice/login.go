package userservice

import (
	"go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"go-backend-clean-arch/internal/domain/dto/userdto"
	"go-backend-clean-arch/internal/domain/entity"
	"go-backend-clean-arch/internal/infrastructure/kind"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/internal/infrastructure/security/bcrypt"
	"go-backend-clean-arch/pkg/message"
)

func (u *UserInteractor) Login(req userdto.LoginRequest) (userdto.LoginResponse, error) {
	const op = message.OpUserUsecaseLogin

	user, err := u.repository.GetByMobile(req.Mobile)
	if err != nil {
		return userdto.LoginResponse{}, err
	}

	err = bcrypt.CompareHashAndSTR(user.Password, req.Password)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(op).WithErr(err).
			WithMessage(message.InvalidCredentials).
			WithKind(kind.KindStatusBadRequest)
	}

	dto := userauthservicedto.CreateTokenRequest{User: entity.User{
		ID:        user.ID,
		Name:      user.Name,
		Mobile:    user.Mobile,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}}

	accessToken, err := u.authInteractor.CreateAccessToken(dto)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	refreshToken, err := u.authInteractor.CreateRefreshToken(dto)
	if err != nil {
		return userdto.LoginResponse{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg500InternalServerError).
			WithKind(kind.KindStatusInternalServerError)
	}

	return userdto.LoginResponse{
		User: userdto.UserInfo{
			ID:        user.ID,
			Name:      user.Name,
			Mobile:    user.Mobile,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
		Token: userdto.Token{
			AccessToken:  accessToken.Token,
			RefreshToken: refreshToken.Token,
		},
	}, nil

}
