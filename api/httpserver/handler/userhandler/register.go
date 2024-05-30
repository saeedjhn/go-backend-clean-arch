package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/bind"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/httpstatus"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/response/httpresponse"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/sanitize"
	"net/http"
)

func (u *UserHandler) Register(c echo.Context) error {
	// Initial
	req := userdto.RegisterRequest{}
	httpRes := httpresponse.New()

	// Bind
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusInternalServerError,
			httpRes.
				WithStatus(false).
				WithStatusCode(http.StatusInternalServerError).
				WithMessage("internal server error").
				WithError(bind.CheckErrFromBind(err).Error()).
				Build(),
		)
	}

	// Validation
	if fieldsErrs, err := u.userValidator.ValidateRegisterRequest(req); err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.GetKind())

		return c.JSON(
			code,
			httpRes.
				WithStatus(false).
				WithStatusCode(code).
				WithMessage(richErr.GetMessage()).
				WithError(fieldsErrs).
				Build(),
		)
	}

	// Sanitize
	sanitize.New().
		SetPolicy(sanitize.StrictPolicy).
		Struct(&req)

	// UseCase
	uRes, err := u.userInteractor.Register(req)
	if err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.GetKind())

		return c.JSON(
			code,
			httpRes.
				WithStatus(false).
				WithStatusCode(code).
				WithMessage(richErr.GetMessage()).
				WithMeta(richErr.GetMeta()).
				Build(),
		)
	}

	return c.JSON(
		http.StatusCreated,
		httpRes.
			WithStatus(true).
			WithStatusCode(http.StatusCreated).
			WithMessage("User is register successfully").
			WithMeta(uRes).
			Build(),
	)
}
