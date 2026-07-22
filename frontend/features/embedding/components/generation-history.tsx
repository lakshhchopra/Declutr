'use client';

import React, { useEffect, useState } from 'react';
import type { EmbeddingHistoryResponse } from '../types/embedding';
import { EmbeddingService } from '../services/embedding-service';

const VAULT_ID = 'vault-demo';

export function GenerationHistory() {
  const [history, setHistory] = useState<EmbeddingHistoryResponse | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    EmbeddingService.getHistory(VAULT_ID).then((res) => {
      setHistory(res);
      setLoading(false);
    });
  }, []);

  if (loading) return <div style={styles.loading}>Loading generation history…</div>;
  if (!history) return null;

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <span style={styles.title}>📜 Generation & Version Upgrade History</span>
        <span style={styles.subtitle}>Audit trail of background vectorization jobs and provider/model upgrades.</span>
      </div>

      {/* Latest Version Card */}
      {history.latestVersion && (
        <div style={styles.verCard}>
          <div style={styles.verHeader}>
            <span style={styles.verTag}>{history.latestVersion.versionTag}</span>
            <span style={styles.verActive}>Active Version</span>
          </div>
          <div style={styles.verBody}>
            <span>Provider: <b>{history.latestVersion.providerName}</b></span>
            <span>Model: <b>{history.latestVersion.modelName}</b></span>
            <span>Dimensions: <b>{history.latestVersion.dimensions}d</b></span>
            <span>Embedded Items: <b>{history.latestVersion.totalEmbeddedItems}</b></span>
          </div>
        </div>
      )}

      {/* Job Log */}
      <div style={styles.sectionHeader}>Background Job Logs</div>
      <div style={styles.table}>
        <div style={styles.tableHead}>
          <span>Job ID</span>
          <span>Target Type</span>
          <span>Target ID</span>
          <span>Chunks</span>
          <span>Status</span>
          <span>Timestamp</span>
        </div>
        {history.jobs.map((j) => (
          <div key={j.jobId} style={styles.tableRow}>
            <span style={styles.jobId}>{j.jobId}</span>
            <span style={styles.targetType}>{j.targetType}</span>
            <span style={styles.targetId}>{j.targetId}</span>
            <span>{j.processedChunks} chunks</span>
            <span style={styles.statusCompleted}>✓ {j.status}</span>
            <span style={styles.time}>{new Date(j.createdAt).toLocaleTimeString()}</span>
          </div>
        ))}
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { padding: '24px', maxWidth: '960px', margin: '0 auto' },
  loading: { textAlign: 'center', padding: '40px', color: '#94a3b8' },
  header: { marginBottom: '24px' },
  title: { fontSize: '20px', fontWeight: 700, color: '#e2e8f0', display: 'block', marginBottom: '4px' },
  subtitle: { fontSize: '13px', color: '#64748b' },
  verCard: { background: 'linear-gradient(135deg, #1e1b4b, #312e81)', border: '1px solid #6366f144', borderRadius: '16px', padding: '20px', marginBottom: '28px' },
  verHeader: { display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginBottom: '12px' },
  verTag: { fontSize: '18px', fontWeight: 800, color: '#e0e7ff' },
  verActive: { background: '#4ade8022', color: '#4ade80', border: '1px solid #4ade8044', borderRadius: '12px', padding: '2px 10px', fontSize: '11px', fontWeight: 700 },
  verBody: { display: 'flex', gap: '20px', fontSize: '13px', color: '#a5b4fc' },
  sectionHeader: { fontSize: '13px', fontWeight: 700, color: '#6366f1', textTransform: 'uppercase' as const, letterSpacing: '0.08em', marginBottom: '12px' },
  table: { background: '#1e293b', border: '1px solid #334155', borderRadius: '14px', overflow: 'hidden' },
  tableHead: { display: 'grid', gridTemplateColumns: '1fr 1fr 1fr 1fr 1fr 1fr', gap: '12px', padding: '12px 16px', background: '#0f172a', fontSize: '11px', fontWeight: 700, color: '#64748b', textTransform: 'uppercase' as const },
  tableRow: { display: 'grid', gridTemplateColumns: '1fr 1fr 1fr 1fr 1fr 1fr', gap: '12px', padding: '14px 16px', borderTop: '1px solid #334155', alignItems: 'center', fontSize: '13px' },
  jobId: { fontFamily: 'monospace', color: '#818cf8' },
  targetType: { color: '#e2e8f0', fontWeight: 600 },
  targetId: { color: '#94a3b8' },
  statusCompleted: { color: '#4ade80', fontWeight: 700 },
  time: { color: '#64748b', fontSize: '12px' },
};
