package models

import "time"

type User struct {
	ID                 string    `json:"id"`
	Email              string    `json:"email"`
	EncryptedMasterKey string    `json:"encryptedMasterKey"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
