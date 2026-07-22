'use client';

import React, { useState, useEffect } from 'react';
import type { ActiveSession, Device } from '../types/security';
import { SecurityService } from '../services/security-service';

export function SessionDeviceManagerComponent() {
  const [sessions, setSessions] = useState<ActiveSession[]>([]);
  const [devices, setDevices] = useState<Device[]>([]);

  const loadData = async () => {
    const [sessList, devList] = await Promise.all([
      SecurityService.getSessions(),
      SecurityService.getDevices(),
    ]);
    setSessions(sessList);
    setDevices(devList);
  };

  useEffect(() => {
    loadData();
  }, []);

  const handleTerminateSession = async (sessionId?: string) => {
    await SecurityService.terminateSession(sessionId);
    loadData();
  };

  const handleToggleTrust = async (deviceId: string, currentTrust: boolean) => {
    await SecurityService.setDeviceTrust(deviceId, !currentTrust);
    loadData();
  };

  return (
    <div style={styles.container}>
      {/* Active Sessions Box */}
      <div style={styles.card}>
        <div style={styles.cardHeader}>
          <h3 style={styles.title}>💻 Active User Sessions ({sessions.length})</h3>
          <button style={styles.btnDanger} onClick={() => handleTerminateSession()}>
            🚫 Terminate All Other Sessions
          </button>
        </div>

        <div style={styles.list}>
          {sessions.map((sess) => (
            <div key={sess.sessionId} style={styles.itemCard}>
              <div style={styles.itemInfo}>
                <div style={styles.itemTitleRow}>
                  <span style={styles.itemName}>{sess.deviceName} ({sess.browser})</span>
                  {sess.isCurrent && <span style={styles.currentBadge}>CURRENT SESSION</span>}
                </div>
                <span style={styles.itemSub}>
                  IP: {sess.ipAddress} • {sess.location} • Last active: {new Date(sess.lastSeenAt).toLocaleTimeString()}
                </span>
              </div>
              {!sess.isCurrent && (
                <button style={styles.btnSmDanger} onClick={() => handleTerminateSession(sess.sessionId)}>
                  Terminate
                </button>
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Device Registry Box */}
      <div style={styles.card}>
        <h3 style={styles.title}>📱 Registered & Trusted Devices ({devices.length})</h3>

        <div style={styles.list}>
          {devices.map((dev) => (
            <div key={dev.deviceId} style={styles.itemCard}>
              <div style={styles.itemInfo}>
                <div style={styles.itemTitleRow}>
                  <span style={styles.itemName}>{dev.deviceName}</span>
                  <span style={dev.isTrusted ? styles.trustedBadge : styles.untrustedBadge}>
                    {dev.isTrusted ? 'TRUSTED DEVICE' : 'UNTRUSTED'}
                  </span>
                </div>
                <span style={styles.itemSub}>
                  {dev.os} • {dev.platform} • IP: {dev.ipAddress} • Last seen: {new Date(dev.lastSeenAt).toLocaleDateString()}
                </span>
              </div>

              <button
                style={dev.isTrusted ? styles.btnSmOutline : styles.btnSmSuccess}
                onClick={() => handleToggleTrust(dev.deviceId, dev.isTrusted)}
              >
                {dev.isTrusted ? 'Revoke Trust' : '✨ Trust Device'}
              </button>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
}

const styles: Record<string, React.CSSProperties> = {
  container: { display: 'flex', flexDirection: 'column', gap: '20px' },
  card: { background: '#1e293b', border: '1px solid #334155', borderRadius: '16px', padding: '24px', display: 'flex', flexDirection: 'column', gap: '16px' },
  cardHeader: { display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  title: { fontSize: '18px', fontWeight: 800, color: '#e2e8f0', margin: 0 },
  list: { display: 'flex', flexDirection: 'column', gap: '10px' },
  itemCard: { background: '#0f172a', border: '1px solid #334155', borderRadius: '12px', padding: '14px 16px', display: 'flex', justifyContent: 'space-between', alignItems: 'center' },
  itemInfo: { display: 'flex', flexDirection: 'column', gap: '4px' },
  itemTitleRow: { display: 'flex', alignItems: 'center', gap: '10px' },
  itemName: { fontSize: '14px', fontWeight: 700, color: '#e2e8f0' },
  currentBadge: { background: '#4ade8015', color: '#4ade80', borderRadius: '6px', padding: '2px 8px', fontSize: '10px', fontWeight: 900, border: '1px solid #4ade8033' },
  trustedBadge: { background: '#38bdf815', color: '#38bdf8', borderRadius: '6px', padding: '2px 8px', fontSize: '10px', fontWeight: 900, border: '1px solid #38bdf833' },
  untrustedBadge: { background: '#ef444415', color: '#ef4444', borderRadius: '6px', padding: '2px 8px', fontSize: '10px', fontWeight: 900, border: '1px solid #ef444433' },
  itemSub: { fontSize: '12px', color: '#94a3b8' },
  btnDanger: { background: 'linear-gradient(135deg, #ef4444, #f87171)', color: '#fff', border: 'none', borderRadius: '10px', padding: '8px 16px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
  btnSmDanger: { background: '#ef444422', border: '1px solid #ef444455', color: '#ef4444', borderRadius: '8px', padding: '6px 12px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
  btnSmSuccess: { background: '#4ade8022', border: '1px solid #4ade8055', color: '#4ade80', borderRadius: '8px', padding: '6px 12px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
  btnSmOutline: { background: '#0f172a', border: '1px solid #334155', color: '#94a3b8', borderRadius: '8px', padding: '6px 12px', fontSize: '12px', fontWeight: 700, cursor: 'pointer' },
};
