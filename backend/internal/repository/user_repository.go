package repository

import "github.com/diablovocado/declutr/internal/models"

type UserRepository interface {
	CreateUser(user models.User) error
	GetUserByEmailHash(emailHash string) (*models.User, error)

	CreateSession(session models.Session) error
	GetSessionByToken(token string) (*models.Session, error)
	DeleteSession(token string) error
}
