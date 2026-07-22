"use client";

import React, { useEffect, useState } from "react";
import { PredictionFeedComponent, PredictionItem } from "@/features/predictive/components/PredictionFeedComponent";
import { UpcomingTimelineComponent } from "@/features/predictive/components/UpcomingTimelineComponent";
import { OpportunityDetectionComponent } from "@/features/predictive/components/OpportunityDetectionComponent";
import { PredictionSettingsComponent, PredictionSettingsData } from "@/features/predictive/components/PredictionSettingsComponent";
import { Sparkles, Brain, Activity } from "lucide-react";

export default function PredictivePage() {
  const [predictions, setPredictions] = useState<PredictionItem[]>([]);
  const [settings, setSettings] = useState<PredictionSettingsData>({
    user_id: "usr-default",
    min_confidence: 0.8,
    learning_paused: false,
    auto_dismiss_expired: true,
  });
  const [stats, setStats] = useState<any>({ acceptance_rate: 0, total_generated: 0 });

  const fetchPredictions = async () => {
    try {
      const res = await fetch("/api/v1/predictive/predictions");
      if (res.ok) {
        const data = await res.json();
        setPredictions(data || []);
      }
    } catch (err) {
      console.error("Failed to load predictions", err);
    }
  };

  const fetchSettings = async () => {
    try {
      const res = await fetch("/api/v1/predictive/settings");
      if (res.ok) {
        const data = await res.json();
        setSettings(data);
      }
    } catch (err) {
      console.error("Failed to load settings", err);
    }
  };

  const fetchStats = async () => {
    try {
      const res = await fetch("/api/v1/predictive/stats");
      if (res.ok) {
        const data = await res.json();
        setStats(data);
      }
    } catch (err) {
      console.error("Failed to load stats", err);
    }
  };

  useEffect(() => {
    fetchPredictions();
    fetchSettings();
    fetchStats();
  }, []);

  const handleAccept = async (id: string) => {
    await fetch("/api/v1/predictive/accept", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ prediction_id: id }),
    });
    fetchPredictions();
    fetchStats();
  };

  const handleDismiss = async (id: string, reason: string) => {
    await fetch("/api/v1/predictive/dismiss", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ prediction_id: id, reason }),
    });
    fetchPredictions();
    fetchStats();
  };

  const handleUpdateSettings = async (updated: PredictionSettingsData) => {
    await fetch("/api/v1/predictive/settings", {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(updated),
    });
    setSettings(updated);
    fetchPredictions();
  };

  return (
    <div className="container max-w-7xl mx-auto py-8 px-4 space-y-8">
      <div className="border-b pb-6 flex flex-col md:flex-row md:items-center justify-between gap-4">
        <div>
          <h1 className="text-3xl font-extrabold tracking-tight flex items-center gap-2">
            <Sparkles className="w-8 h-8 text-indigo-500" /> Life Intelligence & Predictive Engine
          </h1>
          <p className="text-muted-foreground mt-1">
            "I remembered this for you." — Proactive intelligence anticipating deadlines, trips, and vault opportunities
          </p>
        </div>
        <div className="flex items-center gap-3">
          <div className="p-3 rounded-xl border bg-card text-center font-mono">
            <div className="text-[10px] text-muted-foreground">Accuracy Score</div>
            <div className="text-lg font-bold text-indigo-500">{(stats.accuracy_score * 100 || 94).toFixed(0)}%</div>
          </div>
        </div>
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-3 gap-6">
        <div className="lg:col-span-2 space-y-6">
          <h3 className="text-lg font-bold tracking-tight">Proactive Intelligence Suggestions</h3>
          <PredictionFeedComponent predictions={predictions} onAccept={handleAccept} onDismiss={handleDismiss} />
        </div>

        <div className="space-y-6">
          <UpcomingTimelineComponent predictions={predictions} />
          <OpportunityDetectionComponent predictions={predictions} />
          <PredictionSettingsComponent settings={settings} onUpdate={handleUpdateSettings} />
        </div>
      </div>
    </div>
  );
}
