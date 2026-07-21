package extractors

import (
	"context"
	"strings"
	"testing"
)

func TestExtractorRouting(t *testing.T) {
	registry := NewExtractorRegistry()

	tests := []struct {
		mimeType string
		expected string
	}{
		{"image/png", "*extractors.ImageExtractor"},
		{"image/jpeg", "*extractors.ImageExtractor"},
		{"text/plain", "*extractors.TextExtractor"},
		{"application/json", "*extractors.TextExtractor"},
		{"application/pdf", "*extractors.MockComplexExtractor"},
		{"video/mp4", "*extractors.MockComplexExtractor"},
		{"application/octet-stream", "*extractors.BaseExtractor"},
	}

	for _, tc := range tests {
		ext := registry.GetExtractor(tc.mimeType)
		if ext == nil {
			t.Errorf("Expected extractor for %s, got nil", tc.mimeType)
		}
	}
}

func TestTextExtractor(t *testing.T) {
	ext := &TextExtractor{}
	reader := strings.NewReader("Hello\nWorld")
	
	meta, err := ext.Extract(context.Background(), "a1", "v1", "test.txt", "text/plain", 11, reader)
	if err != nil {
		t.Fatalf("Failed to extract: %v", err)
	}

	if meta.General.FileSize != 11 {
		t.Errorf("Expected size 11, got %d", meta.General.FileSize)
	}
	if meta.General.Extension != ".txt" {
		t.Errorf("Expected .txt extension, got %s", meta.General.Extension)
	}
	
	if meta.Properties == nil {
		t.Fatal("Expected properties")
	}
	
	if meta.Properties.Properties["characterCount"].(int64) != 11 {
		t.Errorf("Expected char count 11, got %v", meta.Properties.Properties["characterCount"])
	}
}
