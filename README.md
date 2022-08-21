# api-tasks

#### This api is to account for maintenance tasks performed during a working day.
#### All documentation is located in [Exercise](https://github.com/lucas-simao/api-tasks/blob/staging/EXERCISE.md)
### Requirements
* [Go](https://golang.org/doc/install) >= 1.16
* [Docker](https://docs.docker.com/get-docker/)
* [Docker-compose](https://docs.docker.com/compose/)
* [DBeaver](https://dbeaver.io/download/) <b>or Database Tool of your choice</b>
* [Postman](https://www.postman.com/downloads/) <b><-Import postman collection from /scripts/API TASKS.postman_collection.json</b>

### See all help commands
```
make help
```

### Create database tables
* copy and execute all content inside the scripts/migrations/*.up.sql

### Run api - teste only
```
make copy-env   #Copy .env.example to home 
make api-up     #Run api
```

### Build and run api
```
make storage-up
make build-api
make run-api
```

### Project tree
````
.
├── EXERCISE.md
├── Makefile
├── README.md
├── configs
│   └── container.go
├── coverage.out
├── deployments
│   ├── Dockerfile
│   ├── deployment.yml
│   ├── docker-compose.yml
│   └── service.yml
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
