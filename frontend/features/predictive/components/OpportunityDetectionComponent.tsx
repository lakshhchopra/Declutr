"use client";

import React from "react";
import { AlertCircle, Layers, CheckCircle2 } from "lucide-react";
import { PredictionItem } from "./PredictionFeedComponent";

interface OpportunityDetectionProps {
  predictions: PredictionItem[];
}

export function OpportunityDetectionComponent({ predictions }: OpportunityDetectionProps) {
  const opportunities = predictions.filter((p) => p.type === "OPPORTUNITY_DETECTION" || p.type === "MISSING_DOCUMENT");

  return (
    <div className="p-5 rounded-xl border bg-card space-y-4">
      <div className="flex items-center gap-2 font-bold text-sm">
        <AlertCircle className="w-4 h-4 text-indigo-500" />
        <span>Opportunity & Gap Detection</span>
      </div>

      <div className="space-y-2">
        {opportunities.map((opp) => (
          <div key={opp.id} className="p-3 rounded-lg border bg-secondary/30 space-y-1 text-xs">
            <div className="flex items-center justify-between font-bold">
              <span className="text-foreground">{opp.title}</span>
              <span className="text-[10px] font-mono text-emerald-500 font-bold">Detected</span>
            </div>
            <p className="text-muted-foreground text-[11px]">{opp.description}</p>
          </div>
        ))}
      </div>
    </div>
  );
}
