package handlers

import (
	"context"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/stretchr/testify/suite"
)

type TasksTestSuite struct {
	suite.Suite
	ctx context.Context
}

func TestUTasksTestSuite(t *testing.T) {
	suite.Run(t, new(TasksTestSuite))
}

func (suite *TasksTestSuite) SetupSuite() {
	suite.ctx = context.Background()
}

func (suite *TasksTestSuite) TearDownTest() {
	err := deleteTasks()
	suite.NoError(err)
}

func (suite *TasksTestSuite) TestCreateTask() {
	cases := map[string]struct {
		body       string
		user       entity.User
		statusCode int
	}{
		"1 - Should return 201": {
			body:       `{ "title": "test", "description": "test test"}`,
			user:       TechnicianUser,
			statusCode: http.StatusCreated,
		},
		"2 - Should return 400 - empty title": {
			body:       `{ "title": "", "description": "test test"}`,
			user:       TechnicianUser,
			statusCode: http.StatusBadRequest,
		},
		"3 - Should return 400 - empty description": {
			body:       `{ "title": "test", "description": ""}`,
			user:       TechnicianUser,
			statusCode: http.StatusBadRequest,
		},
		"4 - Should return 403 - user unauthorized to create tasks": {
			body:       `{ "title": "test", "description": "test test"}`,
			user:       ManagerUser,
			statusCode: http.StatusForbidden,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {

			c, rr := createContextAuth(http.MethodPost, "/tasks", strings.NewReader(cases[key].body), cases[key].user)

			handler := CreateTask(TasksService)

			err := handler(c)

			suite.NoError(err)

			suite.Equal(cases[key].statusCode, rr.Code, rr.Body)
		})
	}
}

func (suite *TasksTestSuite) TestGetTasks() {
	_, err := createTask("teste search", "this test should return test", TechnicianUser.Id)
	suite.NoError(err)

	cases := map[string]struct {
		user       entity.User
		statusCode int
	}{
		"1 - Should return 200": {
			user:       TechnicianUser,
			statusCode: http.StatusOK,
		},
		"2 - Should return 200 - Manager can see all tasks": {
			user:       ManagerUser,
			statusCode: http.StatusOK,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {

			c, rr := createContextAuth(http.MethodGet, "/tasks", nil, cases[key].user)

			handler := GetTasks(TasksService)

			err := handler(c)

			suite.NoError(err)

			suite.Equal(cases[key].statusCode, rr.Code, rr.Body)
		})
	}
}

func (suite *TasksTestSuite) TestGetTaskById() {
	taskId, err := createTask("teste search", "this test should return test", TechnicianUser.Id)
	suite.NoError(err)

	cases := map[string]struct {
		taskId     int
		user       entity.User
		statusCode int
	}{
		"1 - Should return 200": {
			taskId:     int(taskId),
			user:       TechnicianUser,
			statusCode: http.StatusOK,
		},
		"2 - Should return 200 - Manager can see any task": {
			user:       ManagerUser,
			taskId:     int(taskId),
			statusCode: http.StatusOK,
		},
		"2 - Should return 204 - task not exist": {
			user:       TechnicianUser,
			taskId:     0,
			statusCode: http.StatusNoContent,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {

			c, rr := createContextAuth(http.MethodGet, "/tasks/:id", nil, cases[key].user)

			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(cases[key].taskId))

			handler := GetTaskById(TasksService)

			err := handler(c)

			suite.NoError(err)

			suite.Equal(cases[key].statusCode, rr.Code, rr.Body)
		})
	}
}

func createTask(title, description string, userId int) (int64, error) {
	result, err := DB.Exec(`INSERT INTO tasks (title, description, created_by_user_id) VALUES(?, ?, ?)`, title, description, userId)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func deleteTasks() error {
	_, err := DB.Exec(`DELETE FROM tasks`)
	if err != nil {
		return err
	}

	return nil
}
