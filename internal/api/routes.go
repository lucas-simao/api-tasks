package api

import (
	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/api-tasks/internal/api/handlers"
)

func addRoutes(e *echo.Echo, s Services) {
	e.POST("/sign-up", handlers.SignUp(s.Users))
	e.POST("/sign-in", handlers.SignIn(s.Users))
}
