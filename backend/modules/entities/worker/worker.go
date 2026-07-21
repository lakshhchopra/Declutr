package worker

import (
	"context"
	"log"

	"github.com/diablovocado/declutr/modules/entities/application"
	processingDomain "github.com/diablovocado/declutr/modules/processing/domain"
)

type EntityExtractionWorker struct {
	service application.EntityService
}

func NewEntityExtractionWorker(service application.EntityService) *EntityExtractionWorker {
	return &EntityExtractionWorker{
		service: service,
	}
}

func (w *EntityExtractionWorker) ProcessJob(ctx context.Context, job *processingDomain.Job) error {
	if job.JobType != processingDomain.TypeEntityExtraction {
		log.Printf("Worker ignores non-entity job: %s", job.JobType)
		return nil
	}

	// 1. Fetch AI Analysis for the Asset (stubbed)
	vaultID := "v_123"
	analysisID := "ai_123"
	analysisText := "This is a document about Google LLC located in NYC from Oct 25th 2023 regarding an amount of $1,500.50."

	// 2. Run Entity Extraction
	err := w.service.ExtractAndStoreEntities(
		ctx,
		vaultID,
		job.AssetID,
		analysisID,
		analysisText,
	)

	if err != nil {
		log.Printf("Failed to run entity extraction for asset %s: %v", job.AssetID, err)
		return err
	}

	log.Printf("Successfully completed entity extraction for asset %s", job.AssetID)
	return nil
}
