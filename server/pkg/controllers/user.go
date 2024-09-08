package controllers

import (
	"net/http"
)

type UserController interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
}

func (u *controller) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	
}