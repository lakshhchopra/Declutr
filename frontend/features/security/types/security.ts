// Security Center, Audit Hub & Trust Platform TypeScript types

export type SecuritySeverity = 'LOW' | 'MEDIUM' | 'HIGH' | 'CRITICAL';

export type AuditCategory =
  | 'AUTH'
  | 'ASSET'
  | 'SHARING'
  | 'WORKFLOW'
  | 'AI'
  | 'SEARCH'
  | 'BACKUP'
  | 'VERSIONING'
  | 'SETTINGS';

export type RiskLevel = 'LOW' | 'MEDIUM' | 'HIGH' | 'CRITICAL';
export type SecurityGrade = 'A' | 'B' | 'C' | 'D' | 'F';

export interface Device {
  deviceId: string;
  vaultId: string;
  deviceName: string;
  browser: string;
  os: string;
  platform: string;
  ipAddress: string;
  location: string;
  firstSeenAt: string;
  lastSeenAt: string;
  isTrusted: boolean;
}

export interface ActiveSession {
  sessionId: string;
  vaultId: string;
  userId: string;
  deviceName: string;
  browser: string;
  ipAddress: string;
  location: string;
  isCurrent: boolean;
  createdAt: string;
  lastSeenAt: string;
}

export interface AuditEvent {
  auditId: string;
  vaultId: string;
  category: AuditCategory;
  action: string;
  actorId: string;
  actorName: string;
  ipAddress: string;
  resourceId?: string;
  details?: Record<string, unknown>;
  createdAt: string;
}

export interface RiskSignal {
  signalId: string;
  signalType: string;
  description: string;
  weight: number;
  detectedAt: string;
}

export interface RiskAssessment {
  assessmentId: string;
  vaultId: string;
  riskScore: number;
  riskLevel: RiskLevel;
  signals: RiskSignal[];
  assessedAt: string;
}

export interface SecurityRecommendation {
  recId: string;
  vaultId: string;
  category: string;
  title: string;
  description: string;
  actionType: string;
  priority: SecuritySeverity;
  isDismissed: boolean;
  createdAt: string;
}

export interface SecurityScore {
  scoreId: string;
  vaultId: string;
  score: number;
  grade: SecurityGrade;
  status: string;
  factors: Record<string, unknown>;
  calculatedAt: string;
}

export interface SecurityDashboard {
  vaultId: string;
  score: SecurityScore;
  risk: RiskAssessment;
  activeSessions: ActiveSession[];
  devices: Device[];
  recentAudit: AuditEvent[];
  recommendations: SecurityRecommendation[];
  mfaEnabled: boolean;
  backupStatus: string;
  encryptedVault: boolean;
}
