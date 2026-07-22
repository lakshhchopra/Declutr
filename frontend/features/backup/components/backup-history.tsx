'use client';

import React, { useState, useEffect } from 'react';
import type { Backup } from '../types/backup';
import { BackupService } from '../services/backup-service';

interface BackupHistoryProps {
  onSelectRestore?: (backupId: string) => void;
}

export function BackupHistory({ onSelectRestore }: BackupHistoryProps) {
  const [backups, setBackups] = useState<Backup[]>([]);
  const [loading, setLoading] = useState(true);
  const [verifyMsg, setVerifyMsg] = useState<Record<string, string>>({});

  const loadData = async () => {
    setLoading(true);
    const list = await BackupService.getBackups();
    setBackups(list);
    setLoading(false);
  };

  useEffect(() => {
    loadData();
  }, []);

  const handleVerify = async (id: string) => {
    const res = await BackupService.verifyIntegrity(id);
    setVerifyMsg((prev) => ({ ...prev, [id]: res.message }));
  };

  return (
    <div style={styles.container}>
      <h3 style={styles.title}>📜 Vault Backup History & Integrity Trail</h3>

      {loading ? (
        <div style={styles.loading}>Loading backup history...</div>
      ) : backups.length === 0 ? (
        <div style={styles.empty}>No backup records found.</div>
      ) : (
        <div style={styles.list}>
          {backups.map((b) => (
            <div key={b.backupId} style={styles.card}>
              <div style={styles.cardHeader}>
                <div style={styles.badgeRow}>
                  <span style={styles.typeBadge}>{b.backupType}</span>
                  <span style={styles.statusBadge}>{b.status}</span>
                  {b.isEncrypted && <span style={styles.encBadge}>🔒 Encrypted</span>}
                </div>
                <span style={styles.date}>{new Date(b.createdAt).toLocaleString()}</span>
              </div>

              <h4 style={styles.backupId}>{b.backupId}</h4>
              <div style={styles.meta}>
                Compressed Size: {(b.compressedSizeBytes / (1024 * 1024)).toFixed(1)} MB • Checksum: {b.checksum}
              </div>

              {verifyMsg[b.backupId] && <div style={styles.verifyMsg}>{verifyMsg[b.backupId]}</div>}

              <div style={styles.actionRow}>
                <button style={styles.verifyBtn} onClick={() => handleVerify(b.backupId)}>
                  🔍 Verify SHA-256 Checksum
                </button>
                <button style={styles.restoreBtn} onClick={() => onSelectRestore && onSelectRestore(b.backupId)}>
                  🚨 Restore From This Backup
                </button>
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
  title: { fontSize: '18px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  loading: { textAlign: 'center', padding: '40px', color: '#94a3b8' },
  empty: { textAlign: 'center', padding: '40px', color: '#64748b' },
  list: { display: 'flex', flexDirection: 'column', gap: '12px' },
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '18px', display: 'flex', flexDirection: 'column', gap: '8px' },
  cardHeader: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  badgeRow: { display: 'flex', alignItems: 'center', gap: '8px' },
  typeBadge: { background: '#6366f122', color: '#818cf8', borderRadius: '6px', padding: '2px 8px', fontSize: '11px', fontWeight: 800, border: '1px solid #6366f144' },
  statusBadge: { background: '#4ade8015', color: '#4ade80', borderRadius: '6px', padding: '2px 8px', fontSize: '10px', fontWeight: 800, border: '1px solid #4ade8033' },
  encBadge: { background: '#f59e0b15', color: '#f59e0b', borderRadius: '6px', padding: '2px 8px', fontSize: '10px', fontWeight: 800, border: '1px solid #f59e0b33' },
  date: { fontSize: '11px', color: '#64748b' },
  backupId: { fontSize: '15px', fontWeight: 700, color: '#e2e8f0', margin: 0 },
  meta: { fontSize: '12px', color: '#94a3b8' },
  verifyMsg: { fontSize: '11px', color: '#4ade80', background: '#4ade8015', padding: '6px 10px', borderRadius: '6px', border: '1px solid #4ade8033' },
  actionRow: { display: 'flex', gap: '10px', marginTop: '4px' },
  verifyBtn: { background: '#0f172a', border: '1px solid #334155', color: '#38bdf8', borderRadius: '8px', padding: '6px 14px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
  restoreBtn: { background: 'linear-gradient(135deg, #ef4444, #f87171)', color: '#fff', border: 'none', borderRadius: '8px', padding: '6px 14px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
};
