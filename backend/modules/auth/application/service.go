package application

import (
	"time"

	"github.com/google/uuid"

	"github.com/diablovocado/declutr/modules/auth/domain"
	"github.com/diablovocado/declutr/modules/auth/repository"
	"github.com/diablovocado/declutr/modules/auth/transport/models"
	"github.com/diablovocado/declutr/shared/crypto"
)

type Service struct {
	UserRepo   repository.UserRepository
	Challenges *ChallengeStore
	SRP        *Engine
}

func (s *Service) Register(req models.RegisterRequest) (string, error) {
	id := uuid.New().String()

	user := domain.User{
		ID: id,

		EmailHash: crypto.HashEmail(req.Email),

		SRPVerifier: req.SRPVerifier,
		SRPSalt:     req.SRPSalt,

		EncryptedMVKCiphertext: req.MVK.Ciphertext,
		EncryptedMVKNonce:      req.MVK.Nonce,
		EncryptedMVKVersion:    req.MVK.Version,

		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.UserRepo.CreateUser(user); err != nil {
		return "", err
	}

	return id, nil
}
