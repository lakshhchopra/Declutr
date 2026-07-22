package repository

import (
	"sync"
	"time"

	"github.com/diablovocado/declutr/modules/search/domain"
)

// IndexableItem represents an item in the search index
type IndexableItem struct {
	AssetID         string
	VaultID         string
	Title           string
	Summary         string
	Content         string
	AssetType       string
	Entities        []string
	Contexts        []string
	Memories        []string
	Tags            []string
	MemoryStrength  float64
	Confidence      float64
	CreatedAt       time.Time
}

// SearchRepository defines the persistence contract for search data
type SearchRepository interface {
	// History
	AddHistory(item *domain.SearchHistoryItem) error
	GetHistory(vaultID string, limit int) ([]*domain.SearchHistoryItem, error)

	// Saved Searches
	SaveSearch(ss *domain.SavedSearch) error
	GetSavedSearches(vaultID string) ([]*domain.SavedSearch, error)
	DeleteSavedSearch(savedID string) error

	// Stats
	GetStats(vaultID string) (*domain.SearchStats, error)

	// Preferences
	GetPreferences(vaultID string) (*domain.SearchPreferences, error)
	UpdatePreferences(prefs *domain.SearchPreferences) error

	// Indexable Items (for hybrid search execution)
	AddIndexItem(item *IndexableItem) error
	ListIndexItems(vaultID string) ([]*IndexableItem, error)
}

// InMemorySearchRepository is a thread-safe in-memory implementation
type InMemorySearchRepository struct {
	mu          sync.RWMutex
	history     map[string][]*domain.SearchHistoryItem // vaultID -> history
	saved       map[string]*domain.SavedSearch        // savedID -> saved
	stats       map[string]*domain.SearchStats        // vaultID -> stats
	preferences map[string]*domain.SearchPreferences  // vaultID -> prefs
	indexItems  map[string][]*IndexableItem           // vaultID -> items
}

// NewInMemorySearchRepository creates a new in-memory search repository
func NewInMemorySearchRepository() *InMemorySearchRepository {
	return &InMemorySearchRepository{
		history:     make(map[string][]*domain.SearchHistoryItem),
		saved:       make(map[string]*domain.SavedSearch),
		stats:       make(map[string]*domain.SearchStats),
		preferences: make(map[string]*domain.SearchPreferences),
		indexItems:  make(map[string][]*IndexableItem),
	}
}

// --- History ---

func (r *InMemorySearchRepository) AddHistory(item *domain.SearchHistoryItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.history[item.VaultID] = append([]*domain.SearchHistoryItem{item}, r.history[item.VaultID]...)
	if len(r.history[item.VaultID]) > 50 {
		r.history[item.VaultID] = r.history[item.VaultID][:50]
	}

	// Update stats
	s, ok := r.stats[item.VaultID]
	if !ok {
		s = &domain.SearchStats{
			VaultID:       item.VaultID,
			StrategyUsage: make(map[string]int),
			UpdatedAt:     time.Now(),
		}
		r.stats[item.VaultID] = s
	}
	s.TotalSearches++
	s.StrategyUsage[item.SearchType]++
	s.UpdatedAt = time.Now()

	return nil
}

func (r *InMemorySearchRepository) GetHistory(vaultID string, limit int) ([]*domain.SearchHistoryItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := r.history[vaultID]
	if limit > 0 && len(items) > limit {
		return items[:limit], nil
	}
	return items, nil
}

// --- Saved Searches ---

func (r *InMemorySearchRepository) SaveSearch(ss *domain.SavedSearch) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	ss.UpdatedAt = time.Now()
	r.saved[ss.SavedID] = ss
	return nil
}

func (r *InMemorySearchRepository) GetSavedSearches(vaultID string) ([]*domain.SavedSearch, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []*domain.SavedSearch
	for _, ss := range r.saved {
		if ss.VaultID == vaultID {
			result = append(result, ss)
		}
	}
	return result, nil
}

func (r *InMemorySearchRepository) DeleteSavedSearch(savedID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.saved, savedID)
	return nil
}

// --- Stats ---

func (r *InMemorySearchRepository) GetStats(vaultID string) (*domain.SearchStats, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if s, ok := r.stats[vaultID]; ok {
		return s, nil
	}
	return &domain.SearchStats{
		VaultID:       vaultID,
		TotalSearches: 0,
		TopQueries:    []string{},
		AvgLatencyMs:  12.5,
		StrategyUsage: map[string]int{"HYBRID": 10, "KEYWORD": 4, "VECTOR": 3},
		UpdatedAt:     time.Now(),
	}, nil
}

// --- Preferences ---

func (r *InMemorySearchRepository) GetPreferences(vaultID string) (*domain.SearchPreferences, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if p, ok := r.preferences[vaultID]; ok {
		return p, nil
	}
	return &domain.SearchPreferences{
		PreferenceID:          "pref-default",
		VaultID:               vaultID,
		RankingWeights:        domain.DefaultRankingWeights(),
		EnableAutoSuggestions: true,
		MaxResultsPerPage:     20,
		CreatedAt:             time.Now(),
		UpdatedAt:             time.Now(),
	}, nil
}

func (r *InMemorySearchRepository) UpdatePreferences(prefs *domain.SearchPreferences) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	prefs.UpdatedAt = time.Now()
	r.preferences[prefs.VaultID] = prefs
	return nil
}

// --- Index Items ---

func (r *InMemorySearchRepository) AddIndexItem(item *IndexableItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.indexItems[item.VaultID] = append(r.indexItems[item.VaultID], item)
	return nil
}

func (r *InMemorySearchRepository) ListIndexItems(vaultID string) ([]*IndexableItem, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	items := r.indexItems[vaultID]
	if len(items) == 0 {
		// Populate default sample index items for demo & testing
		return defaultSampleItems(vaultID), nil
	}
	return items, nil
}

func defaultSampleItems(vaultID string) []*IndexableItem {
	now := time.Now()
	return []*IndexableItem{
		{
			AssetID: "asset-passport-001", VaultID: vaultID, Title: "Japanese Visa & Passport Scan",
			Summary: "Passport photo page and Japanese entry visa for Tokyo vacation 2025.",
			Content: "Passport number A987654321, issued by US Department of State. Entry visa for Japan valid for 90 days.",
			AssetType: "PDF", Entities: []string{"Tokyo", "Japan", "Passport"}, Contexts: []string{"Japan Vacation"},
			Memories: []string{"Japan Vacation 2025"}, Tags: []string{"travel", "passport", "visa"},
			MemoryStrength: 0.88, Confidence: 0.95, CreatedAt: now.Add(-60 * 24 * time.Hour),
		},
		{
			AssetID: "asset-thesis-002", VaultID: vaultID, Title: "PhD Thesis Chapter 4 — Neural Networks",
			Summary: "Deep learning models, PyTorch code snippets, and transformer benchmark results.",
			Content: "Chapter 4 evaluates attention mechanisms and graph neural network embeddings on benchmark datasets.",
			AssetType: "DOCX", Entities: []string{"PyTorch", "Neural Networks", "Dr. Sharma"}, Contexts: []string{"PhD Thesis"},
			Memories: []string{"Thesis Chapter 4 — Neural Networks"}, Tags: []string{"research", "ai", "thesis"},
			MemoryStrength: 0.84, Confidence: 0.90, CreatedAt: now.Add(-30 * 24 * time.Hour),
		},
		{
			AssetID: "asset-medical-003", VaultID: vaultID, Title: "Cardiology Consultation Note — Dr. Sharma",
			Summary: "Annual heart check-up report and prescription refill confirmation.",
			Content: "Patient blood pressure 120/80, normal ECG. Prescription for Atorvastatin 20mg renewed.",
			AssetType: "PDF", Entities: []string{"Dr. Sharma", "Cardiology", "Hospital"}, Contexts: []string{"Medical Treatment"},
			Memories: []string{"Recurring Medical Visits — Dr. Sharma"}, Tags: []string{"health", "medical", "prescription"},
			MemoryStrength: 0.79, Confidence: 0.92, CreatedAt: now.Add(-10 * 24 * time.Hour),
		},
		{
			AssetID: "asset-tax-004", VaultID: vaultID, Title: "Tax Return Filing 2024",
			Summary: "IRS Form 1040, W-2 statements, and investment dividend summaries.",
			Content: "Federal income tax return 2024. Total adjusted gross income reported with W-2 tax withholdings.",
			AssetType: "PDF", Entities: []string{"IRS", "Form 1040", "W-2"}, Contexts: []string{"Tax Filing"},
			Memories: []string{"Tax Filing 2024"}, Tags: []string{"finance", "tax", "irs"},
			MemoryStrength: 0.61, Confidence: 0.85, CreatedAt: now.Add(-120 * 24 * time.Hour),
		},
	}
}

func (r *InMemorySearchRepository) Clear(vaultID string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.history, vaultID)
	delete(r.stats, vaultID)
	delete(r.indexItems, vaultID)
	return nil
}
