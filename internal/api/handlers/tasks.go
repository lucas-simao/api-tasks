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

func DeleteTaskById(s tasks.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		taskId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			result.Message = "error to parse id"
			return c.JSON(http.StatusBadRequest, result)
		}

		session := GetAuthSession(c)

		if session.CodeRole != entity.ManagerRole {
			result.Message = "user don't have permission to delete tasks"
			return c.JSON(http.StatusBadRequest, result)
		}

		err = s.DeleteTaskById(ctx, taskId, session.Id)
		if err != nil {
			if errors.Is(err, repository.ErrNoTaskInResult) {
				return c.JSON(http.StatusNoContent, nil)
			}
			result.Message = fmt.Sprintf("error to delete task: %v", err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		result.Message = fmt.Sprintf("task %d deleted", taskId)
		return c.JSON(http.StatusOK, result)
	}
}

func UpdateTaskById(s tasks.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		p := entity.TaskUpdateRequest{}

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

		taskId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			result.Message = "error to parse id"
			return c.JSON(http.StatusBadRequest, result)
		}

		session := GetAuthSession(c)

		p.Id = taskId
		p.UserId = session.Id

		if session.CodeRole != entity.TechnicianRole {
			result.Message = "user don't have permission to update tasks"
			return c.JSON(http.StatusBadRequest, result)
		}

		task, err := s.UpdateTaskById(ctx, p)
		if err != nil {
			if errors.Is(err, repository.ErrNoTaskInResult) {
				return c.JSON(http.StatusNoContent, nil)
			}
			result.Message = fmt.Sprintf("error to update task: %v", err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		return c.JSON(http.StatusOK, task)
	}
}

func FinishTaskById(s tasks.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		taskId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			result.Message = "error to parse id"
			return c.JSON(http.StatusBadRequest, result)
		}

		session := GetAuthSession(c)

		if session.CodeRole != entity.TechnicianRole {
			result.Message = "user don't have permission to finish tasks"
			return c.JSON(http.StatusBadRequest, result)
		}

		task, err := s.FinishTaskById(ctx, taskId, session.Id)
		if err != nil {
			if errors.Is(err, repository.ErrNoTaskInResult) {
				return c.JSON(http.StatusNoContent, nil)
			}
			result.Message = fmt.Sprintf("error to finish task: %v", err)
			return c.JSON(http.StatusInternalServerError, result)
		}

		return c.JSON(http.StatusOK, task)
	}
}
