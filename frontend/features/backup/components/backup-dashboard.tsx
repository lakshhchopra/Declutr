'use client';

import React, { useState, useEffect } from 'react';
import type { BackupStats, ScheduleFrequency } from '../types/backup';
import { BackupService } from '../services/backup-service';

interface BackupDashboardProps {
  onBackupCreated?: () => void;
}

export function BackupDashboard({ onBackupCreated }: BackupDashboardProps) {
  const [stats, setStats] = useState<BackupStats | null>(null);
  const [passphrase, setPassphrase] = useState('');
  const [frequency, setFrequency] = useState<ScheduleFrequency>('WEEKLY');
  const [retentionDays, setRetentionDays] = useState(30);
  const [loading, setLoading] = useState(false);

  const loadStats = async () => {
    const st = await BackupService.getStats();
    setStats(st);
  };

  useEffect(() => {
    loadStats();
  }, []);

  const handleManualBackup = async () => {
    setLoading(true);
    await BackupService.createBackup('MANUAL', passphrase);
    setPassphrase('');
    setLoading(false);
    loadStats();
    if (onBackupCreated) onBackupCreated();
  };

  const handleSaveSchedule = async () => {
    await BackupService.configureSchedule(frequency, retentionDays, true);
    loadStats();
  };

  return (
    <div style={styles.container}>
      {/* Metrics Header */}
      {stats && (
        <div style={styles.metricsGrid}>
          <div style={styles.metricCard}>
            <div style={styles.metricVal}>{stats.totalBackups}</div>
            <div style={styles.metricLbl}>Total Vault Backups</div>
          </div>
          <div style={styles.metricCard}>
            <div style={{ ...styles.metricVal, color: '#4ade80' }}>
              {(stats.totalSizeBytes / (1024 * 1024)).toFixed(1)} MB
            </div>
            <div style={styles.metricLbl}>Total Backup Storage</div>
          </div>
          <div style={styles.metricCard}>
            <div style={{ ...styles.metricVal, color: '#38bdf8' }}>{stats.compressionRatioPct}%</div>
            <div style={styles.metricLbl}>Compression Savings</div>
          </div>
          <div style={styles.metricCard}>
            <div style={{ ...styles.metricVal, color: stats.lastVerifySuccess ? '#4ade80' : '#ef4444' }}>
              {stats.lastVerifySuccess ? 'VERIFIED' : 'FAILED'}
            </div>
            <div style={styles.metricLbl}>Integrity Status</div>
          </div>
        </div>
      )}

      {/* Manual Backup Trigger Box */}
      <div style={styles.card}>
        <h3 style={styles.cardTitle}>📦 Immediate Manual Snapshot Backup</h3>
        <p style={styles.cardDesc}>
          Create an immediate encrypted, full-vault disaster recovery snapshot payload.
        </p>

        <div style={styles.inputGroup}>
          <label style={styles.label}>Optional Encryption Passphrase</label>
          <input
            type="password"
            value={passphrase}
            onChange={(e) => setPassphrase(e.target.value)}
            placeholder="Enter AES-256 encryption passphrase..."
            style={styles.input}
          />
        </div>

        <button style={styles.btnPrimary} onClick={handleManualBackup} disabled={loading}>
          {loading ? '⏳ Capturing Vault Backup...' : '⚡ Trigger Manual Vault Backup'}
        </button>
      </div>

      {/* Schedule & Retention Policy Box */}
      <div style={styles.card}>
        <h3 style={styles.cardTitle}>🗓️ Automated Backup Scheduler & Retention Policy</h3>
        <div style={styles.row}>
          <div style={styles.flexGroup}>
            <label style={styles.label}>Automated Frequency</label>
            <select value={frequency} onChange={(e) => setFrequency(e.target.value as any)} style={styles.select}>
              <option value="DAILY">Daily Automated Backup</option>
              <option value="WEEKLY">Weekly Automated Backup</option>
              <option value="MONTHLY">Monthly Automated Backup</option>
            </select>
          </div>
          <div style={styles.flexGroup}>
            <label style={styles.label}>Retention Period (Days)</label>
            <input
              type="number"
              value={retentionDays}
              onChange={(e) => setRetentionDays(Number(e.target.value))}
              style={styles.input}
            />
          </div>
        </div>

        <button style={styles.btnSecondary} onClick={handleSaveSchedule}>
          💾 Save Backup Schedule Policy
        </button>
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { display: 'flex', flexDirection: 'column', gap: '20px' },
  metricsGrid: { display: 'grid', gridTemplateColumns: 'repeat(4, 1fr)', gap: '16px' },
  metricCard: { background: '#1e293b', border: '1px solid #334155', borderRadius: '14px', padding: '16px', textAlign: 'center' as const },
  metricVal: { fontSize: '24px', fontWeight: 800, color: '#6366f1' },
  metricLbl: { fontSize: '12px', color: '#94a3b8', marginTop: '4px' },
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '24px', display: 'flex', flexDirection: 'column', gap: '14px' },
  cardTitle: { fontSize: '18px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  cardDesc: { fontSize: '13px', color: '#94a3b8', margin: 0 },
  inputGroup: { display: 'flex', flexDirection: 'column', gap: '6px' },
  flexGroup: { display: 'flex', flexDirection: 'column', gap: '6px', flex: 1 },
  row: { display: 'flex', gap: '16px' },
  label: { fontSize: '12px', fontWeight: 600, color: '#94a3b8' },
  input: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  select: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  btnPrimary: { background: 'linear-gradient(135deg, #6366f1, #818cf8)', color: '#fff', border: 'none', borderRadius: '12px', padding: '12px 24px', fontSize: '14px', fontWeight: 700, cursor: 'pointer', alignSelf: 'flex-start' as const },
  btnSecondary: { background: '#0f172a', border: '1px solid #334155', color: '#38bdf8', borderRadius: '10px', padding: '10px 20px', fontSize: '13px', fontWeight: 700, cursor: 'pointer', alignSelf: 'flex-start' as const },
};
