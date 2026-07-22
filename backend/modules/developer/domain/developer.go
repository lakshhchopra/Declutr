package domain

import (
	"time"
)

// API Scopes
const (
	ScopeVaultRead       = "vault.read"
	ScopeVaultWrite      = "vault.write"
	ScopeAssetRead       = "asset.read"
	ScopeAssetWrite      = "asset.write"
	ScopeWorkflowExecute = "workflow.execute"
	ScopeAIChat          = "ai.chat"
	ScopeSearchQuery     = "search.query"
	ScopeBackupManage    = "backup.manage"
	ScopeAdminManage     = "admin.manage"
)

// Webhook Event Types
const (
	EventAssetUploaded      = "asset.uploaded"
	EventAssetUpdated       = "asset.updated"
	EventContextCreated     = "context.created"
	EventWorkflowFinished   = "workflow.finished"
	EventBackupCompleted    = "backup.completed"
	EventSearchCompleted    = "search.completed"
	EventMemoryCreated      = "memory.created"
	EventRelationshipAdded  = "relationship.added"
	EventUserInvited        = "user.invited"
	EventOrganizationCreated = "organization.created"
)

// APIKey represents a developer scoped API key.
type APIKey struct {
	ID             string    `json:"id"`
	OrganizationID string    `json:"organization_id,omitempty"`
	UserID         string    `json:"user_id"`
	Name           string    `json:"name"`
	KeyPrefix      string    `json:"key_prefix"` // declutr_live_...
	KeyHash        string    `json:"-"`          // SHA-256 hash
	Scopes         []string  `json:"scopes"`
	ExpiresAt      time.Time `json:"expires_at"`
	LastUsedAt     time.Time `json:"last_used_at"`
	CreatedAt      time.Time `json:"created_at"`
}

// OAuthClient represents a registered OAuth 2.1 application.
type OAuthClient struct {
	ID             string   `json:"id"`
	OrganizationID string   `json:"organization_id,omitempty"`
	Name           string   `json:"name"`
	ClientID       string   `json:"client_id"`
	ClientSecret   string   `json:"-"` // Hashed
	RedirectURIs   []string `json:"redirect_uris"`
	Scopes         []string `json:"scopes"`
	CreatedAt      time.Time `json:"created_at"`
}

// OAuthToken represents an issued OAuth access token.
type OAuthToken struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type"`
	ExpiresIn    int       `json:"expires_in"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Scope        string    `json:"scope"`
	UserID       string    `json:"user_id"`
	IssuedAt     time.Time `json:"issued_at"`
}

// WebhookEndpoint represents a registered webhook URL.
type WebhookEndpoint struct {
	ID             string   `json:"id"`
	OrganizationID string   `json:"organization_id,omitempty"`
	UserID         string   `json:"user_id"`
	URL            string   `json:"url"`
	Secret         string   `json:"secret"` // HMAC signing secret
	Events         []string `json:"events"`
	IsEnabled      bool     `json:"is_enabled"`
	CreatedAt      time.Time `json:"created_at"`
}

// WebhookDelivery tracks individual webhook execution delivery logs.
type WebhookDelivery struct {
	ID                string    `json:"id"`
	WebhookID         string    `json:"webhook_id"`
	EventType         string    `json:"event_type"`
	Payload           string    `json:"payload"`
	ResponseStatusCode int      `json:"response_status_code"`
	ResponseBody      string    `json:"response_body,omitempty"`
	LatencyMs         int64     `json:"latency_ms"`
	Attempt           int       `json:"attempt"`
	Success           bool      `json:"success"`
	DeliveredAt       time.Time `json:"delivered_at"`
}

// WebhookDLQItem represents a failed webhook item in the Dead Letter Queue.
type WebhookDLQItem struct {
	ID          string    `json:"id"`
	WebhookID   string    `json:"webhook_id"`
	EventType   string    `json:"event_type"`
	Payload     string    `json:"payload"`
	LastError   string    `json:"last_error"`
	Attempts    int       `json:"attempts"`
	FailedAt    time.Time `json:"failed_at"`
}

// DeveloperApp represents a developer's application metadata.
type DeveloperApp struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	HomepageURL string    `json:"homepage_url"`
	CreatedAt   time.Time `json:"created_at"`
}

// CreateAPIKeyRequest payload.
type CreateAPIKeyRequest struct {
	Name      string   `json:"name"`
	Scopes    []string `json:"scopes"`
	ExpiresIn int      `json:"expires_in_days"` // Days
}

// CreateWebhookRequest payload.
type CreateWebhookRequest struct {
	URL    string   `json:"url"`
	Events []string `json:"events"`
}
