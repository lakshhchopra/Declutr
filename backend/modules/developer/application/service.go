package application

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/diablovocado/declutr/modules/developer/domain"
	"github.com/diablovocado/declutr/modules/developer/repository"
	"github.com/diablovocado/declutr/shared/observability"
)

// DeveloperService manages API keys, OAuth 2.1 apps, webhooks, HMAC signing, and DLQ.
type DeveloperService struct {
	repo repository.DeveloperRepository
}

func NewDeveloperService(repo repository.DeveloperRepository) *DeveloperService {
	return &DeveloperService{repo: repo}
}

// GenerateAPIKey creates a new scoped API key returning full raw secret once to client.
func (s *DeveloperService) GenerateAPIKey(ctx context.Context, userID string, req domain.CreateAPIKeyRequest) (*domain.APIKey, string, error) {
	rawSecret := "declutr_live_" + observability.GenerateID(24)
	prefix := rawSecret[:17] + "..."

	hash := s.HashKey(rawSecret)

	expDays := req.ExpiresIn
	if expDays <= 0 {
		expDays = 365
	}
	exp := time.Now().UTC().AddDate(0, 0, expDays)

	key := &domain.APIKey{
		ID:         "key-" + observability.GenerateID(8),
		UserID:     userID,
		Name:       req.Name,
		KeyPrefix:  prefix,
		KeyHash:    hash,
		Scopes:     req.Scopes,
		ExpiresAt:  exp,
		CreatedAt:  time.Now().UTC(),
		LastUsedAt: time.Now().UTC(),
	}

	if err := s.repo.CreateAPIKey(ctx, key); err != nil {
		return nil, "", err
	}
	return key, rawSecret, nil
}

func (s *DeveloperService) HashKey(raw string) string {
	h := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(h[:])
}

func (s *DeveloperService) ValidateAPIKey(ctx context.Context, rawSecret string, requiredScope string) (*domain.APIKey, bool) {
	hash := s.HashKey(rawSecret)
	key, err := s.repo.GetAPIKeyByHash(ctx, hash)
	if err != nil || key == nil {
		return nil, false
	}

	if time.Now().After(key.ExpiresAt) {
		return nil, false
	}

	if requiredScope != "" {
		hasScope := false
		for _, sc := range key.Scopes {
			if sc == requiredScope || sc == domain.ScopeAdminManage {
				hasScope = true
				break
			}
		}
		if !hasScope {
			return nil, false
		}
	}

	return key, true
}

func (s *DeveloperService) ListAPIKeys(ctx context.Context, userID string) ([]domain.APIKey, error) {
	return s.repo.ListAPIKeys(ctx, userID)
}

func (s *DeveloperService) RevokeAPIKey(ctx context.Context, keyID string) error {
	return s.repo.DeleteAPIKey(ctx, keyID)
}

// OAuth 2.1 App Registration & Token Exchange
func (s *DeveloperService) CreateOAuthApp(ctx context.Context, userID string, name string, redirectURIs []string, scopes []string) (*domain.OAuthClient, string, error) {
	clientID := "client_" + observability.GenerateID(12)
	clientSecret := "secret_" + observability.GenerateID(24)

	client := &domain.OAuthClient{
		ID:           "app-" + observability.GenerateID(8),
		Name:         name,
		ClientID:     clientID,
		ClientSecret: s.HashKey(clientSecret),
		RedirectURIs: redirectURIs,
		Scopes:       scopes,
		CreatedAt:    time.Now().UTC(),
	}

	if err := s.repo.CreateOAuthClient(ctx, client); err != nil {
		return nil, "", err
	}
	return client, clientSecret, nil
}

func (s *DeveloperService) ExchangeOAuthToken(ctx context.Context, clientID string, clientSecret string) (*domain.OAuthToken, error) {
	client, err := s.repo.GetOAuthClientByClientID(ctx, clientID)
	if err != nil {
		return nil, fmt.Errorf("invalid client credentials")
	}

	if client.ClientSecret != s.HashKey(clientSecret) {
		return nil, fmt.Errorf("invalid client secret")
	}

	return &domain.OAuthToken{
		AccessToken:  "declutr_at_" + observability.GenerateID(32),
		TokenType:    "Bearer",
		ExpiresIn:    3600,
		RefreshToken: "declutr_rt_" + observability.GenerateID(32),
		Scope:        "vault.read asset.read search.query",
		IssuedAt:     time.Now().UTC(),
	}, nil
}

// Webhook Engine
func (s *DeveloperService) RegisterWebhook(ctx context.Context, userID string, req domain.CreateWebhookRequest) (*domain.WebhookEndpoint, error) {
	secret := "whsec_" + observability.GenerateID(24)

	hook := &domain.WebhookEndpoint{
		ID:        "wh-" + observability.GenerateID(8),
		UserID:    userID,
		URL:       req.URL,
		Secret:    secret,
		Events:    req.Events,
		IsEnabled: true,
		CreatedAt: time.Now().UTC(),
	}

	if err := s.repo.CreateWebhook(ctx, hook); err != nil {
		return nil, err
	}
	return hook, nil
}

func (s *DeveloperService) ListWebhooks(ctx context.Context, userID string) ([]domain.WebhookEndpoint, error) {
	return s.repo.ListWebhooks(ctx, userID)
}

func (s *DeveloperService) ComputeWebhookSignature(secret string, payload []byte) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(payload)
	return "sha256=" + hex.EncodeToString(mac.Sum(nil))
}

// DispatchEvent sends webhooks to subscribed endpoints with HMAC signing, retries, and DLQ.
func (s *DeveloperService) DispatchEvent(ctx context.Context, userID string, eventType string, eventData interface{}) error {
	webhooks, err := s.repo.ListWebhooks(ctx, userID)
	if err != nil {
		return err
	}

	payloadBytes, err := json.Marshal(map[string]interface{}{
		"event_type": eventType,
		"timestamp":  time.Now().UTC().Format(time.RFC3339),
		"data":       eventData,
	})
	if err != nil {
		return err
	}

	for _, hook := range webhooks {
		if !hook.IsEnabled {
			continue
		}

		subbed := false
		for _, e := range hook.Events {
			if e == eventType || e == "*" {
				subbed = true
				break
			}
		}
		if !subbed {
			continue
		}

		go s.deliverWebhookWithRetries(context.Background(), hook, eventType, payloadBytes)
	}

	return nil
}

func (s *DeveloperService) deliverWebhookWithRetries(ctx context.Context, hook domain.WebhookEndpoint, eventType string, payload []byte) {
	signature := s.ComputeWebhookSignature(hook.Secret, payload)

	maxAttempts := 3
	var lastErr string

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		start := time.Now()

		req, err := http.NewRequestWithContext(ctx, http.MethodPost, hook.URL, bytes.NewBuffer(payload))
		if err != nil {
			lastErr = err.Error()
			continue
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-Declutr-Signature", signature)
		req.Header.Set("X-Declutr-Event", eventType)

		client := &http.Client{Timeout: 5 * time.Second}
		resp, err := client.Do(req)

		latency := time.Since(start).Milliseconds()

		if err == nil && resp.StatusCode >= 200 && resp.StatusCode < 300 {
			// Success delivery
			_ = s.repo.RecordWebhookDelivery(ctx, &domain.WebhookDelivery{
				ID:                 "del-" + observability.GenerateID(8),
				WebhookID:          hook.ID,
				EventType:          eventType,
				Payload:            string(payload),
				ResponseStatusCode: resp.StatusCode,
				LatencyMs:          latency,
				Attempt:            attempt,
				Success:            true,
				DeliveredAt:        time.Now().UTC(),
			})
			return
		}

		if err != nil {
			lastErr = err.Error()
		} else {
			lastErr = fmt.Sprintf("HTTP %d", resp.StatusCode)
		}

		time.Sleep(time.Duration(attempt) * 500 * time.Millisecond)
	}

	// Failed all retries -> Add to Dead Letter Queue (DLQ)
	_ = s.repo.AddToDLQ(ctx, &domain.WebhookDLQItem{
		ID:        "dlq-" + observability.GenerateID(8),
		WebhookID: hook.ID,
		EventType: eventType,
		Payload:   string(payload),
		LastError: lastErr,
		Attempts:  maxAttempts,
		FailedAt:  time.Now().UTC(),
	})
}

func (s *DeveloperService) GetDeliveries(ctx context.Context, webhookID string) ([]domain.WebhookDelivery, error) {
	return s.repo.ListWebhookDeliveries(ctx, webhookID)
}

func (s *DeveloperService) GetDLQ(ctx context.Context) ([]domain.WebhookDLQItem, error) {
	return s.repo.ListDLQ(ctx)
}
