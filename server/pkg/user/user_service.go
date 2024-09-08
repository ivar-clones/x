package user

import "x/pkg/repository"

type Service interface {

}

type service struct {
	db repository.Repository
}

func New(db repository.Repository) Service {
	return &service{
		db: db,
	}
}