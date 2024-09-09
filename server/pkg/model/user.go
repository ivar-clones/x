package model

import "time"

type User struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Bio string `db:"bio" json:"bio"`
	DOB *time.Time `db:"dob" json:"dob"`
	UpsertedAt time.Time `db:"upserted_at" json:"upsertedAt"`
}

type CreateUser struct {
	Name string `json:"name"`
	Bio string `json:"bio"`
	DOB string `json:"dob"`
}

type UpdateUser struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Bio interface{} `json:"bio"`
	DOB string `json:"dob"`
}