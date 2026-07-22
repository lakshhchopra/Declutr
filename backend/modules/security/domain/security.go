package domain

import "time"

// SecuritySeverity specifies severity level of security events
type SecuritySeverity string

const (
	SeverityLow      SecuritySeverity = "LOW"
	SeverityMedium   SecuritySeverity = "MEDIUM"
	SeverityHigh     SecuritySeverity = "HIGH"
	SeverityCritical SecuritySeverity = "CRITICAL"
)

// AuditCategory specifies audit log classification
type AuditCategory string

const (
	AuditAuth       AuditCategory = "AUTH"
	AuditAsset      AuditCategory = "ASSET"
	AuditSharing    AuditCategory = "SHARING"
	AuditWorkflow   AuditCategory = "WORKFLOW"
	AuditAI         AuditCategory = "AI"
	AuditSearch     AuditCategory = "SEARCH"
	AuditBackup     AuditCategory = "BACKUP"
	AuditVersioning AuditCategory = "VERSIONING"
	AuditSettings   AuditCategory = "SETTINGS"
)

// RiskLevel defines risk classification
type RiskLevel string

const (
	RiskLow      RiskLevel = "LOW"
	RiskMedium   RiskLevel = "MEDIUM"
	RiskHigh     RiskLevel = "HIGH"
	RiskCritical RiskLevel = "CRITICAL"
)

// SecurityGrade defines letter grade posture
type SecurityGrade string

const (
	GradeA SecurityGrade = "A"
	GradeB SecurityGrade = "B"
	GradeC SecurityGrade = "C"
	GradeD SecurityGrade = "D"
	GradeF SecurityGrade = "F"
)

// Device model
type Device struct {
	DeviceID    string    `json:"deviceId"`
	VaultID     string    `json:"vaultId"`
	DeviceName  string    `json:"deviceName"`
	Browser     string    `json:"browser"`
	OS          string    `json:"os"`
	Platform    string    `json:"platform"`
	IPAddress   string    `json:"ipAddress"`
	Location    string    `json:"location"`
	FirstSeenAt time.Time `json:"firstSeenAt"`
	LastSeenAt  time.Time `json:"lastSeenAt"`
	IsTrusted   bool      `json:"isTrusted"`
}

// ActiveSession model
type ActiveSession struct {
	SessionID   string    `json:"sessionId"`
	VaultID     string    `json:"vaultId"`
	UserID      string    `json:"userId"`
	DeviceName  string    `json:"deviceName"`
	Browser     string    `json:"browser"`
	IPAddress   string    `json:"ipAddress"`
	Location    string    `json:"location"`
	IsCurrent   bool      `json:"isCurrent"`
	CreatedAt   time.Time `json:"createdAt"`
	LastSeenAt  time.Time `json:"lastSeenAt"`
}

// AuditEvent model
type AuditEvent struct {
	AuditID    string                 `json:"auditId"`
	VaultID    string                 `json:"vaultId"`
	Category   AuditCategory          `json:"category"`
	Action     string                 `json:"action"`
	ActorID    string                 `json:"actorId"`
	ActorName  string                 `json:"actorName"`
	IPAddress  string                 `json:"ipAddress"`
	ResourceID string                 `json:"resourceId,omitempty"`
	Details    map[string]interface{} `json:"details,omitempty"`
	CreatedAt  time.Time              `json:"createdAt"`
}

// RiskSignal detected by risk engine
type RiskSignal struct {
	SignalID    string    `json:"signalId"`
	SignalType  string    `json:"signalType"`
	Description string    `json:"description"`
	Weight      int       `json:"weight"`
	DetectedAt  time.Time `json:"detectedAt"`
}

// RiskAssessment model
type RiskAssessment struct {
	AssessmentID string       `json:"assessmentId"`
	VaultID      string       `json:"vaultId"`
	RiskScore    int          `json:"riskScore"`
	RiskLevel    RiskLevel    `json:"riskLevel"`
	Signals      []RiskSignal `json:"signals"`
	AssessedAt   time.Time    `json:"assessedAt"`
}

// SecurityRecommendation model
type SecurityRecommendation struct {
	RecID       string           `json:"recId"`
	VaultID     string           `json:"vaultId"`
	Category    string           `json:"category"`
	Title       string           `json:"title"`
	Description string           `json:"description"`
	ActionType  string           `json:"actionType"`
	Priority    SecuritySeverity `json:"priority"`
	IsDismissed bool             `json:"isDismissed"`
	CreatedAt   time.Time        `json:"createdAt"`
}

// SecurityScore model
type SecurityScore struct {
	ScoreID      string                 `json:"scoreId"`
	VaultID      string                 `json:"vaultId"`
	Score        int                    `json:"score"`
	Grade        SecurityGrade          `json:"grade"`
	Status       string                 `json:"status"` // HEALTHY, ATTENTION_REQUIRED, CRITICAL
	Factors      map[string]interface{} `json:"factors"`
	CalculatedAt time.Time              `json:"calculatedAt"`
}

// SecurityDashboard complete payload
type SecurityDashboard struct {
	VaultID         string                   `json:"vaultId"`
	Score           *SecurityScore           `json:"score"`
	Risk            *RiskAssessment          `json:"risk"`
	ActiveSessions  []*ActiveSession         `json:"activeSessions"`
	Devices         []*Device                `json:"devices"`
	RecentAudit     []*AuditEvent            `json:"recentAudit"`
	Recommendations []*SecurityRecommendation `json:"recommendations"`
	MFAEnabled      bool                     `json:"mfaEnabled"`
	BackupStatus    string                   `json:"backupStatus"`
	EncryptedVault  bool                     `json:"encryptedVault"`
}

// Request DTOs

type TerminateSessionRequest struct {
	VaultID   string `json:"vaultId"`
	SessionID string `json:"sessionId"` // if empty, terminates all other sessions
}

type TrustDeviceRequest struct {
	VaultID  string `json:"vaultId"`
	DeviceID string `json:"deviceId"`
	Trust    bool   `json:"trust"`
}

type RecordAuditEventRequest struct {
	VaultID    string                 `json:"vaultId"`
	Category   AuditCategory          `json:"category"`
	Action     string                 `json:"action"`
	ActorID    string                 `json:"actorId"`
	ActorName  string                 `json:"actorName"`
	IPAddress  string                 `json:"ipAddress"`
	ResourceID string                 `json:"resourceId,omitempty"`
	Details    map[string]interface{} `json:"details,omitempty"`
}
