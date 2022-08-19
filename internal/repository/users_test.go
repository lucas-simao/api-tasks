package repository

import (
	"context"
	"sort"
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
		user entity.SignUpRequest
		err  error
	}{
		"1 - Should register user": {
			user: suite.userSignUp,
			err:  nil,
		},
		"2 - Should return error duplicate username": {
			user: suite.userSignUp,
			err:  ErrUsernameUnavailable,
		},
		"3 - Should return error": {
			user: entity.SignUpRequest{},
			err:  ErrUsernameUnavailable,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			err := repo.SignUp(suite.ctx, cases[key].user)
			if err != nil {
				suite.Equal(cases[key].err, err)
				return
			}
			suite.NoError(err)
			if cases[key].err == nil {
				userDB, err := repo.SignIn(suite.ctx, cases[key].user.Username)
				suite.NoError(err)

				suite.Equal(cases[key].user.Name, userDB.Name)
				suite.Equal(cases[key].user.Username, userDB.Username)
			}
		})
	}
}

func (suite *UsersTestSuite) TestSignIn() {
	cases := map[string]struct {
		user entity.SignUpRequest
		err  error
	}{
		"1 - Should return user": {
			user: entity.SignUpRequest{
				Name:     "João",
				Username: "joãoOtávio",
				Password: "123456",
			},
			err: nil,
		},
		"2 - Should return error": {
			user: entity.SignUpRequest{},
			err:  ErrUserNotExist,
		},
	}

	keys := make([]string, 0, len(cases))
	for v := range cases {
		keys = append(keys, v)
	}

	sort.Strings(keys)

	for _, key := range keys {
		suite.Run(key, func() {
			if cases[key].err == nil {
				err := repo.SignUp(suite.ctx, cases[key].user)
				suite.NoError(err)
			}

			userDB, err := repo.SignIn(suite.ctx, cases[key].user.Username)
			suite.Equal(cases[key].err, err)

			if cases[key].err == nil {
				suite.Equal(cases[key].user.Name, userDB.Name)
				suite.Equal(cases[key].user.Username, userDB.Username)
			}
		})
	}
}

func (suite *UsersTestSuite) TestGetUserRoleByCode() {
	cases := map[string]struct {
		codeRole int
		name     string
		err      error
	}{
		"1 - Should return role Visitor": {
			codeRole: entity.VisitorRole,
			name:     "visitor",
			err:      nil,
		},
		"1 - Should return role Manager": {
			codeRole: entity.ManagerRole,
			name:     "manager",
			err:      nil,
		},
		"1 - Should return role Technician": {
			codeRole: entity.TechnicianRole,
			name:     "technician",
			err:      nil,
		},
	}

	for name, test := range cases {
		suite.Run(name, func() {
			roleDB, err := repo.GetUserRoleByCode(suite.ctx, test.codeRole)
			suite.NoError(err)

			suite.Equal(test.name, roleDB.Name)
			suite.Equal(test.codeRole, roleDB.Code)
		})
	}
}
