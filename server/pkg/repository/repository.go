package repository

import (
	"context"
	"log"
	"x/pkg/model"

	"github.com/jackc/pgx/v5"
)

type dbConn interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
}

type Repository interface {
	GetAllUsers() ([]model.User, error)
}

type repository struct {
	db dbConn
}

func New(db dbConn) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllUsers() ([]model.User, error) {
	rows, err := r.db.Query(context.Background(), "select id, name, upserted_at from users")
	if err != nil {
		log.Printf("error fetching users: %+v", err)
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