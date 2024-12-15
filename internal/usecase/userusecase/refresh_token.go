package userusecase

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/entity"

	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) RefreshToken(
	ctx context.Context,
	req userdto.RefreshTokenRequest,
) (userdto.RefreshTokenResponse, error) {
	ctx, span := i.trc.Span(ctx, "RefreshToken")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "RefreshToken",
		"usecase.request": req,
	})
	defer span.End()

	resp, err := i.authIntr.ParseToken(i.cfg.Auth.RefreshTokenSecret, req.RefreshToken)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg403Forbidden).
			WithKind(kind.KindStatusBadRequest)
	}

	user, err := i.repository.GetByID(ctx, resp.UserID)
	if err != nil {
		return userdto.RefreshTokenResponse{}, err
	}

	authenticable := entity.Authenticable{ID: user.ID}

	accessToken, err := i.authIntr.CreateAccessToken(authenticable)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg400BadRequest).
			WithKind(kind.KindStatusBadRequest)
	}

	refreshToken, err := i.authIntr.CreateRefreshToken(authenticable)
	if err != nil {
		return userdto.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg400BadRequest).
			WithKind(kind.KindStatusBadRequest)
	}

	return userdto.RefreshTokenResponse{
		Tokens: userdto.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
