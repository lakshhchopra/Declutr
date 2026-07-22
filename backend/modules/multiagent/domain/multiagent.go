package domain

import (
	"time"
)

// AgentRole defines the specialized multi-agent role.
type AgentRole string

const (
	RoleCoordinator AgentRole = "COORDINATOR_AGENT"
	RoleKnowledge   AgentRole = "KNOWLEDGE_AGENT"
	RoleMemory      AgentRole = "MEMORY_AGENT"
	RoleResearch    AgentRole = "RESEARCH_AGENT"
	RoleOrganization AgentRole = "ORGANIZATION_AGENT"
	RoleWorkflow    AgentRole = "WORKFLOW_AGENT"
	RoleSearch      AgentRole = "SEARCH_AGENT"
	RoleSecurity    AgentRole = "SECURITY_AGENT"
	RoleIntegration AgentRole = "INTEGRATION_AGENT"
	RoleTimeline    AgentRole = "TIMELINE_AGENT"
	RoleFinancial   AgentRole = "FINANCIAL_AGENT"
	RoleTravel      AgentRole = "TRAVEL_AGENT"
	RoleLearning    AgentRole = "LEARNING_AGENT"
)

// AgentRegistration represents a registered agent in the system registry.
type AgentRegistration struct {
	ID           string    `json:"id"`
	Role         AgentRole `json:"role"`
	Name         string    `json:"name"`
	Capabilities []string  `json:"capabilities"`
	SupportedTools []string `json:"supported_tools"`
	Status       string    `json:"status"` // READY, BUSY, PAUSED, ERROR
	Version      string    `json:"version"`
	Health       string    `json:"health"` // HEALTHY, DEGRADED, DOWN
	Priority     int       `json:"priority"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// AgentMessage represents a structured message exchanged over the Message Bus.
type AgentMessage struct {
	ID            string                 `json:"id"`
	CorrelationID string                 `json:"correlation_id"`
	GoalID        string                 `json:"goal_id"`
	TaskID        string                 `json:"task_id"`
	Sender        string                 `json:"sender"`   // Sender Agent ID
	Receiver      string                 `json:"receiver"` // Receiver Agent ID or "BROADCAST"
	MessageType   string                 `json:"message_type"` // REQUEST, RESPONSE, STATUS, CONSENSUS_PROPOSAL
	Payload       map[string]interface{} `json:"payload"`
	Context       map[string]interface{} `json:"context"`
	Timestamp     time.Time              `json:"timestamp"`
}

// TaskExecutionMode defines execution strategy (parallel vs sequential).
type TaskExecutionMode string

const (
	ExecSequential TaskExecutionMode = "SEQUENTIAL"
	ExecParallel   TaskExecutionMode = "PARALLEL"
)

// CoordinatorTask represents a task node within a multi-agent task graph.
type CoordinatorTask struct {
	ID             string            `json:"id"`
	GoalID         string            `json:"goal_id"`
	AssignedRole   AgentRole         `json:"assigned_role"`
	AssignedAgentID string           `json:"assigned_agent_id"`
	Action         string            `json:"action"`
	Parameters     map[string]interface{} `json:"parameters"`
	Dependencies   []string          `json:"dependencies"`
	ExecutionMode  TaskExecutionMode `json:"execution_mode"`
	Status         string            `json:"status"` // PENDING, RUNNING, COMPLETED, FAILED, AWAITING_APPROVAL
	ResultPayload  map[string]interface{} `json:"result_payload,omitempty"`
	Confidence     float64           `json:"confidence"`
	ErrorMessage   string            `json:"error_message,omitempty"`
}

// TaskGraph represents a DAG of tasks created by the Task Planner.
type TaskGraph struct {
	GoalID      string            `json:"goal_id"`
	GoalTitle   string            `json:"goal_title"`
	Tasks       []CoordinatorTask `json:"tasks"`
	Status      string            `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
}

// SharedMemoryItem represents items stored in the shared multi-agent memory workspace.
type SharedMemoryItem struct {
	ID        string    `json:"id"`
	GoalID    string    `json:"goal_id"`
	AgentID   string    `json:"agent_id"`
	Key       string    `json:"key"`
	Value     interface{} `json:"value"`
	Category  string    `json:"category"` // TASK_RESULT, USER_PREFERENCE, CONTEXT, STATE
	CreatedAt time.Time `json:"created_at"`
}

// ConsensusResult represents conflict resolution output by the Coordinator.
type ConsensusResult struct {
	GoalID             string    `json:"goal_id"`
	TaskID             string    `json:"task_id"`
	ConsensusAchieved  bool      `json:"consensus_achieved"`
	WinningAgentID     string    `json:"winning_agent_id"`
	WinningConfidence  float64   `json:"winning_confidence"`
	Explanation        string    `json:"explanation"`
	EscalateToUser     bool      `json:"escalate_to_user"`
	ResolvedAt         time.Time `json:"resolved_at"`
}

// AgentHealthMetric represents observability telemetry for an agent.
type AgentHealthMetric struct {
	AgentID      string  `json:"agent_id"`
	Role         AgentRole `json:"role"`
	Status       string  `json:"status"`
	LatencyMs    int64   `json:"latency_ms"`
	SuccessRate  float64 `json:"success_rate"`
	TotalTasks   int     `json:"total_tasks"`
	FailedTasks  int     `json:"failed_tasks"`
}
