package domain

import "time"

type User struct {
	ID string

	EmailHash string

	SRPVerifier string
	SRPSalt     string

	EncryptedMVKCiphertext string
	EncryptedMVKNonce      string
	EncryptedMVKVersion    int

	CreatedAt time.Time
	UpdatedAt time.Time
}
