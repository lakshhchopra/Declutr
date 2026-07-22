"use client";

import React, { useState } from "react";
import { Webhook, Plus, RefreshCw, Check, ShieldAlert, History } from "lucide-react";

export interface WebhookEndpoint {
  id: string;
  url: string;
  secret: string;
  events: string[];
  is_enabled: boolean;
}

export interface WebhookDelivery {
  id: string;
  event_type: string;
  response_status_code: number;
  latency_ms: number;
  success: boolean;
  delivered_at: string;
}

export interface DLQItem {
  id: string;
  event_type: string;
  last_error: string;
  attempts: number;
  failed_at: string;
}

interface WebhookManagerProps {
  webhooks: WebhookEndpoint[];
  deliveries: WebhookDelivery[];
  dlqItems: DLQItem[];
  onRegisterWebhook: (url: string, events: string[]) => void;
}

const AVAILABLE_EVENTS = [
  "asset.uploaded",
  "asset.updated",
  "context.created",
  "workflow.finished",
  "backup.completed",
  "search.completed",
  "memory.created",
  "relationship.added",
  "user.invited",
  "organization.created",
];

export function WebhookManagerComponent({
  webhooks,
  deliveries,
  dlqItems,
  onRegisterWebhook,
}: WebhookManagerProps) {
  const [showModal, setShowModal] = useState(false);
  const [url, setUrl] = useState("");
  const [selectedEvents, setSelectedEvents] = useState<string[]>(["asset.uploaded", "workflow.finished"]);

  const toggleEvent = (ev: string) => {
    if (selectedEvents.includes(ev)) {
      setSelectedEvents(selectedEvents.filter((e) => e !== ev));
    } else {
      setSelectedEvents([...selectedEvents, ev]);
    }
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (url) {
      onRegisterWebhook(url, selectedEvents);
      setUrl("");
      setShowModal(false);
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-xl font-bold tracking-tight">Webhook Event Engine & DLQ</h3>
          <p className="text-sm text-muted-foreground">Register webhook URLs with HMAC-SHA256 signature verification and exponential backoff DLQ</p>
        </div>
        <button
          onClick={() => setShowModal(true)}
          className="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold shadow-sm transition-all"
        >
          <Plus className="w-4 h-4" /> Register Webhook
        </button>
      </div>

      {/* Webhooks Grid */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {webhooks.map((w) => (
          <div key={w.id} className="p-4 rounded-xl border bg-card space-y-3">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-2 truncate">
                <Webhook className="w-4 h-4 text-indigo-500 shrink-0" />
                <h4 className="font-bold text-xs font-mono truncate">{w.url}</h4>
              </div>
              <span className="text-[10px] font-mono px-2 py-0.5 rounded bg-emerald-500/10 text-emerald-500 font-bold border border-emerald-500/20">
                ACTIVE
              </span>
            </div>
            <div className="text-[11px] font-mono text-muted-foreground truncate">
              Secret: <span className="text-foreground">{w.secret}</span>
            </div>
            <div className="flex flex-wrap gap-1">
              {w.events.map((e) => (
                <span key={e} className="px-2 py-0.5 rounded bg-secondary text-muted-foreground text-[10px] font-mono">
                  {e}
                </span>
              ))}
            </div>
          </div>
        ))}
      </div>

      {/* Deliveries & DLQ Inspector */}
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <div className="space-y-3">
          <div className="flex items-center gap-2 text-xs font-bold text-muted-foreground uppercase tracking-wider">
            <History className="w-4 h-4 text-indigo-500" /> Recent Delivery Logs
          </div>
          <div className="rounded-xl border bg-card overflow-hidden">
            <table className="w-full text-left text-[11px] font-mono">
              <thead className="bg-secondary/50 border-b text-muted-foreground">
                <tr>
                  <th className="p-2">Event</th>
                  <th className="p-2">Status</th>
                  <th className="p-2">Latency</th>
                </tr>
              </thead>
              <tbody className="divide-y">
                {deliveries.slice(0, 5).map((d) => (
                  <tr key={d.id}>
                    <td className="p-2 font-bold text-foreground">{d.event_type}</td>
                    <td className="p-2">
                      <span className={`px-2 py-0.5 rounded text-[10px] ${d.success ? "bg-emerald-500/10 text-emerald-500" : "bg-rose-500/10 text-rose-500"}`}>
                        HTTP {d.response_status_code}
                      </span>
                    </td>
                    <td className="p-2 text-muted-foreground">{d.latency_ms} ms</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>

        <div className="space-y-3">
          <div className="flex items-center gap-2 text-xs font-bold text-rose-500 uppercase tracking-wider">
            <ShieldAlert className="w-4 h-4" /> Dead Letter Queue (DLQ)
          </div>
          <div className="rounded-xl border bg-card overflow-hidden">
            <table className="w-full text-left text-[11px] font-mono">
              <thead className="bg-secondary/50 border-b text-muted-foreground">
                <tr>
                  <th className="p-2">Event</th>
                  <th className="p-2">Error</th>
                  <th className="p-2">Attempts</th>
                </tr>
              </thead>
              <tbody className="divide-y">
                {dlqItems.slice(0, 5).map((item) => (
                  <tr key={item.id}>
                    <td className="p-2 font-bold text-rose-500">{item.event_type}</td>
                    <td className="p-2 text-muted-foreground truncate max-w-[120px]">{item.last_error}</td>
                    <td className="p-2 font-bold text-foreground">{item.attempts}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        </div>
      </div>

      {showModal && (
        <div className="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center p-4 z-50">
          <form onSubmit={handleSubmit} className="bg-card border rounded-xl p-6 w-full max-w-lg space-y-4 shadow-xl">
            <h3 className="text-lg font-bold">Register Webhook Endpoint</h3>
            <div>
              <label className="text-xs font-semibold text-muted-foreground">Payload URL</label>
              <input
                type="url"
                required
                value={url}
                onChange={(e) => setUrl(e.target.value)}
                placeholder="https://api.myapp.com/webhooks/declutr"
                className="w-full mt-1 px-3 py-2 border rounded-lg bg-secondary text-xs font-mono"
              />
            </div>
            <div>
              <label className="text-xs font-semibold text-muted-foreground mb-2 block">Subscribe to Events</label>
              <div className="grid grid-cols-2 gap-2">
                {AVAILABLE_EVENTS.map((ev) => (
                  <button
                    key={ev}
                    type="button"
                    onClick={() => toggleEvent(ev)}
                    className={`p-2 rounded-lg border text-[10px] font-mono text-left flex items-center justify-between ${
                      selectedEvents.includes(ev)
                        ? "bg-indigo-500/10 border-indigo-500 text-indigo-500 font-bold"
                        : "bg-secondary text-muted-foreground"
                    }`}
                  >
                    <span>{ev}</span>
                    {selectedEvents.includes(ev) && <Check className="w-3.5 h-3.5" />}
                  </button>
                ))}
              </div>
            </div>
            <div className="flex justify-end gap-2 pt-2">
              <button
                type="button"
                onClick={() => setShowModal(false)}
                className="px-4 py-2 rounded-lg border text-xs font-semibold"
              >
                Cancel
              </button>
              <button
                type="submit"
                className="px-4 py-2 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold"
              >
                Register Webhook
              </button>
            </div>
          </form>
        </div>
      )}
    </div>
  );
}
