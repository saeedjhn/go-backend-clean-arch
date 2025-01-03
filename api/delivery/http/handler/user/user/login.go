package user //nolint:dupl // 1-79 lines are duplicate

import (
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/bind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/sanitize"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/message"
)

func (h *Handler) Login(c echo.Context) error {
	ctx, span := h.trc.Span(
		c.Request().Context(), "HTTP POST login",
	)
	span.SetAttributes(attributes(c))

	defer span.End()

	req := user.LoginRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest,
			echo.Map{
				"status":  false,
				"message": message.ErrorMsg400BadRequest,
				"errors":  bind.CheckErrorFromBind(err),
			})
	}

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

	resp, err := h.userIntr.Login(ctx, req)
	if err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.MapkindToHTTPStatusCode(richErr.Kind())

		if resp.FieldErrors != nil {
			return c.JSON(code, echo.Map{
				"message": richErr.Message(),
				"errors":  resp.FieldErrors,
			})
		}

		return echo.NewHTTPError(code,
			echo.Map{
				"message": richErr.Message(),
				"errors":  richErr.Error(),
			})
	}

	return c.JSON(http.StatusOK, resp)
}
