package tasks

import (
	"context"

	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/lucas-simao/api-tasks/internal/repository"
)

type service struct {
	repository repository.Repository
}

func New(r repository.Repository) Service {
	return service{
		repository: r,
	}
}

func (s service) CreateTask(ctx context.Context, t entity.TaskRequest) (int64, error) {
	return s.repository.CreateTask(ctx, t)
}
