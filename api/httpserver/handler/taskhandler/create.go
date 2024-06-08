package taskhandler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (t *TaskHandler) Create(c echo.Context) error {

	return c.JSON(http.StatusCreated, echo.Map{
		"status":  true,
		"message": "CREATE",
		"data":    "",
	})
}
