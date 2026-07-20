package auth

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	authmodels "github.com/diablovocado/declutr/internal/auth/models"
	"github.com/diablovocado/declutr/internal/auth/session"
	"github.com/diablovocado/declutr/internal/models"
)

func (s *Service) LoginFinish(
	req authmodels.LoginFinishRequest,
) (*authmodels.LoginFinishResponse, error) {

	challenge, ok := s.Challenges.Challenges[req.ChallengeID]
	if !ok {
		return nil, fmt.Errorf("challenge not found")
	}

	if time.Since(challenge.CreatedAt) > 5*time.Minute {
		delete(s.Challenges.Challenges, req.ChallengeID)
		return nil, fmt.Errorf("challenge expired")
	}

	if !s.SRP.VerifyClientProof(
		req.ClientProof,
		challenge.ServerSecret,
	) {
		return nil, fmt.Errorf("invalid proof")
	}

	// Single-use challenge
	delete(s.Challenges.Challenges, req.ChallengeID)

	// Generate a real access token
	accessToken := session.GenerateAccessToken()

	// Create a session
	newSession := models.Session{
		ID:          uuid.New().String(),
		UserID:      challenge.UserID,
		AccessToken: accessToken,
		CreatedAt:   time.Now(),
		ExpiresAt:   time.Now().Add(24 * time.Hour),
	}

	// Store the session
	if err := s.UserRepo.CreateSession(newSession); err != nil {
		return nil, err
	}

	return &authmodels.LoginFinishResponse{
		ServerProof: "temporary-server-proof",
		AccessToken: accessToken,
	}, nil
}
