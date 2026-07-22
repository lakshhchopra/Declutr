import type {
  SearchQueryResponse,
  SavedSearch,
  SearchHistoryItem,
  SearchStats,
  SearchPreferences,
  SearchFilters,
  SearchResultItem,
} from '../types/search';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080/api/v1';

async function apiFetch<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    ...options,
    headers: { 'Content-Type': 'application/json', ...options?.headers },
  });
  if (!res.ok) throw new Error(`Search API error: ${res.status} ${res.statusText}`);
  return res.json();
}

const VAULT_ID = 'vault-demo';

// ─── Mock Data Fallback ────────────────────────────────────────────────────────

const MOCK_RESULTS: SearchResultItem[] = [
  {
    assetId: 'asset-passport-001',
    vaultId: VAULT_ID,
    title: 'Japanese Visa & Passport Scan',
    summary: 'Passport photo page and Japanese entry visa for Tokyo vacation 2025.',
    contentSnippet: 'Passport number A987654321, issued by US Department of State. Entry visa for Japan valid for 90 days.',
    assetType: 'PDF',
    score: 0.94,
    confidence: 0.96,
    whyMatched: 'Matched via exact keyword match in title & matched entity (Tokyo, Japan, Passport) & high semantic similarity.',
    contributingStrategies: ['KEYWORD', 'VECTOR', 'ENTITY', 'CONTEXT', 'MEMORY'],
    matchedEntities: ['Tokyo', 'Japan', 'Passport'],
    matchedContexts: ['Japan Vacation'],
    relatedMemories: ['Japan Vacation 2025'],
    highlightedText: '<mark>Japanese</mark> Visa & <mark>Passport</mark> Scan',
    createdAt: new Date(Date.now() - 30 * 86400000).toISOString(),
  },
  {
    assetId: 'asset-thesis-002',
    vaultId: VAULT_ID,
    title: 'PhD Thesis Chapter 4 — Neural Networks',
    summary: 'Deep learning models, PyTorch code snippets, and transformer benchmark results.',
    contentSnippet: 'Chapter 4 evaluates attention mechanisms and graph neural network embeddings on benchmark datasets.',
    assetType: 'DOCX',
    score: 0.86,
    confidence: 0.90,
    whyMatched: 'Matched via matched entity (PyTorch, Neural Networks) & high semantic similarity.',
    contributingStrategies: ['VECTOR', 'ENTITY', 'CONTEXT', 'MEMORY'],
    matchedEntities: ['PyTorch', 'Neural Networks', 'Dr. Sharma'],
    matchedContexts: ['PhD Thesis'],
    relatedMemories: ['Thesis Chapter 4 — Neural Networks'],
    highlightedText: 'PhD Thesis Chapter 4 — <mark>Neural Networks</mark>',
    createdAt: new Date(Date.now() - 15 * 86400000).toISOString(),
  },
  {
    assetId: 'asset-medical-003',
    vaultId: VAULT_ID,
    title: 'Cardiology Consultation Note — Dr. Sharma',
    summary: 'Annual heart check-up report and prescription refill confirmation.',
    contentSnippet: 'Patient blood pressure 120/80, normal ECG. Prescription for Atorvastatin 20mg renewed.',
    assetType: 'PDF',
    score: 0.79,
    confidence: 0.92,
    whyMatched: 'Matched via matched entity (Dr. Sharma, Cardiology) & matched context (Medical Treatment).',
    contributingStrategies: ['KEYWORD', 'ENTITY', 'CONTEXT', 'MEMORY'],
    matchedEntities: ['Dr. Sharma', 'Cardiology', 'Hospital'],
    matchedContexts: ['Medical Treatment'],
    relatedMemories: ['Recurring Medical Visits — Dr. Sharma'],
    highlightedText: 'Cardiology Consultation Note — <mark>Dr. Sharma</mark>',
    createdAt: new Date(Date.now() - 10 * 86400000).toISOString(),
  },
];

const MOCK_RESPONSE: SearchQueryResponse = {
  results: MOCK_RESULTS,
  total: MOCK_RESULTS.length,
  page: 1,
  pageSize: 20,
  latencyMs: 14,
  parsedQuery: {
    rawText: 'passport Japan 2025',
    cleanedText: 'passport Japan 2025',
    detectedIntent: 'Travel',
    detectedEntities: ['Tokyo', 'Japan', 'Passport'],
    detectedFileTypes: ['PDF'],
  },
  searchPlan: {
    selectedStrategies: ['KEYWORD', 'VECTOR', 'ENTITY', 'CONTEXT', 'MEMORY'],
    weights: { keyword: 0.25, vector: 0.25, entity: 0.15, context: 0.15, relationship: 0.05, memory: 0.10, recency: 0.05 },
    reasoning: 'Executing hybrid retrieval: Keyword (FTS) + Semantic Vector Search + Entity Match (Tokyo, Japan) + Context Intent Match (Travel)',
  },
  suggestions: ['passport Japan 2025 in Japan Vacation', 'passport Japan 2025 PDF documents'],
};

export const SearchService = {
  async search(queryText: string, filters: SearchFilters = {}, vaultId: string = VAULT_ID): Promise<SearchQueryResponse> {
    try {
      return await apiFetch<SearchQueryResponse>(`${BASE_URL}/search/query`, {
        method: 'POST',
        body: JSON.stringify({ vaultId, queryText, filters, page: 1, pageSize: 20 }),
      });
    } catch {
      // Mock filtering if offline
      let filtered = [...MOCK_RESULTS];
      if (filters.fileTypes && filters.fileTypes.length > 0) {
        filtered = filtered.filter((r) => filters.fileTypes?.includes(r.assetType));
      }
      return { ...MOCK_RESPONSE, results: filtered, total: filtered.length };
    }
  },

  async getSuggestions(queryText: string, vaultId: string = VAULT_ID): Promise<string[]> {
    try {
      const res = await apiFetch<{ suggestions: string[] }>(`${BASE_URL}/search/suggestions?vaultId=${vaultId}&q=${encodeURIComponent(queryText)}`);
      return res.suggestions ?? [];
    } catch {
      return [`${queryText} in Japan Vacation`, `${queryText} PDF files`];
    }
  },

  async getSavedSearches(vaultId: string = VAULT_ID): Promise<SavedSearch[]> {
    try {
      const res = await apiFetch<{ savedSearches: SavedSearch[] }>(`${BASE_URL}/search/saved?vaultId=${vaultId}`);
      return res.savedSearches ?? [];
    } catch {
      return [
        { savedId: 's-1', vaultId, searchName: 'Japan Trip Documents', queryText: 'Japan passport visa', filters: {}, isPinned: true, createdAt: new Date().toISOString(), updatedAt: new Date().toISOString() },
        { savedId: 's-2', vaultId, searchName: 'PhD Research Papers', queryText: 'PyTorch thesis neural', filters: {}, isPinned: false, createdAt: new Date().toISOString(), updatedAt: new Date().toISOString() },
      ];
    }
  },

  async saveSearch(searchName: string, queryText: string, filters: SearchFilters, vaultId: string = VAULT_ID): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/search/saved`, {
        method: 'POST',
        body: JSON.stringify({ vaultId, searchName, queryText, filters, isPinned: false }),
      });
    } catch { /* mock */ }
  },

  async deleteSavedSearch(savedId: string): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/search/saved?savedId=${savedId}`, { method: 'DELETE' });
    } catch { /* mock */ }
  },

  async getHistory(vaultId: string = VAULT_ID): Promise<SearchHistoryItem[]> {
    try {
      const res = await apiFetch<{ history: SearchHistoryItem[] }>(`${BASE_URL}/search/history?vaultId=${vaultId}`);
      return res.history ?? [];
    } catch {
      return [
        { historyId: 'h-1', vaultId, queryText: 'passport Japan 2025', parsedQuery: MOCK_RESPONSE.parsedQuery, resultCount: 3, latencyMs: 14, searchType: 'HYBRID', searchedAt: new Date(Date.now() - 300000).toISOString() },
        { historyId: 'h-2', vaultId, queryText: 'Dr. Sharma cardiology', parsedQuery: MOCK_RESPONSE.parsedQuery, resultCount: 1, latencyMs: 9, searchType: 'HYBRID', searchedAt: new Date(Date.now() - 3600000).toISOString() },
      ];
    }
  },

  async getStats(vaultId: string = VAULT_ID): Promise<SearchStats> {
    try {
      return await apiFetch<SearchStats>(`${BASE_URL}/search/stats?vaultId=${vaultId}`);
    } catch {
      return {
        vaultId,
        totalSearches: 48,
        topQueries: ['Japan passport', 'Thesis Chapter 4', 'Dr. Sharma medical', 'Tax return 2024'],
        avgLatencyMs: 12.4,
        strategyUsage: { HYBRID: 34, KEYWORD: 8, VECTOR: 6 },
        updatedAt: new Date().toISOString(),
      };
    }
  },
};
