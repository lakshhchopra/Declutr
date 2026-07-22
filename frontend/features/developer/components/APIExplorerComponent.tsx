"use client";

import React, { useState } from "react";
import { Play, Code, Check } from "lucide-react";

export function APIExplorerComponent() {
  const [selectedEndpoint, setSelectedEndpoint] = useState("POST /api/v1/search/query");
  const [response, setResponse] = useState<any>(null);
  const [loading, setLoading] = useState(false);

  const handleTestEndpoint = async () => {
    setLoading(true);
    try {
      if (selectedEndpoint.includes("search")) {
        const res = await fetch("/api/v1/search/query", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ query: "developer integration" }),
        });
        const data = await res.json();
        setResponse(data);
      } else {
        const res = await fetch("/api/v1/developer/openapi");
        const data = await res.json();
        setResponse(data);
      }
    } catch (err) {
      setResponse({ error: String(err) });
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="space-y-6">
      <div>
        <h3 className="text-xl font-bold tracking-tight">Interactive OpenAPI Explorer</h3>
        <p className="text-sm text-muted-foreground">Test public REST API endpoints directly in your browser with live response payloads</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div className="space-y-2 border-r pr-4">
          <label className="text-xs font-semibold text-muted-foreground uppercase tracking-wider">Select Endpoint</label>
          <div className="space-y-1">
            {["POST /api/v1/search/query", "POST /api/v1/copilot/messages", "GET /api/v1/developer/openapi"].map((ep) => (
              <button
                key={ep}
                onClick={() => setSelectedEndpoint(ep)}
                className={`w-full p-2.5 rounded-lg text-left font-mono text-xs transition-colors ${
                  selectedEndpoint === ep ? "bg-indigo-600 text-white font-bold" : "bg-secondary hover:bg-accent text-foreground"
                }`}
              >
                {ep}
              </button>
            ))}
          </div>

          <button
            onClick={handleTestEndpoint}
            disabled={loading}
            className="w-full mt-4 flex items-center justify-center gap-2 py-2 rounded-lg bg-emerald-600 hover:bg-emerald-700 text-white text-xs font-semibold shadow transition-all"
          >
            <Play className="w-4 h-4" /> {loading ? "Executing..." : "Execute Request"}
          </button>
        </div>

        <div className="md:col-span-2 space-y-2">
          <div className="flex items-center justify-between">
            <span className="text-xs font-semibold text-muted-foreground uppercase tracking-wider">Response Payload</span>
            <span className="text-[10px] font-mono px-2 py-0.5 rounded bg-emerald-500/10 text-emerald-500 font-bold border border-emerald-500/20">
              HTTP 200 OK
            </span>
          </div>

          <div className="p-4 rounded-xl bg-black/90 text-emerald-400 border font-mono text-xs min-h-[300px] overflow-auto">
            {response ? (
              <pre>{JSON.stringify(response, null, 2)}</pre>
            ) : (
              <span className="text-muted-foreground italic">// Click 'Execute Request' to view real-time JSON response</span>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
