package repositories

import (
	"cms/internal/models"
	"errors"
)

var users = []models.User{}

func Register(user models.User) error {
	for _, u := range users {
		if u.Username == user.Username {
			return errors.New("username already exists")
		}
	}
	users = append(users, user)
	return nil
}

func Login(user models.User) error {
	for _, u := range users {
		if u.Username == user.Username && u.Password == user.Password {
			return nil
		}
	}
	return errors.New("invalid username or password")
}
