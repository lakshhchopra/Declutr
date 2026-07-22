"use client";

import React, { useState } from "react";
import { KeyRound, Plus, Trash2, Check, Copy, AlertTriangle } from "lucide-react";

export interface APIKey {
  id: string;
  name: string;
  key_prefix: string;
  scopes: string[];
  expires_at: string;
  created_at: string;
}

interface APIKeyManagerProps {
  apiKeys: APIKey[];
  onCreateKey: (name: string, scopes: string[]) => Promise<string | void>;
  onRevokeKey: (keyId: string) => void;
}

const AVAILABLE_SCOPES = [
  "vault.read",
  "vault.write",
  "asset.read",
  "asset.write",
  "workflow.execute",
  "ai.chat",
  "search.query",
  "backup.manage",
  "admin.manage",
];

export function APIKeyManagerComponent({ apiKeys, onCreateKey, onRevokeKey }: APIKeyManagerProps) {
  const [showModal, setShowModal] = useState(false);
  const [name, setName] = useState("");
  const [selectedScopes, setSelectedScopes] = useState<string[]>(["vault.read", "asset.read"]);
  const [generatedSecret, setGeneratedSecret] = useState<string | null>(null);
  const [copied, setCopied] = useState(false);

  const toggleScope = (scope: string) => {
    if (selectedScopes.includes(scope)) {
      setSelectedScopes(selectedScopes.filter((s) => s !== scope));
    } else {
      setSelectedScopes([...selectedScopes, scope]);
    }
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (name) {
      const rawSecret = await onCreateKey(name, selectedScopes);
      if (rawSecret) {
        setGeneratedSecret(rawSecret);
      }
      setName("");
    }
  };

  const copyToClipboard = () => {
    if (generatedSecret) {
      navigator.clipboard.writeText(generatedSecret);
      setCopied(true);
      setTimeout(() => setCopied(false), 2000);
    }
  };

  return (
    <div className="space-y-6">
      <div className="flex items-center justify-between">
        <div>
          <h3 className="text-xl font-bold tracking-tight">API Key Management</h3>
          <p className="text-sm text-muted-foreground">Manage scoped secret keys for third-party scripts, CLI tools, and AI agents</p>
        </div>
        <button
          onClick={() => {
            setGeneratedSecret(null);
            setShowModal(true);
          }}
          className="inline-flex items-center gap-2 px-4 py-2 rounded-lg bg-indigo-600 hover:bg-indigo-700 text-white text-xs font-semibold shadow-sm transition-all"
        >
          <Plus className="w-4 h-4" /> Create API Key
        </button>
      </div>

      <div className="rounded-xl border bg-card overflow-hidden">
        <table className="w-full text-left text-xs">
          <thead className="bg-secondary/50 border-b text-muted-foreground uppercase font-mono">
            <tr>
              <th className="p-3">Key Name</th>
              <th className="p-3">Prefix</th>
              <th className="p-3">Scopes</th>
              <th className="p-3">Expires</th>
              <th className="p-3 text-right">Actions</th>
            </tr>
          </thead>
          <tbody className="divide-y">
            {apiKeys.map((k) => (
              <tr key={k.id} className="hover:bg-secondary/30 transition-colors">
                <td className="p-3 font-semibold text-foreground">{k.name}</td>
                <td className="p-3 font-mono text-muted-foreground">{k.key_prefix}</td>
                <td className="p-3">
                  <div className="flex flex-wrap gap-1">
                    {k.scopes.map((s) => (
                      <span key={s} className="px-2 py-0.5 rounded bg-indigo-500/10 text-indigo-500 border border-indigo-500/20 text-[10px] font-mono">
                        {s}
                      </span>
                    ))}
                  </div>
                </td>
                <td className="p-3 font-mono text-muted-foreground">
                  {new Date(k.expires_at).toLocaleDateString()}
                </td>
                <td className="p-3 text-right">
                  <button
                    onClick={() => onRevokeKey(k.id)}
                    className="p-1.5 rounded hover:bg-rose-500/10 text-rose-500 transition-colors"
                  >
                    <Trash2 className="w-4 h-4" />
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {showModal && (
        <div className="fixed inset-0 bg-black/60 backdrop-blur-sm flex items-center justify-center p-4 z-50">
          <div className="bg-card border rounded-xl p-6 w-full max-w-lg space-y-4 shadow-xl">
            {generatedSecret ? (
              <div className="space-y-4">
                <div className="flex items-center gap-2 text-amber-500 font-bold">
                  <AlertTriangle className="w-5 h-5" />
                  <span>Save Your API Secret Key</span>
                </div>
                <p className="text-xs text-muted-foreground">
                  Please copy your secret key now. You will not be able to see it again!
                </p>
                <div className="p-3 rounded-lg bg-black/90 font-mono text-xs text-emerald-400 flex items-center justify-between border">
                  <span className="truncate">{generatedSecret}</span>
                  <button onClick={copyToClipboard} className="p-1 text-white hover:text-indigo-400">
                    {copied ? <Check className="w-4 h-4 text-emerald-400" /> : <Copy className="w-4 h-4" />}
                  </button>
                </div>
                <div className="flex justify-end pt-2">
                  <button
                    onClick={() => setShowModal(false)}
                    className="px-4 py-2 rounded-lg bg-indigo-600 text-white text-xs font-semibold"
                  >
                    Done
                  </button>
                </div>
              </div>
            ) : (
              <form onSubmit={handleSubmit} className="space-y-4">
                <h3 className="text-lg font-bold">Generate Scoped API Key</h3>
                <div>
                  <label className="text-xs font-semibold text-muted-foreground">Key Name</label>
                  <input
                    type="text"
                    required
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="e.g. Zapier Integration Worker"
                    className="w-full mt-1 px-3 py-2 border rounded-lg bg-secondary text-xs"
                  />
                </div>
                <div>
                  <label className="text-xs font-semibold text-muted-foreground mb-2 block">Granular Scopes</label>
                  <div className="grid grid-cols-3 gap-2">
                    {AVAILABLE_SCOPES.map((sc) => (
                      <button
                        key={sc}
                        type="button"
                        onClick={() => toggleScope(sc)}
                        className={`p-2 rounded-lg border text-[10px] font-mono text-left flex items-center justify-between ${
                          selectedScopes.includes(sc)
                            ? "bg-indigo-500/10 border-indigo-500 text-indigo-500 font-bold"
                            : "bg-secondary text-muted-foreground"
                        }`}
                      >
                        <span>{sc}</span>
                        {selectedScopes.includes(sc) && <Check className="w-3.5 h-3.5" />}
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
                    Generate Key
                  </button>
                </div>
              </form>
            )}
          </div>
        </div>
      )}
    </div>
  );
}
