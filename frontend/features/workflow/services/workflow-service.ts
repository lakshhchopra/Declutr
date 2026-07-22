import type {
  Workflow,
  WorkflowRun,
  WorkflowLog,
  WorkflowStats,
} from '../types/workflow';

const BASE_URL = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080/api/v1';

async function apiFetch<T>(url: string, options?: RequestInit): Promise<T> {
  const res = await fetch(url, {
    ...options,
    headers: { 'Content-Type': 'application/json', ...options?.headers },
  });
  if (!res.ok) throw new Error(`Workflow API error: ${res.status} ${res.statusText}`);
  return res.json();
}

const VAULT_ID = 'vault-demo';

// ─── Mock Data Fallback ────────────────────────────────────────────────────────

const MOCK_WORKFLOWS: Workflow[] = [
  {
    workflowId: 'wf-travel-001',
    vaultId: VAULT_ID,
    name: 'Auto-tag Travel & Passport Documents',
    description: 'Automatically tags PDF documents matching Japan or Passport entities and moves them into the Japan Vacation collection.',
    enabled: true,
    status: 'IDLE',
    triggers: [
      { triggerId: 't1', workflowId: 'wf-travel-001', triggerType: 'ASSET_UPLOADED', createdAt: new Date().toISOString() },
    ],
    conditions: [
      { conditionId: 'c1', workflowId: 'wf-travel-001', field: 'fileType', operator: 'EQUALS', value: 'PDF', combinator: 'AND', createdAt: new Date().toISOString() },
      { conditionId: 'c2', workflowId: 'wf-travel-001', field: 'entity', operator: 'CONTAINS', value: 'Japan', combinator: 'AND', createdAt: new Date().toISOString() },
    ],
    actions: [
      { actionId: 'a1', workflowId: 'wf-travel-001', actionType: 'APPLY_TAGS', config: { tags: ['Travel', 'Passport'] }, executionOrder: 1, createdAt: new Date().toISOString() },
      { actionId: 'a2', workflowId: 'wf-travel-001', actionType: 'NOTIFY_USER', config: { message: 'Passport document auto-tagged.' }, executionOrder: 2, createdAt: new Date().toISOString() },
    ],
    lastRunAt: new Date(Date.now() - 2 * 3600000).toISOString(),
    runCount: 12,
    failureCount: 0,
    avgDurationMs: 45,
    createdBy: 'USER',
    createdAt: new Date(Date.now() - 7 * 86400000).toISOString(),
    updatedAt: new Date().toISOString(),
  },
  {
    workflowId: 'wf-expiry-002',
    vaultId: VAULT_ID,
    name: 'Document Expiration Alert & Milestone Sync',
    description: 'Fires weekly to scan for documents expiring within 60 days and creates milestone alerts.',
    enabled: true,
    status: 'IDLE',
    triggers: [
      { triggerId: 't2', workflowId: 'wf-expiry-002', triggerType: 'DOCUMENT_EXPIRING', createdAt: new Date().toISOString() },
    ],
    conditions: [
      { conditionId: 'c3', workflowId: 'wf-expiry-002', field: 'confidence', operator: 'GREATER_THAN', value: '0.8', combinator: 'AND', createdAt: new Date().toISOString() },
    ],
    actions: [
      { actionId: 'a3', workflowId: 'wf-expiry-002', actionType: 'CREATE_REMINDER', config: { title: 'Passport Renewal Reminder' }, executionOrder: 1, createdAt: new Date().toISOString() },
    ],
    lastRunAt: new Date(Date.now() - 12 * 3600000).toISOString(),
    runCount: 5,
    failureCount: 0,
    avgDurationMs: 30,
    createdBy: 'SYSTEM',
    createdAt: new Date(Date.now() - 14 * 86400000).toISOString(),
    updatedAt: new Date().toISOString(),
  },
];

const MOCK_RUNS: WorkflowRun[] = [
  {
    runId: 'run-001',
    workflowId: 'wf-travel-001',
    vaultId: VAULT_ID,
    triggerEvent: 'ASSET_UPLOADED',
    status: 'SUCCESS',
    durationMs: 42,
    startedAt: new Date(Date.now() - 2 * 3600000).toISOString(),
    completedAt: new Date(Date.now() - 2 * 3600000 + 42).toISOString(),
  },
  {
    runId: 'run-002',
    workflowId: 'wf-expiry-002',
    vaultId: VAULT_ID,
    triggerEvent: 'DOCUMENT_EXPIRING',
    status: 'SUCCESS',
    durationMs: 30,
    startedAt: new Date(Date.now() - 12 * 3600000).toISOString(),
    completedAt: new Date(Date.now() - 12 * 3600000 + 30).toISOString(),
  },
];

export const WorkflowService = {
  async getWorkflows(vaultId: string = VAULT_ID): Promise<Workflow[]> {
    try {
      const res = await apiFetch<{ workflows: Workflow[] }>(`${BASE_URL}/workflows?vaultId=${vaultId}`);
      return res.workflows ?? [];
    } catch {
      return MOCK_WORKFLOWS;
    }
  },

  async createWorkflow(wf: Partial<Workflow>, vaultId: string = VAULT_ID): Promise<Workflow> {
    try {
      return await apiFetch<Workflow>(`${BASE_URL}/workflows`, {
        method: 'POST',
        body: JSON.stringify({ ...wf, vaultId }),
      });
    } catch {
      const created: Workflow = {
        workflowId: `wf-${Date.now()}`,
        vaultId,
        name: wf.name ?? 'New Automation Workflow',
        description: wf.description ?? '',
        enabled: true,
        status: 'IDLE',
        triggers: wf.triggers ?? [],
        conditions: wf.conditions ?? [],
        actions: wf.actions ?? [],
        runCount: 0,
        failureCount: 0,
        avgDurationMs: 0,
        createdBy: 'USER',
        createdAt: new Date().toISOString(),
        updatedAt: new Date().toISOString(),
      };
      return created;
    }
  },

  async toggleWorkflow(workflowId: string, enabled: boolean): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/workflows/toggle`, {
        method: 'POST',
        body: JSON.stringify({ workflowId, enabled }),
      });
    } catch { /* mock */ }
  },

  async runWorkflow(workflowId: string, vaultId: string = VAULT_ID): Promise<WorkflowRun> {
    try {
      return await apiFetch<WorkflowRun>(`${BASE_URL}/workflows/run`, {
        method: 'POST',
        body: JSON.stringify({ workflowId, vaultId }),
      });
    } catch {
      return {
        runId: `run-${Date.now()}`,
        workflowId,
        vaultId,
        triggerEvent: 'MANUAL_TRIGGER',
        status: 'SUCCESS',
        durationMs: 38,
        startedAt: new Date().toISOString(),
        completedAt: new Date().toISOString(),
      };
    }
  },

  async deleteWorkflow(workflowId: string): Promise<void> {
    try {
      await apiFetch(`${BASE_URL}/workflows?workflowId=${workflowId}`, { method: 'DELETE' });
    } catch { /* mock */ }
  },

  async getRuns(vaultId: string = VAULT_ID): Promise<WorkflowRun[]> {
    try {
      const res = await apiFetch<{ runs: WorkflowRun[] }>(`${BASE_URL}/workflows/history?vaultId=${vaultId}`);
      return res.runs ?? [];
    } catch {
      return MOCK_RUNS;
    }
  },

  async getStats(vaultId: string = VAULT_ID): Promise<WorkflowStats> {
    try {
      return await apiFetch<WorkflowStats>(`${BASE_URL}/workflows/stats?vaultId=${vaultId}`);
    } catch {
      return {
        vaultId,
        totalWorkflows: 2,
        activeWorkflows: 2,
        totalRuns: 17,
        successfulRuns: 17,
        failedRuns: 0,
        successRate: 1.0,
        avgDurationMs: 37.5,
      };
    }
  },
};
