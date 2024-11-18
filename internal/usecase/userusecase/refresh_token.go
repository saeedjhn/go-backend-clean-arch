package userusecase

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/servicedto/userauthservicedto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) RefreshToken(ctx context.Context, req userdto.RefreshTokenRequest) (userdto.RefreshTokenResponse, error) {
	dto := userauthservicedto.ParseTokenRequest{
		Secret: i.config.Auth.RefreshTokenSecret,
		Token:  req.RefreshToken,
	}

	resp, err := i.authIntr.ParseToken(dto)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg403Forbidden).
			WithKind(kind.KindStatusBadRequest)
	}

	user, err := i.repository.GetByID(ctx, resp.Claims.UserID)
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
