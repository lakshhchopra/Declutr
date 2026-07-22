'use client';

import React, { useState, useEffect } from 'react';
import type { AuditEvent, AuditCategory } from '../types/security';
import { SecurityService } from '../services/security-service';

export function AuditViewerComponent() {
  const [events, setEvents] = useState<AuditEvent[]>([]);
  const [selectedCat, setSelectedCat] = useState<AuditCategory | ''>('');
  const [loading, setLoading] = useState(true);

  const loadAudit = async (cat?: AuditCategory) => {
    setLoading(true);
    const list = await SecurityService.getAuditEvents(cat);
    setEvents(list);
    setLoading(false);
  };

  useEffect(() => {
    loadAudit(selectedCat || undefined);
  }, [selectedCat]);

  return (
    <div style={styles.container}>
      {/* Filter Row */}
      <div style={styles.filterRow}>
        <h3 style={styles.title}>📜 Vault Audit Log Hub</h3>
        <select
          value={selectedCat}
          onChange={(e) => setSelectedCat(e.target.value as any)}
          style={styles.select}
        >
          <option value="">All Categories</option>
          <option value="AUTH">Authentication (AUTH)</option>
          <option value="ASSET">Assets & Files (ASSET)</option>
          <option value="SHARING">Sharing & Permissions (SHARING)</option>
          <option value="WORKFLOW">Workflows & Rules (WORKFLOW)</option>
          <option value="AI">AI Requests (AI)</option>
          <option value="SEARCH">Search Queries (SEARCH)</option>
          <option value="BACKUP">Disaster Recovery (BACKUP)</option>
          <option value="VERSIONING">Time Machine (VERSIONING)</option>
        </select>
      </div>

      {/* Events Table / Feed */}
      {loading ? (
        <div style={styles.loading}>Loading audit event log...</div>
      ) : events.length === 0 ? (
        <div style={styles.empty}>No audit log events found for category.</div>
      ) : (
        <div style={styles.feed}>
          {events.map((evt) => (
            <div key={evt.auditId} style={styles.card}>
              <div style={styles.cardHeader}>
                <span style={styles.catBadge}>{evt.category}</span>
                <span style={styles.action}>{evt.action}</span>
                <span style={styles.date}>{new Date(evt.createdAt).toLocaleString()}</span>
              </div>
              <div style={styles.metaRow}>
                <span style={styles.metaItem}>Actor: <strong>{evt.actorName}</strong></span>
                <span style={styles.metaItem}>IP: <strong>{evt.ipAddress}</strong></span>
                {evt.resourceId && <span style={styles.metaItem}>Resource: <strong>{evt.resourceId}</strong></span>}
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { display: 'flex', flexDirection: 'column', gap: '16px' },
  filterRow: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  title: { fontSize: '18px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  select: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '8px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  loading: { textAlign: 'center', padding: '40px', color: '#94a3b8' },
  empty: { textAlign: 'center', padding: '40px', color: '#64748b' },
  feed: { display: 'flex', flexDirection: 'column', gap: '10px' },
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '12px', padding: '14px 16px', display: 'flex', flexDirection: 'column', gap: '6px' },
  cardHeader: { display: 'flex', alignItems: 'center', gap: '10px' },
  catBadge: { background: '#6366f122', color: '#818cf8', borderRadius: '6px', padding: '2px 8px', fontSize: '10px', fontWeight: 900, border: '1px solid #6366f144' },
  action: { fontSize: '14px', fontWeight: 700, color: '#e2e8f0', flex: 1 },
  date: { fontSize: '11px', color: '#64748b' },
  metaRow: { display: 'flex', gap: '16px', fontSize: '12px', color: '#94a3b8' },
  metaItem: { display: 'inline' },
};
