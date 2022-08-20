package repository

import (
	"log"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/lucas-simao/api-tasks/configs"
	"github.com/lucas-simao/api-tasks/internal/entity"
)

var (
	repo           Repository
	DB             *sqlx.DB
	TechnicianUser entity.User
)

func TestMain(m *testing.M) {
	newContainer := configs.ContainerRun()
	newContainer.RunMigrations("../../scripts/migrations")
	repo = New()
	DB = newContainer.DB

	var technicianId int = 3

	TechnicianUser = entity.User{
		Name:     "lucas",
		Username: "lsimaoTasks",
		Password: "123456",
		CodeRole: technicianId,
	}

	TechnicianUser.Id = signUpTechnician(TechnicianUser)

	code := m.Run()
	newContainer.ContainerDown()
	os.Exit(code)
}

func signUpTechnician(t entity.User) int {
	result, err := DB.Exec(sqlSignUp, t.Name, t.Username, t.Password, t.CodeRole)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(id)
}
