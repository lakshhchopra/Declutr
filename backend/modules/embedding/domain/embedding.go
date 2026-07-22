package domain

import "time"

// SourceType categorises what knowledge item is embedded
type SourceType string

const (
	SourceDocument     SourceType = "DOCUMENT"
	SourceSummary      SourceType = "SUMMARY"
	SourceEntity       SourceType = "ENTITY"
	SourceContext      SourceType = "CONTEXT"
	SourceRelationship SourceType = "RELATIONSHIP"
	SourceMemory       SourceType = "MEMORY"
	SourceCollection   SourceType = "COLLECTION"
	SourceNote         SourceType = "NOTE"
	SourceChat         SourceType = "CHAT"
)

// ChunkStrategy defines intelligent chunking algorithms
type ChunkStrategy string

const (
	StrategySemantic     ChunkStrategy = "SEMANTIC"
	StrategyHierarchical ChunkStrategy = "HIERARCHICAL"
	StrategyDocument     ChunkStrategy = "DOCUMENT_AWARE"
	StrategyPage         ChunkStrategy = "PAGE_AWARE"
	StrategyHeading      ChunkStrategy = "HEADING_AWARE"
)

// ProviderName represents supported embedding model providers
type ProviderName string

const (
	ProviderOpenAI ProviderName = "OPENAI"
	ProviderGemini ProviderName = "GEMINI"
	ProviderVoyage ProviderName = "VOYAGE"
	ProviderCohere ProviderName = "COHERE"
	ProviderOllama ProviderName = "OLLAMA"
	ProviderLocal  ProviderName = "LOCAL"
)

// VectorStoreType represents vector storage repository backends
type VectorStoreType string

const (
	VectorStorePGVector VectorStoreType = "PGVECTOR"
	VectorStoreQdrant   VectorStoreType = "QDRANT"
	VectorStoreWeaviate VectorStoreType = "WEAVIATE"
	VectorStorePinecone VectorStoreType = "PINECONE"
	VectorStoreMilvus   VectorStoreType = "MILVUS"
	VectorStoreInMemory VectorStoreType = "INMEMORY"
)

// Vector is a slice of float64 embeddings
type Vector []float64

// StructuredRepresentationInput holds enriched knowledge attributes for rich embedding generation
type StructuredRepresentationInput struct {
	SourceType    SourceType        `json:"sourceType"`
	SourceID      string            `json:"sourceId"`
	VaultID       string            `json:"vaultId"`
	Title         string            `json:"title"`
	Summary       string            `json:"summary"`
	Content       string            `json:"content,omitempty"`
	Entities      []string          `json:"entities,omitempty"`
	Relationships []string          `json:"relationships,omitempty"`
	Contexts      []string          `json:"contexts,omitempty"`
	Intent        string            `json:"intent,omitempty"`
	MemoryScore   float64           `json:"memoryScore,omitempty"`
	Tags          []string          `json:"tags,omitempty"`
	Classification string           `json:"classification,omitempty"`
	Metadata      map[string]string `json:"metadata,omitempty"`
}

// Embedding represents the stored vector embedding for a knowledge item
type Embedding struct {
	EmbeddingID        string     `json:"embeddingId"`
	VaultID            string     `json:"vaultId"`
	SourceType         SourceType `json:"sourceType"`
	SourceID           string     `json:"sourceId"`
	ProviderName       string     `json:"providerName"`
	ModelName          string     `json:"modelName"`
	ModelVersion       string     `json:"modelVersion"`
	Dimensions         int        `json:"dimensions"`
	RepresentationText string     `json:"representationText"`
	ContentHash        string     `json:"contentHash"`
	VectorData         Vector     `json:"vectorData"`
	IsActive           bool       `json:"isActive"`
	CreatedAt          time.Time  `json:"createdAt"`
	UpdatedAt          time.Time  `json:"updatedAt"`
}

// EmbeddingChunk represents a chunked segment of a document with its vector
type EmbeddingChunk struct {
	ChunkID       string        `json:"chunkId"`
	EmbeddingID   string        `json:"embeddingId"`
	VaultID       string        `json:"vaultId"`
	ChunkIndex    int           `json:"chunkIndex"`
	ChunkStrategy ChunkStrategy `json:"chunkStrategy"`
	ChunkText     string        `json:"chunkText"`
	TokenCount    int           `json:"tokenCount"`
	HeadingPath   string        `json:"headingPath,omitempty"`
	PageNumber    int           `json:"pageNumber,omitempty"`
	VectorData    Vector        `json:"vectorData"`
	CreatedAt     time.Time     `json:"createdAt"`
}

// EmbeddingVersion tracks provider and model upgrades
type EmbeddingVersion struct {
	VersionID         string    `json:"versionId"`
	VaultID           string    `json:"vaultId"`
	ProviderName      string    `json:"providerName"`
	ModelName         string    `json:"modelName"`
	Dimensions        int       `json:"dimensions"`
	VersionTag        string    `json:"versionTag"`
	IsActive          bool      `json:"isActive"`
	TotalEmbeddedItems int      `json:"totalEmbeddedItems"`
	UpgradedAt        time.Time `json:"upgradedAt"`
}

// EmbeddingJob tracks background batch processing
type EmbeddingJob struct {
	JobID           string     `json:"jobId"`
	VaultID         string     `json:"vaultId"`
	TargetType      SourceType `json:"targetType"`
	TargetID        string     `json:"targetId"`
	Status          string     `json:"status"` // QUEUED, PROCESSING, COMPLETED, FAILED
	ErrorMessage    string     `json:"errorMessage,omitempty"`
	ProcessedChunks int        `json:"processedChunks"`
	CreatedAt       time.Time  `json:"createdAt"`
	CompletedAt     *time.Time `json:"completedAt,omitempty"`
}

// EmbeddingProviderConfig configures provider per vault
type EmbeddingProviderConfig struct {
	ProviderID   string       `json:"providerId"`
	VaultID      string       `json:"vaultId"`
	ProviderName ProviderName `json:"providerName"`
	ModelName    string       `json:"modelName"`
	Dimensions   int          `json:"dimensions"`
	BatchSize    int          `json:"batchSize"`
	IsDefault    bool         `json:"isDefault"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
}

// VectorMetadata holds extra key-value pairs associated with embeddings
type VectorMetadata struct {
	MetadataID  string    `json:"metadataId"`
	EmbeddingID string    `json:"embeddingId"`
	MetaKey     string    `json:"metaKey"`
	MetaValue   string    `json:"metaValue"`
	CreatedAt   time.Time `json:"createdAt"`
}

// ChunkResult contains the output of chunking a document
type ChunkResult struct {
	Text        string
	Index       int
	Strategy    ChunkStrategy
	TokenCount  int
	HeadingPath string
	PageNumber  int
}

// GenerationOptions configures an embedding generation request
type GenerationOptions struct {
	Provider      ProviderName
	ModelName     string
	ChunkStrategy ChunkStrategy
	ForceRefresh  bool
}

// EmbeddingStats holds aggregate statistics for a vault
type EmbeddingStats struct {
	VaultID            string            `json:"vaultId"`
	TotalEmbeddings    int               `json:"totalEmbeddings"`
	TotalChunks        int               `json:"totalChunks"`
	ActiveProvider     string            `json:"activeProvider"`
	ActiveModel        string            `json:"activeModel"`
	Dimensions         int               `json:"dimensions"`
	SourceTypeBreakdown map[string]int    `json:"sourceTypeBreakdown"`
	StrategyBreakdown   map[string]int    `json:"strategyBreakdown"`
	LatestVersionTag   string            `json:"latestVersionTag"`
}
