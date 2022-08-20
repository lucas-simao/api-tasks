package tasks

import (
	"context"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

type Service interface {
	CreateTask(context.Context, entity.TaskRequest) (int64, error)
}
