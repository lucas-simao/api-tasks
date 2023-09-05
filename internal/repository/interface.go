package repository

import (
	"context"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

type Repository interface {
	// authentication
	SignUp(context.Context, entity.SignUpRequest) error
	SignIn(context.Context, string) (entity.User, error)

	// roles
	GetUserRoleByCode(context.Context, int) (entity.UserRole, error)

	// tasks
	CreateTask(context.Context, entity.TaskRequest) (int64, error)
	GetTasks(context.Context, int, int) ([]entity.TaskResponse, error)
	GetTaskById(context.Context, int, int, int) (entity.TaskResponse, error)
	DeleteTaskById(context.Context, int, int) error
	UpdateTaskById(context.Context, entity.TaskUpdateRequest) (entity.TaskResponse, error)
	FinishTaskById(context.Context, int, int) (entity.TaskResponse, error)
}
