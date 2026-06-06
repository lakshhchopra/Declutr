package auth

	import (
	"time"

	"github.com/google/uuid"

	authmodels "github.com/diablovocado/declutr/internal/auth/models"
	"github.com/diablovocado/declutr/internal/crypto"
	"github.com/diablovocado/declutr/internal/models"
	"github.com/diablovocado/declutr/internal/repository"
)


type Service struct {
	UserRepo repository.UserRepository
}

func (s *Service) Register(req authmodels.RegisterRequest) (string, error) {
	id := uuid.New().String()

	user := models.User{
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
