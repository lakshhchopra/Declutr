package domain

import "time"

type Vault struct {
	ID                string    `json:"id"`
	OwnerID           string    `json:"ownerId"`
	EncryptedVaultKey string    `json:"encryptedVaultKey"`
	Name              string    `json:"name"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}
