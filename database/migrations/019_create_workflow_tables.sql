-- Migration 019: Create Workflow Automation Tables

-- 1. workflows - core workflow definition
CREATE TABLE IF NOT EXISTS workflows (
    workflow_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL DEFAULT '',
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    status VARCHAR(50) NOT NULL DEFAULT 'IDLE', -- IDLE, RUNNING, ERROR
    last_run_at TIMESTAMPTZ,
    run_count INT NOT NULL DEFAULT 0,
    failure_count INT NOT NULL DEFAULT 0,
    avg_duration_ms INT NOT NULL DEFAULT 0,
    created_by VARCHAR(50) NOT NULL DEFAULT 'USER',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 2. workflow_triggers - event trigger definition
CREATE TABLE IF NOT EXISTS workflow_triggers (
    trigger_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(workflow_id) ON DELETE CASCADE,
    trigger_type VARCHAR(50) NOT NULL, -- ASSET_UPLOADED, ASSET_UPDATED, CONTEXT_CREATED, MEMORY_CREATED, DOCUMENT_EXPIRING, DAILY_SCHEDULE, MANUAL_TRIGGER, AI_INSIGHT_CREATED, TIMELINE_EVENT
    config JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 3. workflow_conditions - condition evaluation rules
CREATE TABLE IF NOT EXISTS workflow_conditions (
    condition_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(workflow_id) ON DELETE CASCADE,
    field VARCHAR(100) NOT NULL,
    operator VARCHAR(50) NOT NULL, -- EQUALS, CONTAINS, GREATER_THAN, LESS_THAN, IN
    value VARCHAR(255) NOT NULL,
    combinator VARCHAR(10) NOT NULL DEFAULT 'AND', -- AND, OR, NOT
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 4. workflow_actions - executable action steps
CREATE TABLE IF NOT EXISTS workflow_actions (
    action_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(workflow_id) ON DELETE CASCADE,
    action_type VARCHAR(50) NOT NULL, -- CREATE_COLLECTION, MOVE_ASSET, APPLY_TAGS, GENERATE_SUMMARY, ARCHIVE_ASSET, CREATE_REMINDER, PIN_MEMORY, REFRESH_SEARCH_INDEX, NOTIFY_USER
    config JSONB NOT NULL DEFAULT '{}',
    execution_order INT NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 5. workflow_runs - execution run records
CREATE TABLE IF NOT EXISTS workflow_runs (
    run_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(workflow_id) ON DELETE CASCADE,
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    trigger_event VARCHAR(100) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'RUNNING', -- RUNNING, SUCCESS, FAILED, RETRYING
    duration_ms INT NOT NULL DEFAULT 0,
    error_message TEXT NOT NULL DEFAULT '',
    started_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMPTZ
);

-- 6. workflow_logs - detailed step execution log
CREATE TABLE IF NOT EXISTS workflow_logs (
    log_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    run_id UUID NOT NULL REFERENCES workflow_runs(run_id) ON DELETE CASCADE,
    workflow_id UUID NOT NULL REFERENCES workflows(workflow_id) ON DELETE CASCADE,
    step_type VARCHAR(50) NOT NULL,
    message TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'INFO',
    logged_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 7. workflow_history - aggregated history summary
CREATE TABLE IF NOT EXISTS workflow_history (
    history_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    workflow_id UUID NOT NULL REFERENCES workflows(workflow_id) ON DELETE CASCADE,
    vault_id UUID NOT NULL REFERENCES vaults(vault_id) ON DELETE CASCADE,
    total_runs INT NOT NULL DEFAULT 0,
    success_rate FLOAT NOT NULL DEFAULT 1.0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Indexes
CREATE INDEX IF NOT EXISTS idx_workflows_vault ON workflows(vault_id);
CREATE INDEX IF NOT EXISTS idx_workflow_runs_workflow ON workflow_runs(workflow_id);
