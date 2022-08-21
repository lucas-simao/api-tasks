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
	auth.GET("/tasks", handlers.GetTasks(s.Tasks))
	auth.GET("/tasks/:id", handlers.GetTaskById(s.Tasks))
	auth.DELETE("/tasks/:id", handlers.DeleteTaskById(s.Tasks))
	auth.PUT("/tasks/:id", handlers.UpdateTaskById(s.Tasks))
	auth.PATCH("/tasks/:id", handlers.FinishTaskById(s.Tasks))
}

func JwtConfig() middleware.JWTConfig {
	config := middleware.JWTConfig{
		Claims:     &entity.JwtCustomClaims{},
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}

	return config
}
