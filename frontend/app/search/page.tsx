'use client';

import React, { useState } from 'react';
import type { SearchQueryResponse, SearchFilters } from '../../features/search/types/search';
import { GlobalSearch } from '../../features/search/components/global-search';
import { SearchResults } from '../../features/search/components/search-results';
import { AdvancedFilters } from '../../features/search/components/advanced-filters';
import { SavedSearches } from '../../features/search/components/saved-searches';

export default function SearchPage() {
  const [response, setResponse] = useState<SearchQueryResponse | null>(null);
  const [filters, setFilters] = useState<SearchFilters>({});

  const handleSelectSaved = (query: string) => {
    // Triggers search via GlobalSearch update
  };

  return (
    <div style={styles.page}>
      {/* Page Header */}
      <div style={styles.header}>
        <h1 style={styles.heading}>Hybrid Knowledge Search Engine</h1>
        <p style={styles.subheading}>
          Unified multi-strategy search layer — fusing keyword matching, vector embeddings, canonical entities, contexts, and memory strength into a single explainable result feed.
        </p>
      </div>

      {/* Main Container */}
      <div style={styles.container}>
        {/* Left Search & Results Section */}
        <div style={styles.mainCol}>
          <GlobalSearch onSearchComplete={setResponse} activeFilters={filters} />
          <SearchResults response={response} />
        </div>

        {/* Right Sidebar Filters & Saved Searches */}
        <div style={styles.sidebarCol}>
          <AdvancedFilters filters={filters} onChange={setFilters} />
          <SavedSearches onSelectSearch={handleSelectSaved} />
        </div>
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  page: { minHeight: '100vh', background: '#0f172a', color: '#e2e8f0', fontFamily: 'Inter, system-ui, sans-serif', paddingBottom: '40px' },
  header: { padding: '32px 24px 0', maxWidth: '1080px', margin: '0 auto' },
  heading: { fontSize: '32px', fontWeight: 800, color: '#e0e7ff', marginBottom: '8px', background: 'linear-gradient(135deg, #6366f1 0%, #38bdf8 100%)', WebkitBackgroundClip: 'text', WebkitTextFillColor: 'transparent' },
  subheading: { fontSize: '14px', color: '#64748b', margin: 0, lineHeight: 1.5 },
  container: { display: 'flex', gap: '24px', maxWidth: '1080px', margin: '24px auto 0', padding: '0 24px' },
  mainCol: { flex: 1 },
  sidebarCol: { display: 'flex', flexDirection: 'column', gap: '20px' },
};
