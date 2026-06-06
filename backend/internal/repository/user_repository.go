package repository

import "github.com/diablovocado/declutr/internal/models"

type UserRepository interface {
	CreateUser(user models.User) error
}
