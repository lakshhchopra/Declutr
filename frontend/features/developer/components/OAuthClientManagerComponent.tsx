"use client";

import React, { useState } from "react";
import { Cpu, Plus, Key } from "lucide-react";

export interface OAuthClient {
  id: string;
  name: string;
  client_id: string;
  redirect_uris: string[];
  scopes: string[];
}

interface OAuthClientManagerProps {
  clients: OAuthClient[];
  onCreateClient: (name: string, redirectUris: string[], scopes: string[]) => void;
}

export function OAuthClientManagerComponent({ clients, onCreateClient }: OAuthClientManagerProps) {
  const [showModal, setShowModal] = useState(false);
  const [name, setName] = useState("");
  const [redirectUri, setRedirectUri] = useState("https://myapp.com/oauth/callback");

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (name) {
      onCreateClient(name, [redirectUri], ["vault.read", "asset.read"]);
      setName("");
      setShowModal(false);
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-xl font-bold tracking-tight">OAuth 2.1 Applications</h3>
          <p className="text-sm text-muted-foreground">Register third-party OAuth 2.1 apps with PKCE authorization code grants</p>
        </div>
        <button
          onClick={() => setShowModal(true)}
          className="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold shadow-sm transition-all"
        >
          <Plus className="w-4 h-4" /> Register OAuth App
        </button>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {clients.map((c) => (
          <div key={c.id} className="p-4 rounded-xl border bg-card space-y-3">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-2">
                <Cpu className="w-4 h-4 text-emerald-500" />
                <h4 className="font-bold text-sm">{c.name}</h4>
              </div>
              <span className="text-[10px] font-mono px-2 py-0.5 rounded bg-secondary">
                OAuth 2.1 PKCE
              </span>
            </div>
            <div className="text-xs font-mono text-muted-foreground">
              Client ID: <span className="text-foreground">{c.client_id}</span>
            </div>
            <div className="text-[11px] font-mono text-muted-foreground truncate">
              Redirect URI: {c.redirect_uris?.join(", ")}
            </div>
          </div>
        ))}
      </div>

      {showModal && (
        <div className="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center p-4 z-50">
          <form onSubmit={handleSubmit} className="bg-card border rounded-xl p-6 w-full max-w-md space-y-4 shadow-xl">
            <h3 className="text-lg font-bold">Register OAuth 2.1 App</h3>
            <div>
              <label className="text-xs font-semibold text-muted-foreground">Application Name</label>
              <input
                type="text"
                required
                value={name}
                onChange={(e) => setName(e.target.value)}
                placeholder="e.g. Acme Mobile App"
                className="w-full mt-1 px-3 py-2 border rounded-lg bg-secondary text-xs"
              />
            </div>
            <div>
              <label className="text-xs font-semibold text-muted-foreground">Redirect URI</label>
              <input
                type="url"
                required
                value={redirectUri}
                onChange={(e) => setRedirectUri(e.target.value)}
                placeholder="https://myapp.com/oauth/callback"
                className="w-full mt-1 px-3 py-2 border rounded-lg bg-secondary text-xs font-mono"
              />
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
                Register App
              </button>
            </div>
          </form>
        </div>
      )}
    </div>
  );
}
