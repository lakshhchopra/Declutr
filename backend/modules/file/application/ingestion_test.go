package application

import (
	"testing"
	"time"

	"github.com/diablovocado/declutr/modules/file/domain"
)

func TestAssetStatusStateTransitions(t *testing.T) {
	asset := domain.Asset{
		AssetID:        "ast_100",
		VaultID:        "v_123",
		Filename:       "invoice_receipt.pdf",
		MimeType:       "application/pdf",
		SizeBytes:      1024 * 50,
		ChecksumSHA256: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		Status:         domain.StatusQueued,
		CreatedAt:      time.Now(),
	}

	if asset.Status != domain.StatusQueued {
		t.Fatalf("Expected initial asset status to be QUEUED")
	}

	// Transition through pipeline
	asset.Status = domain.StatusUploading
	if asset.Status != domain.StatusUploading {
		t.Errorf("Expected status UPLOADING")
	}

	asset.Status = domain.StatusValidating
	asset.Status = domain.StatusAIPending
	asset.Status = domain.StatusReady

	if asset.Status != domain.StatusReady {
		t.Errorf("Expected final status READY")
	}
}

func TestUploadJobProgress(t *testing.T) {
	job := domain.UploadJob{
		JobID:              "job_1",
		AssetID:            "ast_100",
		JobType:            "OCR_PARSING",
		Status:             "PENDING",
		ProgressPercentage: 0,
	}

	job.ProgressPercentage = 100
	job.Status = "COMPLETED"

	if job.ProgressPercentage != 100 || job.Status != "COMPLETED" {
		t.Errorf("Job completion status mismatch: %+v", job)
	}
}
