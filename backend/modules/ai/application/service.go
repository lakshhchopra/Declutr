package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/diablovocado/declutr/modules/ai/domain"
	"github.com/diablovocado/declutr/modules/ai/prompts"
	"github.com/diablovocado/declutr/modules/ai/providers"
	"github.com/diablovocado/declutr/modules/ai/repository"
)

type AIAnalysisService interface {
	AnalyzeDocument(ctx context.Context, documentID, assetID, extractedText string) (*domain.AIAnalysis, error)
	GetAnalysis(ctx context.Context, assetID string) (*domain.AIAnalysis, error)
	GetVersionHistory(ctx context.Context, analysisID string) ([]*domain.AnalysisVersion, error)
}

type DefaultAIAnalysisService struct {
	repo          repository.AIAnalysisRepository
	provider      providers.LLMProvider
	promptManager *prompts.PromptManager
}

func NewAIAnalysisService(repo repository.AIAnalysisRepository, provider providers.LLMProvider) *DefaultAIAnalysisService {
	return &DefaultAIAnalysisService{
		repo:          repo,
		provider:      provider,
		promptManager: prompts.NewPromptManager(),
	}
}

func (s *DefaultAIAnalysisService) AnalyzeDocument(ctx context.Context, documentID, assetID, extractedText string) (*domain.AIAnalysis, error) {
	if extractedText == "" {
		return nil, errors.New("cannot analyze empty text")
	}

	sysPrompt := s.promptManager.GetSystemPrompt()
	userPrompt := s.promptManager.BuildUserPrompt(extractedText)

	// Call Provider with retry logic (RetryManager simplified for stub)
	// In reality, this would wrap the s.provider.Analyze call in a loop with exponential backoff
	var analysis *domain.AIAnalysis
	var version *domain.AnalysisVersion
	var err error

	maxRetries := 3
	for i := 0; i < maxRetries; i++ {
		analysis, version, err = s.provider.Analyze(ctx, sysPrompt, userPrompt)
		if err == nil {
			break
		}
		time.Sleep(time.Duration(1<<i) * time.Second)
	}

	if err != nil {
		return nil, fmt.Errorf("AI provider failed after retries: %w", err)
	}

	// Link IDs
	analysis.AnalysisID = "ai_" + fmt.Sprintf("%d", time.Now().UnixNano())
	analysis.DocumentID = documentID
	analysis.AssetID = assetID
	analysis.CreatedAt = time.Now()
	analysis.UpdatedAt = time.Now()

	version.VersionID = "ver_" + fmt.Sprintf("%d", time.Now().UnixNano())
	version.AnalysisID = analysis.AnalysisID
	version.CreatedAt = time.Now()

	// Persist
	if err := s.repo.SaveAnalysis(ctx, analysis); err != nil {
		return nil, err
	}

	if err := s.repo.SaveVersion(ctx, version); err != nil {
		return nil, err
	}

	return analysis, nil
}

func (s *DefaultAIAnalysisService) GetAnalysis(ctx context.Context, assetID string) (*domain.AIAnalysis, error) {
	return s.repo.GetAnalysis(ctx, assetID)
}

func (s *DefaultAIAnalysisService) GetVersionHistory(ctx context.Context, analysisID string) ([]*domain.AnalysisVersion, error) {
	return s.repo.GetVersionHistory(ctx, analysisID)
}
