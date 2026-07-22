package application

import (
	"context"
	"fmt"
	"log"

	"github.com/diablovocado/declutr/modules/embedding/domain"
)

// EmbeddingEngine orchestrates knowledge vectorization across a vault
type EmbeddingEngine struct {
	service *EmbeddingService
}

// NewEmbeddingEngine creates a new EmbeddingEngine
func NewEmbeddingEngine(service *EmbeddingService) *EmbeddingEngine {
	return &EmbeddingEngine{service: service}
}

// ProcessVault executes one incremental embedding cycle for a vault.
// It generates vectors for any un-embedded structured knowledge items and updates statistics.
func (e *EmbeddingEngine) ProcessVault(ctx context.Context, vaultID string) error {
	log.Printf("[EmbeddingEngine] Starting incremental vectorization cycle for vault: %s", vaultID)

	if err := e.service.RefreshEmbeddings(ctx, vaultID); err != nil {
		return fmt.Errorf("embedding engine: refresh failed for vault %s: %w", vaultID, err)
	}

	stats, _ := e.service.GetStats(vaultID)
	if stats != nil {
		log.Printf("[EmbeddingEngine] Vault %s: %d embeddings | %d chunks | Provider=%s | Model=%s (%d dims)",
			vaultID, stats.TotalEmbeddings, stats.TotalChunks, stats.ActiveProvider, stats.ActiveModel, stats.Dimensions)
	}

	return nil
}

// EmbedStructuredItem is a direct helper to vectorise a single structured knowledge representation
func (e *EmbeddingEngine) EmbedStructuredItem(ctx context.Context, input *domain.StructuredRepresentationInput) (*domain.Embedding, error) {
	return e.service.GenerateEmbedding(ctx, input, nil)
}
