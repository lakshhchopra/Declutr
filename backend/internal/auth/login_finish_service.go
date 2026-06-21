package auth

import (
	"fmt"
	"time"

	authmodels "github.com/diablovocado/declutr/internal/auth/models"
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

	_ = challenge

	// Single-use challenge
	delete(s.Challenges.Challenges, req.ChallengeID)

	return &authmodels.LoginFinishResponse{
		ServerProof: "temporary-server-proof",
		AccessToken: "temporary-access-token",
	}, nil
}
