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
	hresp := httpresponse.New().
		WithRequestID(c.Response().Header().Get(echo.HeaderXRequestID)).
		WithPath(c.Path()).
		WithExecutionDuration("123456789")

	// Bind
	if err := c.Bind(&req); err != nil {
		var v *json.UnmarshalTypeError
		if errors.As(err, &v) {
			//m := echo.Map{
			//	"value":  v.Value,
			//	"type":   v.Type,
			//	"field":  v.Field,
			//	"struct": v.Struct,
			//	"offset": v.Offset,
			//	"error":  v.Error(),
			//}

			return c.JSON(http.StatusInternalServerError,
				hresp.WithStatus(false).
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

		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	// Validation
	if _, err := u.userValidator.ValidateRegisterRequest(req); err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.GetKind())

		return c.JSON(
			code,
			hresp.
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
	uresp, err := u.userInteractor.Register(req)
	if err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.GetKind())

		return c.JSON(
			code,
			hresp.
				WithStatus(false).
				WithStatusCode(code).
				WithMessage(richErr.GetMessage()).
				WithMeta(richErr.GetMeta()).
				Build(),
		)
	}

	return c.JSON(
		http.StatusOK,
		hresp.
			WithStatus(true).
			WithStatusCode(http.StatusOK).
			WithMessage("User is register successfully").
			WithMeta(uresp).
			Build(),
	)
}
