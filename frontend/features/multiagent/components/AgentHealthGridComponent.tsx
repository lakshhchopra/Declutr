"use client";

import React from "react";
import { Activity, CheckCircle2, Cpu } from "lucide-react";

export interface HealthMetricItem {
  agent_id: string;
  role: string;
  status: string;
  latency_ms: number;
  success_rate: number;
  total_tasks: number;
}

interface AgentHealthGridProps {
  metrics: HealthMetricItem[];
}

export function AgentHealthGridComponent({ metrics }: AgentHealthGridProps) {
  return (
    <div className="space-y-4">
      <div className="flex items-center gap-2 font-bold text-sm">
        <Activity className="w-4 h-4 text-indigo-500" />
        <span>Specialist Agent Cluster Health & Latency</span>
      </div>

      <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-5 gap-3">
        {metrics.map((m) => (
          <div key={m.agent_id} className="p-3 rounded-lg border bg-card space-y-1 text-xs">
            <div className="flex items-center justify-between">
              <span className="font-bold text-foreground font-mono text-[11px] truncate">{m.role}</span>
              <CheckCircle2 className="w-3.5 h-3.5 text-emerald-500 shrink-0" />
            </div>
            <div className="text-[10px] font-mono text-muted-foreground">
              Latency: <span className="font-bold text-foreground">{m.latency_ms}ms</span>
            </div>
            <div className="text-[10px] font-mono text-emerald-500 font-bold">
              {(m.success_rate * 100).toFixed(0)}% success
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
