package userusecase

import (
	"context"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/kind"
	"github.com/saeedjhn/go-backend-clean-arch/internal/infrastructure/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

func (i *Interactor) RefreshToken(ctx context.Context, req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error) {
	dto := userauthservicedto.ExtractIDFromTokenRequest{Token: req.RefreshToken}

	id, err := i.authIntr.ExtractIDFromRefreshToken(dto)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg403Forbidden).
			WithKind(kind.KindStatusBadRequest)
	}

	user, err := i.repository.GetByID(ctx, id.UserID)
	if err != nil {
		return userdto.RefreshTokenResponse{}, err
	}

	dto2 := userauthservicedto.CreateTokenRequest{User: user}

	accessToken, err := i.authIntr.CreateAccessToken(dto2)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg400BadRequest).
			WithKind(kind.KindStatusBadRequest)
	}

	refreshToken, err := i.authIntr.CreateRefreshToken(dto2)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg400BadRequest).
			WithKind(kind.KindStatusBadRequest)
	}

	return userdto.RefreshTokenResponse{
		AccessToken:  accessToken.Token,
		RefreshToken: refreshToken.Token,
	}, nil
}
