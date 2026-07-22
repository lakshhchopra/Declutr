package application

import (
	"context"
	"fmt"
	"log"

	"github.com/diablovocado/declutr/modules/search/domain"
)

// HybridSearchEngine orchestrates unified knowledge retrieval across all vault sources
type HybridSearchEngine struct {
	service *SearchService
}

// NewHybridSearchEngine creates a new HybridSearchEngine
func NewHybridSearchEngine(service *SearchService) *HybridSearchEngine {
	return &HybridSearchEngine{service: service}
}

// Search executes a hybrid search request combining keyword, vector, entity, context, and memory strategies
func (e *HybridSearchEngine) Search(ctx context.Context, req *domain.SearchQueryRequest) (*domain.SearchQueryResponse, error) {
	log.Printf("[HybridSearchEngine] Executing hybrid search query: %q for vault: %s", req.QueryText, req.VaultID)

	resp, err := e.service.ExecuteSearch(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("hybrid search engine: query failed: %w", err)
	}

	log.Printf("[HybridSearchEngine] Search complete: %d results found in %dms (Plan: %s)",
		resp.Total, resp.LatencyMs, resp.SearchPlan.Reasoning)

	return resp, nil
}
