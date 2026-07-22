'use client';

import React, { useState } from 'react';
import { ShareDialog } from '../../features/collaboration/components/share-dialog';
import { PermissionManager } from '../../features/collaboration/components/permission-manager';
import { CommentPanel } from '../../features/collaboration/components/comment-panel';
import { ActivityFeed } from '../../features/collaboration/components/activity-feed';

export default function CollaborationPage() {
  const [activeTab, setActiveTab] = useState<'DIALOG' | 'PERMISSIONS' | 'COMMENTS' | 'AUDIT'>('DIALOG');

  return (
    <div style={styles.page}>
      {/* Header */}
      <div style={styles.header}>
        <h1 style={styles.heading}>Secure Sharing & Collaboration Platform</h1>
        <p style={styles.subheading}>
          Granular, privacy-first, and auditable resource sharing across collections, assets, folders, contexts, and timeline views.
        </p>
      </div>

      {/* Main Container */}
      <div style={styles.container}>
        {/* Navigation Tabs */}
        <div style={styles.tabsRow}>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'DIALOG' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('DIALOG')}
          >
            🔒 Share Resource
          </button>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'PERMISSIONS' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('PERMISSIONS')}
          >
            👥 Permission Manager
          </button>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'COMMENTS' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('COMMENTS')}
          >
            💬 Comments & Threads
          </button>
          <button
            style={{ ...styles.tabBtn, ...(activeTab === 'AUDIT' ? styles.tabActive : {}) }}
            onClick={() => setActiveTab('AUDIT')}
          >
            📜 Audit Activity Feed
          </button>
        </div>

        {/* Tab Content */}
        {activeTab === 'DIALOG' && <ShareDialog onShareCreated={() => setActiveTab('PERMISSIONS')} />}
        {activeTab === 'PERMISSIONS' && <PermissionManager />}
        {activeTab === 'COMMENTS' && <CommentPanel />}
        {activeTab === 'AUDIT' && <ActivityFeed />}
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
