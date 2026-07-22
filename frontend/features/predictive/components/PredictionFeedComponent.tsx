"use client";

import React from "react";
import { Sparkles, Check, X, ShieldAlert, Clock, ArrowRight } from "lucide-react";

export interface PredictionItem {
  id: string;
  type: string;
  title: string;
  description: string;
  confidence: number;
  priority: string;
  evidence: {
    source_module: string;
    reasoning: string;
    key_facts: string[];
  };
  suggested_action: string;
  status: string;
}

interface PredictionFeedProps {
  predictions: PredictionItem[];
  onAccept: (id: string) => void;
  onDismiss: (id: string, reason: string) => void;
}

export function PredictionFeedComponent({ predictions, onAccept, onDismiss }: PredictionFeedProps) {
  const pendingPreds = predictions.filter((p) => p.status === "PENDING");

  if (pendingPreds.length === 0) {
    return (
      <div className="p-8 rounded-xl border bg-card text-center space-y-2">
        <Sparkles className="w-8 h-8 text-indigo-500 mx-auto" />
        <h4 className="font-bold text-sm">All Life Intelligence Suggestions Reviewed</h4>
        <p className="text-xs text-muted-foreground">Declutr is continuously scanning your vault for proactive insights.</p>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      {pendingPreds.map((p) => (
        <div key={p.id} className="p-5 rounded-xl border bg-card space-y-4 hover:border-indigo-500/50 transition-all shadow-sm">
          <div className="flex items-start justify-between">
            <div className="space-y-1">
              <div className="flex items-center gap-2">
                <span className="p-1.5 rounded-md bg-indigo-500/10 text-indigo-500">
                  <Sparkles className="w-4 h-4" />
                </span>
                <h4 className="font-bold text-sm text-foreground">{p.title}</h4>
                <span className="text-[10px] font-mono px-2 py-0.5 rounded bg-indigo-500/10 text-indigo-500 font-bold">
                  {(p.confidence * 100).toFixed(0)}% Confidence
                </span>
              </div>
              <p className="text-xs text-muted-foreground">{p.description}</p>
            </div>

            <span
              className={`px-2.5 py-0.5 rounded text-[10px] font-mono font-bold ${
                p.priority === "HIGH"
                  ? "bg-rose-500/10 text-rose-500 border border-rose-500/20"
                  : "bg-amber-500/10 text-amber-500 border border-amber-500/20"
              }`}
            >
              {p.priority} PRIORITY
            </span>
          </div>

          <div className="p-3 rounded-lg bg-secondary/50 border text-xs space-y-1">
            <div className="flex items-center justify-between font-mono text-[10px] text-muted-foreground">
              <span>Source: {p.evidence.source_module}</span>
            </div>
            <p className="text-muted-foreground">{p.evidence.reasoning}</p>
          </div>

          <div className="flex items-center justify-between pt-2 border-t text-xs">
            <span className="font-semibold text-indigo-500 flex items-center gap-1">
              <ArrowRight className="w-3.5 h-3.5" /> {p.suggested_action}
            </span>
            <div className="flex items-center gap-2">
              <button
                onClick={() => onDismiss(p.id, "Not relevant")}
                className="px-3 py-1.5 rounded-lg border hover:bg-secondary text-muted-foreground font-semibold flex items-center gap-1"
              >
                <X className="w-3.5 h-3.5" /> Dismiss
              </button>
              <button
                onClick={() => onAccept(p.id)}
                className="px-3.5 py-1.5 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white font-semibold flex items-center gap-1 shadow-sm transition-all"
              >
                <Check className="w-3.5 h-3.5" /> Apply Action
              </button>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
}
