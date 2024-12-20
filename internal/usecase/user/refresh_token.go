package user

import (
	"context"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"github.com/saeedjhn/go-backend-clean-arch/internal/entity"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/kind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (i *Interactor) RefreshToken(
	ctx context.Context,
	req user.RefreshTokenRequest,
) (user.RefreshTokenResponse, error) {
	ctx, span := i.trc.Span(ctx, "RefreshToken")
	span.SetAttributes(map[string]interface{}{
		"usecase.name":    "RefreshToken",
		"usecase.request": req,
	})
	defer span.End()

	if fieldsErrs, err := i.vld.ValidateRefreshTokenRequest(req); err != nil {
		return user.RefreshTokenResponse{FieldErrors: fieldsErrs}, err
	}

	resp, err := i.authIntr.ParseToken(i.cfg.Auth.RefreshTokenSecret, req.RefreshToken)
	if err != nil {
		return user.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg403Forbidden).
			WithKind(kind.KindStatusBadRequest)
	}

	u, err := i.repository.GetByID(ctx, resp.UserID)
	if err != nil {
		return user.RefreshTokenResponse{}, err
	}

	authenticable := entity.Authenticable{ID: u.ID}

	accessToken, err := i.authIntr.CreateAccessToken(authenticable)
	if err != nil {
		return user.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg400BadRequest).
			WithKind(kind.KindStatusBadRequest)
	}

	refreshToken, err := i.authIntr.CreateRefreshToken(authenticable)
	if err != nil {
		return user.RefreshTokenResponse{}, richerror.New(_opUserServiceRefreshToken).WithErr(err).
			WithMessage(message.ErrorMsg400BadRequest).
			WithKind(kind.KindStatusBadRequest)
	}

	return user.RefreshTokenResponse{
		Tokens: user.Tokens{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
