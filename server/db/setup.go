package db

import (
	"database/sql"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	dbRepositories "github.com/wildanpurnomo/go-persistent-chat/server/db/repositories"
)

var (
	AppRepository *dbRepositories.Repository
)

func InitiateAppRepository() {
	AppRepository = &dbRepositories.Repository{}
}

func CreateNewDatabaseClient() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateUp() error {
	migrator, err := migrate.New(
		"file://db/migrations",
		os.Getenv("POSTGRES_URL"),
	)
	if err != nil {
		return err
	}
	if err = migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
