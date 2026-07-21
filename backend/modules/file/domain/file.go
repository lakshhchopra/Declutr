package domain

import "time"

type File struct {
	ID                string    `json:"id"`
	VaultID           string    `json:"vaultId"`
	EncryptedFileKey  string    `json:"encryptedFileKey"`
	EncryptedMetadata string    `json:"encryptedMetadata"`
	StoragePath       string    `json:"storagePath"`
	FileSize          int64     `json:"fileSize"`
	MimeType          string    `json:"mimeType"`
	CreatedAt         time.Time `json:"createdAt"`
}
