package user //nolint:dupl // 1-79 lines are duplicate

import (
	"net/http"

	"github.com/saeedjhn/go-domain-driven-design/internal/dto/user"
	"github.com/saeedjhn/go-domain-driven-design/internal/entity"
	"github.com/saeedjhn/go-domain-driven-design/pkg/bind"
	"github.com/saeedjhn/go-domain-driven-design/pkg/msg"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-domain-driven-design/pkg/httpstatus"
	"github.com/saeedjhn/go-domain-driven-design/pkg/richerror"
)

func (h *Handler) Register(c echo.Context) error {
	ctx, span := h.trc.Span(
		c.Request().Context(), "HTTP POST register",
	)
	span.SetAttributes(attributes(c))

	defer span.End()

	req := user.RegisterRequest{}
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			entity.NewErrorResponse(msg.ErrMsg400BadRequest, bind.CheckErrorFromBind(err).Error()).
				WithMeta(map[string]interface{}{"request": req}),
		)
	}

	resp, err := h.userIntr.Register(ctx, req)
	if err != nil {
		richErr := richerror.Analysis(err)
		code := httpstatus.MapkindToHTTPStatusCode(richErr.Kind())

		if resp.FieldErrors != nil {
			return echo.NewHTTPError(
				code,
				entity.NewErrorResponse(richErr.Message(), resp.FieldErrors).WithMeta(richErr.Meta()),
			)
		}

		return echo.NewHTTPError(
			code,
			entity.NewErrorResponse(richErr.Message(), richErr.Error()).WithMeta(richErr.Meta()),
		)
	}

	return c.JSON(http.StatusOK, entity.NewSuccessResponse(msg.MsgRegister, resp))
}
