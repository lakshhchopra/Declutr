'use client';

import React from 'react';
import type { SearchQueryResponse, SearchStrategy } from '../types/search';

interface SearchResultsProps {
  response: SearchQueryResponse | null;
}

const strategyColors: Record<SearchStrategy, string> = {
  KEYWORD: '#f59e0b',
  VECTOR: '#38bdf8',
  ENTITY: '#a855f7',
  CONTEXT: '#6366f1',
  RELATIONSHIP: '#ec4899',
  MEMORY: '#4ade80',
  METADATA: '#64748b',
  HYBRID: '#818cf8',
};

export function SearchResults({ response }: SearchResultsProps) {
  if (!response) return <div style={styles.loading}>Enter a query to search across all vault knowledge.</div>;

  const { results, total, latencyMs, parsedQuery, searchPlan } = response;

  return (
    <div style={styles.container}>
      {/* Search Plan Banner */}
      <div style={styles.planBanner}>
        <div style={styles.planHeader}>
          <span style={styles.planTitle}>
            🔍 Found {total} results in <span style={styles.latency}>{latencyMs}ms</span>
          </span>
          {parsedQuery.detectedIntent && (
            <span style={styles.intentChip}>Intent: {parsedQuery.detectedIntent}</span>
          )}
        </div>
        <div style={styles.planReasoning}>{searchPlan.reasoning}</div>
        <div style={styles.strategyRow}>
          {searchPlan.selectedStrategies.map((s) => {
            const col = strategyColors[s] ?? '#6366f1';
            return (
              <span key={s} style={{ ...styles.stratBadge, background: col + '22', color: col, borderColor: col + '44' }}>
                ● {s}
              </span>
            );
          })}
        </div>
      </div>

      {/* Results List */}
      {results.length === 0 ? (
        <div style={styles.empty}>No matches found. Try broadening your query or clearing active filters.</div>
      ) : (
        <div style={styles.list}>
          {results.map((item) => (
            <div key={item.assetId} style={styles.card}>
              <div style={styles.cardTop}>
                <div style={styles.cardTitleRow}>
                  <span style={styles.fileTypeBadge}>{item.assetType}</span>
                  <span style={styles.cardTitle} dangerouslySetInnerHTML={{ __html: item.highlightedText || item.title }} />
                </div>
                <div style={styles.scoreBox}>
                  <div style={styles.scoreVal}>{Math.round(item.score * 100)}% Match</div>
                  <div style={styles.scoreBarBg}>
                    <div style={{ ...styles.scoreBarFill, width: `${Math.round(item.score * 100)}%` }} />
                  </div>
                </div>
              </div>

              <div style={styles.summary}>{item.summary}</div>
              <div style={styles.snippet}>"{item.contentSnippet}"</div>

              {/* Explainability Rationale */}
              <div style={styles.explainBox}>
                <span style={styles.explainIcon}>💡</span>
                <span style={styles.explainText}><b>Why matched:</b> {item.whyMatched}</span>
              </div>

              {/* Tags & Pills */}
              <div style={styles.tagsRow}>
                {item.matchedEntities && item.matchedEntities.map((e) => (
                  <span key={e} style={styles.entityTag}>👤 {e}</span>
                ))}
                {item.matchedContexts && item.matchedContexts.map((c) => (
                  <span key={c} style={styles.contextTag}>🌐 {c}</span>
                ))}
                {item.relatedMemories && item.relatedMemories.map((m) => (
                  <span key={m} style={styles.memoryTag}>🧠 {m}</span>
                ))}
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { flex: 1 },
  loading: { textAlign: 'center', padding: '60px', color: '#94a3b8' },
  empty: { textAlign: 'center', padding: '60px', color: '#64748b' },
  planBanner: { background: '#1e293b', border: '1px solid #334155', borderRadius: '14px', padding: '16px 20px', marginBottom: '20px' },
  planHeader: { display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '6px' },
  planTitle: { fontSize: '15px', fontWeight: 700, color: '#e2e8f0' },
  latency: { color: '#4ade80' },
  intentChip: { background: '#6366f122', color: '#818cf8', border: '1px solid #6366f144', borderRadius: '12px', padding: '2px 10px', fontSize: '11px', fontWeight: 700 },
  planReasoning: { fontSize: '12px', color: '#94a3b8', marginBottom: '10px' },
  strategyRow: { display: 'flex', gap: '8px', flexWrap: 'wrap' as const },
  stratBadge: { border: '1px solid', borderRadius: '12px', padding: '2px 8px', fontSize: '10px', fontWeight: 700 },
  list: { display: 'flex', flexDirection: 'column', gap: '16px' },
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px', display: 'flex', flexDirection: 'column', gap: '10px' },
  cardTop: { display: 'flex', justifyContent: 'space-between', alignItems: 'flex-start' },
  cardTitleRow: { display: 'flex', alignItems: 'center', gap: '10px' },
  fileTypeBadge: { background: '#0f172a', border: '1px solid #334155', color: '#38bdf8', borderRadius: '6px', padding: '2px 8px', fontSize: '11px', fontWeight: 700 },
  cardTitle: { fontSize: '16px', fontWeight: 700, color: '#e2e8f0' },
  scoreBox: { textAlign: 'right' as const },
  scoreVal: { fontSize: '13px', fontWeight: 800, color: '#4ade80', marginBottom: '4px' },
  scoreBarBg: { width: '80px', background: '#0f172a', borderRadius: '4px', height: '4px', overflow: 'hidden' },
  scoreBarFill: { height: '100%', background: '#4ade80', borderRadius: '4px' },
  summary: { fontSize: '13px', color: '#cbd5e1', lineHeight: 1.5 },
  snippet: { fontSize: '12px', color: '#64748b', fontStyle: 'italic', background: '#0f172a', padding: '8px 12px', borderRadius: '8px', borderLeft: '3px solid #6366f1' },
  explainBox: { background: '#0f172a', border: '1px solid #334155', borderRadius: '8px', padding: '8px 12px', display: 'flex', alignItems: 'center', gap: '8px' },
  explainIcon: { fontSize: '14px' },
  explainText: { fontSize: '12px', color: '#e2e8f0' },
  tagsRow: { display: 'flex', gap: '8px', flexWrap: 'wrap' as const, marginTop: '4px' },
  entityTag: { background: '#a855f715', color: '#c084fc', border: '1px solid #a855f733', borderRadius: '12px', padding: '2px 10px', fontSize: '11px' },
  contextTag: { background: '#6366f115', color: '#818cf8', border: '1px solid #6366f133', borderRadius: '12px', padding: '2px 10px', fontSize: '11px' },
  memoryTag: { background: '#4ade8015', color: '#4ade80', border: '1px solid #4ade8033', borderRadius: '12px', padding: '2px 10px', fontSize: '11px' },
};
