package application

import (
	"context"
	"fmt"
	"time"

	"github.com/diablovocado/declutr/modules/entities/domain"
	"github.com/diablovocado/declutr/modules/entities/extractors"
	"github.com/diablovocado/declutr/modules/entities/repository"
)

type EntityService interface {
	ExtractAndStoreEntities(ctx context.Context, vaultID, assetID, analysisID, analysisText string) error
	GetEntitiesByVault(ctx context.Context, vaultID string) ([]*domain.Entity, error)
	GetOccurrencesByAsset(ctx context.Context, assetID string) ([]*domain.EntityOccurrence, error)
}

type DefaultEntityService struct {
	repo      repository.EntityRepository
	extractor extractors.EntityExtractor
}

func NewEntityService(repo repository.EntityRepository, extractor extractors.EntityExtractor) *DefaultEntityService {
	return &DefaultEntityService{
		repo:      repo,
		extractor: extractor,
	}
}

func (s *DefaultEntityService) ExtractAndStoreEntities(ctx context.Context, vaultID, assetID, analysisID, analysisText string) error {
	extractedEntities, err := s.extractor.ExtractEntities(ctx, analysisText)
	if err != nil {
		return fmt.Errorf("failed to extract entities: %w", err)
	}

	for _, ext := range extractedEntities {
		// 1. Resolve to canonical entity (or create new)
		entity, err := s.repo.GetEntityByCanonicalName(ctx, vaultID, string(ext.Type), ext.NormalizedValue)
		if err != nil {
			return err
		}

		if entity == nil {
			entity = &domain.Entity{
				EntityID:        "ent_" + fmt.Sprintf("%d", time.Now().UnixNano()),
				VaultID:         vaultID,
				Type:            ext.Type,
				CanonicalName:   ext.NormalizedValue, // Basic resolution logic uses normalized value as canonical
				NormalizedValue: ext.NormalizedValue,
				CreatedAt:       time.Now(),
				UpdatedAt:       time.Now(),
			}
			if err := s.repo.SaveEntity(ctx, entity); err != nil {
				return err
			}
		}

		// 2. Create Occurrence
		occurrence := &domain.EntityOccurrence{
			OccurrenceID:     "occ_" + fmt.Sprintf("%d", time.Now().UnixNano()),
			EntityID:         entity.EntityID,
			AssetID:          assetID,
			AnalysisID:       analysisID,
			OriginalValue:    ext.OriginalValue,
			ConfidenceScore:  ext.ConfidenceScore,
			ExtractorVersion: "mock-v1",
			CreatedAt:        time.Now(),
		}

		if err := s.repo.SaveOccurrence(ctx, occurrence); err != nil {
			return err
		}
	}

	return nil
}

func (s *DefaultEntityService) GetEntitiesByVault(ctx context.Context, vaultID string) ([]*domain.Entity, error) {
	return s.repo.GetEntitiesByVault(ctx, vaultID)
}

func (s *DefaultEntityService) GetOccurrencesByAsset(ctx context.Context, assetID string) ([]*domain.EntityOccurrence, error) {
	return s.repo.GetOccurrencesByAsset(ctx, assetID)
}
