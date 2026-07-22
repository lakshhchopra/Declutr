'use client';

import React, { useState } from 'react';
import { EmbeddingDashboard } from '../../features/embedding/components/embedding-dashboard';
import { EmbeddingStatus } from '../../features/embedding/components/embedding-status';
import { ModelInformation } from '../../features/embedding/components/model-information';
import { GenerationHistory } from '../../features/embedding/components/generation-history';

type Tab = 'dashboard' | 'status' | 'models' | 'history';

const TABS: { id: Tab; label: string; icon: string }[] = [
  { id: 'dashboard', label: 'Embedding Engine', icon: '💎' },
  { id: 'status', label: 'Status & Health', icon: '🟢' },
  { id: 'models', label: 'Providers & Models', icon: '🤖' },
  { id: 'history', label: 'Generation History', icon: '📜' },
];

export default function EmbeddingPage() {
  const [activeTab, setActiveTab] = useState<Tab>('dashboard');

  return (
    <div style={styles.page}>
      {/* Header */}
      <div style={styles.header}>
        <h1 style={styles.heading}>Embedding Engine</h1>
        <p style={styles.subheading}>
          Declutr's Knowledge Representation Layer — transforming structured memories, contexts, entities, and documents into high-dimensional vector representations.
        </p>
      </div>

      {/* Tabs */}
      <div style={styles.tabs}>
        {TABS.map((tab) => (
          <button
            key={tab.id}
            style={{ ...styles.tab, ...(activeTab === tab.id ? styles.tabActive : {}) }}
            onClick={() => setActiveTab(tab.id)}
          >
            {tab.icon} {tab.label}
          </button>
        ))}
      </div>

      {/* Tab Content */}
      <div style={styles.content}>
        {activeTab === 'dashboard' && <EmbeddingDashboard />}
        {activeTab === 'status' && <EmbeddingStatus />}
        {activeTab === 'models' && <ModelInformation />}
        {activeTab === 'history' && <GenerationHistory />}
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  page: { minHeight: '100vh', background: '#0f172a', color: '#e2e8f0', fontFamily: 'Inter, system-ui, sans-serif' },
  header: { padding: '32px 24px 0', maxWidth: '960px', margin: '0 auto' },
  heading: { fontSize: '32px', fontWeight: 800, color: '#e0e7ff', marginBottom: '8px', background: 'linear-gradient(135deg, #10a37f 0%, #38bdf8 100%)', WebkitBackgroundClip: 'text', WebkitTextFillColor: 'transparent' },
  subheading: { fontSize: '15px', color: '#64748b', margin: 0, lineHeight: 1.5 },
  tabs: { display: 'flex', gap: '4px', padding: '24px 24px 0', maxWidth: '960px', margin: '0 auto', borderBottom: '1px solid #1e293b' },
  tab: { background: 'transparent', border: 'none', borderRadius: '8px 8px 0 0', padding: '10px 20px', fontSize: '13px', fontWeight: 600, color: '#64748b', cursor: 'pointer', transition: 'all 0.15s' },
  tabActive: { background: '#1e293b', color: '#e2e8f0', borderBottom: '2px solid #10a37f' },
  content: { maxWidth: '960px', margin: '0 auto' },
};
