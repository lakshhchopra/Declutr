package domain

import "time"

// SearchStrategy represents individual retrieval strategies
type SearchStrategy string

const (
	StrategyKeyword     SearchStrategy = "KEYWORD"
	StrategyVector      SearchStrategy = "VECTOR"
	StrategyEntity      SearchStrategy = "ENTITY"
	StrategyContext     SearchStrategy = "CONTEXT"
	StrategyRelationship SearchStrategy = "RELATIONSHIP"
	StrategyMemory      SearchStrategy = "MEMORY"
	StrategyMetadata    SearchStrategy = "METADATA"
	StrategyHybrid      SearchStrategy = "HYBRID"
)

// SearchFilters specifies parameters for narrowing retrieval results
type SearchFilters struct {
	FileTypes          []string   `json:"fileTypes,omitempty"`
	DateFrom           *time.Time `json:"dateFrom,omitempty"`
	DateTo             *time.Time `json:"dateTo,omitempty"`
	Tags               []string   `json:"tags,omitempty"`
	Collections        []string   `json:"collections,omitempty"`
	Contexts           []string   `json:"contexts,omitempty"`
	Entities           []string   `json:"entities,omitempty"`
	People             []string   `json:"people,omitempty"`
	Locations          []string   `json:"locations,omitempty"`
	MinMemoryStrength  float64    `json:"minMemoryStrength,omitempty"`
	MinConfidence      float64    `json:"minConfidence,omitempty"`
}

// RankingWeights defines the weight configuration for hybrid score fusion
type RankingWeights struct {
	Keyword      float64 `json:"keyword"`
	Vector       float64 `json:"vector"`
	Entity       float64 `json:"entity"`
	Context      float64 `json:"context"`
	Relationship float64 `json:"relationship"`
	Memory       float64 `json:"memory"`
	Recency      float64 `json:"recency"`
}

// DefaultRankingWeights provides balanced fusion weights
func DefaultRankingWeights() RankingWeights {
	return RankingWeights{
		Keyword:      0.25,
		Vector:       0.25,
		Entity:       0.15,
		Context:      0.15,
		Relationship: 0.05,
		Memory:       0.10,
		Recency:      0.05,
	}
}

// ParsedQuery contains the structured intelligence extracted from the raw user query text
type ParsedQuery struct {
	RawText           string     `json:"rawText"`
	CleanedText       string     `json:"cleanedText"`
	DetectedIntent    string     `json:"detectedIntent,omitempty"`
	DetectedEntities  []string   `json:"detectedEntities,omitempty"`
	DetectedLocations []string   `json:"detectedLocations,omitempty"`
	DetectedFileTypes []string   `json:"detectedFileTypes,omitempty"`
	DetectedDateFrom  *time.Time `json:"detectedDateFrom,omitempty"`
	DetectedDateTo    *time.Time `json:"detectedDateTo,omitempty"`
	QuotedTerms       []string   `json:"quotedTerms,omitempty"`
	ExcludedTerms     []string   `json:"excludedTerms,omitempty"`
}

// SearchPlan outlines which retrieval strategies to execute for a query
type SearchPlan struct {
	SelectedStrategies []SearchStrategy  `json:"selectedStrategies"`
	Weights            RankingWeights    `json:"weights"`
	Reasoning          string            `json:"reasoning"`
}

// SearchQueryRequest carries search input from the client
type SearchQueryRequest struct {
	VaultID        string          `json:"vaultId"`
	QueryText      string          `json:"queryText"`
	Page           int             `json:"page"`
	PageSize       int             `json:"pageSize"`
	Filters        SearchFilters   `json:"filters"`
	Weights        *RankingWeights `json:"weights,omitempty"`
	ForceStrategy  SearchStrategy  `json:"forceStrategy,omitempty"`
}

// SearchResultItem represents an enriched matched item with full explainability
type SearchResultItem struct {
	AssetID                string                 `json:"assetId"`
	VaultID                string                 `json:"vaultId"`
	Title                  string                 `json:"title"`
	Summary                string                 `json:"summary"`
	ContentSnippet         string                 `json:"contentSnippet"`
	AssetType              string                 `json:"assetType"`
	Score                  float64                `json:"score"`           // Composite fused score [0, 1]
	Confidence             float64                `json:"confidence"`      // Machine confidence
	WhyMatched             string                 `json:"whyMatched"`     // Clear human-readable match rationale
	ContributingStrategies []SearchStrategy       `json:"contributingStrategies"`
	MatchedEntities        []string               `json:"matchedEntities,omitempty"`
	MatchedContexts        []string               `json:"matchedContexts,omitempty"`
	RelatedMemories        []string               `json:"relatedMemories,omitempty"`
	HighlightedText        string                 `json:"highlightedText,omitempty"`
	CreatedAt              time.Time              `json:"createdAt"`
	Metadata               map[string]interface{} `json:"metadata,omitempty"`
}

// SearchQueryResponse contains the final ranked search results and diagnostic metadata
type SearchQueryResponse struct {
	Results        []*SearchResultItem `json:"results"`
	Total          int                 `json:"total"`
	Page           int                 `json:"page"`
	PageSize       int                 `json:"pageSize"`
	LatencyMs      int64               `json:"latencyMs"`
	ParsedQuery    *ParsedQuery        `json:"parsedQuery"`
	SearchPlan     *SearchPlan         `json:"searchPlan"`
	Suggestions    []string            `json:"suggestions,omitempty"`
}

// SavedSearch represents a bookmarked search with filters
type SavedSearch struct {
	SavedID    string        `json:"savedId"`
	VaultID    string        `json:"vaultId"`
	SearchName string        `json:"searchName"`
	QueryText  string        `json:"queryText"`
	Filters    SearchFilters `json:"filters"`
	IsPinned   bool          `json:"isPinned"`
	CreatedAt  time.Time     `json:"createdAt"`
	UpdatedAt  time.Time     `json:"updatedAt"`
}

// SearchHistoryItem tracks query execution log entries
type SearchHistoryItem struct {
	HistoryID   string       `json:"historyId"`
	VaultID     string       `json:"vaultId"`
	QueryText   string       `json:"queryText"`
	ParsedQuery *ParsedQuery `json:"parsedQuery"`
	ResultCount int          `json:"resultCount"`
	LatencyMs   int64        `json:"latencyMs"`
	SearchType  string       `json:"searchType"`
	SearchedAt  time.Time    `json:"searchedAt"`
}

// SearchStats holds aggregate statistics for search execution
type SearchStats struct {
	VaultID       string            `json:"vaultId"`
	TotalSearches int               `json:"totalSearches"`
	TopQueries    []string          `json:"topQueries"`
	AvgLatencyMs  float64           `json:"avgLatencyMs"`
	StrategyUsage map[string]int    `json:"strategyUsage"`
	UpdatedAt     time.Time         `json:"updatedAt"`
}

// SearchPreferences holds user preferences for search behavior
type SearchPreferences struct {
	PreferenceID          string         `json:"preferenceId"`
	VaultID               string         `json:"vaultId"`
	RankingWeights        RankingWeights `json:"rankingWeights"`
	EnableAutoSuggestions bool           `json:"enableAutoSuggestions"`
	MaxResultsPerPage     int            `json:"maxResultsPerPage"`
	CreatedAt             time.Time      `json:"createdAt"`
	UpdatedAt             time.Time      `json:"updatedAt"`
}
