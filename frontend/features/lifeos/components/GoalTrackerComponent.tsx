"use client";

import React from "react";
import { Target, CheckCircle2, AlertCircle } from "lucide-react";

export interface ProjectGoalItem {
  id: string;
  title: string;
  description: string;
  progress_pct: number;
  is_completed: boolean;
  missing_assets?: string[];
}

interface GoalTrackerProps {
  goals: ProjectGoalItem[];
}

export function GoalTrackerComponent({ goals }: GoalTrackerProps) {
  return (
    <div className="p-5 rounded-xl border bg-card space-y-4">
      <div className="flex items-center gap-2 font-bold text-sm">
        <Target className="w-4 h-4 text-indigo-500" />
        <span>Project Goal Tracker & Milestone Progress</span>
      </div>

      <div className="space-y-3">
        {goals.map((g) => (
          <div key={g.id} className="p-3.5 rounded-lg border bg-secondary/30 space-y-2 text-xs">
            <div className="flex items-center justify-between font-bold">
              <span className="text-foreground">{g.title}</span>
              <span className="font-mono text-indigo-500">{g.progress_pct}% Completed</span>
            </div>
            <p className="text-muted-foreground text-[11px]">{g.description}</p>

            <div className="w-full bg-secondary rounded-full h-1.5 overflow-hidden">
              <div className="bg-indigo-600 h-1.5 rounded-full" style={{ width: `${g.progress_pct}%` }} />
            </div>

            {g.missing_assets && g.missing_assets.length > 0 && (
              <div className="flex items-center gap-1.5 text-[10px] font-mono text-amber-500 pt-1">
                <AlertCircle className="w-3 h-3" /> Missing: {g.missing_assets.join(", ")}
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
}
