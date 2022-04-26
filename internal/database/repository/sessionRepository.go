package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"tcms-auth/internal/database/session"
)

type sessionRepository struct {
	db *sqlx.DB
}

type SessionRepository interface {
	Create(userId int) (string, error)
}

// NewSessionRepository creates new session repository
func NewSessionRepository(db *sqlx.DB) SessionRepository {
	return &sessionRepository{db: db}
}

// Create creates new session for user
func (r *sessionRepository) Create(userId int) (string, error) {
	token := session.GenerateSessionToken()
	_, err := r.db.Exec("INSERT INTO user_session (user_id, token, timestamp) VALUES ($1, $2, $3)", userId, token, time.Now())
	if err != nil {
		return "", err
	}
	return token, nil
}
