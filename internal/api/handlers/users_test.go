package handlers

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sort"
	"strings"
	"testing"

	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/stretchr/testify/suite"
)

type UsersTestSuite struct {
	suite.Suite
	ctx        context.Context
	userSignUp entity.SignUpRequest
	user       entity.User
}

func TestUsersTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}

func (suite *UsersTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	suite.user = entity.User{
		Name:     "lucas",
		Username: "lsimao",
		Password: "123456",
		CodeRole: 1,
	}

	suite.userSignUp = entity.SignUpRequest{
		Name:     suite.userSignUp.Name,
		Username: suite.userSignUp.Username,
		Password: suite.userSignUp.Password,
	}
}

func (suite *UsersTestSuite) TestSignUp() {
	cases := map[string]struct {
		body       string
		statusCode int
	}{
		"1 - Should return 201": {
			body:       `{ "name": "lucas", "username": "lsimao", "password": "123456"}`,
			statusCode: http.StatusCreated,
		},
		"2 - Should return 400 - empty name": {
			body:       `{ "name": "", "username": "lsimao", "password": "123456"}`,
			statusCode: http.StatusBadRequest,
		},
		"3 - Should return 400 - empty username": {
			body:       `{ "name": "lucas", "username": "", "password": "123456"}`,
			statusCode: http.StatusBadRequest,
		},
		"4 - Should return 400 - empty password": {
			body:       `{ "name": "lucas", "username": "lsimao", "password": ""}`,
			statusCode: http.StatusBadRequest,
		},
		"5 - Should return 400 - duplicate username": {
			body:       `{ "name": "lucas", "username": "lsimao", "password": ""}`,
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

			rr, err := singUp(cases[key].body)

			suite.NoError(err)

			suite.Equal(cases[key].statusCode, rr.Code, rr.Body)
		})
	}
}

func (suite *UsersTestSuite) TestSignIn() {
	var username = "lucasSimao"

	rr, err := singUp(fmt.Sprintf(`{ "name": "Lucas S Simao", "username": "%s", "password": "123456"}`, username))

	suite.NoError(err)

	suite.Equal(http.StatusCreated, rr.Code, rr.Body)

	var userId int

	err = DB.Get(&userId, `SELECT id FROM users WHERE username = ?`, username)
	suite.NoError(err)

	cases := map[string]struct {
		body       string
		changeRole bool
		statusCode int
	}{
		"1 - Should return 403 - without permission role": {
			body:       `{ "username": "lucasSimao", "password": "123456"} `,
			statusCode: http.StatusForbidden,
		},
		"2 - Should return 400 - empty username": {
			body:       `{ "username": "", "password": "123456"}`,
			statusCode: http.StatusBadRequest,
		},
		"3 - Should return 400 - empty password": {
			body:       `{ "username": "lucasSimao", "password": ""}`,
			statusCode: http.StatusBadRequest,
		},
		"4 - Should return 400 - wrong password": {
			body:       `{ "username": "lucasSimao", "password": "12345"}`,
			statusCode: http.StatusBadRequest,
		},
		"5 - Should return 200": {
			body:       `{ "username": "lucasSimao", "password": "123456"} `,
			changeRole: true,
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
			c, rr := createContext(http.MethodPost, "/sign-in", strings.NewReader(cases[key].body))

			handler := SignIn(UsersService)

			if cases[key].changeRole {
				_, err = DB.Exec(`UPDATE users SET user_role_id = 2 WHERE id = ?`, userId)
				suite.NoError(err)
			}

			err := handler(c)

			suite.NoError(err)

			suite.Equal(cases[key].statusCode, rr.Code, rr.Body)
		})
	}
}

func singUp(payload string) (*httptest.ResponseRecorder, error) {
	c, rr := createContext(http.MethodPost, "/sign-up", strings.NewReader(payload))

	handler := SignUp(UsersService)

	err := handler(c)

	return rr, err
}
