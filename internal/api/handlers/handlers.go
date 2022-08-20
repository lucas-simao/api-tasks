package handlers

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/api-tasks/internal/entity"
)

type ResultMessage struct {
	Message string `json:"message"`
}

var result ResultMessage

func GetAuthSession(c echo.Context) entity.User {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*entity.JwtCustomClaims)

	return entity.User{
		Id:       claims.Id,
		Name:     claims.Name,
		Username: claims.Username,
		CodeRole: claims.CodeRole,
	}
}
