// Hybrid Knowledge Search Engine TypeScript types

export type SearchStrategy =
  | 'KEYWORD'
  | 'VECTOR'
  | 'ENTITY'
  | 'CONTEXT'
  | 'RELATIONSHIP'
  | 'MEMORY'
  | 'METADATA'
  | 'HYBRID';

export interface SearchFilters {
  fileTypes?: string[];
  dateFrom?: string;
  dateTo?: string;
  tags?: string[];
  collections?: string[];
  contexts?: string[];
  entities?: string[];
  people?: string[];
  locations?: string[];
  minMemoryStrength?: number;
  minConfidence?: number;
}

export interface RankingWeights {
  keyword: number;
  vector: number;
  entity: number;
  context: number;
  relationship: number;
  memory: number;
  recency: number;
}

export interface ParsedQuery {
  rawText: string;
  cleanedText: string;
  detectedIntent?: string;
  detectedEntities?: string[];
  detectedLocations?: string[];
  detectedFileTypes?: string[];
  detectedDateFrom?: string;
  detectedDateTo?: string;
  quotedTerms?: string[];
  excludedTerms?: string[];
}

export interface SearchPlan {
  selectedStrategies: SearchStrategy[];
  weights: RankingWeights;
  reasoning: string;
}

export interface SearchResultItem {
  assetId: string;
  vaultId: string;
  title: string;
  summary: string;
  contentSnippet: string;
  assetType: string;
  score: number;
  confidence: number;
  whyMatched: string;
  contributingStrategies: SearchStrategy[];
  matchedEntities?: string[];
  matchedContexts?: string[];
  relatedMemories?: string[];
  highlightedText?: string;
  createdAt: string;
  metadata?: Record<string, unknown>;
}

export interface SearchQueryResponse {
  results: SearchResultItem[];
  total: number;
  page: number;
  pageSize: number;
  latencyMs: number;
  parsedQuery: ParsedQuery;
  searchPlan: SearchPlan;
  suggestions?: string[];
}

export interface SavedSearch {
  savedId: string;
  vaultId: string;
  searchName: string;
  queryText: string;
  filters: SearchFilters;
  isPinned: boolean;
  createdAt: string;
  updatedAt: string;
}

export interface SearchHistoryItem {
  historyId: string;
  vaultId: string;
  queryText: string;
  parsedQuery: ParsedQuery;
  resultCount: number;
  latencyMs: number;
  searchType: string;
  searchedAt: string;
}

export interface SearchStats {
  vaultId: string;
  totalSearches: number;
  topQueries: string[];
  avgLatencyMs: number;
  strategyUsage: Record<string, number>;
  updatedAt: string;
}

export interface SearchPreferences {
  preferenceId: string;
  vaultId: string;
  rankingWeights: RankingWeights;
  enableAutoSuggestions: boolean;
  maxResultsPerPage: number;
  createdAt: string;
  updatedAt: string;
}
