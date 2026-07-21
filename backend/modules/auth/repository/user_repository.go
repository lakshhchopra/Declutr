package repository

import (
	"github.com/diablovocado/declutr/modules/auth/domain"
)

type UserRepository interface {
	CreateUser(user domain.User) error
	GetUserByEmailHash(emailHash string) (*domain.User, error)

	CreateSession(session domain.Session) error
	GetSessionByToken(token string) (*domain.Session, error)
	DeleteSession(token string) error
}
