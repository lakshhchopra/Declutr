package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/diablovocado/declutr/modules/developer/domain"
)

// DeveloperRepository defines persistence interface for developer resources.
type DeveloperRepository interface {
	CreateAPIKey(ctx context.Context, key *domain.APIKey) error
	ListAPIKeys(ctx context.Context, userID string) ([]domain.APIKey, error)
	GetAPIKeyByHash(ctx context.Context, hash string) (*domain.APIKey, error)
	DeleteAPIKey(ctx context.Context, keyID string) error

	CreateOAuthClient(ctx context.Context, client *domain.OAuthClient) error
	GetOAuthClientByClientID(ctx context.Context, clientID string) (*domain.OAuthClient, error)

	CreateWebhook(ctx context.Context, hook *domain.WebhookEndpoint) error
	ListWebhooks(ctx context.Context, userID string) ([]domain.WebhookEndpoint, error)

	RecordWebhookDelivery(ctx context.Context, delivery *domain.WebhookDelivery) error
	ListWebhookDeliveries(ctx context.Context, webhookID string) ([]domain.WebhookDelivery, error)

	AddToDLQ(ctx context.Context, item *domain.WebhookDLQItem) error
	ListDLQ(ctx context.Context) ([]domain.WebhookDLQItem, error)
}

// InMemoryDeveloperRepository provides a thread-safe in-memory store.
type InMemoryDeveloperRepository struct {
	mu           sync.RWMutex
	apiKeys      map[string]*domain.APIKey
	oauthClients map[string]*domain.OAuthClient
	webhooks     map[string]*domain.WebhookEndpoint
	deliveries   map[string][]domain.WebhookDelivery
	dlq          map[string]*domain.WebhookDLQItem
}

func NewInMemoryDeveloperRepository() *InMemoryDeveloperRepository {
	return &InMemoryDeveloperRepository{
		apiKeys:      make(map[string]*domain.APIKey),
		oauthClients: make(map[string]*domain.OAuthClient),
		webhooks:     make(map[string]*domain.WebhookEndpoint),
		deliveries:   make(map[string][]domain.WebhookDelivery),
		dlq:          make(map[string]*domain.WebhookDLQItem),
	}
}

func (r *InMemoryDeveloperRepository) CreateAPIKey(ctx context.Context, key *domain.APIKey) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.apiKeys[key.ID] = key
	return nil
}

func (r *InMemoryDeveloperRepository) ListAPIKeys(ctx context.Context, userID string) ([]domain.APIKey, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []domain.APIKey
	for _, k := range r.apiKeys {
		if k.UserID == userID {
			result = append(result, *k)
		}
	}
	return result, nil
}

func (r *InMemoryDeveloperRepository) GetAPIKeyByHash(ctx context.Context, hash string) (*domain.APIKey, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, k := range r.apiKeys {
		if k.KeyHash == hash {
			return k, nil
		}
	}
	return nil, fmt.Errorf("API key not found")
}

func (r *InMemoryDeveloperRepository) DeleteAPIKey(ctx context.Context, keyID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.apiKeys, keyID)
	return nil
}

func (r *InMemoryDeveloperRepository) CreateOAuthClient(ctx context.Context, client *domain.OAuthClient) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.oauthClients[client.ClientID] = client
	return nil
}

func (r *InMemoryDeveloperRepository) GetOAuthClientByClientID(ctx context.Context, clientID string) (*domain.OAuthClient, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	client, ok := r.oauthClients[clientID]
	if !ok {
		return nil, fmt.Errorf("OAuth client %s not found", clientID)
	}
	return client, nil
}

func (r *InMemoryDeveloperRepository) CreateWebhook(ctx context.Context, hook *domain.WebhookEndpoint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.webhooks[hook.ID] = hook
	return nil
}

func (r *InMemoryDeveloperRepository) ListWebhooks(ctx context.Context, userID string) ([]domain.WebhookEndpoint, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []domain.WebhookEndpoint
	for _, w := range r.webhooks {
		if w.UserID == userID {
			result = append(result, *w)
		}
	}
	return result, nil
}

func (r *InMemoryDeveloperRepository) RecordWebhookDelivery(ctx context.Context, delivery *domain.WebhookDelivery) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.deliveries[delivery.WebhookID] = append(r.deliveries[delivery.WebhookID], *delivery)
	return nil
}

func (r *InMemoryDeveloperRepository) ListWebhookDeliveries(ctx context.Context, webhookID string) ([]domain.WebhookDelivery, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.deliveries[webhookID], nil
}

func (r *InMemoryDeveloperRepository) AddToDLQ(ctx context.Context, item *domain.WebhookDLQItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.dlq[item.ID] = item
	return nil
}

func (r *InMemoryDeveloperRepository) ListDLQ(ctx context.Context) ([]domain.WebhookDLQItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result []domain.WebhookDLQItem
	for _, item := range r.dlq {
		result = append(result, *item)
	}
	return result, nil
}
