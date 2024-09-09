package repository

import (
	"context"
	"log"
	"x/pkg/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type dbConn interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type Repository interface {
	GetAllUsers() ([]model.User, error)
	CreateUser(name, bio string, dob interface{}) error
	UpdateUser(id int, name, bio string, dob interface{}) error
	GetUser(id int) (*model.User, error)
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
	rows, err := r.db.Query(context.Background(), "select id, name, upserted_at, bio, dob from users")
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

func (r *repository) CreateUser(name, bio string, dob interface{}) error {
	_, err := r.db.Exec(context.Background(), "insert into users (name, bio, dob) values ($1, $2, $3)", name, bio, dob)
	if err != nil {
		log.Printf("error inserting user: %+v", err)
		return err
	}

	return nil
}

func (r *repository) GetUser(id int) (*model.User, error) {
	rows, err := r.db.Query(context.Background(), "select id, name, upserted_at, bio, dob from users where id = $1", id)
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

func (r *repository) UpdateUser(id int, name, bio string, dob interface{}) error {
	_, err := r.db.Exec(context.Background(), "update users set name = $1, bio = $2, dob = $3 where id = $4", name, bio, dob, id)
	if err != nil {
		log.Printf("error inserting user: %+v", err)
		return err
	}

	return nil
}