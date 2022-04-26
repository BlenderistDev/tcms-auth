package model

import "github.com/jackc/pgx/pgtype"

type Session struct {
	Id        int              `db:"id"`
	UserId    int              `db:"user_id"`
	Token     int              `db:"token"`
	Timestamp pgtype.Timestamp `db:"timestamp"`
}
