package worker

import (
	"context"
	"log"

	"github.com/diablovocado/declutr/modules/embedding/application"
	"github.com/diablovocado/declutr/modules/embedding/domain"
	processingDomain "github.com/diablovocado/declutr/modules/processing/domain"
)

// EmbeddingWorker processes EMBEDDING_GENERATION jobs from the processing pipeline.
//
// Pipeline position:
//
//	Memory Engine → Embedding Engine → Vector Storage
type EmbeddingWorker struct {
	engine  *application.EmbeddingEngine
	service *application.EmbeddingService
}

// NewEmbeddingWorker creates a new EmbeddingWorker
func NewEmbeddingWorker(engine *application.EmbeddingEngine, service *application.EmbeddingService) *EmbeddingWorker {
	return &EmbeddingWorker{engine: engine, service: service}
}

// ProcessJob handles EMBEDDING_GENERATION jobs.
// It constructs a rich structured representation for the target asset/knowledge item and generates its vector embedding.
func (w *EmbeddingWorker) ProcessJob(ctx context.Context, job *processingDomain.Job) error {
	if job.JobType != processingDomain.TypeEmbeddingGen {
		log.Printf("[EmbeddingWorker] Ignoring non-embedding job type: %s", job.JobType)
		return nil
	}

	vaultID := job.VaultID
	if vaultID == "" {
		vaultID = "v_default"
	}

	log.Printf("[EmbeddingWorker] Vectorising asset %s in vault %s", job.AssetID, vaultID)

	if job.AssetID != "" {
		input := &domain.StructuredRepresentationInput{
			SourceType:  domain.SourceDocument,
			SourceID:    job.AssetID,
			VaultID:     vaultID,
			Title:       "Processed Asset: " + job.AssetID,
			Summary:     "Vector embedding generated after full AI pipeline completion",
			Content:     "Structured knowledge representation for asset " + job.AssetID,
			MemoryScore: 0.8,
			Tags:        []string{"auto-vectorised", "pipeline"},
		}
		if _, err := w.engine.EmbedStructuredItem(ctx, input); err != nil {
			log.Printf("[EmbeddingWorker] Embedding generation failed for asset %s: %v", job.AssetID, err)
			return err
		}
	}

	// Execute incremental vault vectorization cycle
	if err := w.engine.ProcessVault(ctx, vaultID); err != nil {
		log.Printf("[EmbeddingWorker] Vectorization cycle failed for vault %s: %v", vaultID, err)
		return err
	}

	log.Printf("[EmbeddingWorker] Vectorization complete for vault %s", vaultID)
	return nil
}
