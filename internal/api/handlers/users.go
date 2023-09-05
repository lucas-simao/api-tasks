package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/api-tasks/internal/domain/users"
	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/lucas-simao/api-tasks/internal/repository"
)

func SignUp(u users.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		p := entity.SignUpRequest{}

		err := c.Bind(&p)
		if err != nil {
			result.Message = fmt.Sprintf("error to bind body: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		err = p.Validate()
		if err != nil {
			result.Message = fmt.Sprintf("error to validate: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		err = u.SignUp(ctx, p)
		if err != nil {
			if errors.Is(err, repository.ErrUsernameUnavailable) {
				result.Message = fmt.Sprintf("%v", err)
				return c.JSON(http.StatusBadRequest, result)
			}

			result.Message = fmt.Sprintf("error to sign up: %v", err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Message = "success to sign up"
		return c.JSON(http.StatusCreated, result)
	}
}

func SignIn(u users.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		p := entity.SignInRequest{}

		err := c.Bind(&p)
		if err != nil {
			result.Message = fmt.Sprintf("error to bind body: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		err = p.Validate()
		if err != nil {
			result.Message = fmt.Sprintf("error to validate: %v", err)
			return c.JSON(http.StatusBadRequest, result)
		}

		token, err := u.SignIn(ctx, p)
		if err != nil {
			if errors.Is(err, users.ErrWrongPassword) {
				result.Message = fmt.Sprintf("%v", err)
				return c.JSON(http.StatusBadRequest, result)
			}

			if errors.Is(err, users.ErrUserWithoutValidRole) {
				result.Message = fmt.Sprintf("%v", err)
				return c.JSON(http.StatusForbidden, result)
			}

			result.Message = fmt.Sprintf("error to sign in: %v", err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}
