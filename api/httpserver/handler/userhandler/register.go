package userhandler

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/httpstatus"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/response/httpresponse"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/infrastructure/richerror"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/validator/uservalidator"
	"net/http"
)

func (u *UserHandler) Register(c echo.Context) error {
	req := userdto.RegisterRequest{}

	// Bind
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// Validation
	if fieldErrors, err := uservalidator.New().ValidateRegisterRequest(req); err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.GetKind())

		return c.JSON(
			code,
			httpresponse.New().
				WithStatus(false).
				WithStatusCode(code).
				WithRequestID(c.Response().Header().Get(echo.HeaderXRequestID)).
				WithPath(c.Path()).
				WithExecutionDuration("123456789").
				WithMessage(richErr.GetMessage()).
				WithMeta(fieldErrors, richErr.GetMeta()).
				Build(),
		)
	}

	// Sanitize

	// UseCase
	resp, err := u.userInteractor.Register(req)
	if err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.GetKind())

		return c.JSON(
			code,
			httpresponse.New().
				WithStatus(false).
				WithStatusCode(code).
				WithRequestID(c.Response().Header().Get(echo.HeaderXRequestID)).
				WithPath(c.Path()).
				WithExecutionDuration("123456789").
				WithMessage(richErr.GetMessage()).
				WithMeta(richErr.GetMeta()).
				Build(),
		)
	}

	return c.JSON(
		http.StatusOK,
		httpresponse.New().
			WithStatus(true).
			WithStatusCode(http.StatusOK).
			WithRequestID(c.Response().Header().Get(echo.HeaderXRequestID)).
			WithPath(c.Path()).
			WithExecutionDuration("123456789").
			WithMessage("User is register successfully").
			WithMeta(resp).
			Build(),
	)
}
