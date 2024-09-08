package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"x/pkg/model"
)

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
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

func (u *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest model.CreateUser

	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		log.Printf("error decoding body: %+v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := u.userService.CreateUser(createUserRequest.Name); err != nil {
		log.Printf("error creating user: %+v", err)
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}