package auth

import (
	"time"

	"github.com/google/uuid"

	"github.com/diablovocado/declutr/internal/models"
	"github.com/diablovocado/declutr/internal/repository"
)

type Service struct {
	UserRepo repository.UserRepository
}

func (s *Service) Register() (string, error) {
	id := uuid.New().String()

	user := models.User{
		ID:        id,
		EmailHash: "temporary-email-hash",

		SRPVerifier: "temporary-verifier",
		SRPSalt:     "temporary-salt",

		EncryptedMVKCiphertext: "temporary-ciphertext",
		EncryptedMVKNonce:      "temporary-nonce",
		EncryptedMVKVersion:    1,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.UserRepo.CreateUser(user); err != nil {
		return "", err
	}

	return id, nil
}
