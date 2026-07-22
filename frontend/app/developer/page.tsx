"use client";

import React, { useEffect, useState } from "react";
import { DeveloperDashboardComponent } from "@/features/developer/components/DeveloperDashboardComponent";
import { APIKeyManagerComponent, APIKey } from "@/features/developer/components/APIKeyManagerComponent";
import { WebhookManagerComponent, WebhookEndpoint, WebhookDelivery, DLQItem } from "@/features/developer/components/WebhookManagerComponent";
import { OAuthClientManagerComponent, OAuthClient } from "@/features/developer/components/OAuthClientManagerComponent";
import { APIExplorerComponent } from "@/features/developer/components/APIExplorerComponent";
import { SDKDownloadComponent } from "@/features/developer/components/SDKDownloadComponent";
import { Code, KeyRound, Webhook, Cpu, Play, Download } from "lucide-react";

export default function DeveloperPage() {
  const [activeTab, setActiveTab] = useState<"dashboard" | "keys" | "webhooks" | "oauth" | "explorer" | "sdks">("dashboard");

  const [apiKeys, setApiKeys] = useState<APIKey[]>([]);
  const [webhooks, setWebhooks] = useState<WebhookEndpoint[]>([]);
  const [deliveries, setDeliveries] = useState<WebhookDelivery[]>([]);
  const [dlqItems, setDlqItems] = useState<DLQItem[]>([]);
  const [oauthClients, setOauthClients] = useState<OAuthClient[]>([]);

  const fetchDeveloperData = async () => {
    try {
      const [keysRes, hooksRes, oauthRes, logsRes] = await Promise.all([
        fetch("/api/v1/developer/keys").then((r) => r.json()).catch(() => []),
        fetch("/api/v1/developer/webhooks").then((r) => r.json()).catch(() => []),
        fetch("/api/v1/developer/oauth/apps").then((r) => r.json()).catch(() => []),
        fetch("/api/v1/developer/webhooks/deliveries").then((r) => r.json()).catch(() => ({ deliveries: [], dlq_items: [] })),
      ]);

      if (Array.isArray(keysRes)) setApiKeys(keysRes);
      if (Array.isArray(hooksRes)) setWebhooks(hooksRes);
      if (Array.isArray(oauthRes)) setOauthClients(oauthRes);
      if (logsRes) {
        setDeliveries(logsRes.deliveries || []);
        setDlqItems(logsRes.dlq_items || []);
      }
    } catch (err) {
      console.error("Failed to load developer data", err);
    }
  };

  useEffect(() => {
    fetchDeveloperData();
  }, []);

  const handleCreateKey = async (name: string, scopes: string[]): Promise<string | void> => {
    try {
      const res = await fetch("/api/v1/developer/keys", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, scopes, expires_in_days: 365 }),
      });
      if (res.ok) {
        const data = await res.json();
        await fetchDeveloperData();
        return data.raw_secret;
      }
    } catch (err) {
      console.error("Failed to generate API key", err);
    }
  };

  const handleRevokeKey = async (keyId: string) => {
    await fetch(`/api/v1/developer/keys?id=${keyId}`, { method: "DELETE" });
    fetchDeveloperData();
  };

  const handleRegisterWebhook = async (url: string, events: string[]) => {
    await fetch("/api/v1/developer/webhooks", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ url, events }),
    });
    fetchDeveloperData();
  };

  const handleCreateOAuthClient = async (name: string, redirectUris: string[], scopes: string[]) => {
    await fetch("/api/v1/developer/oauth/apps", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, redirect_uris: redirectUris, scopes }),
    });
    fetchDeveloperData();
  };

  return (
    <div className="container max-w-7xl mx-auto py-8 px-4 space-y-8">
      {/* Header */}
      <div className="border-b pb-6">
        <h1 className="text-3xl font-extrabold tracking-tight">Developer Portal & Public API</h1>
        <p className="text-muted-foreground mt-1">
          Build custom applications, AI agents, automations, and webhooks on top of Declutr
        </p>
      </div>

      {/* Navigation Tabs */}
      <div className="flex border-b space-x-4 overflow-x-auto">
        <button
          onClick={() => setActiveTab("dashboard")}
          className={`pb-3 text-sm font-semibold flex items-center gap-2 border-b-2 transition-all ${
            activeTab === "dashboard" ? "border-accent text-accent" : "border-transparent text-muted-foreground hover:text-foreground"
          }`}
        >
          <Code className="w-4 h-4" /> Overview
        </button>
        <button
          onClick={() => setActiveTab("keys")}
          className={`pb-3 text-sm font-semibold flex items-center gap-2 border-b-2 transition-all ${
            activeTab === "keys" ? "border-accent text-accent" : "border-transparent text-muted-foreground hover:text-foreground"
          }`}
        >
          <KeyRound className="w-4 h-4" /> API Keys
        </button>
        <button
          onClick={() => setActiveTab("webhooks")}
          className={`pb-3 text-sm font-semibold flex items-center gap-2 border-b-2 transition-all ${
            activeTab === "webhooks" ? "border-accent text-accent" : "border-transparent text-muted-foreground hover:text-foreground"
          }`}
        >
          <Webhook className="w-4 h-4" /> Webhooks & DLQ
        </button>
        <button
          onClick={() => setActiveTab("oauth")}
          className={`pb-3 text-sm font-semibold flex items-center gap-2 border-b-2 transition-all ${
            activeTab === "oauth" ? "border-accent text-accent" : "border-transparent text-muted-foreground hover:text-foreground"
          }`}
        >
          <Cpu className="w-4 h-4" /> OAuth 2.1 Apps
        </button>
        <button
          onClick={() => setActiveTab("explorer")}
          className={`pb-3 text-sm font-semibold flex items-center gap-2 border-b-2 transition-all ${
            activeTab === "explorer" ? "border-accent text-accent" : "border-transparent text-muted-foreground hover:text-foreground"
          }`}
        >
          <Play className="w-4 h-4" /> API Explorer
        </button>
        <button
          onClick={() => setActiveTab("sdks")}
          className={`pb-3 text-sm font-semibold flex items-center gap-2 border-b-2 transition-all ${
            activeTab === "sdks" ? "border-accent text-accent" : "border-transparent text-muted-foreground hover:text-foreground"
          }`}
        >
          <Download className="w-4 h-4" /> SDKs & CLI
        </button>
      </div>

      {/* Main Content Areas */}
      <div>
        {activeTab === "dashboard" && (
          <DeveloperDashboardComponent
            apiKeysCount={apiKeys.length}
            webhooksCount={webhooks.length}
            dlqCount={dlqItems.length}
            oauthAppsCount={oauthClients.length}
          />
        )}
        {activeTab === "keys" && (
          <APIKeyManagerComponent
            apiKeys={apiKeys}
            onCreateKey={handleCreateKey}
            onRevokeKey={handleRevokeKey}
          />
        )}
        {activeTab === "webhooks" && (
          <WebhookManagerComponent
            webhooks={webhooks}
            deliveries={deliveries}
            dlqItems={dlqItems}
            onRegisterWebhook={handleRegisterWebhook}
          />
        )}
        {activeTab === "oauth" && (
          <OAuthClientManagerComponent
            clients={oauthClients}
            onCreateClient={handleCreateOAuthClient}
          />
        )}
        {activeTab === "explorer" && <APIExplorerComponent />}
        {activeTab === "sdks" && <SDKDownloadComponent />}
      </div>
    </div>
  );
}
