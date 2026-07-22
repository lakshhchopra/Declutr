"use client";

import React from "react";
import { Compass, Briefcase, Heart, Plane, Home, GraduationCap, DollarSign, Shield, Users, Sparkles } from "lucide-react";

export interface LifeAreaItem {
  id: string;
  name: string;
  description: string;
  icon: string;
  color: string;
}

interface LifeAreaGridProps {
  areas: LifeAreaItem[];
  onSelectArea: (area: LifeAreaItem) => void;
}

export function LifeAreaGridComponent({ areas, onSelectArea }: LifeAreaGridProps) {
  return (
    <div className="space-y-4">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-xl font-bold tracking-tight">Life Areas</h3>
          <p className="text-sm text-muted-foreground">Your digital life organized by domains, not file folders</p>
        </div>
      </div>

      <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-4">
        {areas.map((area) => (
          <div
            key={area.id}
            onClick={() => onSelectArea(area)}
            className="p-4 rounded-xl border bg-card hover:border-indigo-500/50 hover:shadow-md transition-all cursor-pointer space-y-2"
          >
            <div className="p-2.5 rounded-lg bg-indigo-500/10 text-indigo-500 w-fit">
              <Compass className="w-5 h-5" />
            </div>
            <div>
              <h4 className="font-bold text-sm text-foreground">{area.name}</h4>
              <p className="text-[11px] text-muted-foreground line-clamp-1">{area.description}</p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
