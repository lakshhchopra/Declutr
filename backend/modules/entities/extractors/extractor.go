package extractors

import (
	"context"

	"github.com/diablovocado/declutr/modules/entities/domain"
)

type EntityExtractor interface {
	ExtractEntities(ctx context.Context, analysisText string) ([]domain.ExtractedEntity, error)
}

// MockEntityExtractor provides deterministic entity extraction for testing without LLMs.
type MockEntityExtractor struct{}

func NewMockEntityExtractor() *MockEntityExtractor {
	return &MockEntityExtractor{}
}

func (e *MockEntityExtractor) ExtractEntities(ctx context.Context, analysisText string) ([]domain.ExtractedEntity, error) {
	return []domain.ExtractedEntity{
		{
			Type:            domain.TypeOrganization,
			OriginalValue:   "Google LLC",
			NormalizedValue: "Google",
			ConfidenceScore: 0.99,
		},
		{
			Type:            domain.TypeLocation,
			OriginalValue:   "NYC",
			NormalizedValue: "New York City",
			ConfidenceScore: 0.95,
		},
		{
			Type:            domain.TypeDate,
			OriginalValue:   "Oct 25th 2023",
			NormalizedValue: "2023-10-25",
			ConfidenceScore: 0.98,
		},
		{
			Type:            domain.TypeAmount,
			OriginalValue:   "$1,500.50",
			NormalizedValue: "1500.50 USD",
			ConfidenceScore: 0.99,
		},
	}, nil
}
