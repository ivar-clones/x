package user

import (
	"log"
	"x/pkg/model"
	"x/pkg/repository"
)

type Service interface {
	GetAllUsers() ([]model.User, error)
}

type service struct {
	db repository.Repository
}

func New(db repository.Repository) Service {
	return &service{
		db: db,
	}
}

func (s *service) GetAllUsers() ([]model.User, error) {
	users, err := s.db.GetAllUsers()
	if err != nil {
		log.Printf("error fetching users: %+v", err)
		return nil, err
	}

	return users, nil
}