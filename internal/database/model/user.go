package model

type User struct {
	Id                int
	Username          string `db:"username"`
	Password          string `db:"password"`
	TelegramAccessKey string `db:"telegram_access_key"`
}
