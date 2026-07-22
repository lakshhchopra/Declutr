'use client';

import React, { useState, useEffect } from 'react';
import type { ShareActivity } from '../types/collaboration';
import { CollaborationService } from '../services/collaboration-service';

export function ActivityFeed() {
  const [activity, setActivity] = useState<ShareActivity[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    CollaborationService.getActivity().then((res) => {
      setActivity(res);
      setLoading(false);
    });
  }, []);

  return (
    <div style={styles.container}>
      <h4 style={styles.title}>📜 Audit Activity Trail</h4>
      {loading ? (
        <div style={styles.loading}>Loading audit activities...</div>
      ) : activity.length === 0 ? (
        <div style={styles.empty}>No audit log entries recorded.</div>
      ) : (
        <div style={styles.feed}>
          {activity.map((act) => (
            <div key={act.activityId} style={styles.card}>
              <div style={styles.header}>
                <span style={styles.badge}>{act.actionType}</span>
                <span style={styles.date}>{new Date(act.createdAt).toLocaleString()}</span>
              </div>
              <span style={styles.actor}>{act.actorName}</span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px', display: 'flex', flexDirection: 'column', gap: '14px' },
  title: { fontSize: '16px', fontWeight: 700, color: '#e2e8f0', margin: 0 },
  loading: { textAlign: 'center', padding: '40px', color: '#94a3b8' },
  empty: { textAlign: 'center', padding: '40px', color: '#64748b' },
  feed: { display: 'flex', flexDirection: 'column', gap: '10px' },
  card: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '12px 14px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  header: { display: 'flex', alignItems: 'center', gap: '10px' },
  badge: { fontSize: '10px', fontWeight: 800, color: '#4ade80', background: '#4ade8015', padding: '2px 8px', borderRadius: '6px', border: '1px solid #4ade8033' },
  date: { fontSize: '11px', color: '#64748b' },
  actor: { fontSize: '13px', fontWeight: 700, color: '#e2e8f0' },
};
