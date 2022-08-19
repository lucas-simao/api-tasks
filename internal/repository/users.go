package repository

import (
	"context"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

func (r *repository) SignUp(ctx context.Context, u entity.SignUpRequest) error {

	userRoleDefault, err := r.GetUserRoleByCode(ctx, entity.VisitorRole)
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, sqlSignUp, u.Name, u.Username, u.Password, userRoleDefault.Id)
	if err != nil {
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
		return entity.User{}, err
	}

	return u, nil
}
