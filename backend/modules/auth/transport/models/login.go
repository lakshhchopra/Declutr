package models

type LoginStartRequest struct {
	Email string `json:"email"`
}

type LoginStartResponse struct {
	ChallengeID     string `json:"challengeId"`
	Salt            string `json:"salt"`
	ServerPublicKey string `json:"serverPublicKey"`
}

type LoginFinishRequest struct {
	ChallengeID string `json:"challengeId"`

	Email           string `json:"email"`
	ClientPublicKey string `json:"clientPublicKey"`
	ClientProof     string `json:"clientProof"`
}

type LoginFinishResponse struct {
	ServerProof string `json:"serverProof"`
	AccessToken string `json:"accessToken"`
}
