package domain

import (
	"time"
)

type EntityType string

const (
	TypePerson       EntityType = "Person"
	TypeOrganization EntityType = "Organization"
	TypeLocation     EntityType = "Location"
	TypeDate         EntityType = "Date"
	TypeAmount       EntityType = "Amount"
	TypeProduct      EntityType = "Product"
	TypeIdentifier   EntityType = "Identifier"
)

type Entity struct {
	EntityID        string     `json:"entityId"`
	VaultID         string     `json:"vaultId"`
	Type            EntityType `json:"entityType"`
	CanonicalName   string     `json:"canonicalName"`
	NormalizedValue string     `json:"normalizedValue"`
	Description     string     `json:"description"`
	Aliases         []string   `json:"aliases"`
	CreatedAt       time.Time  `json:"createdAt"`
	UpdatedAt       time.Time  `json:"updatedAt"`
}

type EntityOccurrence struct {
	OccurrenceID     string    `json:"occurrenceId"`
	EntityID         string    `json:"entityId"`
	AssetID          string    `json:"assetId"`
	AnalysisID       string    `json:"analysisId"`
	OriginalValue    string    `json:"originalValue"`
	ConfidenceScore  float64   `json:"confidenceScore"`
	ExtractorVersion string    `json:"extractorVersion"`
	CreatedAt        time.Time `json:"createdAt"`
}

// ExtractedEntity represents a raw extraction before resolution
type ExtractedEntity struct {
	Type            EntityType
	OriginalValue   string
	NormalizedValue string
	ConfidenceScore float64
}
