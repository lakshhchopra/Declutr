package domain

import (
	"time"
)

// Default Life Areas
const (
	AreaPersonal  = "Personal"
	AreaWork      = "Work"
	AreaBusiness  = "Business"
	AreaEducation = "Education"
	AreaFinance   = "Finance"
	AreaHealth    = "Health"
	AreaTravel    = "Travel"
	AreaLegal     = "Legal"
	AreaHome      = "Home"
	AreaFamily    = "Family"
	AreaResearch  = "Research"
	AreaHobbies   = "Hobbies"
)

// LifeArea represents a top-level life domain.
type LifeArea struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Color       string    `json:"color"`
	IsCustom    bool      `json:"is_custom"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// ProjectStatus indicates project state.
type ProjectStatus string

const (
	ProjectPlanning   ProjectStatus = "PLANNING"
	ProjectInProgress ProjectStatus = "IN_PROGRESS"
	ProjectBlocked    ProjectStatus = "BLOCKED"
	ProjectCompleted  ProjectStatus = "COMPLETED"
	ProjectArchived   ProjectStatus = "ARCHIVED"
)

// Project represents a first-class user objective hub.
type Project struct {
	ID             string        `json:"id"`
	UserID         string        `json:"user_id"`
	LifeAreaID     string        `json:"life_area_id"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Status         ProjectStatus `json:"status"`
	Budget         float64       `json:"budget,omitempty"`
	AssociatedDocs []string      `json:"associated_docs"`
	People         []string      `json:"people"`
	TargetDate     time.Time     `json:"target_date,omitempty"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

// ProjectGoal represents a milestone goal belonging to a project.
type ProjectGoal struct {
	ID            string    `json:"id"`
	ProjectID     string    `json:"project_id"`
	UserID        string    `json:"user_id"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	ProgressPct   int       `json:"progress_pct"` // 0 to 100
	IsCompleted   bool      `json:"is_completed"`
	MissingAssets []string  `json:"missing_assets,omitempty"`
	DueDate       time.Time `json:"due_date,omitempty"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

// LifeTimelineEvent represents an aggregated life event entry.
type LifeTimelineEvent struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	ProjectID   string    `json:"project_id,omitempty"`
	Title       string    `json:"title"`
	EventType   string    `json:"event_type"` // DOCUMENT, MEMORY, MEETING, WORKFLOW, TRIP
	Description string    `json:"description"`
	Timestamp   time.Time `json:"timestamp"`
}

// LifeDashboard represents unified home view data.
type LifeDashboard struct {
	LifeAreas       []LifeArea          `json:"life_areas"`
	ActiveProjects  []Project           `json:"active_projects"`
	ActiveGoals     []ProjectGoal       `json:"active_goals"`
	PrioritiesToday []string            `json:"priorities_today"`
	RecentTimeline  []LifeTimelineEvent `json:"recent_timeline"`
	HealthScore     int                 `json:"health_score"` // 0 to 100
}

// LifeMetric represents life balance telemetry.
type LifeMetric struct {
	UserID         string  `json:"user_id"`
	ActiveProjects int     `json:"active_projects"`
	GoalCompletion float64 `json:"goal_completion"`
	LifeBalance    int     `json:"life_balance"`
}
