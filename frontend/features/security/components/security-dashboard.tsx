'use client';

import React, { useState, useEffect } from 'react';
import type { SecurityDashboard as DashboardType } from '../types/security';
import { SecurityService } from '../services/security-service';

export function SecurityDashboardComponent() {
  const [data, setData] = useState<DashboardType | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    SecurityService.getDashboard().then((dash) => {
      setData(dash);
      setLoading(false);
    });
  }, []);

  if (loading) return <div style={styles.loading}>Loading Trust Center Dashboard...</div>;
  if (!data) return <div style={styles.empty}>Dashboard unavailable.</div>;

  return (
    <div style={styles.container}>
      {/* Top Banner Grid */}
      <div style={styles.scoreGrid}>
        <div style={styles.scoreCard}>
          <div style={styles.radialScore}>
            <span style={styles.scoreNum}>{data.score.score}</span>
            <span style={styles.scoreMax}>/100</span>
          </div>
          <div style={styles.scoreDetails}>
            <div style={styles.gradeRow}>
              <span style={styles.gradeBadge}>GRADE {data.score.grade}</span>
              <span style={styles.statusBadge}>{data.score.status}</span>
            </div>
            <p style={styles.scoreDesc}>Vault Security Posture Score</p>
          </div>
        </div>

        <div style={styles.postureCard}>
          <h4 style={styles.postureTitle}>🛡️ Trust & Protection Badges</h4>
          <div style={styles.badgeList}>
            <div style={styles.badgeItem}>
              <span>🔒 Zero-Knowledge Encryption</span>
              <span style={styles.badgeVal}>{data.encryptedVault ? 'ACTIVE' : 'OFF'}</span>
            </div>
            <div style={styles.badgeItem}>
              <span>🔑 Multi-Factor Authentication (MFA)</span>
              <span style={{ ...styles.badgeVal, color: data.mfaEnabled ? '#4ade80' : '#f59e0b' }}>
                {data.mfaEnabled ? 'ENABLED' : 'NOT CONFIGURED'}
              </span>
            </div>
            <div style={styles.badgeItem}>
              <span>📦 Disaster Recovery Backups</span>
              <span style={styles.badgeVal}>{data.backupStatus}</span>
            </div>
          </div>
        </div>
      </div>

      {/* Risk Engine Box */}
      <div style={styles.card}>
        <div style={styles.cardHeader}>
          <h3 style={styles.cardTitle}>⚠️ Risk Engine Signal Assessment</h3>
          <span style={styles.riskBadge}>RISK: {data.risk.riskLevel} ({data.risk.riskScore}/100)</span>
        </div>
        <div style={styles.signalList}>
          {data.risk.signals.map((sig) => (
            <div key={sig.signalId} style={styles.signalCard}>
              <span style={styles.sigType}>{sig.signalType}</span>
              <span style={styles.sigDesc}>{sig.description}</span>
              <span style={styles.sigWeight}>Weight +{sig.weight}</span>
            </div>
          ))}
        </div>
      </div>

      {/* Recommendations Box */}
      <div style={styles.card}>
        <h3 style={styles.cardTitle}>✨ Actionable Security Recommendations</h3>
        <div style={styles.recList}>
          {data.recommendations.map((rec) => (
            <div key={rec.recId} style={styles.recCard}>
              <div style={styles.recHeader}>
                <span style={styles.recPriority}>{rec.priority} PRIORITY</span>
                <span style={styles.recCat}>{rec.category}</span>
              </div>
              <h4 style={styles.recTitle}>{rec.title}</h4>
              <p style={styles.recDesc}>{rec.description}</p>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { display: 'flex', flexDirection: 'column', gap: '20px' },
  loading: { textAlign: 'center', padding: '40px', color: '#94a3b8' },
  empty: { textAlign: 'center', padding: '40px', color: '#64748b' },
  scoreGrid: { display: 'grid', gridTemplateColumns: '1fr 1fr', gap: '20px' },
  scoreCard: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '24px', display: 'flex', alignItems: 'center', gap: '24px' },
  radialScore: { background: '#0f172a', border: '3px solid #6366f1', borderRadius: '50%', width: '90px', height: '90px', display: 'flex', flexDirection: 'column', justifyContent: 'center', alignItems: 'center' },
  scoreNum: { fontSize: '32px', fontWeight: 900, color: '#818cf8', lineHeight: 1 },
  scoreMax: { fontSize: '11px', color: '#64748b' },
  scoreDetails: { display: 'flex', flexDirection: 'column', gap: '6px' },
  gradeRow: { display: 'flex', alignItems: 'center', gap: '10px' },
  gradeBadge: { background: '#4ade8015', color: '#4ade80', borderRadius: '8px', padding: '4px 10px', fontSize: '13px', fontWeight: 900, border: '1px solid #4ade8033' },
  statusBadge: { fontSize: '12px', fontWeight: 800, color: '#38bdf8' },
  scoreDesc: { fontSize: '13px', color: '#94a3b8', margin: 0 },
  postureCard: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '24px', display: 'flex', flexDirection: 'column', gap: '14px' },
  postureTitle: { fontSize: '16px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  badgeList: { display: 'flex', flexDirection: 'column', gap: '10px' },
  badgeItem: { display: 'flex', justifyContent: 'space-between', fontSize: '13px', color: '#cbd5e1' },
  badgeVal: { fontWeight: 700, color: '#4ade80' },
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '24px', display: 'flex', flexDirection: 'column', gap: '14px' },
  cardHeader: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  cardTitle: { fontSize: '18px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  riskBadge: { background: '#38bdf815', color: '#38bdf8', borderRadius: '8px', padding: '4px 12px', fontSize: '12px', fontWeight: 800, border: '1px solid #38bdf833' },
  signalList: { display: 'flex', flexDirection: 'column', gap: '10px' },
  signalCard: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '12px 14px', display: 'flex', alignItems: 'center', gap: '12px' },
  sigType: { fontSize: '11px', fontWeight: 800, color: '#f59e0b', background: '#f59e0b15', padding: '2px 8px', borderRadius: '6px' },
  sigDesc: { flex: 1, fontSize: '13px', color: '#e2e8f0' },
  sigWeight: { fontSize: '11px', color: '#64748b' },
  recList: { display: 'flex', flexDirection: 'column', gap: '12px' },
  recCard: { background: '#0f172a', border: '1px solid #334155', borderRadius: '12px', padding: '16px', display: 'flex', flexDirection: 'column', gap: '6px' },
  recHeader: { display: 'flex', gap: '10px', alignItems: 'center' },
  recPriority: { fontSize: '10px', fontWeight: 900, color: '#ef4444', background: '#ef444415', padding: '2px 8px', borderRadius: '6px' },
  recCat: { fontSize: '11px', color: '#64748b' },
  recTitle: { fontSize: '15px', fontWeight: 700, color: '#e2e8f0', margin: 0 },
  recDesc: { fontSize: '13px', color: '#94a3b8', margin: 0, lineHeight: 1.4 },
};
