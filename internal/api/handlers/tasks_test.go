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
		"5 - Should return 400 - description above 2500 characters": {
			body:       `{ "title": "test", "description": "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industrys standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum. Contrary to popular belief, Lorem Ipsum is not simply random text. It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old. Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source. Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of de Finibus Bonorum et Malorum (The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the theory of ethics, very popular during the Renaissance. The first line of Lorem Ipsum, Lorem ipsum dolor sit amet.., comes from a line in section 1.10.32. The standard chunk of Lorem Ipsum used since the 1500s is reproduced below for those interested. Sections 1.10.32 and 1.10.33 from de Finibus Bonorum et Malorum by Cicero are also reproduced in their exact original form, accompanied by English versions from the 1914 translation by H. Rackham. It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout. The point of using Lorem Ipsum is that it has a more-or-less normal distribution of letters, as opposed to using Content here, content here, making it look like readable English. Many desktop publishing packages and web page editors now use Lorem Ipsum as their default model text, and a search for lorem ipsum will uncover many web sites still in their infancy. Various versions have evolved over the years, sometimes by accident, sometimes on purpose (injected humour and the like). There are many variations of passages of Lorem Ipsum available, but the majority have suffered alteration in some form, by injected humour, or randomised words which dont look even slightly believable. If you are going to use a passage of Lorem Ipsum, you need to be sure there isnt anything embarrassing hidden in the middle of text. All the Lorem Ipsum generators on the Internet tend to repeat predefined chunks as necessary, making this the first true generator on the Internet. It uses a dictionary of over 200 Latin words, combined with a handful of model sentence structures, to generate Lorem Ipsum which looks reasonable. The generated Lorem Ipsum is therefore always free from repetition, injected humour, or non-characteristic words etc."}`,
			user:       ManagerUser,
			statusCode: http.StatusBadRequest,
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
		"3 - Should return 204 - task not exist": {
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

func (suite *TasksTestSuite) TestDeleteTaskById() {
	taskId, err := createTask("test delete task", "this test should delete it", TechnicianUser.Id)
	suite.NoError(err)

	cases := map[string]struct {
		taskId     int
		user       entity.User
		statusCode int
	}{
		"1 - Should return 200": {
			user:       ManagerUser,
			taskId:     int(taskId),
			statusCode: http.StatusOK,
		},
		"2 - Should return 400 - Technician can't delete tasks": {
			user:       TechnicianUser,
			taskId:     int(taskId),
			statusCode: http.StatusBadRequest,
		},
		"3 - Should return 204 - task has already been deleted": {
			user:       ManagerUser,
			taskId:     int(taskId),
			statusCode: http.StatusNoContent,
		},
		"4 - Should return 204 - task not exist": {
			user:       ManagerUser,
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

			handler := DeleteTaskById(TasksService)

			err := handler(c)

			suite.NoError(err)

			suite.Equal(cases[key].statusCode, rr.Code, rr.Body)
		})
	}
}

func (suite *TasksTestSuite) TestUpdateTaskById() {
	taskId, err := createTask("test update task", "this test should updated", TechnicianUser.Id)
	suite.NoError(err)

	cases := map[string]struct {
		taskId     int
		body       string
		user       entity.User
		statusCode int
	}{
		"1 - Should return 200": {
			user:       TechnicianUser,
			taskId:     int(taskId),
			body:       `{ "title": "test", "description": "test test"}`,
			statusCode: http.StatusOK,
		},
		"2 - Should return 400 - empty title": {
			user:       TechnicianUser,
			taskId:     int(taskId),
			body:       `{ "title": "", "description": "test test"}`,
			statusCode: http.StatusBadRequest,
		},
		"3 - Should return 400 - empty description": {
			user:       TechnicianUser,
			taskId:     int(taskId),
			body:       `{ "title": "test", "description": ""}`,
			statusCode: http.StatusBadRequest,
		},
		"4 - Should return 204 - task not exist": {
			user:       TechnicianUser,
			taskId:     0,
			body:       `{ "title": "test", "description": "test test"}`,
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

			c, rr := createContextAuth(http.MethodPut, "/tasks/:id", strings.NewReader(cases[key].body), cases[key].user)

			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(cases[key].taskId))

			handler := UpdateTaskById(TasksService)

			err := handler(c)

			suite.NoError(err)

			suite.Equal(cases[key].statusCode, rr.Code, rr.Body)
		})
	}
}

func (suite *TasksTestSuite) TestFinishTaskById() {
	taskId, err := createTask("test finish task", "this test should finish task", TechnicianUser.Id)
	suite.NoError(err)

	cases := map[string]struct {
		taskId     int
		user       entity.User
		statusCode int
	}{
		"1 - Should return 200": {
			user:       TechnicianUser,
			taskId:     int(taskId),
			statusCode: http.StatusOK,
		},
		"2 - Should return 204 - task not exist": {
			user:       TechnicianUser,
			taskId:     0,
			statusCode: http.StatusNoContent,
		},
		"3 - Should return 400 - without permission": {
			user:       ManagerUser,
			taskId:     int(taskId),
			statusCode: http.StatusBadRequest,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {

			c, rr := createContextAuth(http.MethodPatch, "/tasks/:id", nil, cases[key].user)

			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(cases[key].taskId))

			handler := FinishTaskById(TasksService)

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
