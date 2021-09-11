package dbRepositories

import (
	"database/sql"
	"log"
)

type Repository struct {
	DatabaseClient *sql.DB
}

func (r *Repository) InjectDatabaseClient(databaseClient *sql.DB) {
	r.DatabaseClient = databaseClient
}

func (r *Repository) StopDatabaseClientConnection() {
	r.DatabaseClient.Close()
	log.Println("Database client is closed")
}

func (r *Repository) PingDatabaseClient() error {
	if err := r.DatabaseClient.Ping(); err != nil {
		return err
	}

	return nil
}
