package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/api-tasks/internal/domain/tasks"
	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/lucas-simao/api-tasks/internal/repository"
)

func CreateTask(s tasks.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		p := entity.TaskRequest{}

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

		session := GetAuthSession(c)

		p.UserId = session.Id

		if session.CodeRole != entity.TechnicianRole {
			result.Message = "user unauthorized to create tasks"
			return c.JSON(http.StatusForbidden, result)
		}

		id, err := s.CreateTask(ctx, p)
		if err != nil {
			result.Message = fmt.Sprintf("error to create task: %v", err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		return c.JSON(http.StatusCreated, map[string]int64{
			"id": id,
		})
	}
}

func GetTasks(s tasks.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		session := GetAuthSession(c)

		if session.Id == 0 {
			result.Message = "user unauthorized"
			return c.JSON(http.StatusBadRequest, result)
		}

		tasks, err := s.GetTasks(ctx, session.Id, session.CodeRole)
		if err != nil {
			result.Message = fmt.Sprintf("error to get tasks: %v", err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		var httpStatus int = http.StatusNoContent

		if len(tasks) > 0 {
			httpStatus = http.StatusOK
		}

		return c.JSON(httpStatus, tasks)
	}
}

func GetTaskById(s tasks.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		taskId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			result.Message = "error to parse id"
			return c.JSON(http.StatusBadRequest, result)
		}

		session := GetAuthSession(c)

		if session.Id == 0 {
			result.Message = "user unauthorized"
			return c.JSON(http.StatusBadRequest, result)
		}

		task, err := s.GetTaskById(ctx, taskId, session.Id, session.CodeRole)
		if err != nil {
			if errors.Is(err, repository.ErrNoTaskInResult) {
				return c.JSON(http.StatusNoContent, nil)
			}

			result.Message = fmt.Sprintf("error to get task by id: %v", err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		return c.JSON(http.StatusOK, task)
	}
}
