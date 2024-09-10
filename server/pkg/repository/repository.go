package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type dbConn interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

type Repository interface {
	UserRepository
}

type repository struct {
	db dbConn
}

func New(db dbConn) Repository {
	return &repository{
		db: db,
	}
}

