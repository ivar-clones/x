package model

import "time"

type User struct {
	ID int `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Email *string `db:"email" json:"email"`
	Bio *string `db:"bio" json:"bio"`
	DOB *string `db:"dob" json:"dob"`
	UpsertedAt time.Time `db:"upserted_at" json:"upsertedAt"`
}

type CreateUser struct {
	Name string `json:"name" validate:"required"`
	Email *string `json:"email" validate:"omitempty,email"`
	Bio *string `json:"bio"`
	DOB *string `json:"dob" validate:"omitempty,datetime=2006-01-02"`
}

type UpdateUser struct {
	ID int `json:"id" validate:"required"`
	Name *string `json:"name" validate:"min=1"`
	Email *string `json:"email" validate:"omitempty,email"`
	Bio *string `json:"bio"`
	DOB *string `json:"dob" validate:"omitempty,datetime=2006-01-02"`
}