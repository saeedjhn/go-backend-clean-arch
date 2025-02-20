package csrf

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// const formCSRFToken = "_csrf"

// var tokenLookup = fmt.Sprintf("header:%s", echo.HeaderXCSRFToken)
// var tokenLookup = fmt.Sprintf("form:%s", formCSRFToken)

func (h *Handler) SetRoutes(e *echo.Echo) {
	group := e.Group("/csrf")
	{
		group.GET("", func(c echo.Context) error {
			csrfToken, _ := c.Get("csrf").(string)

			return c.JSON(http.StatusOK, echo.Map{
				"csrf_token": csrfToken, // X-CSRF-Token
			})
		}, middleware.CSRF())
	}
}
