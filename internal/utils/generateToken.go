package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/lucas-simao/api-tasks/internal/entity"
)

func GenerateToken(secret string, user entity.User) (string, error) {
	claims := &entity.JwtCustomClaims{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		CodeRole: user.CodeRole,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenSigned, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenSigned, nil
}
