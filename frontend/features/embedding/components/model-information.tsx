'use client';

import React, { useState } from 'react';
import type { ProviderName } from '../types/embedding';
import { EmbeddingService } from '../services/embedding-service';

const VAULT_ID = 'vault-demo';

interface ProviderInfo {
  name: ProviderName;
  label: string;
  defaultModel: string;
  dimensions: number;
  description: string;
  badgeColor: string;
}

const PROVIDERS: ProviderInfo[] = [
  { name: 'OPENAI', label: 'OpenAI', defaultModel: 'text-embedding-3-small', dimensions: 1536, description: 'High-performance, general-purpose dense vector embeddings.', badgeColor: '#10a37f' },
  { name: 'GEMINI', label: 'Google Gemini', defaultModel: 'text-embedding-004', dimensions: 768, description: 'Optimized multilingual multimodal embedding models.', badgeColor: '#4285f4' },
  { name: 'VOYAGE', label: 'Voyage AI', defaultModel: 'voyage-3-lite', dimensions: 1024, description: 'State-of-the-art domain-customized embeddings.', badgeColor: '#8b5cf6' },
  { name: 'COHERE', label: 'Cohere', defaultModel: 'embed-english-v3.0', dimensions: 1024, description: 'Compression-aware compressed vector representations.', badgeColor: '#d97706' },
  { name: 'OLLAMA', label: 'Ollama (Local)', defaultModel: 'nomic-embed-text', dimensions: 768, description: 'Privacy-first fully local open-weights vector models.', badgeColor: '#059669' },
  { name: 'LOCAL', label: 'Local Deterministic', defaultModel: 'local-deterministic-v1', dimensions: 1536, description: 'Zero-dependency local synthetic vector generator for testing.', badgeColor: '#6366f1' },
];

export function ModelInformation() {
  const [activeProvider, setActiveProvider] = useState<ProviderName>('OPENAI');
  const [msg, setMsg] = useState('');

  const handleSelect = async (p: ProviderInfo) => {
    setActiveProvider(p.name);
    await EmbeddingService.updateProvider(VAULT_ID, p.name, p.defaultModel, p.dimensions);
    setMsg(`Active provider updated to ${p.label} (${p.defaultModel})`);
    setTimeout(() => setMsg(''), 3000);
  };

  const handleRebuild = async (p: ProviderInfo) => {
    const tag = `v${Date.now().toString().slice(-4)}`;
    await EmbeddingService.rebuildVersion(VAULT_ID, p.name, p.defaultModel, tag);
    setMsg(`Rebuilt embeddings using ${p.label} (Version ${tag})`);
    setTimeout(() => setMsg(''), 3000);
  };

  return (
    <div style={styles.container}>
      <div style={styles.header}>
        <span style={styles.title}>🤖 Provider & Model Abstraction</span>
        <span style={styles.subtitle}>Switch active vector providers or rebuild embeddings for model upgrades.</span>
      </div>

      {msg && <div style={styles.toast}>{msg}</div>}

      <div style={styles.grid}>
        {PROVIDERS.map((p) => {
          const isSelected = activeProvider === p.name;
          return (
            <div key={p.name} style={{ ...styles.card, borderColor: isSelected ? p.badgeColor : '#334155' }}>
              <div style={styles.cardHeader}>
                <span style={{ ...styles.badge, background: p.badgeColor + '22', color: p.badgeColor, borderColor: p.badgeColor + '44' }}>
                  {p.label}
                </span>
                {isSelected && <span style={styles.activeText}>✓ Active</span>}
              </div>

              <div style={styles.modelName}>{p.defaultModel}</div>
              <div style={styles.dimText}>{p.dimensions} dimensions</div>
              <div style={styles.desc}>{p.description}</div>

              <div style={styles.btnRow}>
                <button
                  style={{ ...styles.btn, background: isSelected ? p.badgeColor : '#1e293b', color: '#fff' }}
                  onClick={() => handleSelect(p)}
                >
                  {isSelected ? 'Active Provider' : 'Select Provider'}
                </button>
                <button style={styles.btnOutline} onClick={() => handleRebuild(p)}>
                  Rebuild v
                </button>
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { padding: '24px', maxWidth: '960px', margin: '0 auto' },
  header: { marginBottom: '24px' },
  title: { fontSize: '20px', fontWeight: 700, color: '#e2e8f0', display: 'block', marginBottom: '4px' },
  subtitle: { fontSize: '13px', color: '#64748b' },
  toast: { background: '#0ea5e922', border: '1px solid #0ea5e944', color: '#38bdf8', borderRadius: '8px', padding: '10px 16px', fontSize: '13px', marginBottom: '20px' },
  grid: { display: 'grid', gridTemplateColumns: 'repeat(auto-fill, minmax(280px, 1fr))', gap: '16px' },
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px', display: 'flex', flexDirection: 'column', gap: '10px' },
  cardHeader: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  badge: { border: '1px solid', borderRadius: '20px', padding: '3px 10px', fontSize: '11px', fontWeight: 700 },
  activeText: { fontSize: '12px', color: '#4ade80', fontWeight: 700 },
  modelName: { fontSize: '15px', fontWeight: 700, color: '#e2e8f0' },
  dimText: { fontSize: '12px', color: '#818cf8', fontWeight: 600 },
  desc: { fontSize: '12px', color: '#94a3b8', lineHeight: 1.5, flex: 1 },
  btnRow: { display: 'flex', gap: '8px', marginTop: '8px' },
  btn: { flex: 1, border: 'none', borderRadius: '8px', padding: '8px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
  btnOutline: { background: 'transparent', border: '1px solid #334155', color: '#94a3b8', borderRadius: '8px', padding: '8px 12px', fontSize: '12px', fontWeight: 600, cursor: 'pointer' },
};
