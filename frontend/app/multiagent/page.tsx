"use client";

import React, { useEffect, useState } from "react";
import { MultiAgentDashboardComponent } from "@/features/multiagent/components/MultiAgentDashboardComponent";
import { TaskGraphVisualizerComponent, TaskGraph } from "@/features/multiagent/components/TaskGraphVisualizerComponent";
import { CoordinatorViewComponent, ConsensusItem } from "@/features/multiagent/components/CoordinatorViewComponent";
import { MessageBusMonitorComponent, AgentMessageItem } from "@/features/multiagent/components/MessageBusMonitorComponent";
import { AgentHealthGridComponent, HealthMetricItem } from "@/features/multiagent/components/AgentHealthGridComponent";
import { Users, Sparkles, Activity } from "lucide-react";

export default function MultiAgentPage() {
  const [activeGraph, setActiveGraph] = useState<TaskGraph | null>(null);
  const [consensus, setConsensus] = useState<ConsensusItem | null>(null);
  const [messages, setMessages] = useState<AgentMessageItem[]>([]);
  const [health, setHealth] = useState<HealthMetricItem[]>([]);

  const fetchHealth = async () => {
    try {
      const res = await fetch("/api/v1/multiagent/health");
      if (res.ok) {
        const data = await res.json();
        setHealth(data || []);
      }
    } catch (err) {
      console.error("Failed to load health metrics", err);
    }
  };

  const fetchMessages = async () => {
    try {
      const res = await fetch("/api/v1/multiagent/messages");
      if (res.ok) {
        const data = await res.json();
        setMessages(data || []);
      }
    } catch (err) {
      console.error("Failed to load message bus logs", err);
    }
  };

  useEffect(() => {
    fetchHealth();
  }, []);

  const handleProcessGoal = async (title: string) => {
    try {
      const res = await fetch("/api/v1/multiagent/goals", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title }),
      });
      if (res.ok) {
        const data = await res.json();
        setActiveGraph(data.task_graph);
        setConsensus(data.consensus);
        fetchMessages();
      }
    } catch (err) {
      console.error("Failed to execute multi-agent goal", err);
    }
  };

  return (
    <div className="container max-w-7xl mx-auto py-8 px-4 space-y-8">
      <div className="border-b pb-6 flex items-center justify-between">
        <div>
          <h1 className="text-3xl font-extrabold tracking-tight flex items-center gap-2">
            <Users className="w-8 h-8 text-indigo-500" /> Multi-Agent Intelligence System
          </h1>
          <p className="text-muted-foreground mt-1">
            Coordinated multi-agent collaboration via structured message bus, task planner, and consensus resolver
          </p>
        </div>
        <span className="px-3 py-1 rounded-full text-xs font-mono font-bold bg-indigo-500/10 text-indigo-500 border border-indigo-500/20">
          13 Specialist Agents Online
        </span>
      </div>

      <MultiAgentDashboardComponent onProcessGoal={handleProcessGoal} />

      <AgentHealthGridComponent metrics={health} />

      {activeGraph && (
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <TaskGraphVisualizerComponent graph={activeGraph} />
          <CoordinatorViewComponent consensus={consensus} />
        </div>
      )}

      <MessageBusMonitorComponent messages={messages} />
    </div>
  );
}
