package user

import (
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"

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

	resp, err := h.userIntr.Profile(ctx, user.ProfileRequest{ID: claims.UserID})
	if err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.MapkindToHTTPStatusCode(richErr.Kind())

		if resp.FieldErrors != nil {
			return echo.NewHTTPError(
				code,
				models.NewErrorResponse(richErr.Message(), resp.FieldErrors).WithMeta(richErr.Meta()),
			)
		}

		return echo.NewHTTPError(
			code,
			models.NewErrorResponse(richErr.Message(), richErr.Error()).WithMeta(richErr.Meta()),
		)
	}

	return c.JSON(http.StatusOK, models.NewSuccessResponse(msg.MsgProfileSeen, resp))
}
