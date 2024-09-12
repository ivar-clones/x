package repository

import (
	"context"
	"log"
	"x/pkg/model"

	"github.com/jackc/pgx/v5"
)

type UserRepository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(name, email, bio string, dob interface{}) error
	UpdateUser(id int, name, email, bio string, dob interface{}) error
	GetUser(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

func (r *repository) GetAllUsers() ([]model.User, error) {
	rows, err := r.db.Query(context.Background(), "select id, name, email, upserted_at, bio, dob from users")
	if err != nil {
		log.Printf("error querying users: %+v", err)
		return nil, err
	}

	defer rows.Close()

	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		log.Printf("error collecting rows: %+v", err)
		return nil, err
	}

	return users, nil
}

func (r *repository) CreateUser(name, email, bio string, dob interface{}) error {
	_, err := r.db.Exec(context.Background(), "insert into users (name, email, bio, dob) values ($1, $2, $3, $4)", name, email, bio, dob)
	if err != nil {
		log.Printf("error inserting user: %+v", err)
		return err
	}

	return nil
}

func (r *repository) GetUser(id int) (*model.User, error) {
	rows, err := r.db.Query(context.Background(), "select id, name, email, upserted_at, bio, dob from users where id = $1", id)
	if err != nil {
		log.Printf("error querying user: %+v", err)
		return nil, err
	}
	
	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		log.Printf("error collecting rows: %+v", err)
		return nil, err
	}

	return &user, nil
}

func (r *repository) GetUserByEmail(email string) (*model.User, error) {
	rows, err := r.db.Query(context.Background(), "select id, name, email, upserted_at, bio, dob from users where email = $1", email)
	if err != nil {
		log.Printf("error querying user: %+v", err)
		return nil, err
	}

	user, err := pgx.CollectExactlyOneRow(rows, pgx.RowToStructByName[model.User])
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		log.Printf("error collecting rows: %+v", err)
		return nil, err
	}

	return &user, nil
}

func (r *repository) UpdateUser(id int, name, email, bio string, dob interface{}) error {
	_, err := r.db.Exec(context.Background(), "update users set name = $1, email = $2, bio = $3, dob = $4 where id = $4", name, email, bio, dob, id)
	if err != nil {
		log.Printf("error inserting user: %+v", err)
		return err
	}

	return nil
}