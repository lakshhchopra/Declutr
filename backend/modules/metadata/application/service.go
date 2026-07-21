package application

import (
	"context"
	"io"
	"fmt"
	"time"

	"github.com/diablovocado/declutr/modules/metadata/domain"
	"github.com/diablovocado/declutr/modules/metadata/extractors"
	"github.com/diablovocado/declutr/modules/metadata/repository"
)

type MetadataService interface {
	ExtractAndSaveMetadata(ctx context.Context, assetID, vaultID, filename, mimeType string, size int64, reader io.Reader) (*domain.CompleteMetadata, error)
	GetMetadata(ctx context.Context, assetID string) (*domain.CompleteMetadata, error)
	GetVersionHistory(ctx context.Context, assetID string) ([]*domain.MetadataVersion, error)
}

type DefaultMetadataService struct {
	repo     repository.MetadataRepository
	registry *extractors.ExtractorRegistry
}

func NewMetadataService(repo repository.MetadataRepository) *DefaultMetadataService {
	return &DefaultMetadataService{
		repo:     repo,
		registry: extractors.NewExtractorRegistry(),
	}
}

func (s *DefaultMetadataService) ExtractAndSaveMetadata(ctx context.Context, assetID, vaultID, filename, mimeType string, size int64, reader io.Reader) (*domain.CompleteMetadata, error) {
	extractor := s.registry.GetExtractor(mimeType)
	
	meta, err := extractor.Extract(ctx, assetID, vaultID, filename, mimeType, size, reader)
	if err != nil {
		return nil, err
	}

	// Persist
	if err := s.repo.SaveMetadata(ctx, meta.General); err != nil {
		return nil, err
	}

	if meta.Properties != nil {
		if err := s.repo.SaveProperties(ctx, meta.Properties); err != nil {
			return nil, err
		}
	}

	if meta.Exif != nil {
		if err := s.repo.SaveExif(ctx, meta.Exif); err != nil {
			return nil, err
		}
	}

	version := &domain.MetadataVersion{
		VersionID: "ver_" + fmt.Sprintf("%d", time.Now().UnixNano()),
		AssetID:   assetID,
		Source:    "SYSTEM_EXTRACTOR",
		ExtractorVersion: "1.0",
		Snapshot:  map[string]interface{}{"metadata_state": "extracted"},
		CreatedAt: time.Now(),
	}

	if err := s.repo.SaveVersion(ctx, version); err != nil {
		return nil, err
	}

	return meta, nil
}

func (s *DefaultMetadataService) GetMetadata(ctx context.Context, assetID string) (*domain.CompleteMetadata, error) {
	general, err := s.repo.GetMetadata(ctx, assetID)
	if err != nil {
		return nil, err
	}

	props, _ := s.repo.GetProperties(ctx, assetID)
	exif, _ := s.repo.GetExif(ctx, assetID)

	return &domain.CompleteMetadata{
		General:    general,
		Properties: props,
		Exif:       exif,
	}, nil
}

func (s *DefaultMetadataService) GetVersionHistory(ctx context.Context, assetID string) ([]*domain.MetadataVersion, error) {
	return s.repo.GetVersionHistory(ctx, assetID)
}
