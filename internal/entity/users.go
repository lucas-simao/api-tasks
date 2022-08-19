package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

var (
	VisitorRole    = 0
	ManagerRole    = 10
	TechnicianRole = 20
)

type SignUpRequest struct {
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"name"`
}

func (c SignUpRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Name, validation.Required, validation.Length(1, 50)),
		validation.Field(&c.Username, validation.Required, validation.Length(1, 30)),
		validation.Field(&c.Password, validation.Required, validation.Length(1, 50)))
}

type UserRole struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
	Code int    `json:"code" db:"code"`
}

type SignInRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"name"`
}

func (c SignInRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Username, validation.Required, validation.Length(1, 30)),
		validation.Field(&c.Password, validation.Required, validation.Length(1, 50)))
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Username string `json:"username" db:"username"`
	CodeRole int    `json:"codeRole" db:"code_role"`
	Password string `json:"-" db:"password"`
}
