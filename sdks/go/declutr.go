package declutr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Config struct {
	APIKey      string
	AccessToken string
	BaseURL     string
	HTTPClient  *http.Client
}

type Client struct {
	apiKey      string
	accessToken string
	baseURL     string
	httpClient  *http.Client
}

func NewClient(cfg Config) *Client {
	baseUrl := cfg.BaseURL
	if baseUrl == "" {
		baseUrl = "http://localhost:8080"
	}
	httpClient := cfg.HTTPClient
	if httpClient == nil {
		httpClient = &http.Client{Timeout: 15 * time.Second}
	}
	return &Client{
		apiKey:      cfg.APIKey,
		accessToken: cfg.AccessToken,
		baseURL:     baseUrl,
		httpClient:  httpClient,
	}
}

func (c *Client) request(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		b, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(b)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	} else if c.accessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.accessToken)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		b, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("declutr error [%d]: %s", resp.StatusCode, string(b))
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (c *Client) Search(ctx context.Context, query string, filters map[string]interface{}) (map[string]interface{}, error) {
	var res map[string]interface{}
	err := c.request(ctx, "POST", "/api/v1/search/query", map[string]interface{}{
		"query":   query,
		"filters": filters,
	}, &res)
	return res, err
}

func (c *Client) SendCopilotMessage(ctx context.Context, conversationID string, content string) (map[string]interface{}, error) {
	var res map[string]interface{}
	err := c.request(ctx, "POST", "/api/v1/copilot/messages", map[string]interface{}{
		"conversation_id": conversationID,
		"content":         content,
	}, &res)
	return res, err
}

func (c *Client) ExecuteWorkflow(ctx context.Context, workflowID string, input interface{}) (map[string]interface{}, error) {
	var res map[string]interface{}
	err := c.request(ctx, "POST", "/api/v1/workflows/run", map[string]interface{}{
		"workflow_id": workflowID,
		"input":       input,
	}, &res)
	return res, err
}
