package model

import "time"

type User struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	UpsertedAt time.Time `db:"upserted_at" json:"upsertedAt"`
}

type CreateUser struct {
	Name string `json:"name"`
}