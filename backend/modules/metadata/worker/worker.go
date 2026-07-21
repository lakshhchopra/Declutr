package worker

import (
	"context"
	"log"

	"github.com/diablovocado/declutr/modules/metadata/application"
	processingDomain "github.com/diablovocado/declutr/modules/processing/domain"
	// Assuming shared asset service exists to get the actual file stream
)

type MetadataExtractionWorker struct {
	service application.MetadataService
	// assetService shared.AssetService // to download/read asset by ID
}

func NewMetadataExtractionWorker(service application.MetadataService) *MetadataExtractionWorker {
	return &MetadataExtractionWorker{
		service: service,
	}
}

func (w *MetadataExtractionWorker) ProcessJob(ctx context.Context, job *processingDomain.Job) error {
	if job.JobType != processingDomain.TypeMetadataExtraction {
		log.Printf("Worker ignores non-metadata job: %s", job.JobType)
		return nil
	}

	// 1. Fetch asset record from DB to get filename and mime-type
	// For stub, using dummy values
	filename := "uploaded_document.pdf"
	mimeType := "application/pdf"
	var size int64 = 1048576 // 1MB

	// 2. Open stream to asset in blob storage
	// reader, err := w.assetService.GetAssetStream(ctx, job.AssetID)
	// if err != nil { return err }
	// defer reader.Close()
	var reader = strings.NewReader("dummy asset content")

	// 3. Extract and save metadata
	_, err := w.service.ExtractAndSaveMetadata(
		ctx,
		job.AssetID,
		job.VaultID,
		filename,
		mimeType,
		size,
		reader,
	)

	if err != nil {
		log.Printf("Failed to extract metadata for asset %s: %v", job.AssetID, err)
		return err
	}

	log.Printf("Successfully extracted metadata for asset %s", job.AssetID)
	return nil
}

// Dummy strings package for the stub above
import "strings"
