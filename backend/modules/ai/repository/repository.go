package repository

import (
	"context"

	"github.com/diablovocado/declutr/modules/ai/domain"
)

type AIAnalysisRepository interface {
	SaveAnalysis(ctx context.Context, analysis *domain.AIAnalysis) error
	GetAnalysis(ctx context.Context, assetID string) (*domain.AIAnalysis, error)
	
	SaveVersion(ctx context.Context, version *domain.AnalysisVersion) error
	GetVersionHistory(ctx context.Context, analysisID string) ([]*domain.AnalysisVersion, error)
}

type DefaultAIAnalysisRepository struct {
	// DB conn
}

func NewAIAnalysisRepository() *DefaultAIAnalysisRepository {
	return &DefaultAIAnalysisRepository{}
}

func (r *DefaultAIAnalysisRepository) SaveAnalysis(ctx context.Context, analysis *domain.AIAnalysis) error {
	// Persist to ai_analysis, ai_tags, ai_classification, ai_topics
	return nil
}

func (r *DefaultAIAnalysisRepository) GetAnalysis(ctx context.Context, assetID string) (*domain.AIAnalysis, error) {
	return nil, nil
}

func (r *DefaultAIAnalysisRepository) SaveVersion(ctx context.Context, version *domain.AnalysisVersion) error {
	return nil
}

func (r *DefaultAIAnalysisRepository) GetVersionHistory(ctx context.Context, analysisID string) ([]*domain.AnalysisVersion, error) {
	return nil, nil
}
