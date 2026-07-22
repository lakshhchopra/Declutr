'use client';

import React, { useState } from 'react';
import { SecurityDashboardComponent } from '../../features/security/components/security-dashboard';
import { AuditViewerComponent } from '../../features/security/components/audit-viewer';
import { SessionDeviceManagerComponent } from '../../features/security/components/session-device-manager';

export default function SecurityPage() {
  const [activeTab, setActiveTab] = useState<'DASHBOARD' | 'AUDIT' | 'SESSIONS'>('DASHBOARD');

  return (
    <div style={styles.page}>
      {/* Header */}
      <div style={styles.header}>
        <h1 style={styles.heading}>Security Center, Audit Hub & Trust Platform</h1>
        <p style={styles.subheading}>
          Central trust center for security posture monitoring, risk signal analysis, active session management, trusted device registry, and comprehensive vault audit history.
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
            🛡️ Security Overview & Posture
          </button>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'AUDIT' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('AUDIT')}
          >
            📜 Audit Log Hub
          </button>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'SESSIONS' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('SESSIONS')}
          >
            💻 Active Sessions & Devices
          </button>
        </div>

        {/* Tab Content */}
        {activeTab === 'DASHBOARD' && <SecurityDashboardComponent />}
        {activeTab === 'AUDIT' && <AuditViewerComponent />}
        {activeTab === 'SESSIONS' && <SessionDeviceManagerComponent />}
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
