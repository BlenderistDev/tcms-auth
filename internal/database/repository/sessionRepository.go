package repository

import "github.com/jmoiron/sqlx"

type sessionRepository struct {
	db *sqlx.DB
}

type SessionRepository interface {
}

// NewSessionRepository creates new session repository
func NewSessionRepository(db *sqlx.DB) SessionRepository {
	return &sessionRepository{db: db}
}
