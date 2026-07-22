'use client';

import React, { useEffect, useState } from 'react';
import type { EmbeddingStatusResponse } from '../types/embedding';
import { EmbeddingService } from '../services/embedding-service';

const VAULT_ID = 'vault-demo';

export function EmbeddingStatus() {
  const [status, setStatus] = useState<EmbeddingStatusResponse | null>(null);
  const [refreshing, setRefreshing] = useState(false);

  useEffect(() => {
    EmbeddingService.getStatus(VAULT_ID).then(setStatus);
  }, []);

  const handleRefresh = async () => {
    setRefreshing(true);
    await EmbeddingService.refreshEmbeddings(VAULT_ID);
    setTimeout(async () => {
      const updated = await EmbeddingService.getStatus(VAULT_ID);
      setStatus(updated);
      setRefreshing(false);
    }, 1000);
  };

  if (!status) return <div style={styles.loading}>Checking status…</div>;

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <span style={styles.title}>🟢 Pipeline Operational Status</span>
        <span style={styles.subtitle}>Real-time health monitoring of the embedding engine and vector repository.</span>
      </div>

      <div style={styles.statusCard}>
        <div style={styles.statusRow}>
          <span style={styles.statusLabel}>Engine Health</span>
          <span style={styles.statusBadge}>● {status.status}</span>
        </div>
        <div style={styles.statusRow}>
          <span style={styles.statusLabel}>Active Provider</span>
          <span style={styles.statusVal}>{status.activeProvider}</span>
        </div>
        <div style={styles.statusRow}>
          <span style={styles.statusLabel}>Active Model</span>
          <span style={styles.statusVal}>{status.activeModel}</span>
        </div>
        <div style={styles.statusRow}>
          <span style={styles.statusLabel}>Vector Dimensions</span>
          <span style={styles.statusVal}>{status.dimensions}d</span>
        </div>
        <div style={styles.statusRow}>
          <span style={styles.statusLabel}>Active Jobs</span>
          <span style={styles.statusVal}>{status.activeJobs} queued</span>
        </div>
      </div>

      <button style={styles.btn} onClick={handleRefresh} disabled={refreshing}>
        {refreshing ? 'Refreshing Pipeline…' : '⚡ Refresh Pipeline'}
      </button>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { padding: '24px', maxWidth: '640px', margin: '0 auto' },
  loading: { textAlign: 'center', padding: '40px', color: '#94a3b8' },
  header: { marginBottom: '24px' },
  title: { fontSize: '20px', fontWeight: 700, color: '#e2e8f0', display: 'block', marginBottom: '4px' },
  subtitle: { fontSize: '13px', color: '#64748b' },
  statusCard: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px', display: 'flex', flexDirection: 'column', gap: '14px', marginBottom: '20px' },
  statusRow: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  statusLabel: { fontSize: '14px', color: '#94a3b8' },
  statusVal: { fontSize: '14px', fontWeight: 700, color: '#e2e8f0' },
  statusBadge: { background: '#4ade8022', color: '#4ade80', border: '1px solid #4ade8044', borderRadius: '12px', padding: '4px 12px', fontSize: '12px', fontWeight: 700 },
  btn: { width: '100%', background: '#6366f1', color: '#fff', border: 'none', borderRadius: '10px', padding: '12px', fontSize: '14px', fontWeight: 700, cursor: 'pointer' },
};
