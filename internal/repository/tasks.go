package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

var (
	ErrTaskWithoutUser = errors.New("user id cannot be empty")
)

func (r *repository) CreateTask(ctx context.Context, t entity.TaskRequest) (int64, error) {
	result, err := r.db.DB.ExecContext(ctx, sqlCreateTask, t.Title, t.Description, t.UserId)
	if err != nil {
		if strings.Contains(err.Error(), "user_id") {
			return 0, ErrTaskWithoutUser
		}
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
