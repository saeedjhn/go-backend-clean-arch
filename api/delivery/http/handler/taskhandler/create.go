package taskhandler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) Create(c echo.Context) error {
	// Tracer
	_, span := h.trc.Span(
		c.Request().Context(), "HTTP POST create",
	)
	span.SetAttributes(attributes(c))

	defer span.End()

	return c.JSON(http.StatusCreated, echo.Map{"data": ""})
}
