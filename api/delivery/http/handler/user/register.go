package user //nolint:dupl // 1-79 lines are duplicate

import (
	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/bind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/sanitize"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

func (h *Handler) Register(c echo.Context) error {
	// Tracer
	ctx, span := h.trc.Span(
		c.Request().Context(), "HTTP POST register",
	)
	span.SetAttributes(attributes(c))

	defer span.End()

	// Bind
	req := user.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			echo.Map{
				"status":  false,
				"message": message.ErrorMsg400BadRequest,
				"errors":  bind.CheckErrorFromBind(err).Error(),
			},
		)
	}

	// Validation
	if fieldsErrs, err := h.vld.ValidateRegisterRequest(req); err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.MapkindToHTTPStatusCode(richErr.Kind())

		return echo.NewHTTPError(code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  fieldsErrs,
			},
		)
	}

	// Sanitize
	err := sanitize.New().
		SetPolicy(sanitize.StrictPolicy).
		Struct(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			echo.Map{
				"status":  false,
				"message": message.ErrorMsg400BadRequest,
				"errors":  nil,
			})
	}

	// Usage Use-case
	resp, err := h.userIntr.Register(ctx, req)
	if err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.MapkindToHTTPStatusCode(richErr.Kind())

		return echo.NewHTTPError(code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  richErr.Error(),
			})
	}

	return c.JSON(http.StatusCreated, resp)
}
