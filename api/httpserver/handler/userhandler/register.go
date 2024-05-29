package userhandler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/dto/userdto"
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
		var v *json.UnmarshalTypeError
		if errors.As(err, &v) {
			return c.JSON(http.StatusInternalServerError,
				httpRes.WithStatus(false).
					WithStatusCode(http.StatusInternalServerError).
					WithMessage("unmarshalling error type").
					WithMeta(echo.Map{
						"error": map[string]string{
							//v.Field: v.Error(), // json: cannot unmarshal number into Go struct field RegisterRequest.name of type string
							v.Field: fmt.Sprintf("cannot convert %s for name of type %s", v.Value, v.Type),
						},
					}).
					Build(),
			)
		}

		return c.JSON(http.StatusInternalServerError,
			httpRes.WithStatus(false).
				WithStatusCode(http.StatusInternalServerError).
				WithMessage("internal server error").
				Build(),
		)
	}

	// Validation
	if _, err := u.userValidator.ValidateRegisterRequest(req); err != nil {
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
