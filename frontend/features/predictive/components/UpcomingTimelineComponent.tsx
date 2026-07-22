"use client";

import React from "react";
import { Calendar, Plane, FileText, Clock } from "lucide-react";
import { PredictionItem } from "./PredictionFeedComponent";

interface UpcomingTimelineProps {
  predictions: PredictionItem[];
}

export function UpcomingTimelineComponent({ predictions }: UpcomingTimelineProps) {
  const upcomingEvents = predictions.filter(
    (p) => p.type === "UPCOMING_TRIP" || p.type === "EXPIRING_DOCUMENT" || p.type === "UPCOMING_DEADLINE"
  );

  return (
    <div className="p-5 rounded-xl border bg-card space-y-4">
      <div className="flex items-center gap-2 font-bold text-sm">
        <Calendar className="w-4 h-4 text-indigo-500" />
        <span>Upcoming Deadlines & Life Events</span>
      </div>

      <div className="space-y-3">
        {upcomingEvents.map((item) => (
          <div key={item.id} className="p-3 rounded-lg border bg-secondary/30 flex items-center justify-between text-xs">
            <div className="flex items-center gap-3">
              <div className="p-2 rounded-lg bg-indigo-500/10 text-indigo-500">
                {item.type === "UPCOMING_TRIP" ? <Plane className="w-4 h-4" /> : <FileText className="w-4 h-4" />}
              </div>
              <div>
                <h5 className="font-bold text-foreground">{item.title}</h5>
                <span className="text-[10px] font-mono text-muted-foreground">{item.suggested_action}</span>
              </div>
            </div>
            <span className="text-[10px] font-mono font-bold text-indigo-500 bg-indigo-500/10 px-2 py-0.5 rounded">
              {(item.confidence * 100).toFixed(0)}% conf
            </span>
          </div>
        ))}
      </div>
    </div>
  );
}
