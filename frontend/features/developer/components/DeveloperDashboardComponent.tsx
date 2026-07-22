"use client";

import React from "react";
import { KeyRound, Webhook, ShieldAlert, Cpu, ArrowUpRight, Terminal } from "lucide-react";

export interface DevDashboardProps {
  apiKeysCount: number;
  webhooksCount: number;
  dlqCount: number;
  oauthAppsCount: number;
}

export function DeveloperDashboardComponent({
  apiKeysCount,
  webhooksCount,
  dlqCount,
  oauthAppsCount,
}: DevDashboardProps) {
  return (
    <div className="space-y-6">
      {/* Overview Stat Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
        <div className="p-5 rounded-xl border bg-card space-y-2">
          <div className="flex items-center justify-between text-muted-foreground">
            <span className="text-xs font-semibold uppercase tracking-wider">Active API Keys</span>
            <KeyRound className="w-4 h-4 text-amber-500" />
          </div>
          <div className="text-2xl font-bold font-mono">{apiKeysCount}</div>
          <div className="text-xs text-muted-foreground">Scoped secret credentials</div>
        </div>

        <div className="p-5 rounded-xl border bg-card space-y-2">
          <div className="flex items-center justify-between text-muted-foreground">
            <span className="text-xs font-semibold uppercase tracking-wider">Registered Webhooks</span>
            <Webhook className="w-4 h-4 text-indigo-500" />
          </div>
          <div className="text-2xl font-bold font-mono">{webhooksCount}</div>
          <div className="text-xs text-muted-foreground">Event bus subscribers</div>
        </div>

        <div className="p-5 rounded-xl border bg-card space-y-2">
          <div className="flex items-center justify-between text-muted-foreground">
            <span className="text-xs font-semibold uppercase tracking-wider">Dead Letter Queue</span>
            <ShieldAlert className="w-4 h-4 text-rose-500" />
          </div>
          <div className="text-2xl font-bold font-mono">{dlqCount}</div>
          <div className="text-xs text-rose-500 font-medium">Failed delivery queue</div>
        </div>

        <div className="p-5 rounded-xl border bg-card space-y-2">
          <div className="flex items-center justify-between text-muted-foreground">
            <span className="text-xs font-semibold uppercase tracking-wider">OAuth 2.1 Clients</span>
            <Cpu className="w-4 h-4 text-emerald-500" />
          </div>
          <div className="text-2xl font-bold font-mono">{oauthAppsCount}</div>
          <div className="text-xs text-emerald-500 font-medium flex items-center gap-1">
            <ArrowUpRight className="w-3 h-3" /> PKCE Standard Enforced
          </div>
        </div>
      </div>

      {/* Quickstart Code Snippet Card */}
      <div className="p-6 rounded-xl border bg-card space-y-4">
        <div className="flex items-center gap-2">
          <Terminal className="w-5 h-5 text-indigo-500" />
          <h3 className="font-bold text-lg">Developer SDK Quickstart</h3>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-4 text-xs font-mono">
          <div className="p-4 rounded-lg bg-black/90 text-emerald-400 border space-y-2 overflow-x-auto">
            <div className="text-muted-foreground font-sans font-semibold text-[11px] uppercase">TypeScript SDK</div>
            <pre>{`import { DeclutrClient } from '@declutr/sdk';

const client = new DeclutrClient({
  apiKey: process.env.DECLUTR_API_KEY
});

const results = await client.search("quarterly report");
console.log(results);`}</pre>
          </div>

          <div className="p-4 rounded-lg bg-black/90 text-emerald-400 border space-y-2 overflow-x-auto">
            <div className="text-muted-foreground font-sans font-semibold text-[11px] uppercase">Python SDK</div>
            <pre>{`from declutr import DeclutrClient

client = DeclutrClient(
    api_key="declutr_live_..."
)

response = client.chat(
    conversation_id="conv-1",
    message="Summarize financial records"
)`}</pre>
          </div>
        </div>
      </div>
    </div>
  );
}
