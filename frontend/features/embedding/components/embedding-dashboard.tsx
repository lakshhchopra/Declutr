'use client';

import React, { useEffect, useState } from 'react';
import type { EmbeddingStats, EmbeddingStatusResponse } from '../types/embedding';
import { EmbeddingService } from '../services/embedding-service';

const VAULT_ID = 'vault-demo';

const providerColors: Record<string, string> = {
  OPENAI: '#10a37f',
  GEMINI: '#4285f4',
  VOYAGE: '#8b5cf6',
  COHERE: '#d97706',
  OLLAMA: '#059669',
  LOCAL: '#6366f1',
};

export function EmbeddingDashboard() {
  const [stats, setStats] = useState<EmbeddingStats | null>(null);
  const [status, setStatus] = useState<EmbeddingStatusResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Promise.all([
      EmbeddingService.getStats(VAULT_ID),
      EmbeddingService.getStatus(VAULT_ID),
    ]).then(([s, st]) => {
      setStats(s);
      setStatus(st);
      setLoading(false);
    });
  }, []);

  if (loading) return <div style={styles.loading}>Loading embedding representations…</div>;
  if (!stats) return null;

  const pColor = providerColors[stats.activeProvider] ?? '#6366f1';

  return (
    <div style={styles.container}>
      {/* Top Banner */}
      <div style={styles.banner}>
        <div style={styles.bannerLeft}>
          <span style={styles.bannerIcon}>💎</span>
          <div>
            <div style={styles.bannerTitle}>Embedding & Knowledge Representation Engine</div>
            <div style={styles.bannerSub}>Transforms structured memories, contexts, entities, and documents into high-dimensional vector representations.</div>
          </div>
        </div>
        <div style={{ ...styles.providerBadge, background: pColor + '22', color: pColor, borderColor: pColor + '44' }}>
          ● {stats.activeProvider} ({stats.dimensions}d)
        </div>
      </div>

      {/* Metric Cards */}
      <div style={styles.statsGrid}>
        {[
          { label: 'Total Vectors', value: stats.totalEmbeddings, icon: '📐', color: '#818cf8' },
          { label: 'Total Chunks', value: stats.totalChunks, icon: '🧩', color: '#38bdf8' },
          { label: 'Active Model', value: stats.activeModel, icon: '🤖', color: '#a855f7' },
          { label: 'Dimensions', value: `${stats.dimensions}d`, icon: '📊', color: '#f59e0b' },
          { label: 'Pipeline Status', value: status?.status ?? 'HEALTHY', icon: '🟢', color: '#4ade80' },
          { label: 'Version Tag', value: stats.latestVersionTag, icon: '🏷️', color: '#ec4899' },
        ].map((card) => (
          <div key={card.label} style={styles.card}>
            <div style={styles.cardIcon}>{card.icon}</div>
            <div style={{ ...styles.cardValue, color: card.color }}>{card.value}</div>
            <div style={styles.cardLabel}>{card.label}</div>
          </div>
        ))}
      </div>

      {/* Breakdown Grids */}
      <div style={styles.twoCol}>
        {/* Source Type Breakdown */}
        <div style={styles.sectionCard}>
          <div style={styles.sectionHeader}>📂 Source Breakdown</div>
          <div style={styles.breakdownList}>
            {Object.entries(stats.sourceTypeBreakdown).map(([source, count]) => {
              const pct = Math.round((count / stats.totalEmbeddings) * 100);
              return (
                <div key={source} style={styles.breakdownRow}>
                  <span style={styles.sourceName}>{source}</span>
                  <div style={styles.barBg}>
                    <div style={{ ...styles.barFill, width: `${pct}%`, background: '#6366f1' }} />
                  </div>
                  <span style={styles.countText}>{count} ({pct}%)</span>
                </div>
              );
            })}
          </div>
        </div>

        {/* Chunk Strategy Breakdown */}
        <div style={styles.sectionCard}>
          <div style={styles.sectionHeader}>🧩 Chunk Strategy Breakdown</div>
          <div style={styles.breakdownList}>
            {Object.entries(stats.strategyBreakdown).map(([strat, count]) => {
              const pct = Math.round((count / stats.totalChunks) * 100);
              return (
                <div key={strat} style={styles.breakdownRow}>
                  <span style={styles.sourceName}>{strat}</span>
                  <div style={styles.barBg}>
                    <div style={{ ...styles.barFill, width: `${pct}%`, background: '#38bdf8' }} />
                  </div>
                  <span style={styles.countText}>{count} ({pct}%)</span>
                </div>
              );
            })}
          </div>
        </div>
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { padding: '24px', maxWidth: '960px', margin: '0 auto' },
  loading: { textAlign: 'center', padding: '60px', color: '#94a3b8' },
  banner: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px 24px', display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '28px' },
  bannerLeft: { display: 'flex', alignItems: 'center', gap: '16px' },
  bannerIcon: { fontSize: '32px' },
  bannerTitle: { fontSize: '18px', fontWeight: 800, color: '#e0e7ff', marginBottom: '4px' },
  bannerSub: { fontSize: '13px', color: '#64748b' },
  providerBadge: { border: '1px solid', borderRadius: '20px', padding: '6px 16px', fontSize: '13px', fontWeight: 700, whiteSpace: 'nowrap' as const },
  statsGrid: { display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(140px, 1fr))', gap: '14px', marginBottom: '28px' },
  card: { background: '#1e293b', borderRadius: '14px', padding: '16px', border: '1px solid #334155', display: 'flex', flexDirection: 'column', alignItems: 'center', gap: '6px' },
  cardIcon: { fontSize: '24px' },
  cardValue: { fontSize: '20px', fontWeight: 800, textAlign: 'center' as const },
  cardLabel: { fontSize: '11px', fontWeight: 600, color: '#64748b', textTransform: 'uppercase' as const, letterSpacing: '0.08em', textAlign: 'center' as const },
  twoCol: { display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(380px, 1fr))', gap: '20px' },
  sectionCard: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px' },
  sectionHeader: { fontSize: '14px', fontWeight: 700, color: '#e2e8f0', marginBottom: '16px' },
  breakdownList: { display: 'flex', flexDirection: 'column', gap: '12px' },
  breakdownRow: { display: 'flex', alignItems: 'center', gap: '12px' },
  sourceName: { fontSize: '12px', fontWeight: 600, color: '#94a3b8', width: '130px' },
  barBg: { flex: 1, background: '#0f172a', borderRadius: '4px', height: '6px', overflow: 'hidden' },
  barFill: { height: '100%', borderRadius: '4px' },
  countText: { fontSize: '11px', color: '#64748b', width: '80px', textAlign: 'right' as const },
};
