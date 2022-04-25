package database

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

// GetConnection returns database connection
func GetConnection(url string) (*sqlx.DB, error) {
	return sqlx.Connect("pgx", url)
}
