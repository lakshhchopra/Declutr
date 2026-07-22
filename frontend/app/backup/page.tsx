'use client';

import React, { useState } from 'react';
import { BackupDashboard } from '../../features/backup/components/backup-dashboard';
import { RestoreWizard } from '../../features/backup/components/restore-wizard';
import { BackupHistory } from '../../features/backup/components/backup-history';

export default function BackupPage() {
  const [activeTab, setActiveTab] = useState<'DASHBOARD' | 'RESTORE' | 'HISTORY'>('DASHBOARD');
  const [selectedBackupId, setSelectedBackupId] = useState<string | undefined>(undefined);

  const handleSelectRestore = (backupId: string) => {
    setSelectedBackupId(backupId);
    setActiveTab('RESTORE');
  };

  return (
    <div style={styles.page}>
      {/* Header */}
      <div style={styles.header}>
        <h1 style={styles.heading}>Backup, Disaster Recovery & Business Continuity</h1>
        <p style={styles.subheading}>
          Protect your entire digital vault with encrypted full/incremental snapshots, automated retention schedules, SHA-256 integrity verification, and catastrophe recovery.
        </p>
      </div>

      {/* Main Container */}
      <div style={styles.container}>
        {/* Navigation Tabs */}
        <div style={styles.tabsRow}>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'DASHBOARD' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('DASHBOARD')}
          >
            📦 Backup Dashboard
          </button>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'RESTORE' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('RESTORE')}
          >
            🚨 Disaster Recovery Wizard
          </button>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'HISTORY' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('HISTORY')}
          >
            📜 Backup History & Integrity
          </button>
        </div>

        {/* Tab Content */}
        {activeTab === 'DASHBOARD' && <BackupDashboard onBackupCreated={() => setActiveTab('HISTORY')} />}
        {activeTab === 'RESTORE' && <RestoreWizard selectedBackupId={selectedBackupId} />}
        {activeTab === 'HISTORY' && <BackupHistory onSelectRestore={handleSelectRestore} />}
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  page: { minHeight: '100vh', background: '#0f172a', color: '#e2e8f0', fontFamily: 'Inter, system-ui, sans-serif', paddingBottom: '40px' },
  header: { padding: '32px 24px 0', maxWidth: '1080px', margin: '0 auto' },
  heading: { fontSize: '30px', fontWeight: 800, color: '#e0e7ff', marginBottom: '8px', background: 'linear-gradient(135deg, #6366f1 0%, #38bdf8 100%)', WebkitBackgroundClip: 'text', WebkitTextFillColor: 'transparent' },
  subheading: { fontSize: '14px', color: '#64748b', margin: 0, lineHeight: 1.5 },
  container: { maxWidth: '1080px', margin: '24px auto 0', padding: '0 24px' },
  tabsRow: { display: 'flex', gap: '12px', borderBottom: '1px solid #334155', paddingBottom: '12px', marginBottom: '20px' },
  tabBtn: { background: 'transparent', border: 'none', color: '#64748b', fontSize: '14px', fontWeight: 700, cursor: 'pointer', padding: '6px 12px', borderRadius: '8px' },
  tabActive: { background: '#6366f122', color: '#818cf8' },
};
