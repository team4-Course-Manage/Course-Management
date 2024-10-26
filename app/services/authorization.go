package services

import (
	"cms/internal/models"
	"cms/internal/repositories"
)

func Register(user models.User) error {
	return repositories.Register(user)
}

func Login(user models.User) error {
	return repositories.Login(user)
}
