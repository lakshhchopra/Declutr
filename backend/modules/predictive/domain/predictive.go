package domain

import (
	"time"
)

// PredictionType defines the category of predictive intelligence.
type PredictionType string

const (
	PredUpcomingDeadline    PredictionType = "UPCOMING_DEADLINE"
	PredExpiringDocument    PredictionType = "EXPIRING_DOCUMENT"
	PredUpcomingTrip        PredictionType = "UPCOMING_TRIP"
	PredUpcomingMeeting     PredictionType = "UPCOMING_MEETING"
	PredMissingDocument     PredictionType = "MISSING_DOCUMENT"
	PredSuggestedUpload     PredictionType = "SUGGESTED_UPLOAD"
	PredSuggestedOrg        PredictionType = "SUGGESTED_ORGANIZATION"
	PredSuggestedArchive    PredictionType = "SUGGESTED_ARCHIVE"
	PredSuggestedDeletion   PredictionType = "SUGGESTED_DELETION"
	PredSuggestedWorkflow   PredictionType = "SUGGESTED_WORKFLOW"
	PredSuggestedCollection PredictionType = "SUGGESTED_COLLECTION"
	PredSuggestedSummary    PredictionType = "SUGGESTED_SUMMARY"
	PredRecurringTask       PredictionType = "RECURRING_TASK"
	PredRecurringExpense    PredictionType = "RECURRING_EXPENSE"
	PredKnowledgeGap        PredictionType = "KNOWLEDGE_GAP"
	PredOpportunityDetect   PredictionType = "OPPORTUNITY_DETECTION"
)

// PredictionPriority indicates urgency level.
type PredictionPriority string

const (
	PriorityHigh   PredictionPriority = "HIGH"
	PriorityMedium PredictionPriority = "MEDIUM"
	PriorityLow    PredictionPriority = "LOW"
)

// PredictionStatus indicates user interaction status.
type PredictionStatus string

const (
	StatusPending   PredictionStatus = "PENDING"
	StatusAccepted  PredictionStatus = "ACCEPTED"
	StatusDismissed PredictionStatus = "DISMISSED"
	StatusExpired   PredictionStatus = "EXPIRED"
)

// PredictionEvidence holds explanation details and rationales.
type PredictionEvidence struct {
	SourceModule string   `json:"source_module"` // MEMORY, TIMELINE, KNOWLEDGE_GRAPH, REVERSE_PERSONA
	Reasoning    string   `json:"reasoning"`
	KeyFacts     []string `json:"key_facts"`
}

// Prediction represents a proactive life intelligence insight model.
type Prediction struct {
	ID               string             `json:"id"`
	UserID           string             `json:"user_id"`
	Type             PredictionType     `json:"type"`
	Title            string             `json:"title"`
	Description      string             `json:"description"`
	Confidence       float64            `json:"confidence"` // 0.0 to 1.0
	Priority         PredictionPriority `json:"priority"`
	Evidence         PredictionEvidence `json:"evidence"`
	AffectedAssets   []string           `json:"affected_assets"`
	SuggestedAction  string             `json:"suggested_action"`
	Status           PredictionStatus   `json:"status"`
	Expiration       time.Time          `json:"expiration"`
	CreatedAt        time.Time          `json:"created_at"`
	UpdatedAt        time.Time          `json:"updated_at"`
}

// PredictionSettings defines user-configurable thresholds and category controls.
type PredictionSettings struct {
	UserID              string   `json:"user_id"`
	MinConfidence       float64  `json:"min_confidence"` // Default 0.80
	EnabledCategories   []string `json:"enabled_categories"`
	LearningPaused      bool     `json:"learning_paused"`
	AutoDismissExpired  bool     `json:"auto_dismiss_expired"`
}

// PredictionFeedback captures user approval/dismissal feedback for ML tuning.
type PredictionFeedback struct {
	ID           string    `json:"id"`
	PredictionID string    `json:"prediction_id"`
	UserID       string    `json:"user_id"`
	Action       string    `json:"action"` // ACCEPTED, DISMISSED
	Reason       string    `json:"reason,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
}

// PredictionStats holds observability metrics.
type PredictionStats struct {
	TotalGenerated int     `json:"total_generated"`
	AcceptedCount  int     `json:"accepted_count"`
	DismissedCount int     `json:"dismissed_count"`
	AcceptanceRate float64 `json:"acceptance_rate"`
	AccuracyScore  float64 `json:"accuracy_score"`
}
