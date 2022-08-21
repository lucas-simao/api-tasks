package repository

import (
	"context"
	"errors"
	"strings"

	"github.com/lucas-simao/api-tasks/internal/entity"
)

var (
	ErrTaskWithoutUser = errors.New("user id cannot be empty")
	ErrNoTaskInResult  = errors.New("task not found")
)

func (r *repository) CreateTask(ctx context.Context, t entity.TaskRequest) (int64, error) {
	result, err := r.db.ExecContext(ctx, sqlCreateTask, t.Title, t.Description, t.UserId)
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

func (r *repository) GetTasks(ctx context.Context, userId, roleCode int) ([]entity.TaskResponse, error) {
	var args []interface{}

	sql := sqlGetTasks

	if roleCode == entity.TechnicianRole {
		sql += ` AND created_by_user_id=?`
		args = append(args, userId)
	}

	rows, err := r.db.QueryContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	var tasks = []entity.TaskResponse{}

	for rows.Next() {
		t := entity.TaskResponse{}

		err := rows.Scan(
			&t.Id,
			&t.Title,
			&t.Description,
			&t.CreatedBy.Id,
			&t.CreatedBy.Name,
			&t.CreatedBy.Date,
			&t.DeletedBy.Id,
			&t.DeletedBy.Name,
			&t.DeletedBy.Date,
			&t.UpdatedAt,
			&t.FinishedAt,
		)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *repository) GetTaskById(ctx context.Context, taskId, userId, roleCode int) (entity.TaskResponse, error) {
	var args []interface{}

	sql := sqlGetTasks

	sql += ` AND t.id=?`
	args = append(args, taskId)

	if roleCode == entity.TechnicianRole {
		sql += ` AND created_by_user_id=?`
		args = append(args, userId)
	}

	t := entity.TaskResponse{}

	err := r.db.QueryRowContext(ctx, sql, args...).Scan(
		&t.Id,
		&t.Title,
		&t.Description,
		&t.CreatedBy.Id,
		&t.CreatedBy.Name,
		&t.CreatedBy.Date,
		&t.DeletedBy.Id,
		&t.DeletedBy.Name,
		&t.DeletedBy.Date,
		&t.UpdatedAt,
		&t.FinishedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return entity.TaskResponse{}, ErrNoTaskInResult
		}
		return entity.TaskResponse{}, err
	}

	return t, nil
}
