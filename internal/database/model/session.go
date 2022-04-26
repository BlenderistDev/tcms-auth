package model

import (
	"time"
)

type Session struct {
	Id        int       `db:"id"`
	UserId    int       `db:"user_id"`
	Token     int       `db:"token"`
	Timestamp time.Time `db:"timestamp"`
}
