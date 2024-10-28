package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"x/pkg/model"

	"github.com/go-playground/validator/v10"
)

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
}

func (u *controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.userService.GetAllUsers()
	if err != nil {
		log.Printf("error fetching users: %+v", err)
		http.Error(w, "error fetching users", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		log.Printf("error marshalling users: %+v", err)
		http.Error(w, "error fetching users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *controller) GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.PathValue("email")
	
	user, err := u.userService.GetUserByEmail(email)
	if err != nil {
		log.Printf("error fetching users: %+v", err)
		http.Error(w, "error fetching users", http.StatusInternalServerError)
		return
	}

	if user == nil {
		log.Printf("no user found");
		http.Error(w, "no user found", http.StatusNoContent)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		log.Printf("error marshalling users: %+v", err)
		http.Error(w, "error fetching users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (u *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest model.CreateUser
	validate := validator.New()

	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		log.Printf("error decoding body: %+v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(createUserRequest); err != nil {
		log.Printf("error decoding body: %+v", err)
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	if err := u.userService.CreateUser(createUserRequest.Name, createUserRequest.Email, createUserRequest.Bio, createUserRequest.DOB); err != nil {
		log.Printf("error creating user: %+v", err)
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (u *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var updateUserRequest model.UpdateUser
	validate := validator.New()

	err := json.NewDecoder(r.Body).Decode(&updateUserRequest)
	if err != nil {
		log.Printf("error decoding body: %+v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := validate.Struct(updateUserRequest); err != nil {
		log.Printf("error decoding body: %+v", err)
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	if err := u.userService.UpdateUser(updateUserRequest.ID, updateUserRequest.Name, updateUserRequest.Email, updateUserRequest.Bio, updateUserRequest.DOB); err != nil {
		log.Printf("error creating user: %+v", err)
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}