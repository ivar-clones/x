package repository

import (
	"context"
	"log"
	"x/pkg/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetAllUsers() ([]model.User, error)
}

type repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	rows, err := r.db.Query(context.Background(), "select * from users")
	if err != nil {
		log.Printf("error fetching users: %+v", err)
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user model.User
		err = rows.Scan(&user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if rows.Err() != nil {
		log.Printf("error scanning rows: %+v", rows.Err())
		return nil, rows.Err()
	}

	return users, nil
}