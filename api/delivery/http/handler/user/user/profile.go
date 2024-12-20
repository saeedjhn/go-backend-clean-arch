package user

import (
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (h *Handler) Profile(c echo.Context) error {
	ctx, span := h.trc.Span(
		c.Request().Context(), "HTTP POST profile",
	)
	span.SetAttributes(attributes(c))

	defer span.End()

	claims := claim.GetClaimsFromEchoContext(c)

	resp, err := h.userIntr.Profile(
		ctx, user.ProfileRequest{ID: claims.UserID},
	)
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

	return c.JSON(http.StatusOK, resp)
}
