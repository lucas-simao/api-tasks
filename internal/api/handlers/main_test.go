package handlers

import (
	"io"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/api-tasks/configs"
	"github.com/lucas-simao/api-tasks/internal/domain/tasks"
	"github.com/lucas-simao/api-tasks/internal/domain/users"
	"github.com/lucas-simao/api-tasks/internal/entity"
	"github.com/lucas-simao/api-tasks/internal/repository"
	"github.com/lucas-simao/api-tasks/internal/utils"
)

var (
	repo           repository.Repository
	DB             *sqlx.DB
	UsersService   users.Service
	TasksService   tasks.Service
	TechnicianUser = entity.User{
		Id:       1,
		Name:     "lucas",
		Username: "lsimaoTasksHandlers",
		Password: "123456",
		CodeRole: entity.TechnicianRole,
	}
	ManagerUser = entity.User{
		Id:       2,
		Name:     "joão",
		Username: "joaoTasksHandlers",
		Password: "123456789",
		CodeRole: entity.ManagerRole,
	}
)

func TestMain(m *testing.M) {
	port := "3321"
	newContainer := configs.ContainerRun(port)
	newContainer.RunMigrations("../../../scripts/migrations")
	repo = repository.New()
	DB = newContainer.DB

	UsersService = users.New(repo)
	TasksService = tasks.New(repo)

	// Register Technician
	signUpTechnician(TechnicianUser)
	// Register Manager
	signUpTManager(ManagerUser)

	code := m.Run()
	newContainer.ContainerDown()
	os.Exit(code)
}

func createContext(method, url string, body io.Reader) (c echo.Context, responseRecorder *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, url, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c = e.NewContext(req, rec)

	return c, rec
}

func createContextAuth(method, url string, body io.Reader, user entity.User) (c echo.Context, responseRecorder *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, url, body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c = e.NewContext(req, rec)

	secret := os.Getenv("JWT_SECRET")

	tokenSigned, err := utils.GenerateToken(secret, user)
	if err != nil {
		log.Fatal(err)
	}

	claims := entity.JwtCustomClaims{
		Id:       user.Id,
		Name:     user.Name,
		Username: user.Username,
		CodeRole: user.CodeRole,
	}

	token, err := jwt.ParseWithClaims(tokenSigned, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		log.Fatal(err)
	}

	c.Set("user", token)

	return c, rec
}

func signUpTechnician(t entity.User) {
	var technicianRoleId = 3
	_, err := DB.Exec(`INSERT INTO users (name, username, password, user_role_id) VALUES(?, ?, ?, ?)`, t.Name, t.Username, t.Password, technicianRoleId)
	if err != nil {
		log.Fatal(err)
	}
}

func signUpTManager(t entity.User) {
	var managerRoleId = 2
	_, err := DB.Exec(`INSERT INTO users (name, username, password, user_role_id) VALUES(?, ?, ?, ?)`, t.Name, t.Username, t.Password, managerRoleId)
	if err != nil {
		log.Fatal(err)
	}
}
