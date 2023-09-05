package users

import (
	"context"
	"errors"
	"os"

	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/lucas-simao/api-tasks/internal/repository"
	"github.com/lucas-simao/api-tasks/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repository repository.Repository
}

func New(r repository.Repository) Service {
	return service{
		repository: r,
	}
}

var (
	ErrWrongPassword        = errors.New("wrong password")
	ErrUserWithoutValidRole = errors.New("user don't have valid role")
)

func (s service) SignUp(ctx context.Context, u entity.SignUpRequest) error {
	password, err := s.EncryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = password

	return s.repository.SignUp(ctx, u)
}

func (s service) SignIn(ctx context.Context, u entity.SignInRequest) (string, error) {
	userDB, err := s.repository.SignIn(ctx, u.Username)
	if err != nil {
		return "", err
	}

	if err := s.ComparePassword(userDB.Password, u.Password); err != nil {
		return "", ErrWrongPassword
	}

	if userDB.CodeRole == entity.VisitorRole {
		return "", ErrUserWithoutValidRole
	}

	secret := os.Getenv("JWT_SECRET")

	token, err := utils.GenerateToken(secret, userDB)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s service) EncryptPassword(password string) (string, error) {
	passByte := []byte(password)

	hashedPassword, err := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (s service) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
