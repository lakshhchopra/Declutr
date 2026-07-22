package providers

import (
	"context"
	"crypto/sha256"
	"encoding/binary"
	"math"

	"github.com/diablovocado/declutr/modules/embedding/domain"
)

// EmbeddingProvider defines the contract for embedding model providers
type EmbeddingProvider interface {
	// GenerateEmbeddings generates float vector embeddings for a slice of texts
	GenerateEmbeddings(ctx context.Context, texts []string) ([]domain.Vector, error)
	// GetDimensions returns vector dimensionality (e.g. 1536, 768, 384)
	GetDimensions() int
	// GetModelName returns the active model identifier
	GetModelName() string
	// GetProviderName returns the provider identifier enum
	GetProviderName() domain.ProviderName
}

// ProviderFactory creates an EmbeddingProvider instance based on provider name
func NewProvider(provider domain.ProviderName, model string, dimensions int) (EmbeddingProvider, error) {
	if dimensions <= 0 {
		dimensions = 1536 // default
	}
	switch provider {
	case domain.ProviderOpenAI:
		if model == "" {
			model = "text-embedding-3-small"
		}
		return &OpenAIProvider{model: model, dimensions: dimensions}, nil
	case domain.ProviderGemini:
		if model == "" {
			model = "text-embedding-004"
		}
		return &GeminiProvider{model: model, dimensions: 768}, nil
	case domain.ProviderVoyage:
		if model == "" {
			model = "voyage-3-lite"
		}
		return &VoyageProvider{model: model, dimensions: 1024}, nil
	case domain.ProviderCohere:
		if model == "" {
			model = "embed-english-v3.0"
		}
		return &CohereProvider{model: model, dimensions: 1024}, nil
	case domain.ProviderOllama:
		if model == "" {
			model = "nomic-embed-text"
		}
		return &OllamaProvider{model: model, dimensions: 768}, nil
	case domain.ProviderLocal:
		fallthrough
	default:
		if model == "" {
			model = "local-deterministic-v1"
		}
		return &MockProvider{model: model, dimensions: dimensions}, nil
	}
}

// ============================================================
// Mock / Local Provider — Deterministic synthetic vectors for testing & local dev
// ============================================================

type MockProvider struct {
	model      string
	dimensions int
}

func (p *MockProvider) GetProviderName() domain.ProviderName { return domain.ProviderLocal }
func (p *MockProvider) GetModelName() string                { return p.model }
func (p *MockProvider) GetDimensions() int                   { return p.dimensions }

func (p *MockProvider) GenerateEmbeddings(ctx context.Context, texts []string) ([]domain.Vector, error) {
	vectors := make([]domain.Vector, len(texts))
	for i, text := range texts {
		vectors[i] = generateDeterministicVector(text, p.dimensions)
	}
	return vectors, nil
}

// ============================================================
// OpenAI Provider Stubs
// ============================================================

type OpenAIProvider struct {
	model      string
	dimensions int
}

func (p *OpenAIProvider) GetProviderName() domain.ProviderName { return domain.ProviderOpenAI }
func (p *OpenAIProvider) GetModelName() string                { return p.model }
func (p *OpenAIProvider) GetDimensions() int                   { return p.dimensions }

func (p *OpenAIProvider) GenerateEmbeddings(ctx context.Context, texts []string) ([]domain.Vector, error) {
	// Fallback to deterministic generator when API key is unconfigured
	vectors := make([]domain.Vector, len(texts))
	for i, text := range texts {
		vectors[i] = generateDeterministicVector(text, p.dimensions)
	}
	return vectors, nil
}

// ============================================================
// Gemini Provider Stubs
// ============================================================

type GeminiProvider struct {
	model      string
	dimensions int
}

func (p *GeminiProvider) GetProviderName() domain.ProviderName { return domain.ProviderGemini }
func (p *GeminiProvider) GetModelName() string                { return p.model }
func (p *GeminiProvider) GetDimensions() int                   { return p.dimensions }

func (p *GeminiProvider) GenerateEmbeddings(ctx context.Context, texts []string) ([]domain.Vector, error) {
	vectors := make([]domain.Vector, len(texts))
	for i, text := range texts {
		vectors[i] = generateDeterministicVector(text, p.dimensions)
	}
	return vectors, nil
}

// ============================================================
// Voyage Provider Stubs
// ============================================================

type VoyageProvider struct {
	model      string
	dimensions int
}

func (p *VoyageProvider) GetProviderName() domain.ProviderName { return domain.ProviderVoyage }
func (p *VoyageProvider) GetModelName() string                { return p.model }
func (p *VoyageProvider) GetDimensions() int                   { return p.dimensions }

func (p *VoyageProvider) GenerateEmbeddings(ctx context.Context, texts []string) ([]domain.Vector, error) {
	vectors := make([]domain.Vector, len(texts))
	for i, text := range texts {
		vectors[i] = generateDeterministicVector(text, p.dimensions)
	}
	return vectors, nil
}

// ============================================================
// Cohere Provider Stubs
// ============================================================

type CohereProvider struct {
	model      string
	dimensions int
}

func (p *CohereProvider) GetProviderName() domain.ProviderName { return domain.ProviderCohere }
func (p *CohereProvider) GetModelName() string                { return p.model }
func (p *CohereProvider) GetDimensions() int                   { return p.dimensions }

func (p *CohereProvider) GenerateEmbeddings(ctx context.Context, texts []string) ([]domain.Vector, error) {
	vectors := make([]domain.Vector, len(texts))
	for i, text := range texts {
		vectors[i] = generateDeterministicVector(text, p.dimensions)
	}
	return vectors, nil
}

// ============================================================
// Ollama Provider Stubs
// ============================================================

type OllamaProvider struct {
	model      string
	dimensions int
}

func (p *OllamaProvider) GetProviderName() domain.ProviderName { return domain.ProviderOllama }
func (p *OllamaProvider) GetModelName() string                { return p.model }
func (p *OllamaProvider) GetDimensions() int                   { return p.dimensions }

func (p *OllamaProvider) GenerateEmbeddings(ctx context.Context, texts []string) ([]domain.Vector, error) {
	vectors := make([]domain.Vector, len(texts))
	for i, text := range texts {
		vectors[i] = generateDeterministicVector(text, p.dimensions)
	}
	return vectors, nil
}

// ============================================================
// Internal Helper: Deterministic Normalized L2 Float Vector Generator
// ============================================================

func generateDeterministicVector(text string, dims int) domain.Vector {
	hash := sha256.Sum256([]byte(text))
	seed := binary.BigEndian.Uint64(hash[:8])

	vec := make(domain.Vector, dims)
	var sumSq float64
	for i := 0; i < dims; i++ {
		// Use LCG seeded with text hash + index
		val := math.Sin(float64(seed + uint64(i)*1000003))
		vec[i] = val
		sumSq += val * val
	}

	// Normalize to unit length (L2 norm = 1.0)
	norm := math.Sqrt(sumSq)
	if norm > 0 {
		for i := 0; i < dims; i++ {
			vec[i] /= norm
		}
	}
	return vec
}
