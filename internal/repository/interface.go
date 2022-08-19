package repository

import (
	"context"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

type Repository interface {
	SignUp(context.Context, entity.SignUpRequest) error
	SignIn(context.Context, string) (entity.User, error)
}
