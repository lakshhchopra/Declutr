'use client';

import React, { useState, useEffect } from 'react';
import type { WorkflowRun } from '../types/workflow';
import { WorkflowService } from '../services/workflow-service';

export function ExecutionHistory() {
  const [runs, setRuns] = useState<WorkflowRun[]>([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    WorkflowService.getRuns().then((res) => {
      setRuns(res);
      setLoading(false);
    });
  }, []);

  return (
    <div style={styles.container}>
      <h3 style={styles.title}>📜 Execution History & Run Logs</h3>
      {loading ? (
        <div style={styles.loading}>Loading execution runs...</div>
      ) : runs.length === 0 ? (
        <div style={styles.empty}>No execution runs recorded yet.</div>
      ) : (
        <table style={styles.table}>
          <thead>
            <tr style={styles.thRow}>
              <th style={styles.th}>Run ID</th>
              <th style={styles.th}>Trigger Event</th>
              <th style={styles.th}>Status</th>
              <th style={styles.th}>Duration</th>
              <th style={styles.th}>Executed At</th>
            </tr>
          </thead>
          <tbody>
            {runs.map((r) => {
              const isSucc = r.status === 'SUCCESS';
              return (
                <tr key={r.runId} style={styles.tr}>
                  <td style={styles.tdId}>{r.runId}</td>
                  <td style={styles.td}>{r.triggerEvent}</td>
                  <td style={styles.td}>
                    <span style={{ ...styles.statusBadge, color: isSucc ? '#4ade80' : '#ef4444', borderColor: isSucc ? '#4ade8044' : '#ef444444', background: isSucc ? '#4ade8015' : '#ef444415' }}>
                      {r.status}
                    </span>
                  </td>
                  <td style={styles.td}>{r.durationMs}ms</td>
                  <td style={styles.td}>{new Date(r.startedAt).toLocaleString()}</td>
                </tr>
              );
            })}
          </tbody>
        </table>
      )}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px' },
  title: { fontSize: '16px', fontWeight: 700, color: '#e2e8f0', marginBottom: '16px' },
  loading: { textAlign: 'center', padding: '40px', color: '#94a3b8' },
  empty: { textAlign: 'center', padding: '40px', color: '#64748b' },
  table: { width: '100%', borderCollapse: 'collapse' as const },
  thRow: { borderBottom: '1px solid #334155', textAlign: 'left' as const },
  th: { padding: '10px', fontSize: '11px', color: '#64748b', fontWeight: 700, textTransform: 'uppercase' as const },
  tr: { borderBottom: '1px solid #0f172a' },
  td: { padding: '12px 10px', fontSize: '13px', color: '#cbd5e1' },
  tdId: { padding: '12px 10px', fontSize: '13px', color: '#38bdf8', fontWeight: 700 },
  statusBadge: { border: '1px solid', borderRadius: '6px', padding: '2px 8px', fontSize: '10px', fontWeight: 800 },
};
