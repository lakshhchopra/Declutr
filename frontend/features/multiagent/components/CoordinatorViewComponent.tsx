"use client";

import React from "react";
import { CheckCircle2, ShieldCheck, Scale, Award } from "lucide-react";

export interface ConsensusItem {
  goal_id: string;
  consensus_achieved: boolean;
  winning_agent_id: string;
  winning_confidence: number;
  explanation: string;
  escalate_to_user: boolean;
}

interface CoordinatorViewProps {
  consensus: ConsensusItem | null;
}

export function CoordinatorViewComponent({ consensus }: CoordinatorViewProps) {
  if (!consensus) return null;

  return (
    <div className="p-5 rounded-xl border bg-card space-y-3">
      <div className="flex items-center gap-2 font-bold text-sm text-foreground">
        <Scale className="w-5 h-5 text-indigo-500" />
        <span>Coordinator Consensus & Conflict Resolution Output</span>
      </div>

      <div className="p-4 rounded-lg bg-emerald-500/10 border border-emerald-500/20 text-xs space-y-2">
        <div className="flex items-center justify-between font-mono font-bold text-emerald-500">
          <span>Consensus Status: ACHIEVED</span>
          <span>Winning Confidence: {(consensus.winning_confidence * 100).toFixed(0)}%</span>
        </div>
        <p className="text-muted-foreground leading-relaxed">{consensus.explanation}</p>
      </div>
    </div>
  );
}
