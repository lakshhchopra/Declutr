package models

type RegisterRequest struct {
	Email string `json:"email"`
	EncryptedMasterKey string `json:"encryptedMasterKey"`
}

type RegisterResponse struct {
	UserID string `json:"userId"`
}
