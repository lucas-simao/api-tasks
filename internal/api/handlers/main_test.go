package handlers

import (
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/lucas-simao/api-tasks/configs"
	"github.com/lucas-simao/api-tasks/internal/domain/users"
	"github.com/lucas-simao/api-tasks/internal/repository"
)

var (
	repo         repository.Repository
	DB           *sqlx.DB
	UsersService users.Service
)

func TestMain(m *testing.M) {
	newContainer := configs.ContainerRun()
	newContainer.RunMigrations("../../../scripts/migrations")
	repo = repository.New()
	DB = newContainer.DB

	UsersService = users.New(repo)

	code := m.Run()
	newContainer.ContainerDown()
	os.Exit(code)
}

func RecordedEchoContext(method, url string, body io.Reader) (c echo.Context, responseRecorder *httptest.ResponseRecorder) {
	e := echo.New()
	request := httptest.NewRequest(method, url, body)
	request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := request.Context()

	// if userID != uuid.Nil {
	// 	session := auth.AuthSession{UserID: userID, Phone: "12345"}
	// 	ctx = auth.SetAuthSession(request.Context(), session)
	// }

	responseRecorder = httptest.NewRecorder()
	c = e.NewContext(request.WithContext(ctx), responseRecorder)

	return c, responseRecorder
}
