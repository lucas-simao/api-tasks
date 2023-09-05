package entity

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/golang-jwt/jwt"
)

var (
	VisitorRole    = 0
	ManagerRole    = 10
	TechnicianRole = 20
)

type JwtCustomClaims struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	CodeRole int    `json:"codeRole"`
	jwt.StandardClaims
}

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

type TaskRequest struct {
	Title       string `json:"title" db:"title"`
	Description string `json:"description" db:"description"`
	UserId      int    `json:"-" db:"user_id"`
}

func (c TaskRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 100)),
		validation.Field(&c.Description, validation.Required, validation.Length(1, 2500)))
}

type TaskResponse struct {
	Id          int                       `json:"id"`
	Title       string                    `json:"title"`
	Description string                    `json:"description"`
	UpdatedAt   string                    `json:"updatedAt"`
	FinishedAt  string                    `json:"finishedAt"`
	CreatedBy   TaskUserOperationResponse `json:"createdBy"`
	DeletedBy   TaskUserOperationResponse `json:"deletedBy"`
}

type TaskUserOperationResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Date string `json:"date"`
}

type TaskUpdateRequest struct {
	Id          int    `json:"-"`
	UserId      int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (c TaskUpdateRequest) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Title, validation.Required, validation.Length(1, 100)),
		validation.Field(&c.Description, validation.Required, validation.Length(1, 2500)))
}
