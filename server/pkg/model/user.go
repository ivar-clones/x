package model

import "time"

type User struct {
	ID int `db:"id"`
	Name string `db:"name"`
	UpsertedAt time.Time `db:"upserted_at"`
}