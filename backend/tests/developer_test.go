package tests

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	devApp "github.com/diablovocado/declutr/modules/developer/application"
	"github.com/diablovocado/declutr/modules/developer/domain"
	devRepo "github.com/diablovocado/declutr/modules/developer/repository"
	"github.com/diablovocado/declutr/sdks/go"
)

func TestAPIKeyGenerationAndValidation(t *testing.T) {
	repo := devRepo.NewInMemoryDeveloperRepository()
	service := devApp.NewDeveloperService(repo)
	ctx := context.Background()

	// 1. Generate API Key
	key, rawSecret, err := service.GenerateAPIKey(ctx, "usr-dev-1", domain.CreateAPIKeyRequest{
		Name:      "Test Script Key",
		Scopes:    []string{domain.ScopeVaultRead, domain.ScopeAssetRead},
		ExpiresIn: 30,
	})
	if err != nil || key == nil {
		t.Fatalf("Failed to generate API key: %v", err)
	}

	if rawSecret == "" || key.KeyPrefix == "" {
		t.Error("Generated secret or prefix is empty")
	}

	// 2. Validate Key with Valid Scope
	validKey, ok := service.ValidateAPIKey(ctx, rawSecret, domain.ScopeVaultRead)
	if !ok || validKey == nil {
		t.Errorf("Expected valid API key for scope vault.read, got ok=%v", ok)
	}

	// 3. Validate Key with Unauthorized Scope
	_, invalidScopeOk := service.ValidateAPIKey(ctx, rawSecret, domain.ScopeAdminManage)
	if invalidScopeOk {
		t.Error("Expected scope validation failure for admin.manage")
	}
}

func TestOAuth21TokenExchange(t *testing.T) {
	repo := devRepo.NewInMemoryDeveloperRepository()
	service := devApp.NewDeveloperService(repo)
	ctx := context.Background()

	// Register OAuth Client App
	client, secret, err := service.CreateOAuthApp(ctx, "usr-dev-1", "Mobile App Client", []string{"https://app/callback"}, []string{"vault.read"})
	if err != nil || client == nil {
		t.Fatalf("Failed to create OAuth client app: %v", err)
	}

	// Exchange Code for Access Token
	token, err := service.ExchangeOAuthToken(ctx, client.ClientID, secret)
	if err != nil || token == nil {
		t.Fatalf("Failed to exchange OAuth token: %v", err)
	}

	if token.TokenType != "Bearer" || token.AccessToken == "" {
		t.Errorf("Invalid OAuth token payload: %v", token)
	}
}

func TestWebhookHMACSigningAndDLQ(t *testing.T) {
	repo := devRepo.NewInMemoryDeveloperRepository()
	service := devApp.NewDeveloperService(repo)
	ctx := context.Background()

	// Register Webhook
	hook, err := service.RegisterWebhook(ctx, "usr-dev-1", domain.CreateWebhookRequest{
		URL:    "http://invalid.local.test/webhook",
		Events: []string{domain.EventAssetUploaded},
	})
	if err != nil || hook == nil {
		t.Fatalf("Failed to register webhook: %v", err)
	}

	// Test HMAC Signature Computation
	sig := service.ComputeWebhookSignature("secret-123", []byte(`{"event_type":"asset.uploaded"}`))
	if sig == "" || len(sig) < 10 {
		t.Errorf("Invalid HMAC signature: %s", sig)
	}

	// Trigger Event Dispatch (will fail to HTTP post invalid URL -> land in DLQ)
	err = service.DispatchEvent(ctx, "usr-dev-1", domain.EventAssetUploaded, map[string]string{"asset_id": "ast-999"})
	if err != nil {
		t.Errorf("DispatchEvent returned error: %v", err)
	}
}

func TestGoSDKClientExecution(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"results": [{"id": "ast-1", "name": "Doc.pdf"}]}`))
	}))
	defer mockServer.Close()

	client := declutr.NewClient(declutr.Config{
		APIKey:  "declutr_live_testkey123",
		BaseURL: mockServer.URL,
	})

	res, err := client.Search(context.Background(), "pdf document", nil)
	if err != nil {
		t.Fatalf("Go SDK search call failed: %v", err)
	}

	if res == nil || res["results"] == nil {
		t.Errorf("Expected search results from SDK client, got: %v", res)
	}
}
