'use client';

import React from 'react';
import type { SearchFilters } from '../types/search';

interface AdvancedFiltersProps {
  filters: SearchFilters;
  onChange: (filters: SearchFilters) => void;
}

const FILE_TYPES = ['PDF', 'DOCX', 'PNG', 'JPG', 'MP4', 'MP3'];

export function AdvancedFilters({ filters, onChange }: AdvancedFiltersProps) {
  const toggleFileType = (ft: string) => {
    const current = filters.fileTypes ?? [];
    const next = current.includes(ft)
      ? current.filter((x) => x !== ft)
      : [...current, ft];
    onChange({ ...filters, fileTypes: next });
  };

  const handleStrengthChange = (val: number) => {
    onChange({ ...filters, minMemoryStrength: val });
  };

  const handleReset = () => {
    onChange({});
  };

  return (
    <div style={styles.card}>
      <div style={styles.header}>
        <span style={styles.title}>⚙️ Filters</span>
        <button style={styles.resetBtn} onClick={handleReset}>Reset</button>
      </div>

      {/* File Type Filter */}
      <div style={styles.section}>
        <div style={styles.label}>File Types</div>
        <div style={styles.chipGrid}>
          {FILE_TYPES.map((ft) => {
            const active = filters.fileTypes?.includes(ft);
            return (
              <button
                key={ft}
                style={{ ...styles.chip, ...(active ? styles.chipActive : {}) }}
                onClick={() => toggleFileType(ft)}
              >
                {ft}
              </button>
            );
          })}
        </div>
      </div>

      {/* Minimum Memory Strength Slider */}
      <div style={styles.section}>
        <div style={styles.labelRow}>
          <span>Min Memory Strength</span>
          <span style={styles.valText}>{Math.round((filters.minMemoryStrength ?? 0) * 100)}%</span>
        </div>
        <input
          type="range"
          min="0"
          max="1"
          step="0.05"
          value={filters.minMemoryStrength ?? 0}
          onChange={(e) => handleStrengthChange(parseFloat(e.target.value))}
          style={styles.slider}
        />
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '18px', width: '240px', display: 'flex', flexDirection: 'column', gap: '16px', height: 'fit-content' },
  header: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  title: { fontSize: '14px', fontWeight: 700, color: '#e2e8f0' },
  resetBtn: { background: 'transparent', border: 'none', color: '#64748b', fontSize: '12px', cursor: 'pointer' },
  section: { display: 'flex', flexDirection: 'column', gap: '8px' },
  label: { fontSize: '12px', fontWeight: 600, color: '#94a3b8' },
  labelRow: { display: 'flex', justifyContent: 'space-between', fontSize: '12px', color: '#94a3b8', fontWeight: 600 },
  valText: { color: '#4ade80', fontWeight: 700 },
  chipGrid: { display: 'flex', gap: '6px', flexWrap: 'wrap' as const },
  chip: { background: '#0f172a', border: '1px solid #334155', color: '#94a3b8', borderRadius: '8px', padding: '4px 10px', fontSize: '11px', fontWeight: 700, cursor: 'pointer' },
  chipActive: { background: '#6366f122', borderColor: '#6366f1', color: '#818cf8' },
  slider: { width: '100%', accentColor: '#6366f1' },
};
