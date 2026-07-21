package application

import (
	"github.com/diablovocado/declutr/modules/auth/domain"
)

type ChallengeStore struct {
	Challenges map[string]domain.Challenge
}

func NewChallengeStore() *ChallengeStore {
	return &ChallengeStore{
		Challenges: make(map[string]domain.Challenge),
	}
}
