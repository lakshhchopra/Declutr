"use client";

import React, { useEffect, useState } from "react";
import { LifeAreaGridComponent, LifeAreaItem } from "@/features/lifeos/components/LifeAreaGridComponent";
import { ProjectHubComponent, ProjectItem } from "@/features/lifeos/components/ProjectHubComponent";
import { GoalTrackerComponent, ProjectGoalItem } from "@/features/lifeos/components/GoalTrackerComponent";
import { LifeTimelineComponent, LifeTimelineEventItem } from "@/features/lifeos/components/LifeTimelineComponent";
import { Compass, Sparkles, CheckCircle2, Calendar } from "lucide-react";

export default function LifeOSPage() {
  const [dashboard, setDashboard] = useState<any>(null);

  const fetchDashboard = async () => {
    try {
      const res = await fetch("/api/v1/lifeos/dashboard");
      if (res.ok) {
        const data = await res.json();
        setDashboard(data);
      }
    } catch (err) {
      console.error("Failed to load LifeOS dashboard", err);
    }
  };

  useEffect(() => {
    fetchDashboard();
  }, []);

  const handleCreateProject = async (name: string, desc: string) => {
    await fetch("/api/v1/lifeos/projects", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ life_area_id: "area-Work", name, description: desc }),
    });
    fetchDashboard();
  };

  if (!dashboard) {
    return <div className="container max-w-7xl mx-auto py-12 px-4 text-center">Loading LifeOS...</div>;
  }

  return (
    <div className="container max-w-7xl mx-auto py-8 px-4 space-y-8">
      <div className="border-b pb-6 flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 className="text-3xl font-extrabold tracking-tight flex items-center gap-2">
            <Compass className="w-8 h-8 text-indigo-500" /> Life Operating System (LifeOS)
          </h1>
          <p className="text-muted-foreground mt-1">
            Operating system for your digital life — structured around Life Areas, Projects, and Goals instead of folders
          </p>
        </div>
        <div className="flex items-center gap-3 font-mono">
          <div className="p-3 rounded-xl border bg-card text-center">
            <div className="text-[10px] text-muted-foreground">Life Health Score</div>
            <div className="text-xl font-bold text-emerald-500">{dashboard.health_score}/100</div>
          </div>
        </div>
      </div>

      {/* Today's Priorities */}
      <div className="p-5 rounded-xl border bg-card space-y-2">
        <h4 className="font-bold text-sm text-foreground flex items-center gap-2">
          <CheckCircle2 className="w-4 h-4 text-indigo-500" /> Today's Priorities
        </h4>
        <div className="flex flex-wrap gap-2">
          {dashboard.priorities_today.map((item: string, idx: number) => (
            <span key={idx} className="px-3 py-1 rounded-lg bg-secondary border text-xs font-semibold text-foreground">
              {item}
            </span>
          ))}
        </div>
      </div>

      <LifeAreaGridComponent areas={dashboard.life_areas || []} onSelectArea={() => {}} />

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div className="lg:col-span-2 space-y-6">
          <ProjectHubComponent projects={dashboard.active_projects || []} onCreateProject={handleCreateProject} />
        </div>
        <div className="space-y-6">
          <GoalTrackerComponent goals={dashboard.active_goals || []} />
          <LifeTimelineComponent events={dashboard.recent_timeline || []} />
        </div>
      </div>
    </div>
  );
}
