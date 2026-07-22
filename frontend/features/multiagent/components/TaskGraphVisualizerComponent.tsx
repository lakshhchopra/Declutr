"use client";

import React from "react";
import { CheckCircle2, ArrowRight, Layers, Bot } from "lucide-react";

export interface CoordinatorTaskItem {
  id: string;
  assigned_role: string;
  action: string;
  execution_mode: string;
  status: string;
  confidence: number;
}

export interface TaskGraph {
  goal_id: string;
  goal_title: string;
  status: string;
  tasks: CoordinatorTaskItem[];
}

interface TaskGraphVisualizerProps {
  graph: TaskGraph | null;
}

export function TaskGraphVisualizerComponent({ graph }: TaskGraphVisualizerProps) {
  if (!graph) return null;

  return (
    <div className="p-5 rounded-xl border bg-card space-y-4">
      <div className="flex items-center justify-between">
        <div>
          <h4 className="font-bold text-sm text-foreground">Orchestrated Task Execution DAG Graph</h4>
          <p className="text-xs text-muted-foreground">Goal: {graph.goal_title}</p>
        </div>
        <span className="px-2.5 py-0.5 rounded text-[10px] font-mono font-bold bg-emerald-500/10 text-emerald-500 border border-emerald-500/20">
          {graph.status}
        </span>
      </div>

      <div className="space-y-3">
        {graph.tasks.map((task, idx) => (
          <div key={task.id} className="p-3 rounded-lg border bg-secondary/30 flex items-center justify-between text-xs">
            <div className="flex items-center gap-3">
              <span className="w-6 h-6 rounded-full bg-indigo-500/10 text-indigo-500 border border-indigo-500/20 flex items-center justify-center font-mono font-bold text-[10px]">
                {idx + 1}
              </span>
              <div>
                <span className="font-bold text-foreground font-mono">{task.assigned_role}</span>
                <div className="text-muted-foreground flex items-center gap-1.5 font-mono text-[11px] mt-0.5">
                  <span>{task.action}</span>
                  <span className="text-indigo-400">({task.execution_mode})</span>
                </div>
              </div>
            </div>

            <div className="flex items-center gap-3">
              <span className="font-mono text-[10px] text-muted-foreground">{(task.confidence * 100).toFixed(0)}% conf</span>
              <span className="inline-flex items-center gap-1 px-2 py-0.5 rounded text-[10px] font-mono font-bold bg-emerald-500/10 text-emerald-500">
                <CheckCircle2 className="w-3 h-3" /> {task.status}
              </span>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
