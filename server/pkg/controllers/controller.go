package controllers

import (
	"x/pkg/user"
)

type Controller interface {
	UserController
}

type controller struct {
	userService user.Service
}

func New(userService user.Service) Controller {
	return &controller{
		userService: userService,
	}
}