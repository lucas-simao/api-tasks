package configs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

var (
	user       = "root"
	password   = "secret"
	database   = "api"
	dataSource = "%s:%s@tcp(localhost:%s)/%s?parseTime=true"
)

type Container struct {
	DB       *sqlx.DB
	pool     *dockertest.Pool
	resource *dockertest.Resource
}

func ContainerRun(port string) *Container {
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	opts := dockertest.RunOptions{
		Repository: "mysql",
		Tag:        "8",
		Env: []string{
			"MYSQL_DATABASE=" + database,
			"MYSQL_ROOT_PASSWORD=" + password,
		},
		ExposedPorts: []string{"3306"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"3306": {
				{HostIP: "0.0.0.0", HostPort: port},
			},
		},
		Cmd: []string{"--default-authentication-plugin=mysql_native_password"},
	}

	resource, err := pool.RunWithOptions(&opts, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err.Error())
	}

	time.Sleep(10 * time.Second)

	var mysqlUrl = fmt.Sprintf(dataSource, user, password, port, database)

	db, err := sqlx.Open("mysql", mysqlUrl)
	if err != nil {
		log.Panic(err)
		return &Container{}
	}

	err = resource.Expire(2 * 60)
	if err != nil {
		log.Print(err)
	}

	os.Setenv("DATABASE_URL", mysqlUrl)

	return &Container{
		db, pool, resource,
	}
}

func (c *Container) ContainerDown() {
	if err := c.pool.Purge(c.resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
}

func (c *Container) RunMigrations(migrationsDir string) {
	files, err := os.ReadDir(migrationsDir)
	if err != nil {
		log.Fatalf("Error read migrations dir: %v", err)
	}

	var extensionNameUpMigrations = ".up.sql"

	for i := range files {
		if strings.Contains(files[i].Name(), extensionNameUpMigrations) {
			fileData, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", migrationsDir, files[i].Name()))
			if err != nil {
				log.Fatalf("Error read file: %s, %v", files[i].Name(), err)
			}

			requests := strings.Split(strings.TrimSpace(string(fileData)), ";")

			for _, request := range requests {
				if len(request) > 5 {
					_, err = c.DB.Exec(request)
					if err != nil {
						log.Fatalf("Error run migrations: %v", err)
					}
				}
			}
		}
	}
}
