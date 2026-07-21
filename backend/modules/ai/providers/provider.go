package providers

import (
	"context"

	"github.com/diablovocado/declutr/modules/ai/domain"
)

type LLMProvider interface {
	Analyze(ctx context.Context, systemPrompt, userPrompt string) (*domain.AIAnalysis, *domain.AnalysisVersion, error)
}

// MockProvider generates deterministic output for local dev without requiring an API key.
type MockProvider struct{}

func NewMockProvider() *MockProvider {
	return &MockProvider{}
}

func (p *MockProvider) Analyze(ctx context.Context, systemPrompt, userPrompt string) (*domain.AIAnalysis, *domain.AnalysisVersion, error) {
	analysis := &domain.AIAnalysis{
		Title:                "Extracted Project Notes",
		ShortSummary:         "A brief overview of project tasks.",
		DetailedSummary:      "This document outlines the Alpha project requirements, including security and performance constraints.",
		Language:             "en",
		WritingStyle:         "Technical",
		Sentiment:            "Neutral",
		Complexity:           "Medium",
		ReadingLevel:         "College",
		EstimatedReadingTime: 60,
		DocumentPurpose:      "Project Planning",
		ConfidenceScore:      0.95,
		Classification: domain.Classification{
			DocumentCategory: "General Note",
			DocumentType:     "Markdown Document",
			IsScanned:        false,
			IsCorrupted:      false,
			IsIncomplete:     false,
			QualityScore:     0.99,
			ConfidenceScore:  0.98,
		},
		Tags: []domain.Tag{
			{Name: "Project", ConfidenceScore: 0.9},
			{Name: "Planning", ConfidenceScore: 0.85},
		},
		Topics: []domain.Topic{
			{Name: "Software Engineering", ConfidenceScore: 0.99},
		},
	}

	version := &domain.AnalysisVersion{
		Provider:      "mock",
		ModelName:     "mock-deterministic-v1",
		PromptVersion: "1.0.0",
		TokenUsage: domain.TokenUsage{
			PromptTokens:     150,
			CompletionTokens: 350,
			TotalTokens:      500,
			EstimatedCostUSD: 0.001,
		},
		LatencyMs: 250,
		RawOutput: map[string]interface{}{"status": "mocked"},
	}

	return analysis, version, nil
}

// Skeleton for OpenAI
type OpenAIProvider struct {
	apiKey string
}

func NewOpenAIProvider(apiKey string) *OpenAIProvider {
	return &OpenAIProvider{apiKey: apiKey}
}

func (p *OpenAIProvider) Analyze(ctx context.Context, systemPrompt, userPrompt string) (*domain.AIAnalysis, *domain.AnalysisVersion, error) {
	// 1. Build OpenAI HTTP Request (or use go-openai)
	// 2. Specify ResponseFormat: JSON_SCHEMA
	// 3. Unmarshal structured output into domain.AIAnalysis
	// 4. Calculate token cost
	return nil, nil, nil
}
