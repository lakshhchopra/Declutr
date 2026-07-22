"use client";

import React from "react";
import { MessageSquare, ArrowRight } from "lucide-react";

export interface AgentMessageItem {
  id: string;
  correlation_id: string;
  sender: string;
  receiver: string;
  message_type: string;
  timestamp: string;
}

interface MessageBusMonitorProps {
  messages: AgentMessageItem[];
}

export function MessageBusMonitorComponent({ messages }: MessageBusMonitorProps) {
  return (
    <div className="p-5 rounded-xl border bg-card space-y-4">
      <div className="flex items-center gap-2 font-bold text-sm">
        <MessageSquare className="w-4 h-4 text-indigo-500" />
        <span>Structured Message Bus Live Telemetry Log</span>
      </div>

      <div className="space-y-2 max-h-64 overflow-y-auto pr-1">
        {messages.map((msg) => (
          <div key={msg.id} className="p-2.5 rounded bg-secondary/50 border text-[11px] font-mono flex items-center justify-between">
            <div className="flex items-center gap-2">
              <span className="font-bold text-indigo-500">{msg.sender}</span>
              <ArrowRight className="w-3 h-3 text-muted-foreground" />
              <span className="font-bold text-foreground">{msg.receiver}</span>
            </div>
            <span className="px-2 py-0.5 rounded bg-card text-muted-foreground">{msg.message_type}</span>
          </div>
        ))}
      </div>
    </div>
  );
}
