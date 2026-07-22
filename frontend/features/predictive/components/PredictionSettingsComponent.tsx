"use client";

import React, { useState } from "react";
import { Settings, Shield, Sliders } from "lucide-react";

export interface PredictionSettingsData {
  user_id: string;
  min_confidence: number;
  learning_paused: boolean;
  auto_dismiss_expired: boolean;
}

interface PredictionSettingsProps {
  settings: PredictionSettingsData;
  onUpdate: (updated: PredictionSettingsData) => void;
}

export function PredictionSettingsComponent({ settings, onUpdate }: PredictionSettingsProps) {
  const [minConf, setMinConf] = useState(settings.min_confidence || 0.8);
  const [learningPaused, setLearningPaused] = useState(settings.learning_paused || false);

  const handleSave = (e: React.FormEvent) => {
    e.preventDefault();
    onUpdate({
      ...settings,
      min_confidence: parseFloat(minConf.toString()),
      learning_paused: learningPaused,
    });
  };

  return (
    <div className="p-5 rounded-xl border bg-card space-y-4">
      <div className="flex items-center gap-2 font-bold text-sm">
        <Sliders className="w-4 h-4 text-indigo-500" />
        <span>Predictive Intelligence Controls & Confidence Thresholds</span>
      </div>

      <form onSubmit={handleSave} className="space-y-4 text-xs">
        <div>
          <div className="flex justify-between font-semibold">
            <label className="text-muted-foreground">Minimum Confidence Threshold</label>
            <span className="font-mono text-indigo-500 font-bold">{(minConf * 100).toFixed(0)}%</span>
          </div>
          <input
            type="range"
            min="0.5"
            max="0.99"
            step="0.05"
            value={minConf}
            onChange={(e) => setMinConf(parseFloat(e.target.value))}
            className="w-full mt-2 accent-indigo-500"
          />
        </div>

        <div className="flex items-center justify-between pt-2 border-t">
          <div>
            <span className="font-bold text-foreground block">Pause Pattern Learning</span>
            <span className="text-[10px] text-muted-foreground">Temporarily halt scanning vault history for predictive recommendations</span>
          </div>
          <input
            type="checkbox"
            checked={learningPaused}
            onChange={(e) => setLearningPaused(e.target.checked)}
            className="w-4 h-4 accent-indigo-500"
          />
        </div>

        <button
          type="submit"
          className="w-full py-2 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white font-semibold shadow-sm transition-all"
        >
          Save Intelligence Preferences
        </button>
      </form>
    </div>
  );
}
