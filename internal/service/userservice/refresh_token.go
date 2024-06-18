package userservice

import (
	"go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"go-backend-clean-arch/internal/domain/dto/userdto"
	"go-backend-clean-arch/internal/infrastructure/kind"
	"go-backend-clean-arch/internal/infrastructure/richerror"
	"go-backend-clean-arch/pkg/message"
)

func (u *UserInteractor) RefreshToken(req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error) {
	const op = message.OpUserUsecaseRefreshToken

	dto := userauthservicedto.ExtractIDFromTokenRequest{Token: req.RefreshToken}

	id, err := u.authInteractor.ExtractIDFromRefreshToken(dto)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsgUserNotFound).
			WithKind(kind.KindStatusBadRequest)
	}

	user, err := u.repository.GetByID(id.UserID)
	if err != nil {
		return userdto.RefreshTokenResponse{}, err
	}

	dto2 := userauthservicedto.CreateTokenRequest{User: user}

	accessToken, err := u.authInteractor.CreateAccessToken(dto2)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg400BadRequest).
			WithKind(kind.KindStatusBadRequest)
	}

	refreshToken, err := u.authInteractor.CreateRefreshToken(dto2)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(op).WithErr(err).
			WithMessage(message.ErrorMsg400BadRequest).
			WithKind(kind.KindStatusBadRequest)
	}

	return userdto.RefreshTokenResponse{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	}, nil
}
