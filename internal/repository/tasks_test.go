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

func (suite *TasksTestSuite) TestGetTasks() {

	_, err := repo.CreateTask(suite.ctx, entity.TaskRequest{
		Title:       "test1",
		Description: "test1 for test1",
		UserId:      TechnicianUser.Id,
	})
	suite.NoError(err)

	cases := map[string]struct {
		userId, roleCode int
		haveTasks        bool
		err              error
	}{
		"1 - Should return tasks": {
			userId:    TechnicianUser.Id,
			roleCode:  TechnicianUser.CodeRole,
			haveTasks: true,
			err:       nil,
		},
		"2 - Shouldn't return": {
			userId:   ManagerUser.Id,
			roleCode: ManagerUser.CodeRole,
			err:      nil,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			tasks, err := repo.GetTasks(suite.ctx, cases[key].userId, cases[key].roleCode)
			if err != nil {
				suite.Equal(cases[key].err, err)
				return
			}

			suite.NoError(err)

			if cases[key].haveTasks {
				suite.GreaterOrEqual(len(tasks), 1)
				suite.Equal(tasks[0].CreatedBy.Id, cases[key].userId)
			}
		})
	}
}

func (suite *TasksTestSuite) TestGetTaskById() {
	title := "test1"
	description := "test1 for test1"
	taskId, err := repo.CreateTask(suite.ctx, entity.TaskRequest{
		Title:       title,
		Description: description,
		UserId:      TechnicianUser.Id,
	})
	suite.NoError(err)

	cases := map[string]struct {
		userId, roleCode, taskId int
		err                      error
	}{
		"1 - Should return tasks": {
			userId:   TechnicianUser.Id,
			roleCode: TechnicianUser.CodeRole,
			taskId:   int(taskId),
			err:      nil,
		},
		"2 - Shouldn't return": {
			userId:   TechnicianUser.Id,
			roleCode: TechnicianUser.CodeRole,
			taskId:   0,
			err:      ErrNoTaskInResult,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			task, err := repo.GetTaskById(suite.ctx, cases[key].taskId, cases[key].userId, cases[key].roleCode)
			if cases[key].err != nil {
				suite.Equal(cases[key].err, err)
				return
			}

			suite.NoError(err)
			suite.Equal(cases[key].taskId, task.Id)
			suite.Equal(title, task.Title)
			suite.Equal(description, task.Description)
			suite.Equal(TechnicianUser.Id, task.CreatedBy.Id)
			suite.Equal(TechnicianUser.Name, task.CreatedBy.Name)
			suite.Equal(0, task.DeletedBy.Id)
			suite.Equal("", task.FinishedAt)
		})
	}
}

func (suite *TasksTestSuite) TestDeleteTaskById() {
	title := "test delete"
	description := "test2 for test2"
	taskId, err := repo.CreateTask(suite.ctx, entity.TaskRequest{
		Title:       title,
		Description: description,
		UserId:      TechnicianUser.Id,
	})
	suite.NoError(err)

	cases := map[string]struct {
		userId, taskId int
		err            error
	}{
		"1 - Should delete tasks": {
			userId: ManagerUser.Id,
			taskId: int(taskId),
			err:    nil,
		},
		"2 - Shouldn't delete - return error": {
			userId: ManagerUser.Id,
			taskId: 0,
			err:    ErrNoTaskInResult,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			err := repo.DeleteTaskById(suite.ctx, cases[key].taskId, cases[key].userId)
			if cases[key].err != nil {
				suite.Equal(cases[key].err, err)
				return
			}

			suite.NoError(err)
		})
	}
}

func (suite *TasksTestSuite) TestUpdateTaskById() {
	title := "test update"
	description := "test3 for test3"
	taskId, err := repo.CreateTask(suite.ctx, entity.TaskRequest{
		Title:       title,
		Description: description,
		UserId:      TechnicianUser.Id,
	})
	suite.NoError(err)

	cases := map[string]struct {
		updateData entity.TaskUpdateRequest
		err        error
	}{
		"1 - Should update task": {
			updateData: entity.TaskUpdateRequest{
				Id:          int(taskId),
				UserId:      TechnicianUser.Id,
				Title:       "New title",
				Description: "New description",
			},
			err: nil,
		},
		"2 - Shouldn't update - without valid user": {
			updateData: entity.TaskUpdateRequest{
				Id:          int(taskId),
				UserId:      0,
				Title:       "New title",
				Description: "New description",
			},
			err: ErrNoTaskInResult,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			task, err := repo.UpdateTaskById(suite.ctx, cases[key].updateData)
			if cases[key].err != nil {
				suite.Equal(cases[key].err, err)
				return
			}

			suite.NoError(err)
			suite.Equal(cases[key].updateData.Id, task.Id)
			suite.Equal(cases[key].updateData.Title, task.Title)
			suite.Equal(cases[key].updateData.Description, task.Description)
			suite.Equal(cases[key].updateData.UserId, task.CreatedBy.Id)
		})
	}
}

func (suite *TasksTestSuite) TestFinishTaskById() {
	title := "test update"
	description := "test3 for test3"
	taskId, err := repo.CreateTask(suite.ctx, entity.TaskRequest{
		Title:       title,
		Description: description,
		UserId:      TechnicianUser.Id,
	})
	suite.NoError(err)

	cases := map[string]struct {
		taskId, userId int
		err            error
	}{
		"1 - Should finish task": {
			taskId: int(taskId),
			userId: TechnicianUser.Id,
			err:    nil,
		},
		"2 - Shouldn't finish - is not same user task": {
			taskId: int(taskId),
			userId: 0,
			err:    ErrNoTaskInResult,
		},
		"3 - Shouldn't finish - without valid task id": {
			taskId: 0,
			userId: TechnicianUser.Id,
			err:    ErrNoTaskInResult,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			task, err := repo.FinishTaskById(suite.ctx, cases[key].taskId, cases[key].userId)
			if cases[key].err != nil {
				suite.Equal(cases[key].err, err)
				return
			}

			suite.NoError(err)
			suite.NotEmpty(task.FinishedAt)
		})
	}
}
