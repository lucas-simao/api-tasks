package handlers

import (
	"context"
	"net/http"
	"sort"
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
