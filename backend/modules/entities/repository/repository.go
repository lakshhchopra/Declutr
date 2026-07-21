package repository

import (
	"context"

	"github.com/diablovocado/declutr/modules/entities/domain"
)

type EntityRepository interface {
	SaveEntity(ctx context.Context, entity *domain.Entity) error
	GetEntityByCanonicalName(ctx context.Context, vaultID, entityType, canonicalName string) (*domain.Entity, error)
	GetEntitiesByVault(ctx context.Context, vaultID string) ([]*domain.Entity, error)
	
	SaveOccurrence(ctx context.Context, occurrence *domain.EntityOccurrence) error
	GetOccurrencesByAsset(ctx context.Context, assetID string) ([]*domain.EntityOccurrence, error)
	GetOccurrencesByEntity(ctx context.Context, entityID string) ([]*domain.EntityOccurrence, error)
}

type DefaultEntityRepository struct {
	// DB conn
}

func NewEntityRepository() *DefaultEntityRepository {
	return &DefaultEntityRepository{}
}

func (r *DefaultEntityRepository) SaveEntity(ctx context.Context, entity *domain.Entity) error {
	// Upsert entity, returning ID
	return nil
}

func (r *DefaultEntityRepository) GetEntityByCanonicalName(ctx context.Context, vaultID, entityType, canonicalName string) (*domain.Entity, error) {
	return nil, nil
}

func (r *DefaultEntityRepository) GetEntitiesByVault(ctx context.Context, vaultID string) ([]*domain.Entity, error) {
	return nil, nil
}

func (r *DefaultEntityRepository) SaveOccurrence(ctx context.Context, occurrence *domain.EntityOccurrence) error {
	return nil
}

func (r *DefaultEntityRepository) GetOccurrencesByAsset(ctx context.Context, assetID string) ([]*domain.EntityOccurrence, error) {
	return nil, nil
}

func (r *DefaultEntityRepository) GetOccurrencesByEntity(ctx context.Context, entityID string) ([]*domain.EntityOccurrence, error) {
	return nil, nil
}
