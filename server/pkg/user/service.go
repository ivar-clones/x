package user

import (
	"log"
	"x/pkg/model"
	"x/pkg/repository"
)

type Service interface {
	GetAllUsers() ([]model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(name string, email, bio, dob *string) error
	UpdateUser(id int, name, email, bio, dob *string) error
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

func (s *service) GetUserByEmail(email string) (*model.User, error) {
	user, err := s.db.GetUserByEmail(email)
	if err != nil {
		log.Printf("error fetching user: %+v", err)
		return nil, err
	}

	return user, nil
}

func (s *service) CreateUser(name string, email, bio, dob *string) error {
	if err := s.db.CreateUser(name, email, bio, dob); err != nil {
		log.Printf("error creating user: %+v", err)
		return err
	}

	return nil
}

func (s *service) UpdateUser(id int, name, email, bio, dob *string) error {
	currentUser, err := s.db.GetUser(id)
	if err != nil {
		log.Printf("error fetching user: %+v", err)
		return err
	}

	if name == nil {
		name = &currentUser.Name
	}

	if email == nil {
		email = currentUser.Email
	}

	if bio == nil {
		bio = currentUser.Bio
	} else if (*bio == "") {
		bio = nil
	}


	if dob == nil {
		dob = currentUser.DOB
	}

	if err := s.db.UpdateUser(id, name, email, bio, dob); err != nil {
		log.Printf("error creating user: %+v", err)
		return err
	}

	return nil
}