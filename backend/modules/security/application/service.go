package application

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"

	"github.com/diablovocado/declutr/modules/security/domain"
	"github.com/diablovocado/declutr/modules/security/repository"
)

// SecurityCenterService manages vault security posture, audit log hub, risk assessments, sessions, and device trust
type SecurityCenterService struct {
	repo repository.SecurityRepository
}

// NewSecurityCenterService creates a new SecurityCenterService instance
func NewSecurityCenterService(repo repository.SecurityRepository) *SecurityCenterService {
	return &SecurityCenterService{repo: repo}
}

// GetDashboard aggregates security score, risk level, active sessions, devices, audit events, and recommendations
func (s *SecurityCenterService) GetDashboard(vaultID string) (*domain.SecurityDashboard, error) {
	if vaultID == "" {
		return nil, fmt.Errorf("security: vaultId is required")
	}

	score, err := s.repo.GetSecurityScore(vaultID)
	if err != nil {
		return nil, err
	}
	risk, _ := s.repo.GetLatestRiskAssessment(vaultID)
	sessions, _ := s.repo.ListSessions(vaultID)
	devices, _ := s.repo.ListDevices(vaultID)
	recentAudit, _ := s.repo.ListAuditEvents(vaultID, "", 10)
	recs, _ := s.repo.ListRecommendations(vaultID)

	return &domain.SecurityDashboard{
		VaultID:         vaultID,
		Score:           score,
		Risk:            risk,
		ActiveSessions:  sessions,
		Devices:         devices,
		RecentAudit:     recentAudit,
		Recommendations: recs,
		MFAEnabled:      false,
		BackupStatus:    "HEALTHY (Weekly Auto)",
		EncryptedVault:  true,
	}, nil
}

// RecordAuditEvent logs an asynchronous audit event across any vault module
func (s *SecurityCenterService) RecordAuditEvent(ctx context.Context, req *domain.RecordAuditEventRequest) error {
	if req.VaultID == "" || req.Action == "" {
		return fmt.Errorf("security: vaultId and action are required")
	}

	evt := &domain.AuditEvent{
		AuditID:    "audit-" + uuid.New().String()[:8],
		VaultID:    req.VaultID,
		Category:   req.Category,
		Action:     req.Action,
		ActorID:    req.ActorID,
		ActorName:  req.ActorName,
		IPAddress:  req.IPAddress,
		ResourceID: req.ResourceID,
		Details:    req.Details,
		CreatedAt:  time.Now(),
	}

	if err := s.repo.RecordAuditEvent(evt); err != nil {
		return err
	}
	log.Printf("[AuditEngine] Audit Event logged: %s (%s) by %s", req.Action, req.Category, req.ActorName)
	return nil
}

// ListAuditEvents retrieves filtered audit log entries
func (s *SecurityCenterService) ListAuditEvents(vaultID string, category domain.AuditCategory, limit int) ([]*domain.AuditEvent, error) {
	if vaultID == "" {
		return nil, fmt.Errorf("security: vaultId is required")
	}
	return s.repo.ListAuditEvents(vaultID, category, limit)
}

// ListSessions returns active user sessions
func (s *SecurityCenterService) ListSessions(vaultID string) ([]*domain.ActiveSession, error) {
	if vaultID == "" {
		return nil, fmt.Errorf("security: vaultId is required")
	}
	return s.repo.ListSessions(vaultID)
}

// TerminateSession terminates a single or all other active user sessions
func (s *SecurityCenterService) TerminateSession(req *domain.TerminateSessionRequest) error {
	if req.VaultID == "" {
		return fmt.Errorf("security: vaultId is required")
	}
	if req.SessionID != "" {
		return s.repo.TerminateSession(req.SessionID)
	}
	return s.repo.TerminateAllOtherSessions(req.VaultID, "sess-mac-current")
}

// ListDevices returns registered user devices
func (s *SecurityCenterService) ListDevices(vaultID string) ([]*domain.Device, error) {
	if vaultID == "" {
		return nil, fmt.Errorf("security: vaultId is required")
	}
	return s.repo.ListDevices(vaultID)
}

// SetDeviceTrust toggles device trust status
func (s *SecurityCenterService) SetDeviceTrust(req *domain.TrustDeviceRequest) error {
	if req.DeviceID == "" {
		return fmt.Errorf("security: deviceId is required")
	}
	return s.repo.SetDeviceTrust(req.DeviceID, req.Trust)
}

// GetRiskAssessment calculates or returns the latest risk engine assessment
func (s *SecurityCenterService) GetRiskAssessment(vaultID string) (*domain.RiskAssessment, error) {
	if vaultID == "" {
		return nil, fmt.Errorf("security: vaultId is required")
	}
	return s.repo.GetLatestRiskAssessment(vaultID)
}

// GetRecommendations returns actionable security posture recommendations
func (s *SecurityCenterService) GetRecommendations(vaultID string) ([]*domain.SecurityRecommendation, error) {
	if vaultID == "" {
		return nil, fmt.Errorf("security: vaultId is required")
	}
	return s.repo.ListRecommendations(vaultID)
}
