package models

import "github.com/diablovocado/declutr/shared/crypto"

type RegisterRequest struct {
	Email       string              `json:"email"`
	SRPVerifier string              `json:"srpVerifier"`
	SRPSalt     string              `json:"srpSalt"`
	MVK         crypto.EncryptedMVK `json:"mvk"`
}

type RegisterResponse struct {
	UserID string `json:"userId"`
}
