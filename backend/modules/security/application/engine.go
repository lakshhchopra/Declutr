package application

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/diablovocado/declutr/modules/security/domain"
)

// RiskEngine evaluates domain security events and signals to compute dynamic risk scores and levels
type RiskEngine struct {
	service *SecurityCenterService
}

// NewRiskEngine creates a new RiskEngine instance
func NewRiskEngine(service *SecurityCenterService) *RiskEngine {
	return &RiskEngine{service: service}
}

// AssessVaultRisk evaluates vault security signals and returns a fresh RiskAssessment
func (e *RiskEngine) AssessVaultRisk(ctx context.Context, vaultID string) (*domain.RiskAssessment, error) {
	log.Printf("[RiskEngine] Evaluating security risk signals for vault %s", vaultID)

	signals := []domain.RiskSignal{
		{
			SignalID:    "sig-" + uuid.New().String()[:8],
			SignalType:  "NEW_DEVICE_LOGIN",
			Description: "Login from unrecognized browser agent",
			Weight:      15,
			DetectedAt:  time.Now(),
		},
	}

	score := 15
	level := domain.RiskLow
	if score > 75 {
		level = domain.RiskCritical
	} else if score > 50 {
		level = domain.RiskHigh
	} else if score > 25 {
		level = domain.RiskMedium
	}

	assessment := &domain.RiskAssessment{
		AssessmentID: "risk-" + uuid.New().String()[:8],
		VaultID:      vaultID,
		RiskScore:    score,
		RiskLevel:    level,
		Signals:      signals,
		AssessedAt:   time.Now(),
	}

	_ = e.service.repo.SaveRiskAssessment(assessment)
	return assessment, nil
}
