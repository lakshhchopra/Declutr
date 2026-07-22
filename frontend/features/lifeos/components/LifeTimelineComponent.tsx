"use client";

import React from "react";
import { Clock, FileText, Calendar, Plane, MessageSquare } from "lucide-react";

export interface LifeTimelineEventItem {
  id: string;
  title: string;
  event_type: string;
  description: string;
  timestamp: string;
}

interface LifeTimelineProps {
  events: LifeTimelineEventItem[];
}

export function LifeTimelineComponent({ events }: LifeTimelineProps) {
  return (
    <div className="p-5 rounded-xl border bg-card space-y-4">
      <div className="flex items-center gap-2 font-bold text-sm">
        <Clock className="w-4 h-4 text-indigo-500" />
        <span>Unified Life Timeline Stream</span>
      </div>

      <div className="space-y-2">
        {events.map((ev) => (
          <div key={ev.id} className="p-3 rounded-lg bg-secondary/40 border text-xs flex items-center justify-between font-mono">
            <div className="flex items-center gap-2">
              <span className="px-2 py-0.5 rounded bg-indigo-500/10 text-indigo-500 font-bold text-[10px]">
                {ev.event_type}
              </span>
              <span className="font-bold text-foreground">{ev.title}</span>
            </div>
            <span className="text-[10px] text-muted-foreground">{ev.description}</span>
          </div>
        ))}
      </div>
    </div>
  );
}
