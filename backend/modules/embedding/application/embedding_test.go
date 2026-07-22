package application_test

import (
	"context"
	"testing"

	"github.com/diablovocado/declutr/modules/embedding/application"
	"github.com/diablovocado/declutr/modules/embedding/chunking"
	"github.com/diablovocado/declutr/modules/embedding/domain"
	"github.com/diablovocado/declutr/modules/embedding/providers"
	"github.com/diablovocado/declutr/modules/embedding/repository"
)

const testVaultID = "vault-test-001"

func setupService() *application.EmbeddingService {
	repo := repository.NewInMemoryVectorRepository()
	return application.NewEmbeddingService(repo)
}

// TestStructuredEmbeddingGeneration validates generating rich structured vector embeddings
func TestStructuredEmbeddingGeneration(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	input := &domain.StructuredRepresentationInput{
		SourceType:     domain.SourceDocument,
		SourceID:       "doc-001",
		VaultID:        testVaultID,
		Title:          "Japan Travel Itinerary 2025",
		Summary:        "Detailed schedule covering Tokyo, Kyoto, and Osaka.",
		Classification: "Travel Document",
		Intent:         "Vacation Planning",
		Contexts:       []string{"Japan Vacation", "Family Trip"},
		Entities:       []string{"Tokyo", "Kyoto", "Narita Airport"},
		Relationships:  []string{"Tokyo MENTIONS Narita Airport"},
		MemoryScore:    0.88,
		Tags:           []string{"travel", "japan", "flights"},
	}

	emb, err := svc.GenerateEmbedding(ctx, input, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if emb == nil {
		t.Fatal("expected embedding to be returned, got nil")
	}
	if emb.EmbeddingID == "" {
		t.Error("expected non-empty embedding ID")
	}
	if len(emb.VectorData) != 1536 {
		t.Errorf("expected 1536 dimensions, got %d", len(emb.VectorData))
	}
	if emb.ContentHash == "" {
		t.Error("expected non-empty SHA-256 content hash")
	}

	// Test Deduplication — generating exact same input without force refresh should return same embedding
	emb2, err := svc.GenerateEmbedding(ctx, input, nil)
	if err != nil {
		t.Fatalf("deduplication check failed: %v", err)
	}
	if emb2.EmbeddingID != emb.EmbeddingID {
		t.Errorf("expected duplicate to return existing ID %s, got %s", emb.EmbeddingID, emb2.EmbeddingID)
	}

	t.Logf("PASS: Structured Embedding Generated — ID=%s, Hash=%s, Dims=%d",
		emb.EmbeddingID, emb.ContentHash[:8], len(emb.VectorData))
}

// TestIntelligentChunking validates all 5 chunking strategies
func TestIntelligentChunking(t *testing.T) {
	sampleDoc := `# Executive Summary
Declutr is an intelligent digital vault.

## Architecture
The system uses modular monorepo architecture.

### Storage
PostgreSQL and pgvector handle metadata and embeddings.
---
Page 2
## Features
Context Engine and Memory Engine automate knowledge organisation.`

	strategies := []domain.ChunkStrategy{
		domain.StrategySemantic,
		domain.StrategyHeading,
		domain.StrategyPage,
		domain.StrategyHierarchical,
		domain.StrategyDocument,
	}

	for _, strat := range strategies {
		chunker := chunking.GetChunker(strat)
		results, err := chunker.Chunk(sampleDoc)
		if err != nil {
			t.Fatalf("chunking failed for strategy %s: %v", strat, err)
		}
		if len(results) == 0 {
			t.Errorf("expected chunks for strategy %s, got 0", strat)
		}
		t.Logf("PASS: Strategy %s produced %d chunks (first chunk text: %q)",
			strat, len(results), truncate(results[0].Text, 40))
	}
}

// TestProviderSwitching validates instantiating and running all provider abstractions
func TestProviderSwitching(t *testing.T) {
	ctx := context.Background()
	providersList := []domain.ProviderName{
		domain.ProviderOpenAI,
		domain.ProviderGemini,
		domain.ProviderVoyage,
		domain.ProviderCohere,
		domain.ProviderOllama,
		domain.ProviderLocal,
	}

	for _, pName := range providersList {
		p, err := providers.NewProvider(pName, "", 0)
		if err != nil {
			t.Fatalf("failed to create provider %s: %v", pName, err)
		}
		vectors, err := p.GenerateEmbeddings(ctx, []string{"Test structured knowledge text"})
		if err != nil {
			t.Fatalf("provider %s failed to generate vector: %v", pName, err)
		}
		if len(vectors) != 1 || len(vectors[0]) != p.GetDimensions() {
			t.Errorf("provider %s returned invalid vector length: expected %d, got %d",
				pName, p.GetDimensions(), len(vectors[0]))
		}
		t.Logf("PASS: Provider %s — Model=%s, Dims=%d", pName, p.GetModelName(), p.GetDimensions())
	}
}

// TestVectorStoreRepository validates storing embeddings and Cosine Similarity nearest neighbor search
func TestVectorStoreRepository(t *testing.T) {
	repo := repository.NewInMemoryVectorRepository()
	ctx := context.Background()
	provider, _ := providers.NewProvider(domain.ProviderLocal, "", 128)

	// Create 3 vectors
	texts := []string{"Medical prescription report", "Medical doctor consultation", "Flight ticket booking to Tokyo"}
	for i, text := range texts {
		vecs, _ := provider.GenerateEmbeddings(ctx, []string{text})
		_ = repo.StoreEmbedding(&domain.Embedding{
			EmbeddingID:        []string{"emb-med-1", "emb-med-2", "emb-fly-1"}[i],
			VaultID:            testVaultID,
			SourceType:         domain.SourceDocument,
			SourceID:           []string{"doc-1", "doc-2", "doc-3"}[i],
			ProviderName:       "LOCAL",
			ModelName:          "local-v1",
			Dimensions:         128,
			RepresentationText: text,
			ContentHash:        text,
			VectorData:         vecs[0],
			IsActive:           true,
		})
	}

	// Query with text similar to medical
	queryVecs, _ := provider.GenerateEmbeddings(ctx, []string{"Medical prescription report"})
	nearest, err := repo.FindNearest(testVaultID, queryVecs[0], 2)
	if err != nil {
		t.Fatalf("FindNearest failed: %v", err)
	}
	if len(nearest) == 0 {
		t.Fatal("expected nearest vectors, got 0")
	}
	// The top result should be exact match "emb-med-1"
	if nearest[0].EmbeddingID != "emb-med-1" {
		t.Errorf("expected top result 'emb-med-1', got %s", nearest[0].EmbeddingID)
	}
	t.Logf("PASS: Vector Store Nearest — Top result: %s (%s)", nearest[0].EmbeddingID, nearest[0].RepresentationText)
}

// TestIncrementalUpdate validates background job creation and refreshing
func TestIncrementalUpdate(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	// Insert an embedding first
	_, _ = svc.GenerateEmbedding(ctx, &domain.StructuredRepresentationInput{
		SourceType: domain.SourceMemory,
		SourceID:   "mem-001",
		VaultID:    testVaultID,
		Title:      "Thesis Research",
		Summary:    "Deep learning notes",
	}, nil)

	err := svc.RefreshEmbeddings(ctx, testVaultID)
	if err != nil {
		t.Fatalf("refresh failed: %v", err)
	}

	stats, err := svc.GetStats(testVaultID)
	if err != nil {
		t.Fatalf("get stats failed: %v", err)
	}
	if stats.TotalEmbeddings != 1 {
		t.Errorf("expected 1 total embedding, got %d", stats.TotalEmbeddings)
	}
	t.Logf("PASS: Incremental Update — %d embeddings active in vault %s", stats.TotalEmbeddings, testVaultID)
}

// TestVersionUpgrade validates model/version upgrades and re-indexing
func TestVersionUpgrade(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	// Initial embedding with v1
	_, _ = svc.GenerateEmbedding(ctx, &domain.StructuredRepresentationInput{
		SourceType: domain.SourceDocument,
		SourceID:   "doc-v1",
		VaultID:    testVaultID,
		Title:      "Tax Filing 2024",
		Summary:    "Tax forms and receipts",
	}, nil)

	// Rebuild for version v2.0.0 using GEMINI provider
	ver, err := svc.RebuildForVersion(ctx, testVaultID, domain.ProviderGemini, "text-embedding-004", "v2.0.0")
	if err != nil {
		t.Fatalf("rebuild for version failed: %v", err)
	}

	if ver.VersionTag != "v2.0.0" {
		t.Errorf("expected version tag 'v2.0.0', got %s", ver.VersionTag)
	}
	if ver.ProviderName != "GEMINI" {
		t.Errorf("expected provider GEMINI, got %s", ver.ProviderName)
	}
	t.Logf("PASS: Version Upgrade — Upgraded to %s tag %s (%d items rebuilt)",
		ver.ProviderName, ver.VersionTag, ver.TotalEmbeddedItems)
}

func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max] + "…"
	}
	return s
}
