package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/lucas-simao/api-tasks/internal/api"
	"github.com/lucas-simao/api-tasks/internal/domain/tasks"
	"github.com/lucas-simao/api-tasks/internal/domain/users"
	"github.com/lucas-simao/api-tasks/internal/gateway/notifications"
	"github.com/lucas-simao/api-tasks/internal/repository"
)

func main() {
	var isProduction = os.Getenv("IS_PRODUCTION")
	if isProduction == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Panic("Error to load .env in the root directory")
		}
	}

	// Database
	repo := repository.New()

	notifications := notifications.New()

	// Domains
	tasks := tasks.New(repo, notifications)
	users := users.New(repo)

	// Api
	a := api.New(api.Services{
		Tasks: tasks,
		Users: users,
	})
	api.Start(a)
}
