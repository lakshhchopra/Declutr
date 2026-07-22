package repository

import (
	"fmt"
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/security/domain"
)

// SecurityRepository defines persistence contract for security posture, audit log, sessions, devices, and risk
type SecurityRepository interface {
	RecordAuditEvent(evt *domain.AuditEvent) error
	ListAuditEvents(vaultID string, category domain.AuditCategory, limit int) ([]*domain.AuditEvent, error)

	SaveDevice(dev *domain.Device) error
	ListDevices(vaultID string) ([]*domain.Device, error)
	SetDeviceTrust(deviceID string, trust bool) error

	CreateSession(session *domain.ActiveSession) error
	ListSessions(vaultID string) ([]*domain.ActiveSession, error)
	TerminateSession(sessionID string) error
	TerminateAllOtherSessions(vaultID string, currentSessionID string) error

	SaveRiskAssessment(risk *domain.RiskAssessment) error
	GetLatestRiskAssessment(vaultID string) (*domain.RiskAssessment, error)

	SaveRecommendations(recs []*domain.SecurityRecommendation) error
	ListRecommendations(vaultID string) ([]*domain.SecurityRecommendation, error)

	SaveSecurityScore(score *domain.SecurityScore) error
	GetSecurityScore(vaultID string) (*domain.SecurityScore, error)

	ClearAllData(vaultID string) error
}

// InMemorySecurityRepository is a thread-safe in-memory store
type InMemorySecurityRepository struct {
	mu              sync.RWMutex
	auditEvents     map[string]*domain.AuditEvent            // auditID -> Event
	devices         map[string]*domain.Device                // deviceID -> Device
	sessions        map[string]*domain.ActiveSession         // sessionID -> Session
	riskAssessments map[string]*domain.RiskAssessment        // vaultID -> RiskAssessment
	recommendations map[string][]*domain.SecurityRecommendation // vaultID -> Recs
	scores          map[string]*domain.SecurityScore         // vaultID -> Score
}

// NewInMemorySecurityRepository creates a new in-memory security repository
func NewInMemorySecurityRepository() *InMemorySecurityRepository {
	return &InMemorySecurityRepository{
		auditEvents:     make(map[string]*domain.AuditEvent),
		devices:         make(map[string]*domain.Device),
		sessions:        make(map[string]*domain.ActiveSession),
		riskAssessments: make(map[string]*domain.RiskAssessment),
		recommendations: make(map[string][]*domain.SecurityRecommendation),
		scores:          make(map[string]*domain.SecurityScore),
	}
}

func (r *InMemorySecurityRepository) RecordAuditEvent(evt *domain.AuditEvent) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.auditEvents[evt.AuditID] = evt
	return nil
}

func (r *InMemorySecurityRepository) ListAuditEvents(vaultID string, category domain.AuditCategory, limit int) ([]*domain.AuditEvent, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []*domain.AuditEvent
	for _, evt := range r.auditEvents {
		if evt.VaultID == vaultID {
			if category == "" || evt.Category == category {
				list = append(list, evt)
			}
		}
	}
	if len(list) == 0 {
		list = defaultSampleAuditEvents(vaultID)
		for _, evt := range list {
			r.auditEvents[evt.AuditID] = evt
		}
	}
	if limit > 0 && len(list) > limit {
		return list[:limit], nil
	}
	return list, nil
}

func (r *InMemorySecurityRepository) SaveDevice(dev *domain.Device) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.devices[dev.DeviceID] = dev
	return nil
}

func (r *InMemorySecurityRepository) ListDevices(vaultID string) ([]*domain.Device, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []*domain.Device
	for _, d := range r.devices {
		if d.VaultID == vaultID {
			list = append(list, d)
		}
	}
	if len(list) == 0 {
		list = defaultSampleDevices(vaultID)
		for _, d := range list {
			r.devices[d.DeviceID] = d
		}
	}
	return list, nil
}

func (r *InMemorySecurityRepository) SetDeviceTrust(deviceID string, trust bool) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	dev, ok := r.devices[deviceID]
	if !ok {
		return fmt.Errorf("device %s not found", deviceID)
	}
	dev.IsTrusted = trust
	return nil
}

func (r *InMemorySecurityRepository) CreateSession(session *domain.ActiveSession) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.sessions[session.SessionID] = session
	return nil
}

func (r *InMemorySecurityRepository) ListSessions(vaultID string) ([]*domain.ActiveSession, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var list []*domain.ActiveSession
	for _, s := range r.sessions {
		if s.VaultID == vaultID {
			list = append(list, s)
		}
	}
	if len(list) == 0 {
		list = defaultSampleSessions(vaultID)
		for _, s := range list {
			r.sessions[s.SessionID] = s
		}
	}
	return list, nil
}

func (r *InMemorySecurityRepository) TerminateSession(sessionID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.sessions, sessionID)
	return nil
}

func (r *InMemorySecurityRepository) TerminateAllOtherSessions(vaultID string, currentSessionID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for id, s := range r.sessions {
		if s.VaultID == vaultID && id != currentSessionID {
			delete(r.sessions, id)
		}
	}
	return nil
}

func (r *InMemorySecurityRepository) SaveRiskAssessment(risk *domain.RiskAssessment) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.riskAssessments[risk.VaultID] = risk
	return nil
}

func (r *InMemorySecurityRepository) GetLatestRiskAssessment(vaultID string) (*domain.RiskAssessment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	risk, ok := r.riskAssessments[vaultID]
	if !ok {
		return &domain.RiskAssessment{
			AssessmentID: "risk-default-001",
			VaultID:      vaultID,
			RiskScore:    12,
			RiskLevel:    domain.RiskLow,
			Signals: []domain.RiskSignal{
				{SignalID: "sig-1", SignalType: "NEW_DEVICE", Description: "New browser login detected from Chrome macOS", Weight: 12, DetectedAt: time.Now().Add(-2 * time.Hour)},
			},
			AssessedAt: time.Now(),
		}, nil
	}
	return risk, nil
}

func (r *InMemorySecurityRepository) SaveRecommendations(recs []*domain.SecurityRecommendation) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(recs) > 0 {
		r.recommendations[recs[0].VaultID] = recs
	}
	return nil
}

func (r *InMemorySecurityRepository) ListRecommendations(vaultID string) ([]*domain.SecurityRecommendation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	list, ok := r.recommendations[vaultID]
	if !ok || len(list) == 0 {
		list = defaultSampleRecommendations(vaultID)
		r.recommendations[vaultID] = list
	}
	return list, nil
}

func (r *InMemorySecurityRepository) SaveSecurityScore(score *domain.SecurityScore) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.scores[score.VaultID] = score
	return nil
}

func (r *InMemorySecurityRepository) GetSecurityScore(vaultID string) (*domain.SecurityScore, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	score, ok := r.scores[vaultID]
	if !ok {
		return &domain.SecurityScore{
			ScoreID:      "score-001",
			VaultID:      vaultID,
			Score:        92,
			Grade:        domain.GradeA,
			Status:       "HEALTHY",
			Factors:      map[string]interface{}{"mfa": false, "encrypted": true, "backups": true},
			CalculatedAt: time.Now(),
		}, nil
	}
	return score, nil
}

func (r *InMemorySecurityRepository) ClearAllData(vaultID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.riskAssessments, vaultID)
	delete(r.recommendations, vaultID)
	delete(r.scores, vaultID)
	for id, evt := range r.auditEvents {
		if evt.VaultID == vaultID {
			delete(r.auditEvents, id)
		}
	}
	for id, d := range r.devices {
		if d.VaultID == vaultID {
			delete(r.devices, id)
		}
	}
	for id, s := range r.sessions {
		if s.VaultID == vaultID {
			delete(r.sessions, id)
		}
	}
	return nil
}

// Sample Data Generators
func defaultSampleAuditEvents(vaultID string) []*domain.AuditEvent {
	now := time.Now()
	return []*domain.AuditEvent{
		{
			AuditID:    "audit-1",
			VaultID:    vaultID,
			Category:   domain.AuditAuth,
			Action:     "USER_LOGIN_SUCCESS",
			ActorID:    "usr-owner",
			ActorName:  "Vault Owner",
			IPAddress:  "192.168.1.45",
			Details:    map[string]interface{}{"browser": "Chrome 125.0"},
			CreatedAt:  now.Add(-10 * time.Minute),
		},
		{
			AuditID:    "audit-2",
			VaultID:    vaultID,
			Category:   domain.AuditAsset,
			Action:     "ASSET_UPLOAD",
			ActorID:    "usr-owner",
			ActorName:  "Vault Owner",
			IPAddress:  "192.168.1.45",
			ResourceID: "asset-passport-001",
			Details:    map[string]interface{}{"filename": "Japanese_Visa.pdf"},
			CreatedAt:  now.Add(-2 * time.Hour),
		},
		{
			AuditID:    "audit-3",
			VaultID:    vaultID,
			Category:   domain.AuditSharing,
			Action:     "SHARE_CREATED",
			ActorID:    "usr-owner",
			ActorName:  "Vault Owner",
			IPAddress:  "192.168.1.45",
			ResourceID: "share-japan-001",
			Details:    map[string]interface{}{"title": "Japan Trip Photos"},
			CreatedAt:  now.Add(-24 * time.Hour),
		},
	}
}

func defaultSampleDevices(vaultID string) []*domain.Device {
	now := time.Now()
	return []*domain.Device{
		{
			DeviceID:    "dev-macbook-pro",
			VaultID:    vaultID,
			DeviceName:  "MacBook Pro 16\"",
			Browser:     "Chrome 125.0",
			OS:          "macOS Sonoma",
			Platform:    "WEB",
			IPAddress:   "192.168.1.45",
			Location:    "Tokyo, Japan",
			FirstSeenAt: now.Add(-30 * 24 * time.Hour),
			LastSeenAt:  now.Add(-5 * time.Minute),
			IsTrusted:   true,
		},
		{
			DeviceID:    "dev-iphone-15",
			VaultID:    vaultID,
			DeviceName:  "iPhone 15 Pro",
			Browser:     "Declutr Mobile 1.0",
			OS:          "iOS 17.5",
			Platform:    "MOBILE",
			IPAddress:   "192.168.1.88",
			Location:    "Tokyo, Japan",
			FirstSeenAt: now.Add(-14 * 24 * time.Hour),
			LastSeenAt:  now.Add(-1 * time.Hour),
			IsTrusted:   true,
		},
	}
}

func defaultSampleSessions(vaultID string) []*domain.ActiveSession {
	now := time.Now()
	return []*domain.ActiveSession{
		{
			SessionID:  "sess-mac-current",
			VaultID:    vaultID,
			UserID:     "usr-owner",
			DeviceName: "MacBook Pro 16\"",
			Browser:    "Chrome 125.0",
			IPAddress:  "192.168.1.45",
			Location:   "Tokyo, Japan",
			IsCurrent:  true,
			CreatedAt:  now.Add(-2 * time.Hour),
			LastSeenAt: now.Add(-1 * time.Minute),
		},
		{
			SessionID:  "sess-iphone-mobile",
			VaultID:    vaultID,
			UserID:     "usr-owner",
			DeviceName: "iPhone 15 Pro",
			Browser:    "Declutr Mobile",
			IPAddress:  "192.168.1.88",
			Location:   "Tokyo, Japan",
			IsCurrent:  false,
			CreatedAt:  now.Add(-14 * 24 * time.Hour),
			LastSeenAt: now.Add(-1 * time.Hour),
		},
	}
}

func defaultSampleRecommendations(vaultID string) []*domain.SecurityRecommendation {
	now := time.Now()
	return []*domain.SecurityRecommendation{
		{
			RecID:       "rec-mfa-001",
			VaultID:     vaultID,
			Category:    "AUTHENTICATION",
			Title:       "Enable Multi-Factor Authentication (MFA)",
			Description: "Add WebAuthn or TOTP authenticator app protection to prevent unauthorized access.",
			ActionType:  "ENABLE_MFA",
			Priority:    domain.SeverityHigh,
			IsDismissed: false,
			CreatedAt:   now,
		},
		{
			RecID:       "rec-share-review-002",
			VaultID:     vaultID,
			Category:    "COLLABORATION",
			Title:       "Review Active Public Share Links",
			Description: "You have 1 public password-protected link active. Verify expiration date.",
			ActionType:  "REVIEW_SHARES",
			Priority:    domain.SeverityMedium,
			IsDismissed: false,
			CreatedAt:   now,
		},
	}
}
