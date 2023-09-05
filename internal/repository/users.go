package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

var (
	ErrUsernameUnavailable = errors.New("username is unavailable")
	ErrUserNotExist        = errors.New("user don't exist")
)

func (r *repository) SignUp(ctx context.Context, u entity.SignUpRequest) error {

	userRoleDefault, err := r.GetUserRoleByCode(ctx, entity.VisitorRole)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlSignUp, u.Name, u.Username, u.Password, userRoleDefault.Id)
	if err != nil {
		if strings.Contains(err.Error(), "users.username") {
			return ErrUsernameUnavailable
		}
		return err
	}

	return nil
}

func (r *repository) GetUserRoleByCode(ctx context.Context, code int) (entity.UserRole, error) {

	var u = entity.UserRole{}

	err := r.db.GetContext(ctx, &u, sqlGetUserRoleByCode, code)
	if err != nil {
		return entity.UserRole{}, err
	}

	return u, nil
}

func (r *repository) SignIn(ctx context.Context, username string) (entity.User, error) {

	var u = entity.User{}

	err := r.db.GetContext(ctx, &u, sqlGetUserByUsername, username)
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {
			return entity.User{}, ErrUserNotExist
		}
		return entity.User{}, err
	}

	return u, nil
}
