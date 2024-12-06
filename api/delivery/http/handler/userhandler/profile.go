package userhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-backend-clean-arch/configs"
	"github.com/saeedjhn/go-backend-clean-arch/internal/domain/dto/userdto"
	"github.com/saeedjhn/go-backend-clean-arch/internal/presenter/httppresenter"
	"github.com/saeedjhn/go-backend-clean-arch/internal/usecase/authusecase"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/claim"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/httpstatus"
	"github.com/saeedjhn/go-backend-clean-arch/pkg/richerror"
)

func (h *Handler) Profile(c echo.Context) error {
	// Tracer
	ctx, span := h.trc.Span(
		c.Request().Context(), "HTTP POST profile",
	)
	span.SetAttributes(attributes(c))

	defer span.End()

	// Give claims
	claims := claim.GetClaimsFromEchoContext[authusecase.Claims](c, configs.AuthMiddlewareContextKey)

	// Usage Use-case
	resp, err := h.userIntr.Profile(
		ctx, userdto.ProfileRequest{ID: claims.UserID},
	)
	if err != nil {
		richErr, _ := richerror.Analysis(err)
		code := httpstatus.FromKind(richErr.Kind())

		h.app.Logger.Errorf("Error.Profile: %v", err)

		return echo.NewHTTPError(code,
			echo.Map{
				"status":  false,
				"message": richErr.Message(),
				"errors":  richErr.Error(),
			})
	}

	return c.JSON(http.StatusOK, httppresenter.New(
		httppresenter.WithData(resp.User),
	).ToMap())

	// return c.JSON(http.StatusOK, httppresenter.New().WithData(resp).ToMap())

	// return c.JSON(http.StatusOK, h.present.WithData(resp))

	// return c.JSON(http.StatusOK, h.present.Ok(resp))
}
