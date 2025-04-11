package user //nolint:dupl // 1-79 lines are duplicate

import (
	"net/http"

	"github.com/saeedjhn/go-backend-clean-arch/internal/sharedkernel/models"

	"github.com/saeedjhn/go-backend-clean-arch/internal/dto/user"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/bind"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/msg"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (h *Handler) RefreshToken(c echo.Context) error {
	ctx, span := h.trc.Span(
		c.Request().Context(), "HTTP POST refresh-token",
	)
	span.SetAttributes(attributes(c))

	defer span.End()

	req := user.RefreshTokenRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewErrorResponse(msg.ErrMsg400BadRequest, bind.CheckErrorFromBind(err).Error()).
				WithMeta(map[string]interface{}{"request": req}),
		)
	}

	resp, err := h.userIntr.RefreshToken(ctx, req)
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

	return c.JSON(http.StatusOK, models.NewSuccessResponse(msg.MsgRefreshTokenRecreated, resp))
}
