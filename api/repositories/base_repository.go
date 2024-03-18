package repositories

import (
	"database/sql"
	"os"
)

type BaseRepository struct {
}

func (r *BaseRepository) GetConnection() (*sql.DB, error) {
	connection_string := os.Getenv("DB_DSN")

	db, err := sql.Open("postgres", connection_string)

	return db, err
}
