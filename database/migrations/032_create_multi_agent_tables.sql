-- Migration: 032_create_multi_agent_tables.sql
-- Description: Multi-Agent Intelligence Platform Database Tables

-- 1. Agent Registry Table
CREATE TABLE IF NOT EXISTS agent_registry (
    id VARCHAR(255) PRIMARY KEY,
    role VARCHAR(50) NOT NULL,
    name VARCHAR(255) NOT NULL,
    capabilities JSONB NOT NULL DEFAULT '[]'::jsonb,
    supported_tools JSONB NOT NULL DEFAULT '[]'::jsonb,
    status VARCHAR(50) NOT NULL DEFAULT 'READY', -- READY, BUSY, PAUSED, ERROR
    version VARCHAR(50) NOT NULL DEFAULT '2.0.0',
    health VARCHAR(50) NOT NULL DEFAULT 'HEALTHY',
    priority INT DEFAULT 1,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_agent_reg_role ON agent_registry (role);

-- 2. Multi-Agent Coordinator Tasks Table
CREATE TABLE IF NOT EXISTS agent_tasks (
    id VARCHAR(255) PRIMARY KEY,
    goal_id VARCHAR(255) NOT NULL,
    assigned_role VARCHAR(50) NOT NULL,
    assigned_agent_id VARCHAR(255) REFERENCES agent_registry(id) ON DELETE SET NULL,
    action VARCHAR(100) NOT NULL,
    parameters JSONB NOT NULL DEFAULT '{}'::jsonb,
    dependencies JSONB NOT NULL DEFAULT '[]'::jsonb,
    execution_mode VARCHAR(50) NOT NULL DEFAULT 'PARALLEL',
    status VARCHAR(50) NOT NULL DEFAULT 'PENDING',
    result_payload JSONB,
    confidence NUMERIC(3,2) DEFAULT 0.95,
    error_message TEXT
);

CREATE INDEX IF NOT EXISTS idx_multi_tasks_goal ON agent_tasks (goal_id);

-- 3. Structured Message Bus Messages Audit Table
CREATE TABLE IF NOT EXISTS agent_messages (
    id VARCHAR(255) PRIMARY KEY,
    correlation_id VARCHAR(255) NOT NULL,
    goal_id VARCHAR(255) NOT NULL,
    task_id VARCHAR(255) NOT NULL,
    sender VARCHAR(255) NOT NULL,
    receiver VARCHAR(255) NOT NULL,
    message_type VARCHAR(50) NOT NULL,
    payload JSONB NOT NULL DEFAULT '{}'::jsonb,
    context JSONB NOT NULL DEFAULT '{}'::jsonb,
    timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_multi_msg_corr ON agent_messages (correlation_id);

-- 4. Shared Multi-Agent Memory Table
CREATE TABLE IF NOT EXISTS agent_memory (
    id VARCHAR(255) PRIMARY KEY,
    goal_id VARCHAR(255) NOT NULL,
    agent_id VARCHAR(255) NOT NULL,
    key VARCHAR(255) NOT NULL,
    value JSONB NOT NULL,
    category VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_multi_mem_goal ON agent_memory (goal_id);

-- 5. Agent Observability Telemetry & Health Table
CREATE TABLE IF NOT EXISTS agent_health (
    agent_id VARCHAR(255) PRIMARY KEY REFERENCES agent_registry(id) ON DELETE CASCADE,
    role VARCHAR(50) NOT NULL,
    status VARCHAR(50) NOT NULL,
    latency_ms INT8 DEFAULT 0,
    success_rate NUMERIC(3,2) DEFAULT 1.0,
    total_tasks INT DEFAULT 0,
    failed_tasks INT DEFAULT 0,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- 6. Agent Capabilities Table
CREATE TABLE IF NOT EXISTS agent_capabilities (
    id VARCHAR(255) PRIMARY KEY,
    agent_id VARCHAR(255) NOT NULL REFERENCES agent_registry(id) ON DELETE CASCADE,
    capability_name VARCHAR(100) NOT NULL
);
