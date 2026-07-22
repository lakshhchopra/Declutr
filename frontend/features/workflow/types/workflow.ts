// Workflow Automation Engine TypeScript types

export type TriggerType =
  | 'ASSET_UPLOADED'
  | 'ASSET_UPDATED'
  | 'ASSET_DELETED'
  | 'CONTEXT_CREATED'
  | 'MEMORY_CREATED'
  | 'ENTITY_FOUND'
  | 'RELATIONSHIP_CREATED'
  | 'DOCUMENT_EXPIRING'
  | 'DAILY_SCHEDULE'
  | 'MANUAL_TRIGGER'
  | 'AI_INSIGHT_CREATED'
  | 'TIMELINE_EVENT';

export type ConditionOperator = 'EQUALS' | 'CONTAINS' | 'GREATER_THAN' | 'LESS_THAN' | 'IN';
export type ConditionCombinator = 'AND' | 'OR' | 'NOT';

export type ActionType =
  | 'CREATE_COLLECTION'
  | 'MOVE_ASSET'
  | 'APPLY_TAGS'
  | 'GENERATE_SUMMARY'
  | 'ARCHIVE_ASSET'
  | 'CREATE_REMINDER'
  | 'PIN_MEMORY'
  | 'REFRESH_SEARCH_INDEX'
  | 'NOTIFY_USER';

export type RunStatus = 'PENDING' | 'RUNNING' | 'SUCCESS' | 'FAILED' | 'RETRYING';

export interface WorkflowTrigger {
  triggerId: string;
  workflowId: string;
  triggerType: TriggerType;
  config?: Record<string, unknown>;
  createdAt: string;
}

export interface WorkflowCondition {
  conditionId: string;
  workflowId: string;
  field: string;
  operator: ConditionOperator;
  value: string;
  combinator: ConditionCombinator;
  createdAt: string;
}

export interface WorkflowAction {
  actionId: string;
  workflowId: string;
  actionType: ActionType;
  config?: Record<string, unknown>;
  executionOrder: number;
  createdAt: string;
}

export interface Workflow {
  workflowId: string;
  vaultId: string;
  name: string;
  description: string;
  enabled: boolean;
  status: string;
  triggers: WorkflowTrigger[];
  conditions: WorkflowCondition[];
  actions: WorkflowAction[];
  lastRunAt?: string;
  runCount: number;
  failureCount: number;
  avgDurationMs: number;
  createdBy: string;
  createdAt: string;
  updatedAt: string;
}

export interface WorkflowRun {
  runId: string;
  workflowId: string;
  vaultId: string;
  triggerEvent: string;
  status: RunStatus;
  durationMs: number;
  errorMessage?: string;
  startedAt: string;
  completedAt?: string;
}

export interface WorkflowLog {
  logId: string;
  runId: string;
  workflowId: string;
  stepType: string;
  message: string;
  status: string;
  loggedAt: string;
}

export interface WorkflowStats {
  vaultId: string;
  totalWorkflows: number;
  activeWorkflows: number;
  totalRuns: number;
  successfulRuns: number;
  failedRuns: number;
  successRate: number;
  avgDurationMs: number;
}
