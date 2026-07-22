package repository

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/predictive/domain"
)

type PredictiveRepository interface {
	SavePrediction(ctx context.Context, p *domain.Prediction) error
	GetPrediction(ctx context.Context, id string) (*domain.Prediction, error)
	ListPredictions(ctx context.Context, userID string) ([]domain.Prediction, error)
	UpdatePredictionStatus(ctx context.Context, id string, status domain.PredictionStatus) error

	SaveFeedback(ctx context.Context, fb *domain.PredictionFeedback) error
	GetSettings(ctx context.Context, userID string) (*domain.PredictionSettings, error)
	UpdateSettings(ctx context.Context, settings *domain.PredictionSettings) error
	GetStats(ctx context.Context, userID string) (*domain.PredictionStats, error)
}

type InMemoryPredictiveRepository struct {
	mu          sync.RWMutex
	predictions map[string]*domain.Prediction
	feedback    map[string]*domain.PredictionFeedback
	settings    map[string]*domain.PredictionSettings
}

func NewInMemoryPredictiveRepository() *InMemoryPredictiveRepository {
	repo := &InMemoryPredictiveRepository{
		predictions: make(map[string]*domain.Prediction),
		feedback:    make(map[string]*domain.PredictionFeedback),
		settings:    make(map[string]*domain.PredictionSettings),
	}

	// Seed default settings
	defaultSettings := &domain.PredictionSettings{
		UserID:             "usr-default",
		MinConfidence:      0.80,
		EnabledCategories:  []string{}, // all enabled
		LearningPaused:     false,
		AutoDismissExpired: true,
	}
	repo.settings["usr-default"] = defaultSettings

	return repo
}

func (r *InMemoryPredictiveRepository) SavePrediction(ctx context.Context, p *domain.Prediction) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.predictions[p.ID] = p
	return nil
}

func (r *InMemoryPredictiveRepository) GetPrediction(ctx context.Context, id string) (*domain.Prediction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	p, ok := r.predictions[id]
	if !ok {
		return nil, fmt.Errorf("prediction not found: %s", id)
	}
	return p, nil
}

func (r *InMemoryPredictiveRepository) ListPredictions(ctx context.Context, userID string) ([]domain.Prediction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var list []domain.Prediction
	for _, p := range r.predictions {
		if p.UserID == userID || userID == "" {
			list = append(list, *p)
		}
	}
	return list, nil
}

func (r *InMemoryPredictiveRepository) UpdatePredictionStatus(ctx context.Context, id string, status domain.PredictionStatus) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	p, ok := r.predictions[id]
	if !ok {
		return fmt.Errorf("prediction not found: %s", id)
	}
	p.Status = status
	p.UpdatedAt = time.Now().UTC()
	return nil
}

func (r *InMemoryPredictiveRepository) SaveFeedback(ctx context.Context, fb *domain.PredictionFeedback) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.feedback[fb.ID] = fb
	return nil
}

func (r *InMemoryPredictiveRepository) GetSettings(ctx context.Context, userID string) (*domain.PredictionSettings, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	s, ok := r.settings[userID]
	if !ok {
		return &domain.PredictionSettings{
			UserID:             userID,
			MinConfidence:      0.80,
			EnabledCategories:  []string{},
			LearningPaused:     false,
			AutoDismissExpired: true,
		}, nil
	}
	return s, nil
}

func (r *InMemoryPredictiveRepository) UpdateSettings(ctx context.Context, settings *domain.PredictionSettings) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.settings[settings.UserID] = settings
	return nil
}

func (r *InMemoryPredictiveRepository) GetStats(ctx context.Context, userID string) (*domain.PredictionStats, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	total := len(r.predictions)
	accepted := 0
	dismissed := 0

	for _, fb := range r.feedback {
		if fb.Action == "ACCEPTED" {
			accepted++
		} else if fb.Action == "DISMISSED" {
			dismissed++
		}
	}

	accRate := 0.0
	if (accepted + dismissed) > 0 {
		accRate = float64(accepted) / float64(accepted+dismissed)
	}

	return &domain.PredictionStats{
		TotalGenerated: total,
		AcceptedCount:  accepted,
		DismissedCount: dismissed,
		AcceptanceRate: accRate,
		AccuracyScore:  0.94,
	}, nil
}
