package domain

import (
	"time"
)

type TokenUsage struct {
	PromptTokens     int     `json:"promptTokens"`
	CompletionTokens int     `json:"completionTokens"`
	TotalTokens      int     `json:"totalTokens"`
	EstimatedCostUSD float64 `json:"estimatedCostUsd"`
}

type AIAnalysis struct {
	AnalysisID           string         `json:"analysisId"`
	DocumentID           string         `json:"documentId"`
	AssetID              string         `json:"assetId"`
	Title                string         `json:"title"`
	ShortSummary         string         `json:"shortSummary"`
	DetailedSummary      string         `json:"detailedSummary"`
	Language             string         `json:"language"`
	WritingStyle         string         `json:"writingStyle"`
	Sentiment            string         `json:"sentiment"`
	Complexity           string         `json:"complexity"`
	ReadingLevel         string         `json:"readingLevel"`
	EstimatedReadingTime int            `json:"estimatedReadingTime"`
	DocumentPurpose      string         `json:"documentPurpose"`
	ConfidenceScore      float64        `json:"confidenceScore"`
	Classification       Classification `json:"classification"`
	Tags                 []Tag          `json:"tags"`
	Topics               []Topic        `json:"topics"`
	CreatedAt            time.Time      `json:"createdAt"`
	UpdatedAt            time.Time      `json:"updatedAt"`
}

type Classification struct {
	DocumentCategory string  `json:"documentCategory"`
	DocumentType     string  `json:"documentType"`
	IsScanned        bool    `json:"isScanned"`
	IsCorrupted      bool    `json:"isCorrupted"`
	IsIncomplete     bool    `json:"isIncomplete"`
	QualityScore     float64 `json:"qualityScore"`
	ConfidenceScore  float64 `json:"confidenceScore"`
}

type Tag struct {
	Name            string  `json:"name"`
	ConfidenceScore float64 `json:"confidenceScore"`
}

type Topic struct {
	Name            string  `json:"name"`
	ConfidenceScore float64 `json:"confidenceScore"`
}

type AnalysisVersion struct {
	VersionID     string                 `json:"versionId"`
	AnalysisID    string                 `json:"analysisId"`
	Provider      string                 `json:"provider"`
	ModelName     string                 `json:"modelName"`
	PromptVersion string                 `json:"promptVersion"`
	TokenUsage    TokenUsage             `json:"tokenUsage"`
	LatencyMs     int                    `json:"latencyMs"`
	RawOutput     map[string]interface{} `json:"rawOutput"`
	CreatedAt     time.Time              `json:"createdAt"`
}
