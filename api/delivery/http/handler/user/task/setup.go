package task

import (
	"github.com/labstack/echo/v4"
	"github.com/saeedjhn/go-domain-driven-design/internal/bootstrap"
	mysqltask "github.com/saeedjhn/go-domain-driven-design/internal/repository/mysql/task"
	authusecase "github.com/saeedjhn/go-domain-driven-design/internal/usecase/authentication"
	taskusecase "github.com/saeedjhn/go-domain-driven-design/internal/usecase/task"
)

func Setup(app *bootstrap.Application, e *echo.Echo) {
	// Dependencies
	repo := mysqltask.New(app.DB.MySQL)

	authIntr := authusecase.New(app.Config.Auth)
	taskIntr := taskusecase.New(app.Config, repo)

	New(app.Trc, authIntr, taskIntr).SetRoutes(e)
}
