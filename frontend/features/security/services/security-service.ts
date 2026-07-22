import type {
  SecurityDashboard,
  AuditEvent,
  ActiveSession,
  Device,
  RiskAssessment,
  SecurityRecommendation,
  AuditCategory,
} from '../types/security';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080/api/v1';

async function apiFetch<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    ...options,
    headers: { 'Content-Type': 'application/json', ...options?.headers },
  });
  if (!res.ok) throw new Error(`Security API error: ${res.status} ${res.statusText}`);
  return res.json();
}

const VAULT_ID = 'vault-demo';

// ─── Mock Data Fallback ────────────────────────────────────────────────────────

const MOCK_DASHBOARD: SecurityDashboard = {
  vaultId: VAULT_ID,
  score: {
    scoreId: 'sc-1',
    vaultId: VAULT_ID,
    score: 92,
    grade: 'A',
    status: 'HEALTHY',
    factors: { mfa: false, encrypted: true },
    calculatedAt: new Date().toISOString(),
  },
  risk: {
    assessmentId: 'risk-1',
    vaultId: VAULT_ID,
    riskScore: 12,
    riskLevel: 'LOW',
    signals: [
      { signalId: 's1', signalType: 'NEW_DEVICE', description: 'New browser login detected from Chrome macOS', weight: 12, detectedAt: new Date(Date.now() - 3600000).toISOString() },
    ],
    assessedAt: new Date().toISOString(),
  },
  activeSessions: [
    { sessionId: 's1', vaultId: VAULT_ID, userId: 'usr-owner', deviceName: 'MacBook Pro 16"', browser: 'Chrome 125.0', ipAddress: '192.168.1.45', location: 'Tokyo, Japan', isCurrent: true, createdAt: new Date(Date.now() - 7200000).toISOString(), lastSeenAt: new Date().toISOString() },
    { sessionId: 's2', vaultId: VAULT_ID, userId: 'usr-owner', deviceName: 'iPhone 15 Pro', browser: 'Declutr Mobile 1.0', ipAddress: '192.168.1.88', location: 'Tokyo, Japan', isCurrent: false, createdAt: new Date(Date.now() - 86400000).toISOString(), lastSeenAt: new Date(Date.now() - 3600000).toISOString() },
  ],
  devices: [
    { deviceId: 'd1', vaultId: VAULT_ID, deviceName: 'MacBook Pro 16"', browser: 'Chrome 125.0', os: 'macOS Sonoma', platform: 'WEB', ipAddress: '192.168.1.45', location: 'Tokyo, Japan', firstSeenAt: new Date(Date.now() - 30 * 86400000).toISOString(), lastSeenAt: new Date().toISOString(), isTrusted: true },
    { deviceId: 'd2', vaultId: VAULT_ID, deviceName: 'iPhone 15 Pro', browser: 'Declutr Mobile 1.0', os: 'iOS 17.5', platform: 'MOBILE', ipAddress: '192.168.1.88', location: 'Tokyo, Japan', firstSeenAt: new Date(Date.now() - 14 * 86400000).toISOString(), lastSeenAt: new Date(Date.now() - 3600000).toISOString(), isTrusted: true },
  ],
  recentAudit: [
    { auditId: 'a1', vaultId: VAULT_ID, category: 'AUTH', action: 'USER_LOGIN_SUCCESS', actorId: 'usr-owner', actorName: 'Vault Owner', ipAddress: '192.168.1.45', createdAt: new Date(Date.now() - 600000).toISOString() },
    { auditId: 'a2', vaultId: VAULT_ID, category: 'ASSET', action: 'ASSET_UPLOAD', actorId: 'usr-owner', actorName: 'Vault Owner', ipAddress: '192.168.1.45', resourceId: 'asset-passport-001', createdAt: new Date(Date.now() - 7200000).toISOString() },
  ],
  recommendations: [
    { recId: 'r1', vaultId: VAULT_ID, category: 'AUTHENTICATION', title: 'Enable Multi-Factor Authentication (MFA)', description: 'Add TOTP authenticator protection to prevent unauthorized logins.', actionType: 'ENABLE_MFA', priority: 'HIGH', isDismissed: false, createdAt: new Date().toISOString() },
  ],
  mfaEnabled: false,
  backupStatus: 'HEALTHY (Weekly Auto)',
  encryptedVault: true,
};

export const SecurityService = {
  async getDashboard(vaultId: string = VAULT_ID): Promise<SecurityDashboard> {
    try {
      return await apiFetch<SecurityDashboard>(`${BASE_URL}/security/dashboard?vaultId=${vaultId}`);
    } catch {
      return MOCK_DASHBOARD;
    }
  },

  async getAuditEvents(category?: AuditCategory, vaultId: string = VAULT_ID): Promise<AuditEvent[]> {
    try {
      const url = category
        ? `${BASE_URL}/security/audit?vaultId=${vaultId}&category=${category}`
        : `${BASE_URL}/security/audit?vaultId=${vaultId}`;
      const res = await apiFetch<{ events: AuditEvent[] }>(url);
      return res.events ?? [];
    } catch {
      return MOCK_DASHBOARD.recentAudit;
    }
  },

  async getSessions(vaultId: string = VAULT_ID): Promise<ActiveSession[]> {
    try {
      const res = await apiFetch<{ sessions: ActiveSession[] }>(`${BASE_URL}/security/sessions?vaultId=${vaultId}`);
      return res.sessions ?? [];
    } catch {
      return MOCK_DASHBOARD.activeSessions;
    }
  },

  async terminateSession(sessionId?: string, vaultId: string = VAULT_ID): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/security/sessions/terminate`, {
        method: 'POST',
        body: JSON.stringify({ vaultId, sessionId }),
      });
    } catch { /* mock */ }
  },

  async getDevices(vaultId: string = VAULT_ID): Promise<Device[]> {
    try {
      const res = await apiFetch<{ devices: Device[] }>(`${BASE_URL}/security/devices?vaultId=${vaultId}`);
      return res.devices ?? [];
    } catch {
      return MOCK_DASHBOARD.devices;
    }
  },

  async setDeviceTrust(deviceId: string, trust: boolean, vaultId: string = VAULT_ID): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/security/devices/trust`, {
        method: 'POST',
        body: JSON.stringify({ vaultId, deviceId, trust }),
      });
    } catch { /* mock */ }
  },

  async getRiskAssessment(vaultId: string = VAULT_ID): Promise<RiskAssessment> {
    try {
      return await apiFetch<RiskAssessment>(`${BASE_URL}/security/risk?vaultId=${vaultId}`);
    } catch {
      return MOCK_DASHBOARD.risk;
    }
  },

  async getRecommendations(vaultId: string = VAULT_ID): Promise<SecurityRecommendation[]> {
    try {
      const res = await apiFetch<{ recommendations: SecurityRecommendation[] }>(`${BASE_URL}/security/recommendations?vaultId=${vaultId}`);
      return res.recommendations ?? [];
    } catch {
      return MOCK_DASHBOARD.recommendations;
    }
  },
};
