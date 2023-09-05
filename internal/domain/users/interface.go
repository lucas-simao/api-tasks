package users

import (
	"context"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

type Service interface {
	SignUp(context.Context, entity.SignUpRequest) error
	SignIn(context.Context, entity.SignInRequest) (string, error)
}
