package tests

import (
	"context"
	"testing"

	predApp "github.com/diablovocado/declutr/modules/predictive/application"
	predDomain "github.com/diablovocado/declutr/modules/predictive/domain"
	predRepo "github.com/diablovocado/declutr/modules/predictive/repository"
)

func TestPredictiveEngineAndRecommendationPlanner(t *testing.T) {
	repo := predRepo.NewInMemoryPredictiveRepository()
	service := predApp.NewPredictiveService(repo)
	ctx := context.Background()

	// 1. Generate Predictions
	preds, err := service.GenerateAndGetPredictions(ctx, "usr-test-1")
	if err != nil || len(preds) == 0 {
		t.Fatalf("Failed to generate predictions: %v", err)
	}

	if len(preds) < 3 {
		t.Errorf("Expected at least 3 proactive predictions, got %d", len(preds))
	}

	targetPredID := preds[0].ID

	// 2. Accept Prediction
	err = service.AcceptPrediction(ctx, "usr-test-1", targetPredID)
	if err != nil {
		t.Fatalf("Failed to accept prediction: %v", err)
	}

	// 3. Verify Stats Updated
	stats, err := service.GetStats(ctx, "usr-test-1")
	if err != nil || stats.AcceptedCount != 1 {
		t.Errorf("Expected 1 accepted prediction in stats, got %v", stats)
	}

	// 4. Test Confidence Threshold Filtering
	planner := predApp.NewRecommendationPlanner()
	highConfSettings := &predDomain.PredictionSettings{
		UserID:        "usr-test-1",
		MinConfidence: 0.95,
	}

	filtered := planner.RankAndFilterRecommendations(ctx, preds, highConfSettings)
	for _, p := range filtered {
		if p.Confidence < 0.95 {
			t.Errorf("Expected filtered predictions to have confidence >= 0.95, got %f", p.Confidence)
		}
	}
}

func TestPredictiveDismissalFeedbackLearning(t *testing.T) {
	repo := predRepo.NewInMemoryPredictiveRepository()
	service := predApp.NewPredictiveService(repo)
	ctx := context.Background()

	preds, _ := service.GenerateAndGetPredictions(ctx, "usr-test-2")
	targetID := preds[0].ID

	err := service.DismissPrediction(ctx, "usr-test-2", targetID, "Not relevant to my trip")
	if err != nil {
		t.Fatalf("Failed to dismiss prediction: %v", err)
	}

	stats, _ := service.GetStats(ctx, "usr-test-2")
	if stats.DismissedCount != 1 {
		t.Errorf("Expected 1 dismissed prediction, got %d", stats.DismissedCount)
	}
}
