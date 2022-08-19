package repository

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func New() Repository {
	dataSource, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		log.Panic("Error to get DATABASE_URL")
	}

	driverName, ok := os.LookupEnv("DATABASE_TYPE")
	if !ok {
		log.Panic("Error to get DATABASE_TYPE")
	}

	newDb, err := sqlx.Connect(driverName, dataSource)
	if err != nil {
		log.Panic(err)
		return &repository{}
	}

	return &repository{
		db: newDb,
	}
}
