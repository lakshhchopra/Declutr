package domain

import (
	"time"
)

type AssetMetadata struct {
	AssetID         string     `json:"assetId"`
	VaultID         string     `json:"vaultId"`
	Filename        string     `json:"filename"`
	Extension       string     `json:"extension,omitempty"`
	MimeType        string     `json:"mimeType,omitempty"`
	FileSize        int64      `json:"fileSize"`
	Checksum        string     `json:"checksum,omitempty"`
	Hash            string     `json:"hash,omitempty"`
	Encoding        string     `json:"encoding,omitempty"`
	CreatedDate     *time.Time `json:"createdDate,omitempty"`
	ModifiedDate    *time.Time `json:"modifiedDate,omitempty"`
	UploadDate      time.Time  `json:"uploadDate"`
	LastExtractedAt time.Time  `json:"lastExtractedAt"`
}

type AssetProperties struct {
	AssetID    string                 `json:"assetId"`
	Properties map[string]interface{} `json:"properties"`
}

type AssetExif struct {
	AssetID     string     `json:"assetId"`
	CameraMake  string     `json:"cameraMake,omitempty"`
	CameraModel string     `json:"cameraModel,omitempty"`
	Lens        string     `json:"lens,omitempty"`
	GPSLat      *float64   `json:"gpsLat,omitempty"`
	GPSLong     *float64   `json:"gpsLong,omitempty"`
	ISO         *int       `json:"iso,omitempty"`
	Exposure    string     `json:"exposure,omitempty"`
	FStop       *float64   `json:"fStop,omitempty"`
	FocalLength *float64   `json:"focalLength,omitempty"`
	DateTaken   *time.Time `json:"dateTaken,omitempty"`
	RawData     map[string]interface{} `json:"rawData,omitempty"`
}

type MetadataVersion struct {
	VersionID        string                 `json:"versionId"`
	AssetID          string                 `json:"assetId"`
	Source           string                 `json:"source"`
	ExtractorVersion string                 `json:"extractorVersion"`
	Confidence       *float64               `json:"confidence,omitempty"`
	Snapshot         map[string]interface{} `json:"snapshot"`
	CreatedAt        time.Time              `json:"createdAt"`
}

type CompleteMetadata struct {
	General    *AssetMetadata   `json:"general"`
	Properties *AssetProperties `json:"properties,omitempty"`
	Exif       *AssetExif       `json:"exif,omitempty"`
}
