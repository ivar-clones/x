package user

import (
	"log"
	"strconv"
	"strings"
	"time"
	"x/pkg/model"
	"x/pkg/repository"
)

type Service interface {
	GetAllUsers() ([]model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(name, email, bio, dob string) error
	UpdateUser(id int, name, email string, bio interface{}, dob string) error
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

func (s *service) CreateUser(name, email, bio, dob string) error {
	var validatedDob interface{}
	if dob == "" {
		validatedDob = nil
	} else {
		parsedDob := strings.Split(dob, "-")
		day, _ := strconv.Atoi(parsedDob[0])
		month, _ := strconv.Atoi(parsedDob[1])
		year, _ := strconv.Atoi(parsedDob[2])
		validatedDob = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	if err := s.db.CreateUser(name, email, bio, validatedDob); err != nil {
		log.Printf("error creating user: %+v", err)
		return err
	}

	return nil
}

func (s *service) UpdateUser(id int, name, email string, bio interface{}, dob string) error {
	var validatedDob interface{}
	if dob == "" {
		validatedDob = nil
	} else {
		parsedDob := strings.Split(dob, "-")
		day, _ := strconv.Atoi(parsedDob[0])
		month, _ := strconv.Atoi(parsedDob[1])
		year, _ := strconv.Atoi(parsedDob[2])
		validatedDob = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	}

	currentUser, err := s.db.GetUser(id)
	if err != nil {
		log.Printf("error fetching user: %+v", err)
		return err
	}

	if name == "" {
		name = currentUser.Name
	}

	if email == "" {
		email = currentUser.Email
	}

	if bio == nil {
		bio = currentUser.Bio
	}

	if validatedDob == nil {
		validatedDob = currentUser.DOB
	}

	if err := s.db.UpdateUser(id, name, email, bio.(string), validatedDob); err != nil {
		log.Printf("error creating user: %+v", err)
		return err
	}

	return nil
}