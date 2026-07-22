package repository

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/embedding/domain"
)

// VectorStoreRepository defines the vendor-independent database interface
type VectorStoreRepository interface {
	// Embedding operations
	StoreEmbedding(emb *domain.Embedding) error
	GetEmbedding(embeddingID string) (*domain.Embedding, error)
	GetEmbeddingBySource(vaultID string, sourceType domain.SourceType, sourceID string) (*domain.Embedding, error)
	GetEmbeddingByHash(vaultID, hash string) (*domain.Embedding, error)
	ListEmbeddings(vaultID string) ([]*domain.Embedding, error)
	DeleteEmbedding(embeddingID string) error

	// Chunks
	StoreChunk(chunk *domain.EmbeddingChunk) error
	GetChunks(embeddingID string) ([]*domain.EmbeddingChunk, error)

	// Versions
	RecordVersion(version *domain.EmbeddingVersion) error
	GetLatestVersion(vaultID string) (*domain.EmbeddingVersion, error)

	// Jobs
	CreateJob(job *domain.EmbeddingJob) error
	UpdateJob(job *domain.EmbeddingJob) error
	GetJob(jobID string) (*domain.EmbeddingJob, error)
	ListJobs(vaultID string) ([]*domain.EmbeddingJob, error)

	// Providers
	UpsertProviderConfig(cfg *domain.EmbeddingProviderConfig) error
	GetProviderConfig(vaultID string) (*domain.EmbeddingProviderConfig, error)

	// Vector Similarity Search (foundation interface, not RAG search API)
	FindNearest(vaultID string, target domain.Vector, limit int) ([]*domain.Embedding, error)

	// Vault operations
	DeleteAllVectorData(vaultID string) error
}

// InMemoryVectorRepository is a thread-safe in-memory vector store driver
type InMemoryVectorRepository struct {
	mu          sync.RWMutex
	embeddings  map[string]*domain.Embedding         // embeddingID → embedding
	chunks      map[string][]*domain.EmbeddingChunk // embeddingID → chunks
	versions    map[string][]*domain.EmbeddingVersion// vaultID → versions
	jobs        map[string]*domain.EmbeddingJob     // jobID → job
	providers   map[string]*domain.EmbeddingProviderConfig // vaultID → provider
}

// NewInMemoryVectorRepository creates a new in-memory vector repository
func NewInMemoryVectorRepository() *InMemoryVectorRepository {
	return &InMemoryVectorRepository{
		embeddings: make(map[string]*domain.Embedding),
		chunks:     make(map[string][]*domain.EmbeddingChunk),
		versions:   make(map[string][]*domain.EmbeddingVersion),
		jobs:       make(map[string]*domain.EmbeddingJob),
		providers:  make(map[string]*domain.EmbeddingProviderConfig),
	}
}

// --- Embeddings ---

func (r *InMemoryVectorRepository) StoreEmbedding(emb *domain.Embedding) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	emb.UpdatedAt = time.Now()
	r.embeddings[emb.EmbeddingID] = emb
	return nil
}

func (r *InMemoryVectorRepository) GetEmbedding(embeddingID string) (*domain.Embedding, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if emb, ok := r.embeddings[embeddingID]; ok {
		return emb, nil
	}
	return nil, fmt.Errorf("embedding %s not found", embeddingID)
}

func (r *InMemoryVectorRepository) GetEmbeddingBySource(vaultID string, sourceType domain.SourceType, sourceID string) (*domain.Embedding, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, emb := range r.embeddings {
		if emb.VaultID == vaultID && emb.SourceType == sourceType && emb.SourceID == sourceID && emb.IsActive {
			return emb, nil
		}
	}
	return nil, fmt.Errorf("embedding not found for source %s/%s in vault %s", sourceType, sourceID, vaultID)
}

func (r *InMemoryVectorRepository) GetEmbeddingByHash(vaultID, hash string) (*domain.Embedding, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, emb := range r.embeddings {
		if emb.VaultID == vaultID && emb.ContentHash == hash && emb.IsActive {
			return emb, nil
		}
	}
	return nil, fmt.Errorf("embedding with hash %s not found in vault %s", hash, vaultID)
}

func (r *InMemoryVectorRepository) ListEmbeddings(vaultID string) ([]*domain.Embedding, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.Embedding
	for _, emb := range r.embeddings {
		if emb.VaultID == vaultID {
			result = append(result, emb)
		}
	}
	return result, nil
}

func (r *InMemoryVectorRepository) DeleteEmbedding(embeddingID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.embeddings, embeddingID)
	delete(r.chunks, embeddingID)
	return nil
}

// --- Chunks ---

func (r *InMemoryVectorRepository) StoreChunk(chunk *domain.EmbeddingChunk) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.chunks[chunk.EmbeddingID] = append(r.chunks[chunk.EmbeddingID], chunk)
	return nil
}

func (r *InMemoryVectorRepository) GetChunks(embeddingID string) ([]*domain.EmbeddingChunk, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.chunks[embeddingID], nil
}

// --- Versions ---

func (r *InMemoryVectorRepository) RecordVersion(version *domain.EmbeddingVersion) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.versions[version.VaultID] = append(r.versions[version.VaultID], version)
	return nil
}

func (r *InMemoryVectorRepository) GetLatestVersion(vaultID string) (*domain.EmbeddingVersion, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	versions := r.versions[vaultID]
	if len(versions) == 0 {
		return &domain.EmbeddingVersion{
			VersionID:  "v1-init",
			VaultID:    vaultID,
			ProviderName: "LOCAL",
			ModelName:  "local-deterministic-v1",
			Dimensions: 1536,
			VersionTag: "v1.0.0",
			IsActive:   true,
			UpgradedAt: time.Now(),
		}, nil
	}
	return versions[len(versions)-1], nil
}

// --- Jobs ---

func (r *InMemoryVectorRepository) CreateJob(job *domain.EmbeddingJob) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.jobs[job.JobID] = job
	return nil
}

func (r *InMemoryVectorRepository) UpdateJob(job *domain.EmbeddingJob) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.jobs[job.JobID] = job
	return nil
}

func (r *InMemoryVectorRepository) GetJob(jobID string) (*domain.EmbeddingJob, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if job, ok := r.jobs[jobID]; ok {
		return job, nil
	}
	return nil, fmt.Errorf("job %s not found", jobID)
}

func (r *InMemoryVectorRepository) ListJobs(vaultID string) ([]*domain.EmbeddingJob, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.EmbeddingJob
	for _, j := range r.jobs {
		if j.VaultID == vaultID {
			result = append(result, j)
		}
	}
	return result, nil
}

// --- Providers ---

func (r *InMemoryVectorRepository) UpsertProviderConfig(cfg *domain.EmbeddingProviderConfig) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cfg.UpdatedAt = time.Now()
	r.providers[cfg.VaultID] = cfg
	return nil
}

func (r *InMemoryVectorRepository) GetProviderConfig(vaultID string) (*domain.EmbeddingProviderConfig, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if cfg, ok := r.providers[vaultID]; ok {
		return cfg, nil
	}
	return &domain.EmbeddingProviderConfig{
		ProviderID:   "p-default",
		VaultID:      vaultID,
		ProviderName: domain.ProviderLocal,
		ModelName:    "local-deterministic-v1",
		Dimensions:   1536,
		BatchSize:    32,
		IsDefault:    true,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil
}

// --- Vector Similarity Search ---

type searchItem struct {
	emb   *domain.Embedding
	score float64
}

func (r *InMemoryVectorRepository) FindNearest(vaultID string, target domain.Vector, limit int) ([]*domain.Embedding, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var items []searchItem
	for _, emb := range r.embeddings {
		if emb.VaultID == vaultID && emb.IsActive && len(emb.VectorData) == len(target) {
			sim := cosineSimilarity(target, emb.VectorData)
			items = append(items, searchItem{emb: emb, score: sim})
		}
	}

	sort.Slice(items, func(i, j int) bool {
		return items[i].score > items[j].score
	})

	var result []*domain.Embedding
	for i := 0; i < len(items) && i < limit; i++ {
		result = append(result, items[i].emb)
	}
	return result, nil
}

func cosineSimilarity(a, b domain.Vector) float64 {
	if len(a) != len(b) || len(a) == 0 {
		return 0.0
	}
	var dot, normA, normB float64
	for i := 0; i < len(a); i++ {
		dot += a[i] * b[i]
		normA += a[i] * a[i]
		normB += b[i] * b[i]
	}
	if normA == 0 || normB == 0 {
		return 0.0
	}
	return dot / (math.Sqrt(normA) * math.Sqrt(normB))
}

// --- Delete All ---

func (r *InMemoryVectorRepository) DeleteAllVectorData(vaultID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for id, emb := range r.embeddings {
		if emb.VaultID == vaultID {
			delete(r.embeddings, id)
			delete(r.chunks, id)
		}
	}
	for id, j := range r.jobs {
		if j.VaultID == vaultID {
			delete(r.jobs, id)
		}
	}
	delete(r.versions, vaultID)
	delete(r.providers, vaultID)
	return nil
}
