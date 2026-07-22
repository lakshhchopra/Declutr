'use client';

import React, { useState } from 'react';
import type { TriggerType, ActionType } from '../types/workflow';
import { WorkflowService } from '../services/workflow-service';

interface VisualRuleBuilderProps {
  onSaveComplete?: () => void;
}

const TRIGGERS: { label: string; value: TriggerType }[] = [
  { label: '📄 Asset Uploaded', value: 'ASSET_UPLOADED' },
  { label: '✏️ Asset Updated', value: 'ASSET_UPDATED' },
  { label: '⚠️ Document Expiring Soon', value: 'DOCUMENT_EXPIRING' },
  { label: '🧠 Memory Created', value: 'MEMORY_CREATED' },
  { label: '📅 Daily Schedule', value: 'DAILY_SCHEDULE' },
  { label: '⚡ Manual Trigger', value: 'MANUAL_TRIGGER' },
];

const ACTIONS: { label: string; value: ActionType }[] = [
  { label: '🏷️ Apply Tags', value: 'APPLY_TAGS' },
  { label: '📁 Create Collection', value: 'CREATE_COLLECTION' },
  { label: '📦 Archive Asset', value: 'ARCHIVE_ASSET' },
  { label: '⏰ Create Reminder', value: 'CREATE_REMINDER' },
  { label: '📌 Pin Memory', value: 'PIN_MEMORY' },
  { label: '🔔 Notify User', value: 'NOTIFY_USER' },
];

export function VisualRuleBuilder({ onSaveComplete }: VisualRuleBuilderProps) {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [selectedTrigger, setSelectedTrigger] = useState<TriggerType>('ASSET_UPLOADED');
  const [selectedAction, setSelectedAction] = useState<ActionType>('APPLY_TAGS');
  const [field, setField] = useState('fileType');
  const [value, setValue] = useState('PDF');

  const handleSave = async () => {
    if (!name.trim()) return;
    await WorkflowService.createWorkflow({
      name,
      description,
      triggers: [{ triggerId: 't-new', workflowId: '', triggerType: selectedTrigger, createdAt: new Date().toISOString() }],
      conditions: [{ conditionId: 'c-new', workflowId: '', field, operator: 'EQUALS', value, combinator: 'AND', createdAt: new Date().toISOString() }],
      actions: [{ actionId: 'a-new', workflowId: '', actionType: selectedAction, executionOrder: 1, createdAt: new Date().toISOString() }],
    });
    if (onSaveComplete) onSaveComplete();
  };

  return (
    <div style={styles.card}>
      <h3 style={styles.title}>🛠️ Visual Rule Builder</h3>

      {/* Basic Info */}
      <div style={styles.fieldGroup}>
        <label style={styles.label}>Workflow Name</label>
        <input
          type="text"
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="e.g. Auto-tag Tax Receipts"
          style={styles.input}
        />
      </div>

      <div style={styles.fieldGroup}>
        <label style={styles.label}>Description</label>
        <input
          type="text"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
          placeholder="Optional explanation..."
          style={styles.input}
        />
      </div>

      {/* Trigger Step */}
      <div style={styles.stepBox}>
        <div style={styles.stepHeader}>1️⃣ Select Trigger Event</div>
        <div style={styles.chipsRow}>
          {TRIGGERS.map((t) => (
            <button
              key={t.value}
              style={{ ...styles.chip, ...(selectedTrigger === t.value ? styles.chipActive : {}) }}
              onClick={() => setSelectedTrigger(t.value)}
            >
              {t.label}
            </button>
          ))}
        </div>
      </div>

      {/* Condition Step */}
      <div style={styles.stepBox}>
        <div style={styles.stepHeader}>2️⃣ Set Filter Condition (If...)</div>
        <div style={styles.ruleRow}>
          <input
            type="text"
            value={field}
            onChange={(e) => setField(e.target.value)}
            placeholder="Field (e.g. fileType, entity)"
            style={{ ...styles.input, width: '160px' }}
          />
          <span style={styles.operatorText}>EQUALS</span>
          <input
            type="text"
            value={value}
            onChange={(e) => setValue(e.target.value)}
            placeholder="Value (e.g. PDF, Japan)"
            style={{ ...styles.input, width: '160px' }}
          />
        </div>
      </div>

      {/* Action Step */}
      <div style={styles.stepBox}>
        <div style={styles.stepHeader}>3️⃣ Select Action Step (Then...)</div>
        <div style={styles.chipsRow}>
          {ACTIONS.map((a) => (
            <button
              key={a.value}
              style={{ ...styles.chip, ...(selectedAction === a.value ? styles.chipActive : {}) }}
              onClick={() => setSelectedAction(a.value)}
            >
              {a.label}
            </button>
          ))}
        </div>
      </div>

      <button style={styles.saveBtn} onClick={handleSave}>
        💾 Save Automation Rule
      </button>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '24px', display: 'flex', flexDirection: 'column', gap: '20px' },
  title: { fontSize: '18px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  fieldGroup: { display: 'flex', flexDirection: 'column', gap: '6px' },
  label: { fontSize: '12px', fontWeight: 600, color: '#94a3b8' },
  input: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  stepBox: { background: '#0f172a', border: '1px solid #334155', borderRadius: '12px', padding: '16px', display: 'flex', flexDirection: 'column', gap: '12px' },
  stepHeader: { fontSize: '13px', fontWeight: 700, color: '#6366f1' },
  chipsRow: { display: 'flex', gap: '8px', flexWrap: 'wrap' as const },
  chip: { background: '#1e293b', border: '1px solid #334155', color: '#94a3b8', borderRadius: '10px', padding: '8px 14px', fontSize: '12px', fontWeight: 600, cursor: 'pointer' },
  chipActive: { background: '#6366f122', borderColor: '#6366f1', color: '#818cf8' },
  ruleRow: { display: 'flex', alignItems: 'center', gap: '12px' },
  operatorText: { fontSize: '12px', fontWeight: 800, color: '#4ade80' },
  saveBtn: { background: 'linear-gradient(135deg, #6366f1, #818cf8)', color: '#fff', border: 'none', borderRadius: '12px', padding: '12px 24px', fontSize: '14px', fontWeight: 700, cursor: 'pointer', alignSelf: 'flex-start' as const },
};
