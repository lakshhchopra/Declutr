package repository

import (
	"context"

	"github.com/diablovocado/declutr/modules/metadata/domain"
)

type MetadataRepository interface {
	SaveMetadata(ctx context.Context, metadata *domain.AssetMetadata) error
	GetMetadata(ctx context.Context, assetID string) (*domain.AssetMetadata, error)
	
	SaveProperties(ctx context.Context, props *domain.AssetProperties) error
	GetProperties(ctx context.Context, assetID string) (*domain.AssetProperties, error)
	
	SaveExif(ctx context.Context, exif *domain.AssetExif) error
	GetExif(ctx context.Context, assetID string) (*domain.AssetExif, error)
	
	SaveVersion(ctx context.Context, version *domain.MetadataVersion) error
	GetVersionHistory(ctx context.Context, assetID string) ([]*domain.MetadataVersion, error)
	GetVersion(ctx context.Context, versionID string) (*domain.MetadataVersion, error)
}

type DefaultMetadataRepository struct {
	// Add DB connection here (e.g. pgxpool)
}

func NewMetadataRepository() *DefaultMetadataRepository {
	return &DefaultMetadataRepository{}
}

func (r *DefaultMetadataRepository) SaveMetadata(ctx context.Context, metadata *domain.AssetMetadata) error {
	// Implementation would persist to asset_metadata table
	return nil
}

func (r *DefaultMetadataRepository) GetMetadata(ctx context.Context, assetID string) (*domain.AssetMetadata, error) {
	// Dummy
	return nil, nil
}

func (r *DefaultMetadataRepository) SaveProperties(ctx context.Context, props *domain.AssetProperties) error {
	return nil
}

func (r *DefaultMetadataRepository) GetProperties(ctx context.Context, assetID string) (*domain.AssetProperties, error) {
	return nil, nil
}

func (r *DefaultMetadataRepository) SaveExif(ctx context.Context, exif *domain.AssetExif) error {
	return nil
}

func (r *DefaultMetadataRepository) GetExif(ctx context.Context, assetID string) (*domain.AssetExif, error) {
	return nil, nil
}

func (r *DefaultMetadataRepository) SaveVersion(ctx context.Context, version *domain.MetadataVersion) error {
	return nil
}

func (r *DefaultMetadataRepository) GetVersionHistory(ctx context.Context, assetID string) ([]*domain.MetadataVersion, error) {
	return nil, nil
}

func (r *DefaultMetadataRepository) GetVersion(ctx context.Context, versionID string) (*domain.MetadataVersion, error) {
	return nil, nil
}
