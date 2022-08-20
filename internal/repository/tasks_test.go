package repository

import (
	"context"
	"sort"
	"testing"

	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TasksTestSuite struct {
	suite.Suite
	ctx context.Context
}

func TestTasksTestSuite(t *testing.T) {
	suite.Run(t, new(TasksTestSuite))
}

func (suite *TasksTestSuite) SetupSuite() {
	suite.ctx = context.Background()
}

func (suite *TasksTestSuite) TestCreateTask() {
	cases := map[string]struct {
		task entity.TaskRequest
		err  error
	}{
		"1 - Should register task": {
			task: entity.TaskRequest{
				Title:       "test",
				Description: "test for test",
				UserId:      TechnicianUser.Id,
			},
			err: nil,
		},
		"2 - Should return error - empty userId": {
			task: entity.TaskRequest{
				Title:       "test",
				Description: "test for test",
			},
			err: ErrTaskWithoutUser,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			id, err := repo.CreateTask(suite.ctx, cases[key].task)
			if err != nil {
				suite.Equal(cases[key].err, err)
				return
			}

			suite.NoError(err)
			suite.Greater(id, int64(0))
		})
	}
}
