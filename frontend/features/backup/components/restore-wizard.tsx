'use client';

import React, { useState, useEffect } from 'react';
import type { Backup, RestoreMode, RestoreStrategy } from '../types/backup';
import { BackupService } from '../services/backup-service';

interface RestoreWizardProps {
  selectedBackupId?: string;
  onRestored?: () => void;
}

export function RestoreWizard({ selectedBackupId, onRestored }: RestoreWizardProps) {
  const [backups, setBackups] = useState<Backup[]>([]);
  const [targetBackupId, setTargetBackupId] = useState(selectedBackupId || '');
  const [restoreMode, setRestoreMode] = useState<RestoreMode>('FULL_VAULT');
  const [restoreStrategy, setRestoreStrategy] = useState<RestoreStrategy>('MERGE_RESTORE');
  const [passphrase, setPassphrase] = useState('');
  const [statusMsg, setStatusMsg] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    BackupService.getBackups().then((list) => {
      setBackups(list);
      if (list.length > 0 && !targetBackupId) {
        setTargetBackupId(list[0].backupId);
      }
    });
  }, []);

  const handleRestore = async () => {
    if (!targetBackupId) return;
    setLoading(true);
    setStatusMsg('');
    const job = await BackupService.restoreBackup(targetBackupId, restoreMode, restoreStrategy, passphrase);
    setLoading(false);
    setStatusMsg(`✅ Restoration Job ${job.jobId} executed successfully in ${restoreMode} mode!`);
    if (onRestored) onRestored();
  };

  return (
    <div style={styles.card}>
      <h3 style={styles.title}>🚨 Disaster Recovery Restore Wizard</h3>
      <p style={styles.desc}>
        Restore full or selective vault data from a historical backup package with overwrite or merge strategy.
      </p>

      {/* Target Backup */}
      <div style={styles.fieldGroup}>
        <label style={styles.label}>Select Backup Package</label>
        <select value={targetBackupId} onChange={(e) => setTargetBackupId(e.target.value)} style={styles.select}>
          {backups.map((b) => (
            <option key={b.backupId} value={b.backupId}>
              {b.backupId} ({b.backupType} • {(b.compressedSizeBytes / (1024 * 1024)).toFixed(1)} MB • {new Date(b.createdAt).toLocaleDateString()})
            </option>
          ))}
        </select>
      </div>

      <div style={styles.row}>
        {/* Mode */}
        <div style={styles.flexGroup}>
          <label style={styles.label}>Recovery Scope (Mode)</label>
          <select value={restoreMode} onChange={(e) => setRestoreMode(e.target.value as any)} style={styles.select}>
            <option value="FULL_VAULT">Full Vault Recovery (Everything)</option>
            <option value="ASSETS_ONLY">Assets & Files Only</option>
            <option value="METADATA_ONLY">Metadata & Entities Only</option>
            <option value="AI_STATE_ONLY">AI Analysis & Memories Only</option>
            <option value="WORKFLOWS_ONLY">Workflows & Rules Only</option>
          </select>
        </div>

        {/* Strategy */}
        <div style={styles.flexGroup}>
          <label style={styles.label}>Recovery Strategy</label>
          <select value={restoreStrategy} onChange={(e) => setRestoreStrategy(e.target.value as any)} style={styles.select}>
            <option value="MERGE_RESTORE">Merge Restore (Preserve Non-Conflicting Data)</option>
            <option value="OVERWRITE_EXISTING">Overwrite Existing (Replace Conflicts)</option>
            <option value="RESTORE_AS_NEW_VAULT">Restore as New Vault Container</option>
            <option value="DRY_RUN">Dry Run (Preview & Validate Only)</option>
          </select>
        </div>
      </div>

      <div style={styles.fieldGroup}>
        <label style={styles.label}>Passphrase (If Encrypted)</label>
        <input
          type="password"
          value={passphrase}
          onChange={(e) => setPassphrase(e.target.value)}
          placeholder="Required if backup is encrypted..."
          style={styles.input}
        />
      </div>

      <button style={styles.btnDanger} onClick={handleRestore} disabled={loading}>
        {loading ? '⚡ Restoring Vault Data...' : '🚨 Execute Disaster Recovery Restore'}
      </button>

      {statusMsg && <div style={styles.statusSuccess}>{statusMsg}</div>}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '24px', display: 'flex', flexDirection: 'column', gap: '16px' },
  title: { fontSize: '18px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  desc: { fontSize: '13px', color: '#94a3b8', margin: 0 },
  fieldGroup: { display: 'flex', flexDirection: 'column', gap: '6px' },
  flexGroup: { display: 'flex', flexDirection: 'column', gap: '6px', flex: 1 },
  row: { display: 'flex', gap: '16px' },
  label: { fontSize: '12px', fontWeight: 600, color: '#94a3b8' },
  input: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  select: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  btnDanger: { background: 'linear-gradient(135deg, #ef4444, #f87171)', color: '#fff', border: 'none', borderRadius: '12px', padding: '12px 24px', fontSize: '14px', fontWeight: 700, cursor: 'pointer', alignSelf: 'flex-start' as const },
  statusSuccess: { background: '#4ade8015', border: '1px solid #4ade8044', color: '#4ade80', borderRadius: '10px', padding: '12px 16px', fontSize: '13px', fontWeight: 600 },
};
