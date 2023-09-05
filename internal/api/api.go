package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/lucas-simao/api-tasks/internal/domain/tasks"
	"github.com/lucas-simao/api-tasks/internal/domain/users"

	"os"
)

type Services struct {
	Tasks tasks.Service
	Users users.Service
}

var port string = "9000"

func New(s Services) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	addRoutes(e, s)

	return e
}

func Start(e *echo.Echo) {
	value, ok := os.LookupEnv("PORT")
	if ok {
		port = value
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
