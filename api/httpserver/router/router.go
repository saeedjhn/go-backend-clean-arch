package router

import (
	"github.com/labstack/echo/v4"
	"go-backend-clean-arch-according-to-go-standards-project-layout/api/httpserver/router/userrouter"
	"go-backend-clean-arch-according-to-go-standards-project-layout/internal/bootstrap"
)

//func Setup(
//	cfg *configs.Config,
//	mysqlDB *mysql.MySqlDB,
//	postgresqlDB *postgresql.PostgresqlDB,
//	e *echo.Echo,
//) {
//	userrouter.New(cfg, postgresqlDB, e)
//	//taskrouter.New(cfg, mysqlDB, e)
//}

func Setup(
	app *bootstrap.Application,
	e *echo.Echo,
) {
	userrouter.New(app, e)
	//taskrouter.New(cfg, mysqlDB, e)
}
