package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"x/pkg/model"
)

var badDobError = errors.New("Format for date of birth should be DD-MM-YYYY")

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
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

func (u *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest model.CreateUser

	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		log.Printf("error decoding body: %+v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := validatedDob(createUserRequest.DOB); errors.Is(err, badDobError) {
		log.Printf("bad format for dob: %+v", createUserRequest.DOB)
		http.Error(w, err.Error(), http.StatusBadRequest)
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

	err := json.NewDecoder(r.Body).Decode(&updateUserRequest)
	if err != nil {
		log.Printf("error decoding body: %+v", err)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	if err := validatedDob(updateUserRequest.DOB); errors.Is(err, badDobError) {
		log.Printf("bad format for dob: %+v", updateUserRequest.DOB)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := u.userService.UpdateUser(updateUserRequest.ID, updateUserRequest.Name, updateUserRequest.Email, updateUserRequest.Bio, updateUserRequest.DOB); err != nil {
		log.Printf("error creating user: %+v", err)
		http.Error(w, "error creating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func validatedDob(dob string) error {
	if dob == "" {
		return nil
	}
	
	parsedDob := strings.Split(dob, "-")
	if len(parsedDob) != 3 || (parsedDob[0] == "" || parsedDob[1] == "" || parsedDob[2] == "") {
		return badDobError
	}

	_, err := strconv.Atoi(parsedDob[0])
	if err != nil {
		return badDobError
	}

	_, err = strconv.Atoi(parsedDob[1])
	if err != nil {
		return badDobError
	}

	_, err = strconv.Atoi(parsedDob[2])
	if err != nil {
		return badDobError
	}
	
	return nil
}