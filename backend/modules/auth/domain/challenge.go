package domain

import "time"

type ChallengeID string

type Challenge struct {
	ID ChallengeID

	UserID string

	ServerSecret    string
	ServerPublicKey string

	CreatedAt time.Time
}
