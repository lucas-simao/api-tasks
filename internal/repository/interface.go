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
}
