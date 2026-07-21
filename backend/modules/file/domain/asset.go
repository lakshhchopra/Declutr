package domain

import "time"

type AssetStatus string

const (
	StatusQueued           AssetStatus = "QUEUED"
	StatusUploading        AssetStatus = "UPLOADING"
	StatusUploaded         AssetStatus = "UPLOADED"
	StatusValidating       AssetStatus = "VALIDATING"
	StatusMetadataPending  AssetStatus = "METADATA_PENDING"
	StatusAIPending        AssetStatus = "AI_PENDING"
	StatusIndexedPending   AssetStatus = "INDEXED_PENDING"
	StatusReady            AssetStatus = "READY"
	StatusFailed           AssetStatus = "FAILED"
)

type Asset struct {
	AssetID        string      `json:"assetId"`
	VaultID        string      `json:"vaultId"`
	OwnerID        string      `json:"ownerId"`
	Filename       string      `json:"filename"`
	MimeType       string      `json:"mimeType"`
	SizeBytes      int64       `json:"sizeBytes"`
	ChecksumSHA256 string      `json:"checksumSha256"`
	StorageKey     string      `json:"storageKey"`
	Status         AssetStatus `json:"status"`
	ErrorMessage   string      `json:"errorMessage,omitempty"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
}

type UploadJob struct {
	JobID              string    `json:"jobId"`
	AssetID            string    `json:"assetId"`
	JobType            string    `json:"jobType"`
	Status             string    `json:"status"`
	ProgressPercentage int       `json:"progressPercentage"`
	Attempts           int       `json:"attempts"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}
