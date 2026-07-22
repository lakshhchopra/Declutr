'use client';

import React, { useState, useEffect } from 'react';
import type { Workflow, WorkflowStats } from '../types/workflow';
import { WorkflowService } from '../services/workflow-service';

interface WorkflowDashboardProps {
  onEditWorkflow?: (wf: Workflow) => void;
  onCreateNew?: () => void;
}

export function WorkflowDashboard({ onCreateNew }: WorkflowDashboardProps) {
  const [workflows, setWorkflows] = useState<Workflow[]>([]);
  const [stats, setStats] = useState<WorkflowStats | null>(null);
  const [loading, setLoading] = useState(true);

  const loadData = async () => {
    setLoading(true);
    const [wfList, st] = await Promise.all([
      WorkflowService.getWorkflows(),
      WorkflowService.getStats(),
    ]);
    setWorkflows(wfList);
    setStats(st);
    setLoading(false);
  };

  useEffect(() => {
    loadData();
  }, []);

  const handleToggle = async (id: string, currentEnabled: boolean) => {
    await WorkflowService.toggleWorkflow(id, !currentEnabled);
    setWorkflows((prev) =>
      prev.map((w) => (w.workflowId === id ? { ...w, enabled: !currentEnabled } : w))
    );
  };

  const handleRunManual = async (id: string) => {
    await WorkflowService.runWorkflow(id);
    loadData();
  };

  return (
    <div style={styles.container}>
      {/* Metrics Row */}
      {stats && (
        <div style={styles.metricsGrid}>
          <div style={styles.metricCard}>
            <div style={styles.metricVal}>{stats.totalWorkflows}</div>
            <div style={styles.metricLbl}>Total Automations</div>
          </div>
          <div style={styles.metricCard}>
            <div style={{ ...styles.metricVal, color: '#4ade80' }}>{stats.activeWorkflows}</div>
            <div style={styles.metricLbl}>Active Rules</div>
          </div>
          <div style={styles.metricCard}>
            <div style={styles.metricVal}>{stats.totalRuns}</div>
            <div style={styles.metricLbl}>Total Executions</div>
          </div>
          <div style={styles.metricCard}>
            <div style={{ ...styles.metricVal, color: '#38bdf8' }}>{Math.round(stats.successRate * 100)}%</div>
            <div style={styles.metricLbl}>Success Rate</div>
          </div>
        </div>
      )}

      {/* Header & Create Button */}
      <div style={styles.sectionHeader}>
        <span style={styles.sectionTitle}>⚡ Active Workflow Rules</span>
        <button style={styles.createBtn} onClick={onCreateNew}>
          + Create Workflow Rule
        </button>
      </div>

      {/* Workflow List */}
      {loading ? (
        <div style={styles.loading}>Loading workflows...</div>
      ) : (
        <div style={styles.list}>
          {workflows.map((wf) => (
            <div key={wf.workflowId} style={styles.card}>
              <div style={styles.cardHeader}>
                <div style={styles.titleRow}>
                  <span style={styles.wfName}>{wf.name}</span>
                  <span style={{ ...styles.statusBadge, background: wf.enabled ? '#4ade8015' : '#64748b15', color: wf.enabled ? '#4ade80' : '#64748b' }}>
                    {wf.enabled ? 'ACTIVE' : 'DISABLED'}
                  </span>
                </div>
                <div style={styles.actionBtns}>
                  <button style={styles.runBtn} onClick={() => handleRunManual(wf.workflowId)}>
                    ▶ Run Now
                  </button>
                  <button
                    style={{ ...styles.toggleBtn, background: wf.enabled ? '#ef444422' : '#4ade8022', color: wf.enabled ? '#ef4444' : '#4ade80' }}
                    onClick={() => handleToggle(wf.workflowId, wf.enabled)}
                  >
                    {wf.enabled ? 'Disable' : 'Enable'}
                  </button>
                </div>
              </div>

              <div style={styles.description}>{wf.description}</div>

              {/* Triggers & Actions summary */}
              <div style={styles.stepsRow}>
                <div style={styles.stepBox}>
                  <span style={styles.stepLabel}>Trigger:</span>
                  <span style={styles.stepVal}>{wf.triggers[0]?.triggerType || 'ASSET_UPLOADED'}</span>
                </div>
                <span style={styles.arrow}>→</span>
                <div style={styles.stepBox}>
                  <span style={styles.stepLabel}>Conditions:</span>
                  <span style={styles.stepVal}>{wf.conditions.length} Rules</span>
                </div>
                <span style={styles.arrow}>→</span>
                <div style={styles.stepBox}>
                  <span style={styles.stepLabel}>Actions:</span>
                  <span style={styles.stepVal}>{wf.actions.map((a) => a.actionType).join(', ')}</span>
                </div>
              </div>

              <div style={styles.footerRow}>
                <span>Runs: {wf.runCount}</span>
                <span>Avg Duration: {wf.avgDurationMs}ms</span>
                {wf.lastRunAt && <span>Last run: {new Date(wf.lastRunAt).toLocaleTimeString()}</span>}
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { display: 'flex', flexDirection: 'column', gap: '20px' },
  metricsGrid: { display: 'grid', gridTemplateColumns: 'repeat(4, 1fr)', gap: '16px' },
  metricCard: { background: '#1e293b', border: '1px solid #334155', borderRadius: '14px', padding: '16px', textAlign: 'center' as const },
  metricVal: { fontSize: '24px', fontWeight: 800, color: '#6366f1' },
  metricLbl: { fontSize: '12px', color: '#94a3b8', marginTop: '4px' },
  sectionHeader: { display: 'flex', justifyContent: 'space-between', alignItems: 'center', marginTop: '8px' },
  sectionTitle: { fontSize: '16px', fontWeight: 700, color: '#e2e8f0' },
  createBtn: { background: 'linear-gradient(135deg, #6366f1, #818cf8)', color: '#fff', border: 'none', borderRadius: '10px', padding: '8px 18px', fontSize: '13px', fontWeight: 700, cursor: 'pointer' },
  loading: { textAlign: 'center', padding: '40px', color: '#94a3b8' },
  list: { display: 'flex', flexDirection: 'column', gap: '16px' },
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '20px', display: 'flex', flexDirection: 'column', gap: '12px' },
  cardHeader: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  titleRow: { display: 'flex', alignItems: 'center', gap: '10px' },
  wfName: { fontSize: '16px', fontWeight: 700, color: '#e2e8f0' },
  statusBadge: { borderRadius: '6px', padding: '2px 8px', fontSize: '10px', fontWeight: 800, border: '1px solid #ffffff15' },
  actionBtns: { display: 'flex', gap: '8px' },
  runBtn: { background: '#0f172a', border: '1px solid #334155', color: '#38bdf8', borderRadius: '8px', padding: '6px 12px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
  toggleBtn: { border: 'none', borderRadius: '8px', padding: '6px 12px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
  description: { fontSize: '13px', color: '#94a3b8', lineHeight: 1.4 },
  stepsRow: { display: 'flex', alignItems: 'center', gap: '12px', background: '#0f172a', borderRadius: '10px', padding: '10px 14px', border: '1px solid #334155' },
  stepBox: { display: 'flex', gap: '6px', fontSize: '12px' },
  stepLabel: { color: '#64748b', fontWeight: 600 },
  stepVal: { color: '#e2e8f0', fontWeight: 700 },
  arrow: { color: '#6366f1', fontWeight: 700 },
  footerRow: { display: 'flex', gap: '20px', fontSize: '11px', color: '#64748b' },
};
