package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"tcms-auth/internal/database/model"
	"tcms-auth/internal/database/session"
)

type sessionRepository struct {
	db *sqlx.DB
}

type SessionRepository interface {
	Create(userId int) (string, error)
	GetUser(token string) (*model.AuthUser, error)
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

func (r sessionRepository) GetUser(token string) (*model.AuthUser, error) {
	rows, err := r.db.Queryx(`
		SELECT u.username, u.telegram_access_key
		FROM user_session
		LEFT JOIN users u on u.id = user_session.user_id
		WHERE user_session.token = $1;
	`, token)

	if err != nil {
		return nil, err
	}

	if !rows.Next() {
		return nil, nil
	}

	var authUser model.AuthUser
	err = rows.StructScan(&authUser)
	if err != nil {
		return nil, err
	}

	return &authUser, nil
}
