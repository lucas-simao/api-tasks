package tasks

import (
	"context"

	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/lucas-simao/api-tasks/internal/gateway/notifications"
	"github.com/lucas-simao/api-tasks/internal/repository"
)

type service struct {
	repository    repository.Repository
	notifications notifications.Notifications
}

func New(r repository.Repository, n notifications.Notifications) Service {
	return service{
		repository:    r,
		notifications: n,
	}
}

func (s service) CreateTask(ctx context.Context, t entity.TaskRequest) (int64, error) {
	return s.repository.CreateTask(ctx, t)
}

func (s service) GetTasks(ctx context.Context, userId, roleCode int) ([]entity.TaskResponse, error) {
	return s.repository.GetTasks(ctx, userId, roleCode)
}

func (s service) GetTaskById(ctx context.Context, taskId, userId, roleCode int) (entity.TaskResponse, error) {
	return s.repository.GetTaskById(ctx, taskId, userId, roleCode)
}

func (s service) DeleteTaskById(ctx context.Context, taskId, userId int) error {
	return s.repository.DeleteTaskById(ctx, taskId, userId)
}

func (s service) UpdateTaskById(ctx context.Context, task entity.TaskUpdateRequest) (entity.TaskResponse, error) {
	return s.repository.UpdateTaskById(ctx, task)
}

func (s service) FinishTaskById(ctx context.Context, taskId, userId int) (entity.TaskResponse, error) {
	task, err := s.repository.FinishTaskById(ctx, taskId, userId)

	go func() {
		if err == nil {
			s.notifications.NotifyManager(task)
		}
	}()

	return task, err
}
