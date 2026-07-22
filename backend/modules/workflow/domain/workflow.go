package domain

import "time"

// TriggerType specifies what event initiates workflow execution
type TriggerType string

const (
	TriggerAssetUploaded   TriggerType = "ASSET_UPLOADED"
	TriggerAssetUpdated    TriggerType = "ASSET_UPDATED"
	TriggerAssetDeleted    TriggerType = "ASSET_DELETED"
	TriggerContextCreated  TriggerType = "CONTEXT_CREATED"
	TriggerMemoryCreated   TriggerType = "MEMORY_CREATED"
	TriggerEntityFound     TriggerType = "ENTITY_FOUND"
	TriggerRelCreated      TriggerType = "RELATIONSHIP_CREATED"
	TriggerDocExpiring     TriggerType = "DOCUMENT_EXPIRING"
	TriggerDailySchedule   TriggerType = "DAILY_SCHEDULE"
	TriggerManual          TriggerType = "MANUAL_TRIGGER"
	TriggerInsightCreated  TriggerType = "AI_INSIGHT_CREATED"
	TriggerTimelineEvent   TriggerType = "TIMELINE_EVENT"
)

// ConditionOperator defines evaluation operators
type ConditionOperator string

const (
	OpEquals      ConditionOperator = "EQUALS"
	OpContains    ConditionOperator = "CONTAINS"
	OpGreaterThan ConditionOperator = "GREATER_THAN"
	OpLessThan    ConditionOperator = "LESS_THAN"
	OpIn          ConditionOperator = "IN"
)

// ConditionCombinator defines logical rules (AND, OR, NOT)
type ConditionCombinator string

const (
	CombinatorAND ConditionCombinator = "AND"
	CombinatorOR  ConditionCombinator = "OR"
	CombinatorNOT ConditionCombinator = "NOT"
)

// ActionType defines executable actions
type ActionType string

const (
	ActionCreateCollection ActionType = "CREATE_COLLECTION"
	ActionMoveAsset        ActionType = "MOVE_ASSET"
	ActionApplyTags        ActionType = "APPLY_TAGS"
	ActionGenerateSummary  ActionType = "GENERATE_SUMMARY"
	ActionArchiveAsset     ActionType = "ARCHIVE_ASSET"
	ActionCreateReminder   ActionType = "CREATE_REMINDER"
	ActionPinMemory        ActionType = "PIN_MEMORY"
	ActionRefreshSearch    ActionType = "REFRESH_SEARCH_INDEX"
	ActionNotifyUser       ActionType = "NOTIFY_USER"
)

// RunStatus represents the state of a workflow execution run
type RunStatus string

const (
	RunPending   RunStatus = "PENDING"
	RunRunning   RunStatus = "RUNNING"
	RunSuccess   RunStatus = "SUCCESS"
	RunFailed    RunStatus = "FAILED"
	RunRetrying  RunStatus = "RETRYING"
)

// WorkflowTrigger model
type WorkflowTrigger struct {
	TriggerID   string                 `json:"triggerId"`
	WorkflowID  string                 `json:"workflowId"`
	TriggerType TriggerType            `json:"triggerType"`
	Config      map[string]interface{} `json:"config,omitempty"`
	CreatedAt   time.Time              `json:"createdAt"`
}

// WorkflowCondition model
type WorkflowCondition struct {
	ConditionID string              `json:"conditionId"`
	WorkflowID  string              `json:"workflowId"`
	Field       string              `json:"field"`
	Operator    ConditionOperator   `json:"operator"`
	Value       string              `json:"value"`
	Combinator  ConditionCombinator `json:"combinator"`
	CreatedAt   time.Time           `json:"createdAt"`
}

// WorkflowAction model
type WorkflowAction struct {
	ActionID       string                 `json:"actionId"`
	WorkflowID     string                 `json:"workflowId"`
	ActionType     ActionType             `json:"actionType"`
	Config         map[string]interface{} `json:"config,omitempty"`
	ExecutionOrder int                    `json:"executionOrder"`
	CreatedAt      time.Time              `json:"createdAt"`
}

// Workflow model
type Workflow struct {
	WorkflowID    string               `json:"workflowId"`
	VaultID       string               `json:"vaultId"`
	Name          string               `json:"name"`
	Description   string               `json:"description"`
	Enabled       bool                 `json:"enabled"`
	Status        string               `json:"status"` // IDLE, RUNNING, ERROR
	Triggers      []WorkflowTrigger    `json:"triggers"`
	Conditions    []WorkflowCondition  `json:"conditions"`
	Actions       []WorkflowAction     `json:"actions"`
	LastRunAt     *time.Time           `json:"lastRunAt,omitempty"`
	RunCount      int                  `json:"runCount"`
	FailureCount  int                  `json:"failureCount"`
	AvgDurationMs int                  `json:"avgDurationMs"`
	CreatedBy     string               `json:"createdBy"`
	CreatedAt     time.Time            `json:"createdAt"`
	UpdatedAt     time.Time            `json:"updatedAt"`
}

// WorkflowRun record
type WorkflowRun struct {
	RunID        string    `json:"runId"`
	WorkflowID   string    `json:"workflowId"`
	VaultID      string    `json:"vaultId"`
	TriggerEvent string    `json:"triggerEvent"`
	Status       RunStatus `json:"status"`
	DurationMs   int64     `json:"durationMs"`
	ErrorMessage string    `json:"errorMessage,omitempty"`
	StartedAt    time.Time `json:"startedAt"`
	CompletedAt  *time.Time`json:"completedAt,omitempty"`
}

// WorkflowLog entry
type WorkflowLog struct {
	LogID      string    `json:"logId"`
	RunID      string    `json:"runId"`
	WorkflowID string    `json:"workflowId"`
	StepType   string    `json:"stepType"`
	Message    string    `json:"message"`
	Status     string    `json:"status"` // INFO, SUCCESS, ERROR
	LoggedAt   time.Time `json:"loggedAt"`
}

// WorkflowStats metrics
type WorkflowStats struct {
	VaultID       string  `json:"vaultId"`
	TotalWorkflows int    `json:"totalWorkflows"`
	ActiveWorkflows int   `json:"activeWorkflows"`
	TotalRuns     int     `json:"totalRuns"`
	SuccessfulRuns int    `json:"successfulRuns"`
	FailedRuns    int     `json:"failedRuns"`
	SuccessRate   float64 `json:"successRate"`
	AvgDurationMs float64 `json:"avgDurationMs"`
}

// RunWorkflowRequest payload
type RunWorkflowRequest struct {
	WorkflowID string                 `json:"workflowId"`
	VaultID    string                 `json:"vaultId"`
	Context    map[string]interface{} `json:"context,omitempty"`
}

// ToggleWorkflowRequest payload
type ToggleWorkflowRequest struct {
	WorkflowID string `json:"workflowId"`
	Enabled    bool   `json:"enabled"`
}
