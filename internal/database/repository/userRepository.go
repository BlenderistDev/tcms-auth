package repository

import (
	"github.com/jmoiron/sqlx"
	"tcms-auth/internal/database/model"
)

type UserRepository interface {
	Create(user *model.User) error
	GetUser(username string) (*model.User, error)
	UpdateTelegramAccessKey(id int, key string) error
}

type userRepository struct {
	db *sqlx.DB
}

// Create creates user
func (u *userRepository) Create(user *model.User) error {
	_, err := u.db.NamedExec("INSERT INTO users (username, password, telegram_access_key) VALUES (:username, :password, :telegram_access_key)", user)
	return err
}

func (u *userRepository) GetUser(username string) (*model.User, error) {
	rows, err := u.db.Queryx("SELECT * FROM users WHERE username = $1;", username)
	if err != nil {
		return nil, err
	}
	var user model.User
	if !rows.Next() {
		return nil, nil
	}
	err = rows.StructScan(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateTelegramAccessKey updates telegram access key for specified user id
func (u *userRepository) UpdateTelegramAccessKey(id int, key string) error {
	_, err := u.db.Exec("UPDATE users SET telegram_access_key = $1 WHERE id = $2", key, id)
	return err
}

// NewUserRepository creates new user repository
func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}
