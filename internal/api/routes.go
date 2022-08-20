package api

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lucas-simao/api-tasks/internal/api/handlers"
	"github.com/lucas-simao/api-tasks/internal/entity"
)

func addRoutes(e *echo.Echo, s Services) {
	// public
	public := e.Group("")
	public.POST("/sign-up", handlers.SignUp(s.Users))
	public.POST("/sign-in", handlers.SignIn(s.Users))

	// authenticated
	auth := e.Group("")
	auth.Use(middleware.JWTWithConfig(JwtConfig()))

	auth.POST("/tasks", handlers.CreateTask(s.Tasks))
}

func JwtConfig() middleware.JWTConfig {
	secret := os.Getenv("JWT_SECRET")

	signingKey := []byte(secret)

	config := middleware.JWTConfig{
		Claims:     &entity.JwtCustomClaims{},
		SigningKey: signingKey,
	}

	return config
}
