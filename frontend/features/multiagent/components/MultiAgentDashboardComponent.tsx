"use client";

import React, { useState } from "react";
import { Users, Bot, Send, Sparkles, Activity } from "lucide-react";

interface MultiAgentDashboardProps {
  onProcessGoal: (title: string) => void;
}

export function MultiAgentDashboardComponent({ onProcessGoal }: MultiAgentDashboardProps) {
  const [title, setTitle] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (title) {
      onProcessGoal(title);
      setTitle("");
    }
  };

  return (
    <div className="p-6 rounded-xl border bg-card space-y-4">
      <div className="flex items-center gap-2 font-bold text-base">
        <Users className="w-5 h-5 text-indigo-500" />
        <span>Submit Goal to Multi-Agent Coordinator</span>
      </div>

      <form onSubmit={handleSubmit} className="flex gap-2">
        <input
          type="text"
          required
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          placeholder="e.g. Research and summarize Q3 financial receipts and check compliance"
          className="flex-1 px-4 py-2.5 border rounded-xl bg-secondary text-xs"
        />
        <button
          type="submit"
          className="px-5 py-2.5 rounded-xl bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold flex items-center gap-1.5 shadow-sm transition-all whitespace-nowrap"
        >
          <Send className="w-4 h-4" /> Orchestrate Multi-Agent Execution
        </button>
      </form>
    </div>
  );
}
