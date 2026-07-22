'use client';

import React, { useState, useEffect } from 'react';
import type { Share, MemberRole } from '../types/collaboration';
import { CollaborationService } from '../services/collaboration-service';

export function PermissionManager() {
  const [shares, setShares] = useState<Share[]>([]);
  const [inviteEmail, setInviteEmail] = useState('');
  const [selectedRole, setSelectedRole] = useState<MemberRole>('READ_ONLY');
  const [selectedShareId, setSelectedShareId] = useState('');

  const loadData = async () => {
    const list = await CollaborationService.getShares();
    setShares(list);
    if (list.length > 0 && !selectedShareId) {
      setSelectedShareId(list[0].shareId);
    }
  };

  useEffect(() => {
    loadData();
  }, []);

  const handleInvite = async () => {
    if (!inviteEmail.trim() || !selectedShareId) return;
    await CollaborationService.inviteUser(selectedShareId, inviteEmail, selectedRole);
    setInviteEmail('');
    loadData();
  };

  const selectedShare = shares.find((s) => s.shareId === selectedShareId);

  return (
    <div style={styles.container}>
      {/* Share Selector */}
      <div style={styles.selectorRow}>
        <span style={styles.label}>Select Shared Resource:</span>
        <select value={selectedShareId} onChange={(e) => setSelectedShareId(e.target.value)} style={styles.select}>
          {shares.map((s) => (
            <option key={s.shareId} value={s.shareId}>
              {s.title} ({s.resourceType})
            </option>
          ))}
        </select>
      </div>

      {/* Invite Member Box */}
      <div style={styles.box}>
        <h4 style={styles.boxTitle}>✉️ Invite Member</h4>
        <div style={styles.inviteRow}>
          <input
            type="email"
            value={inviteEmail}
            onChange={(e) => setInviteEmail(e.target.value)}
            placeholder="colleague@domain.com"
            style={styles.input}
          />
          <select value={selectedRole} onChange={(e) => setSelectedRole(e.target.value as any)} style={styles.select}>
            <option value="READ_ONLY">Read Only (Viewer)</option>
            <option value="COMMENT_ONLY">Comment Only</option>
            <option value="EDIT">Editor (View/Edit/Comment)</option>
            <option value="CO_OWNER">Co-Owner (Full Control)</option>
          </select>
          <button style={styles.btn} onClick={handleInvite}>
            Send Invite
          </button>
        </div>
      </div>

      {/* Members List */}
      {selectedShare && (
        <div style={styles.box}>
          <h4 style={styles.boxTitle}>👥 Members & Roles ({selectedShare.members?.length || 0})</h4>
          <div style={styles.memberList}>
            {selectedShare.members?.map((m) => (
              <div key={m.memberId} style={styles.memberCard}>
                <div style={styles.memInfo}>
                  <span style={styles.memEmail}>{m.email}</span>
                  <span style={styles.memRole}>{m.role}</span>
                </div>
                <span style={styles.memDate}>Joined: {new Date(m.joinedAt).toLocaleDateString()}</span>
              </div>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { display: 'flex', flexDirection: 'column', gap: '20px' },
  selectorRow: { display: 'flex', alignItems: 'center', gap: '12px' },
  label: { fontSize: '13px', fontWeight: 600, color: '#94a3b8' },
  select: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '8px 12px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  box: { background: '#1e293b', border: '1px solid #334155', borderRadius: '14px', padding: '20px', display: 'flex', flexDirection: 'column', gap: '14px' },
  boxTitle: { fontSize: '15px', fontWeight: 700, color: '#e2e8f0', margin: 0 },
  inviteRow: { display: 'flex', gap: '12px' },
  input: { flex: 1, background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '8px 14px', color: '#e2e8f0', fontSize: '13px', outline: 'none' },
  btn: { background: 'linear-gradient(135deg, #6366f1, #818cf8)', color: '#fff', border: 'none', borderRadius: '10px', padding: '8px 18px', fontSize: '13px', fontWeight: 700, cursor: 'pointer' },
  memberList: { display: 'flex', flexDirection: 'column', gap: '10px' },
  memberCard: { background: '#0f172a', border: '1px solid #334155', borderRadius: '10px', padding: '12px 14px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  memInfo: { display: 'flex', alignItems: 'center', gap: '10px' },
  memEmail: { fontSize: '14px', fontWeight: 600, color: '#e2e8f0' },
  memRole: { fontSize: '11px', fontWeight: 800, color: '#38bdf8', background: '#38bdf815', padding: '2px 8px', borderRadius: '6px', border: '1px solid #38bdf833' },
  memDate: { fontSize: '11px', color: '#64748b' },
};
