'use client';

import React, { useState } from 'react';
import type { ResourceType, AccessType } from '../types/collaboration';
import { CollaborationService } from '../services/collaboration-service';

interface ShareDialogProps {
  onShareCreated?: () => void;
}

export function ShareDialog({ onShareCreated }: ShareDialogProps) {
  const [title, setTitle] = useState('');
  const [resourceType, setResourceType] = useState<ResourceType>('COLLECTION');
  const [resourceId, setResourceId] = useState('col-demo-123');
  const [accessType, setAccessType] = useState<AccessType>('INVITE_ONLY');

  const handleCreate = async () => {
    if (!title.trim()) return;
    await CollaborationService.createShare(title, resourceType, resourceId, accessType);
    if (onShareCreated) onShareCreated();
  };

  return (
    <div style={styles.card}>
      <h3 style={styles.title}>🔒 Secure Share Resource</h3>

      <div style={styles.fieldGroup}>
        <label style={styles.label}>Resource Title</label>
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="e.g. Japan Vacation Collection"
          style={styles.input}
        />
      </div>

      <div style={styles.row}>
        <div style={styles.fieldGroupFlex}>
          <label style={styles.label}>Resource Type</label>
          <select value={resourceType} onChange={(e) => setResourceType(e.target.value as any)} style={styles.select}>
            <option value="ASSET">Asset (File/PDF)</option>
            <option value="FOLDER">Folder</option>
            <option value="COLLECTION">Collection</option>
            <option value="CONTEXT">Context Stream</option>
            <option value="PROJECT">Project Workspace</option>
            <option value="TIMELINE_VIEW">Timeline View</option>
          </select>
        </div>

        <div style={styles.fieldGroupFlex}>
          <label style={styles.label}>Sharing Privacy Mode</label>
          <select value={accessType} onChange={(e) => setAccessType(e.target.value as any)} style={styles.select}>
            <option value="PRIVATE">Private (Vault Owner Only)</option>
            <option value="INVITE_ONLY">Invite Only (Explicit Members)</option>
            <option value="LINK_SHARING">Link Sharing (Protected Link)</option>
            <option value="TEMPORARY_ACCESS">Temporary Access (Expires)</option>
          </select>
        </div>
      </div>

      <div style={styles.fieldGroup}>
        <label style={styles.label}>Resource ID</label>
        <input type="text" value={resourceId} onChange={(e) => setResourceId(e.target.value)} style={styles.input} />
      </div>

      <button style={styles.saveBtn} onClick={handleCreate}>
        ✨ Create Shared Resource Container
      </button>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '24px', display: 'flex', flexDirection: 'column', gap: '18px' },
  title: { fontSize: '18px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  fieldGroup: { display: 'flex', flexDirection: 'column', gap: '6px' },
  fieldGroupFlex: { display: 'flex', flexDirection: 'column', gap: '6px', flex: 1 },
  row: { display: 'flex', gap: '16px' },
  label: { fontSize: '12px', fontWeight: 600, color: '#94a3b8' },
  input: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  select: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '10px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  saveBtn: { background: 'linear-gradient(135deg, #6366f1, #818cf8)', color: '#fff', border: 'none', borderRadius: '12px', padding: '12px 24px', fontSize: '14px', fontWeight: 700, cursor: 'pointer', alignSelf: 'flex-start' as const },
};
