'use client';

import React, { useEffect, useState } from 'react';
import type { SavedSearch } from '../types/search';
import { SearchService } from '../services/search-service';

const VAULT_ID = 'vault-demo';

interface SavedSearchesProps {
  onSelectSearch: (query: string) => void;
}

export function SavedSearches({ onSelectSearch }: SavedSearchesProps) {
  const [saved, setSaved] = useState<SavedSearch[]>([]);

  useEffect(() => {
    SearchService.getSavedSearches(VAULT_ID).then(setSaved);
  }, []);

  const handleDelete = async (id: string) => {
    await SearchService.deleteSavedSearch(id);
    setSaved((prev) => prev.filter((s) => s.savedId !== id));
  };

  return (
    <div style={styles.card}>
      <div style={styles.header}>📌 Saved Searches</div>
      {saved.length === 0 ? (
        <div style={styles.empty}>No saved searches.</div>
      ) : (
        <div style={styles.list}>
          {saved.map((s) => (
            <div key={s.savedId} style={styles.item} onClick={() => onSelectSearch(s.queryText)}>
              <div>
                <div style={styles.name}>{s.searchName}</div>
                <div style={styles.query}>"{s.queryText}"</div>
              </div>
              <button
                style={styles.delBtn}
                onClick={(e) => {
                  e.stopPropagation();
                  handleDelete(s.savedId);
                }}
              >
                ✕
              </button>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '18px', width: '240px' },
  header: { fontSize: '14px', fontWeight: 700, color: '#e2e8f0', marginBottom: '12px' },
  empty: { fontSize: '12px', color: '#64748b' },
  list: { display: 'flex', flexDirection: 'column', gap: '8px' },
  item: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 12px', display: 'flex', justifyContent: 'space-between', alignItems: 'center', cursor: 'pointer' },
  name: { fontSize: '13px', fontWeight: 700, color: '#e2e8f0' },
  query: { fontSize: '11px', color: '#94a3b8' },
  delBtn: { background: 'transparent', border: 'none', color: '#64748b', fontSize: '12px', cursor: 'pointer' },
};
