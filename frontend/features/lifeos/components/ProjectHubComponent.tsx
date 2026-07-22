"use client";

import React, { useState } from "react";
import { FolderKanban, Plus, Calendar, DollarSign, Users, ArrowRight } from "lucide-react";

export interface ProjectItem {
  id: string;
  name: string;
  description: string;
  status: string;
  budget?: number;
  people: string[];
  target_date?: string;
}

interface ProjectHubProps {
  projects: ProjectItem[];
  onCreateProject: (name: string, desc: string) => void;
}

export function ProjectHubComponent({ projects, onCreateProject }: ProjectHubProps) {
  const [name, setName] = useState("");
  const [desc, setDesc] = useState("");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (name) {
      onCreateProject(name, desc);
      setName("");
      setDesc("");
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-xl font-bold tracking-tight">Active Projects Hub</h3>
          <p className="text-sm text-muted-foreground">First-class project hubs containing knowledge, timeline, and workflow context</p>
        </div>
      </div>

      <form onSubmit={handleSubmit} className="p-4 rounded-xl border bg-card flex flex-col md:flex-row gap-3">
        <input
          type="text"
          required
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="New Project Title (e.g. Tax Filing 2027, Masters Application)"
          className="flex-1 px-3 py-2 border rounded-lg bg-secondary text-xs"
        />
        <input
          type="text"
          value={desc}
          onChange={(e) => setDesc(e.target.value)}
          placeholder="Description..."
          className="flex-1 px-3 py-2 border rounded-lg bg-secondary text-xs"
        />
        <button
          type="submit"
          className="px-4 py-2 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold flex items-center gap-1.5 whitespace-nowrap shadow-sm transition-all"
        >
          <Plus className="w-4 h-4" /> Create Project
        </button>
      </form>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {projects.map((prj) => (
          <div key={prj.id} className="p-5 rounded-xl border bg-card space-y-3 hover:border-indigo-500/50 transition-all shadow-sm">
            <div className="flex items-start justify-between">
              <div className="flex items-center gap-2">
                <div className="p-2 rounded-lg bg-indigo-500/10 text-indigo-500">
                  <FolderKanban className="w-5 h-5" />
                </div>
                <div>
                  <h4 className="font-bold text-sm text-foreground">{prj.name}</h4>
                  <span className="text-[10px] font-mono text-emerald-500 font-bold">{prj.status}</span>
                </div>
              </div>
            </div>

            <p className="text-xs text-muted-foreground">{prj.description}</p>

            <div className="pt-2 border-t flex items-center justify-between text-xs text-muted-foreground font-mono">
              {prj.budget ? <span>Budget: ${prj.budget.toLocaleString()}</span> : <span />}
              {prj.people.length > 0 && <span>People: {prj.people.join(", ")}</span>}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
