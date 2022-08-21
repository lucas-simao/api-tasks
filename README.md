# api-tasks

#### This api is to account for maintenance tasks performed during a working day.
#### All documentation is located in [Exercise](https://github.com/lucas-simao/api-tasks/blob/staging/EXERCISE.md)
### Requirements
* [Go](https://golang.org/doc/install) >= 1.16
* [Docker](https://docs.docker.com/get-docker/)
* [Docker-compose](https://docs.docker.com/compose/)
* [Postman](https://www.postman.com/downloads/) <b><-Import postman collection from /scripts/API TASKS.postman_collection.json</b>
* MySQL 8
### See all help commands
```
make help
```

### Run api
```
make copy-env
make api-up
```

### Project tree
````
├── EXERCISE.md
├── Makefile
├── README.md
├── configs
│   └── container.go
├── coverage.out
├── deployments
│   └── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── api
│   │   ├── api.go
│   │   ├── handlers
│   │   │   ├── handlers.go
│   │   │   ├── main_test.go
│   │   │   ├── tasks.go
│   │   │   ├── tasks_test.go
│   │   │   ├── users.go
│   │   │   └── users_test.go
│   │   └── routes.go
│   ├── domain
│   │   ├── tasks
│   │   │   ├── interface.go
│   │   │   └── tasks.go
│   │   └── users
│   │       ├── interface.go
│   │       └── users.go
│   ├── entity
│   │   └── users.go
│   ├── gateway
│   │   └── notifications
│   │       └── notifications.go
│   ├── repository
│   │   ├── interface.go
│   │   ├── main_test.go
│   │   ├── repository.go
│   │   ├── sql.go
│   │   ├── tasks.go
│   │   ├── tasks_test.go
│   │   ├── users.go
│   │   └── users_test.go
│   └── utils
│       └── generateToken.go
├── main.go
└── scripts
    ├── API TASKS.postman_collection.json
    └── migrations
        ├── 0001.down.sql
        └── 0001.up.sql
````
