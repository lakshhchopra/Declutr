'use client';

import React, { useState, useEffect } from 'react';
import type { SearchQueryResponse, SearchFilters } from '../types/search';
import { SearchService } from '../services/search-service';

interface GlobalSearchProps {
  onSearchComplete: (response: SearchQueryResponse) => void;
  activeFilters: SearchFilters;
}

export function GlobalSearch({ onSearchComplete, activeFilters }: GlobalSearchProps) {
  const [query, setQuery] = useState('passport Japan 2025');
  const [suggestions, setSuggestions] = useState<string[]>([]);
  const [showSuggestions, setShowSuggestions] = useState(false);
  const [loading, setLoading] = useState(false);

  const runSearch = async (text: string) => {
    if (!text.trim()) return;
    setLoading(true);
    const res = await SearchService.search(text, activeFilters);
    onSearchComplete(res);
    setLoading(false);
    setShowSuggestions(false);
  };

  useEffect(() => {
    runSearch(query);
  }, [activeFilters]);

  const handleInputChange = (text: string) => {
    setQuery(text);
    if (text.length > 2) {
      SearchService.getSuggestions(text).then(setSuggestions);
      setShowSuggestions(true);
    } else {
      setShowSuggestions(false);
    }
  };

  return (
    <div style={styles.container}>
      <div style={styles.inputWrapper}>
        <span style={styles.searchIcon}>🔍</span>
        <input
          type="text"
          value={query}
          onChange={(e) => handleInputChange(e.target.value)}
          onKeyDown={(e) => e.key === 'Enter' && runSearch(query)}
          placeholder="Search files, memories, contexts, entities… (e.g. 'passport from Japan trip')"
          style={styles.input}
        />
        {loading && <span style={styles.spinner}>⚡</span>}
        <button style={styles.searchBtn} onClick={() => runSearch(query)}>
          Search
        </button>
      </div>

      {/* Autocomplete Dropdown */}
      {showSuggestions && suggestions.length > 0 && (
        <div style={styles.suggestionsDropdown}>
          {suggestions.map((sug, i) => (
            <div
              key={i}
              style={styles.suggestionItem}
              onClick={() => {
                setQuery(sug);
                runSearch(sug);
              }}
            >
              💡 {sug}
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { position: 'relative' as const, marginBottom: '24px' },
  inputWrapper: { display: 'flex', alignItems: 'center', background: '#1e293b', border: '1px solid #334155', borderRadius: '14px', padding: '6px 14px', gap: '10px', boxShadow: '0 4px 20px rgba(0,0,0,0.2)' },
  searchIcon: { fontSize: '18px', color: '#64748b' },
  input: { flex: 1, background: 'transparent', border: 'none', color: '#e2e8f0', fontSize: '15px', outline: 'none' },
  spinner: { fontSize: '16px', animation: 'spin 1s linear infinite' },
  searchBtn: { background: 'linear-gradient(135deg, #6366f1, #818cf8)', color: '#fff', border: 'none', borderRadius: '10px', padding: '8px 20px', fontSize: '13px', fontWeight: 700, cursor: 'pointer' },
  suggestionsDropdown: { position: 'absolute' as const, top: '100%', left: 0, right: 0, marginTop: '6px', background: '#1e293b', border: '1px solid #334155', borderRadius: '12px', overflow: 'hidden', zIndex: 10, boxShadow: '0 8px 32px rgba(0,0,0,0.4)' },
  suggestionItem: { padding: '10px 16px', fontSize: '13px', color: '#cbd5e1', cursor: 'pointer', borderBottom: '1px solid #0f172a' },
};
