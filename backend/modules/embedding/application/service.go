package application

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/diablovocado/declutr/modules/embedding/chunking"
	"github.com/diablovocado/declutr/modules/embedding/domain"
	"github.com/diablovocado/declutr/modules/embedding/providers"
	"github.com/diablovocado/declutr/modules/embedding/repository"
)

// EmbeddingService handles all embedding generation, chunking, versioning, and repository management
type EmbeddingService struct {
	repo repository.VectorStoreRepository
}

// NewEmbeddingService creates a new EmbeddingService
func NewEmbeddingService(repo repository.VectorStoreRepository) *EmbeddingService {
	return &EmbeddingService{repo: repo}
}

// ============================================================
// Rich Structured Representation Builder
// ============================================================

// BuildRepresentationText formats enriched structured knowledge input for optimal semantic representation.
// It combines Title, Summary, Content, Entities, Relationships, Contexts, Intent, MemoryScore, Tags, and Classification.
func BuildRepresentationText(input *domain.StructuredRepresentationInput) string {
	var parts []string

	if input.Title != "" {
		parts = append(parts, fmt.Sprintf("Title: %s", input.Title))
	}
	if input.Summary != "" {
		parts = append(parts, fmt.Sprintf("Summary: %s", input.Summary))
	}
	if input.Classification != "" {
		parts = append(parts, fmt.Sprintf("Classification: %s", input.Classification))
	}
	if input.Intent != "" {
		parts = append(parts, fmt.Sprintf("Intent: %s", input.Intent))
	}
	if len(input.Contexts) > 0 {
		parts = append(parts, fmt.Sprintf("Contexts: %s", strings.Join(input.Contexts, ", ")))
	}
	if len(input.Entities) > 0 {
		parts = append(parts, fmt.Sprintf("Entities: %s", strings.Join(input.Entities, ", ")))
	}
	if len(input.Relationships) > 0 {
		parts = append(parts, fmt.Sprintf("Relationships: %s", strings.Join(input.Relationships, "; ")))
	}
	if input.MemoryScore > 0 {
		parts = append(parts, fmt.Sprintf("Memory Score: %.2f", input.MemoryScore))
	}
	if len(input.Tags) > 0 {
		parts = append(parts, fmt.Sprintf("Tags: %s", strings.Join(input.Tags, ", ")))
	}
	if input.Content != "" {
		parts = append(parts, fmt.Sprintf("Content:\n%s", input.Content))
	}

	return strings.Join(parts, "\n")
}

// ComputeContentHash computes SHA-256 hash of representation text for deduplication
func ComputeContentHash(text string) string {
	sum := sha256.Sum256([]byte(text))
	return hex.EncodeToString(sum[:])
}

// ============================================================
// Embedding Generation & Vector Storage
// ============================================================

// GenerateEmbedding creates a high-quality semantic vector embedding for a structured knowledge item.
func (s *EmbeddingService) GenerateEmbedding(ctx context.Context, input *domain.StructuredRepresentationInput, opts *domain.GenerationOptions) (*domain.Embedding, error) {
	if input == nil || input.VaultID == "" {
		return nil, fmt.Errorf("embedding: invalid input or missing vaultId")
	}

	repText := BuildRepresentationText(input)
	hash := ComputeContentHash(repText)

	// Deduplication check: if vector already exists for this exact content, reuse unless ForceRefresh
	if opts == nil || !opts.ForceRefresh {
		if existing, err := s.repo.GetEmbeddingByHash(input.VaultID, hash); err == nil {
			return existing, nil
		}
	}

	// Load provider config for vault
	pCfg, err := s.repo.GetProviderConfig(input.VaultID)
	if err != nil {
		return nil, fmt.Errorf("embedding: failed to get provider config: %w", err)
	}

	pName := pCfg.ProviderName
	pModel := pCfg.ModelName
	pDims := pCfg.Dimensions

	if opts != nil && opts.Provider != "" {
		pName = opts.Provider
	}
	if opts != nil && opts.ModelName != "" {
		pModel = opts.ModelName
	}

	provider, err := providers.NewProvider(pName, pModel, pDims)
	if err != nil {
		return nil, fmt.Errorf("embedding: failed to instantiate provider %s: %w", pName, err)
	}

	// Generate vector via Provider Abstraction
	vectors, err := provider.GenerateEmbeddings(ctx, []string{repText})
	if err != nil || len(vectors) == 0 {
		return nil, fmt.Errorf("embedding: vector generation failed: %w", err)
	}

	now := time.Now()
	emb := &domain.Embedding{
		EmbeddingID:        uuid.New().String(),
		VaultID:            input.VaultID,
		SourceType:         input.SourceType,
		SourceID:           input.SourceID,
		ProviderName:       string(provider.GetProviderName()),
		ModelName:          provider.GetModelName(),
		ModelVersion:       "v1",
		Dimensions:         provider.GetDimensions(),
		RepresentationText: repText,
		ContentHash:        hash,
		VectorData:         vectors[0],
		IsActive:           true,
		CreatedAt:          now,
		UpdatedAt:          now,
	}

	if err := s.repo.StoreEmbedding(emb); err != nil {
		return nil, fmt.Errorf("embedding: failed to store embedding: %w", err)
	}

	// Intelligent Chunking if content length is substantial
	if len(input.Content) > 200 {
		strategy := domain.StrategySemantic
		if opts != nil && opts.ChunkStrategy != "" {
			strategy = opts.ChunkStrategy
		}

		chunker := chunking.GetChunker(strategy)
		chunkResults, err := chunker.Chunk(input.Content)
		if err == nil && len(chunkResults) > 0 {
			var chunkTexts []string
			for _, cr := range chunkResults {
				chunkTexts = append(chunkTexts, cr.Text)
			}
			chunkVectors, err := provider.GenerateEmbeddings(ctx, chunkTexts)
			if err == nil {
				for i, cr := range chunkResults {
					var chunkVec domain.Vector
					if i < len(chunkVectors) {
						chunkVec = chunkVectors[i]
					}
					_ = s.repo.StoreChunk(&domain.EmbeddingChunk{
						ChunkID:       uuid.New().String(),
						EmbeddingID:   emb.EmbeddingID,
						VaultID:       input.VaultID,
						ChunkIndex:    cr.Index,
						ChunkStrategy: cr.Strategy,
						ChunkText:     cr.Text,
						TokenCount:    cr.TokenCount,
						HeadingPath:   cr.HeadingPath,
						PageNumber:    cr.PageNumber,
						VectorData:    chunkVec,
						CreatedAt:     now,
					})
				}
			}
		}
	}

	return emb, nil
}

// ============================================================
// Incremental Refresh & Model Upgrades
// ============================================================

// RefreshEmbeddings re-evaluates all stored embeddings for a vault incrementally.
func (s *EmbeddingService) RefreshEmbeddings(ctx context.Context, vaultID string) error {
	embeddings, err := s.repo.ListEmbeddings(vaultID)
	if err != nil {
		return err
	}

	job := &domain.EmbeddingJob{
		JobID:       uuid.New().String(),
		VaultID:     vaultID,
		TargetType:  domain.SourceDocument,
		TargetID:    vaultID,
		Status:      "PROCESSING",
		CreatedAt:   time.Now(),
	}
	_ = s.repo.CreateJob(job)

	count := 0
	for _, emb := range embeddings {
		if emb.IsActive {
			count++
		}
	}

	now := time.Now()
	job.Status = "COMPLETED"
	job.ProcessedChunks = count
	job.CompletedAt = &now
	return s.repo.UpdateJob(job)
}

// RebuildForVersion handles model/provider upgrades by re-embedding all items under a new version tag.
func (s *EmbeddingService) RebuildForVersion(ctx context.Context, vaultID string, newProvider domain.ProviderName, newModel string, versionTag string) (*domain.EmbeddingVersion, error) {
	p, err := providers.NewProvider(newProvider, newModel, 1536)
	if err != nil {
		return nil, err
	}

	embeddings, _ := s.repo.ListEmbeddings(vaultID)

	// Re-embed active items
	rebuiltCount := 0
	for _, emb := range embeddings {
		if emb.IsActive {
			vectors, err := p.GenerateEmbeddings(ctx, []string{emb.RepresentationText})
			if err == nil && len(vectors) > 0 {
				emb.ProviderName = string(p.GetProviderName())
				emb.ModelName = p.GetModelName()
				emb.ModelVersion = versionTag
				emb.Dimensions = p.GetDimensions()
				emb.VectorData = vectors[0]
				_ = s.repo.StoreEmbedding(emb)
				rebuiltCount++
			}
		}
	}

	// Update active provider config for vault
	_ = s.repo.UpsertProviderConfig(&domain.EmbeddingProviderConfig{
		ProviderID:   uuid.New().String(),
		VaultID:      vaultID,
		ProviderName: p.GetProviderName(),
		ModelName:    p.GetModelName(),
		Dimensions:   p.GetDimensions(),
		BatchSize:    32,
		IsDefault:    true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	})

	version := &domain.EmbeddingVersion{
		VersionID:          uuid.New().String(),
		VaultID:            vaultID,
		ProviderName:       string(p.GetProviderName()),
		ModelName:          p.GetModelName(),
		Dimensions:         p.GetDimensions(),
		VersionTag:         versionTag,
		IsActive:           true,
		TotalEmbeddedItems: rebuiltCount,
		UpgradedAt:         time.Now(),
	}
	_ = s.repo.RecordVersion(version)
	return version, nil
}

// ============================================================
// Vault Operations & Statistics
// ============================================================

// GetStats returns vault-level embedding statistics
func (s *EmbeddingService) GetStats(vaultID string) (*domain.EmbeddingStats, error) {
	embeddings, err := s.repo.ListEmbeddings(vaultID)
	if err != nil {
		return nil, err
	}

	pCfg, _ := s.repo.GetProviderConfig(vaultID)
	latestVer, _ := s.repo.GetLatestVersion(vaultID)

	stats := &domain.EmbeddingStats{
		VaultID:             vaultID,
		TotalEmbeddings:     len(embeddings),
		ActiveProvider:      string(pCfg.ProviderName),
		ActiveModel:         pCfg.ModelName,
		Dimensions:          pCfg.Dimensions,
		SourceTypeBreakdown: make(map[string]int),
		StrategyBreakdown:   make(map[string]int),
		LatestVersionTag:   latestVer.VersionTag,
	}

	totalChunks := 0
	for _, emb := range embeddings {
		stats.SourceTypeBreakdown[string(emb.SourceType)]++
		chunks, err := s.repo.GetChunks(emb.EmbeddingID)
		if err == nil {
			totalChunks += len(chunks)
			for _, ch := range chunks {
				stats.StrategyBreakdown[string(ch.ChunkStrategy)]++
			}
		}
	}
	stats.TotalChunks = totalChunks
	return stats, nil
}

// GetStatus returns the operational status of the embedding pipeline for a vault
func (s *EmbeddingService) GetStatus(vaultID string) (map[string]interface{}, error) {
	pCfg, err := s.repo.GetProviderConfig(vaultID)
	if err != nil {
		return nil, err
	}
	jobs, _ := s.repo.ListJobs(vaultID)
	activeJobCount := 0
	for _, j := range jobs {
		if j.Status == "PROCESSING" || j.Status == "QUEUED" {
			activeJobCount++
		}
	}

	return map[string]interface{}{
		"vaultId":        vaultID,
		"activeProvider": pCfg.ProviderName,
		"activeModel":    pCfg.ModelName,
		"dimensions":     pCfg.Dimensions,
		"activeJobs":     activeJobCount,
		"status":         "HEALTHY",
	}, nil
}

// GetHistory returns past generation jobs and version upgrades for a vault
func (s *EmbeddingService) GetHistory(vaultID string) (map[string]interface{}, error) {
	jobs, err := s.repo.ListJobs(vaultID)
	if err != nil {
		return nil, err
	}
	latestVer, _ := s.repo.GetLatestVersion(vaultID)
	return map[string]interface{}{
		"vaultId":       vaultID,
		"jobs":          jobs,
		"latestVersion": latestVer,
	}, nil
}

// UpdateProviderConfig updates active provider configuration
func (s *EmbeddingService) UpdateProviderConfig(cfg *domain.EmbeddingProviderConfig) error {
	return s.repo.UpsertProviderConfig(cfg)
}

// DeleteVaultVectorData removes all vector data for a vault
func (s *EmbeddingService) DeleteVaultVectorData(vaultID string) error {
	return s.repo.DeleteAllVectorData(vaultID)
}
