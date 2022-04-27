package model

type User struct {
	Id                int    `db:"id"`
	Username          string `db:"username"`
	Password          string `db:"password"`
	TelegramAccessKey string `db:"telegram_access_key"`
}

type AuthUser struct {
	Username          string `db:"username"`
	TelegramAccessKey string `db:"telegram_access_key"`
}
