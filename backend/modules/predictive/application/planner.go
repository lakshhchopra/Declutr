package application

import (
	"context"

	"github.com/diablovocado/declutr/modules/predictive/domain"
)

type RecommendationPlanner struct{}

func NewRecommendationPlanner() *RecommendationPlanner {
	return &RecommendationPlanner{}
}

func (rp *RecommendationPlanner) RankAndFilterRecommendations(
	ctx context.Context,
	predictions []domain.Prediction,
	settings *domain.PredictionSettings,
) []domain.Prediction {
	var filtered []domain.Prediction

	categoryMap := make(map[string]bool)
	for _, cat := range settings.EnabledCategories {
		categoryMap[cat] = true
	}

	for _, p := range predictions {
		// 1. Confidence threshold check
		if p.Confidence < settings.MinConfidence {
			continue
		}

		// 2. Category filter check
		if len(settings.EnabledCategories) > 0 && !categoryMap[string(p.Type)] {
			continue
		}

		// 3. Status filter check
		if p.Status != domain.StatusPending {
			continue
		}

		filtered = append(filtered, p)
	}

	return filtered
}
