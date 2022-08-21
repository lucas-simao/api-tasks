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
	TechnicianUser = entity.User{
		Name:     "lucas",
		Username: "lsimaoTasks",
		Password: "123456",
		CodeRole: entity.TechnicianRole,
	}
	ManagerUser = entity.User{
		Name:     "joão",
		Username: "joaoTasks",
		Password: "12345",
		CodeRole: entity.ManagerRole,
	}
	ManagerRoleId    int = 2
	TechnicianRoleId int = 3
)

func TestMain(m *testing.M) {
	port := "3322"
	newContainer := configs.ContainerRun(port)
	newContainer.RunMigrations("../../scripts/migrations")
	repo = New()
	DB = newContainer.DB

	TechnicianUser.Id = SignUpTechnician(TechnicianUser)
	ManagerUser.Id = SignUpManager(ManagerUser)

	code := m.Run()
	newContainer.ContainerDown()
	os.Exit(code)
}

func SignUpTechnician(t entity.User) int {
	result, err := DB.Exec(sqlSignUp, t.Name, t.Username, t.Password, TechnicianRoleId)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(id)
}

func SignUpManager(t entity.User) int {
	result, err := DB.Exec(sqlSignUp, t.Name, t.Username, t.Password, ManagerRoleId)
	if err != nil {
		log.Fatal(err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return int(id)
}
