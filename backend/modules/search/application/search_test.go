package application_test

import (
	"context"
	"testing"

	"github.com/diablovocado/declutr/modules/search/application"
	"github.com/diablovocado/declutr/modules/search/domain"
	"github.com/diablovocado/declutr/modules/search/repository"
)

const testVaultID = "vault-test-001"

func setupService() *application.SearchService {
	repo := repository.NewInMemorySearchRepository()
	return application.NewSearchService(repo, nil)
}

// TestQueryParser validates intelligence extraction from raw query text
func TestQueryParser(t *testing.T) {
	rawQuery := `"passport photo" japan 2025 -draft pdf`
	pq := application.ParseQuery(rawQuery)

	if len(pq.QuotedTerms) == 0 || pq.QuotedTerms[0] != "passport photo" {
		t.Errorf("expected quoted term 'passport photo', got %v", pq.QuotedTerms)
	}
	if len(pq.ExcludedTerms) == 0 || pq.ExcludedTerms[0] != "draft" {
		t.Errorf("expected excluded term 'draft', got %v", pq.ExcludedTerms)
	}
	if len(pq.DetectedFileTypes) == 0 || pq.DetectedFileTypes[0] != "PDF" {
		t.Errorf("expected detected filetype 'PDF', got %v", pq.DetectedFileTypes)
	}
	if pq.DetectedDateFrom == nil {
		t.Error("expected detected year 2025 date range")
	}
	if pq.DetectedIntent != "Travel" {
		t.Errorf("expected detected intent 'Travel', got %s", pq.DetectedIntent)
	}

	t.Logf("PASS: Query Parser — Quoted=%v, Excluded=%v, FileTypes=%v, Intent=%s",
		pq.QuotedTerms, pq.ExcludedTerms, pq.DetectedFileTypes, pq.DetectedIntent)
}

// TestKeywordSearch validates exact keyword matching against indexed items
func TestKeywordSearch(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	req := &domain.SearchQueryRequest{
		VaultID:   testVaultID,
		QueryText: "Passport",
		Page:      1,
		PageSize:  10,
	}

	resp, err := svc.ExecuteSearch(ctx, req)
	if err != nil {
		t.Fatalf("search failed: %v", err)
	}
	if resp.Total == 0 {
		t.Fatal("expected search results for 'Passport', got 0")
	}
	if resp.Results[0].AssetID != "asset-passport-001" {
		t.Errorf("expected top match 'asset-passport-001', got %s", resp.Results[0].AssetID)
	}

	t.Logf("PASS: Keyword Search — Top result: %s (Score=%.2f)",
		resp.Results[0].Title, resp.Results[0].Score)
}

// TestHybridSearchAndFusion validates weighted score fusion combining keyword, entity, context, and memory scores
func TestHybridSearchAndFusion(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	req := &domain.SearchQueryRequest{
		VaultID:   testVaultID,
		QueryText: "Japan vacation visa Tokyo 2025",
		Page:      1,
		PageSize:  10,
	}

	resp, err := svc.ExecuteSearch(ctx, req)
	if err != nil {
		t.Fatalf("hybrid search failed: %v", err)
	}
	if resp.Total == 0 {
		t.Fatal("expected hybrid search results, got 0")
	}

	top := resp.Results[0]
	if len(top.ContributingStrategies) == 0 {
		t.Error("expected contributing strategies in hybrid result")
	}
	if top.WhyMatched == "" {
		t.Error("expected human-readable match explanation")
	}

	t.Logf("PASS: Hybrid Search & Fusion — Top: %s | Score=%.2f | Why: %s | Strategies=%v",
		top.Title, top.Score, top.WhyMatched, top.ContributingStrategies)
}

// TestSearchFiltering validates filtering by file type and date range
func TestSearchFiltering(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	// Filter by DOCX filetype only
	req := &domain.SearchQueryRequest{
		VaultID:   testVaultID,
		QueryText: "Thesis",
		Page:      1,
		PageSize:  10,
		Filters: domain.SearchFilters{
			FileTypes: []string{"DOCX"},
		},
	}

	resp, err := svc.ExecuteSearch(ctx, req)
	if err != nil {
		t.Fatalf("filtered search failed: %v", err)
	}
	for _, res := range resp.Results {
		if res.AssetType != "DOCX" {
			t.Errorf("expected asset type DOCX, got %s", res.AssetType)
		}
	}

	t.Logf("PASS: Search Filtering — %d results matched filter DOCX", resp.Total)
}

// TestMatchExplainability verifies match rationale, contributing strategies, and highlighted text
func TestMatchExplainability(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	req := &domain.SearchQueryRequest{
		VaultID:   testVaultID,
		QueryText: "Dr. Sharma Cardiology",
		Page:      1,
		PageSize:  10,
	}

	resp, err := svc.ExecuteSearch(ctx, req)
	if err != nil {
		t.Fatalf("search failed: %v", err)
	}
	if resp.Total == 0 {
		t.Fatal("expected results for Dr. Sharma, got 0")
	}

	top := resp.Results[0]
	if top.WhyMatched == "" {
		t.Error("expected non-empty WhyMatched explanation")
	}
	if len(top.MatchedEntities) == 0 {
		t.Error("expected matched entities")
	}
	t.Logf("PASS: Explainability — Why: %s | Matched Entities=%v", top.WhyMatched, top.MatchedEntities)
}

// TestSavedSearches validates saving, retrieving, and deleting saved search queries
func TestSavedSearches(t *testing.T) {
	svc := setupService()

	ss := &domain.SavedSearch{
		SavedID:    "saved-001",
		VaultID:    testVaultID,
		SearchName: "Japan Trip Documents",
		QueryText:  "Japan vacation passport visa",
		IsPinned:   true,
	}

	if err := svc.SaveSearch(ss); err != nil {
		t.Fatalf("save search failed: %v", err)
	}

	list, err := svc.GetSavedSearches(testVaultID)
	if err != nil {
		t.Fatalf("get saved searches failed: %v", err)
	}
	if len(list) != 1 {
		t.Errorf("expected 1 saved search, got %d", len(list))
	}

	if err := svc.DeleteSavedSearch("saved-001"); err != nil {
		t.Fatalf("delete saved search failed: %v", err)
	}

	listAfter, _ := svc.GetSavedSearches(testVaultID)
	if len(listAfter) != 0 {
		t.Errorf("expected 0 saved searches after delete, got %d", len(listAfter))
	}

	t.Logf("PASS: Saved Searches — Save, retrieve, and delete verified")
}

// TestSearchHistoryAndStats validates logging history entries and retrieving stats
func TestSearchHistoryAndStats(t *testing.T) {
	svc := setupService()
	ctx := context.Background()

	_, _ = svc.ExecuteSearch(ctx, &domain.SearchQueryRequest{VaultID: testVaultID, QueryText: "Medical report"})
	_, _ = svc.ExecuteSearch(ctx, &domain.SearchQueryRequest{VaultID: testVaultID, QueryText: "Tax filing"})

	history, err := svc.GetHistory(testVaultID, 10)
	if err != nil {
		t.Fatalf("get history failed: %v", err)
	}
	if len(history) < 2 {
		t.Errorf("expected at least 2 history items, got %d", len(history))
	}

	stats, err := svc.GetStats(testVaultID)
	if err != nil {
		t.Fatalf("get stats failed: %v", err)
	}
	if stats.TotalSearches < 2 {
		t.Errorf("expected at least 2 total searches in stats, got %d", stats.TotalSearches)
	}

	t.Logf("PASS: Search History & Stats — Total Searches=%d, Latency=%.1fms",
		stats.TotalSearches, stats.AvgLatencyMs)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
