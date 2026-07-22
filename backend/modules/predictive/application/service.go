package application

import (
	"context"
	"time"

	"github.com/diablovocado/declutr/modules/predictive/domain"
	"github.com/diablovocado/declutr/modules/predictive/repository"
	"github.com/diablovocado/declutr/shared/observability"
)

type PredictiveService struct {
	repo    repository.PredictiveRepository
	engine  *PredictiveEngine
	planner *RecommendationPlanner
}

func NewPredictiveService(repo repository.PredictiveRepository) *PredictiveService {
	return &PredictiveService{
		repo:    repo,
		engine:  NewPredictiveEngine(),
		planner: NewRecommendationPlanner(),
	}
}

func (s *PredictiveService) GenerateAndGetPredictions(ctx context.Context, userID string) ([]domain.Prediction, error) {
	settings, err := s.repo.GetSettings(ctx, userID)
	if err != nil {
		return nil, err
	}

	if settings.LearningPaused {
		return s.repo.ListPredictions(ctx, userID)
	}

	// Run pattern analysis engine
	rawPredictions, err := s.engine.AnalyzeUserPatterns(ctx, userID)
	if err != nil {
		return nil, err
	}

	for i := range rawPredictions {
		_ = s.repo.SavePrediction(ctx, &rawPredictions[i])
	}

	// Filter and rank via recommendation planner
	recommendations := s.planner.RankAndFilterRecommendations(ctx, rawPredictions, settings)
	return recommendations, nil
}

func (s *PredictiveService) AcceptPrediction(ctx context.Context, userID string, predictionID string) error {
	err := s.repo.UpdatePredictionStatus(ctx, predictionID, domain.StatusAccepted)
	if err != nil {
		return err
	}

	feedback := &domain.PredictionFeedback{
		ID:           "fb-" + observability.GenerateID(8),
		PredictionID: predictionID,
		UserID:       userID,
		Action:       "ACCEPTED",
		CreatedAt:    time.Now().UTC(),
	}
	return s.repo.SaveFeedback(ctx, feedback)
}

func (s *PredictiveService) DismissPrediction(ctx context.Context, userID string, predictionID string, reason string) error {
	err := s.repo.UpdatePredictionStatus(ctx, predictionID, domain.StatusDismissed)
	if err != nil {
		return err
	}

	feedback := &domain.PredictionFeedback{
		ID:           "fb-" + observability.GenerateID(8),
		PredictionID: predictionID,
		UserID:       userID,
		Action:       "DISMISSED",
		Reason:       reason,
		CreatedAt:    time.Now().UTC(),
	}
	return s.repo.SaveFeedback(ctx, feedback)
}

func (s *PredictiveService) GetSettings(ctx context.Context, userID string) (*domain.PredictionSettings, error) {
	return s.repo.GetSettings(ctx, userID)
}

func (s *PredictiveService) UpdateSettings(ctx context.Context, settings *domain.PredictionSettings) error {
	return s.repo.UpdateSettings(ctx, settings)
}

func (s *PredictiveService) GetStats(ctx context.Context, userID string) (*domain.PredictionStats, error) {
	return s.repo.GetStats(ctx, userID)
}

func (s *PredictiveService) GetHistory(ctx context.Context, userID string) ([]domain.Prediction, error) {
	return s.repo.ListPredictions(ctx, userID)
}
