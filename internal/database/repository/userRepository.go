package repository

import (
	"github.com/jmoiron/sqlx"
	"tcms-auth/internal/database/model"
)

type UserRepository interface {
	Create(user *model.User) error
}

type userRepository struct {
	db *sqlx.DB
}

func (u *userRepository) Create(user *model.User) error {
	_, err := u.db.NamedExec("INSERT INTO users (username, password, telegram_access_key) VALUES (:username, :password, :telegram_access_key)", user)
	return err
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
