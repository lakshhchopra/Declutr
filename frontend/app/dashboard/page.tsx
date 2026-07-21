"use client";

import React from "react";
import { FolderKey, Search, ShieldCheck, Activity, Plus } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "../../shared/components/ui/card";
import { Badge } from "../../shared/components/ui/badge";
import { Button } from "../../shared/components/ui/button";
import { Grid } from "../../shared/components/layout/layout-primitives";
import { EmptyState } from "../../shared/components/feedback/empty-state";

export default function DashboardPage() {
  return (
    <PageShell
      title="Dashboard"
      subtitle="Overview of your zero-knowledge vaults, intelligent indexing, and session telemetry."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Dashboard" }]}
      actions={
        <Button variant="default" leftIcon={<Plus className="h-4 w-4" />}>
          New Vault
        </Button>
      }
    >
      <Grid cols={3} className="mb-8">
        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle>Active Vaults</CardTitle>
              <FolderKey className="h-5 w-5 text-emerald-400" />
            </div>
            <CardDescription>AES-256 Encrypted Workspaces</CardDescription>
          </CardHeader>
          <CardContent>
            <span className="text-3xl font-extrabold text-white">1</span>
            <Badge variant="emerald" className="ml-3">Online</Badge>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle>Indexed Vectors</CardTitle>
              <Search className="h-5 w-5 text-blue-400" />
            </div>
            <CardDescription>512-dim pgvector HNSW Index</CardDescription>
          </CardHeader>
          <CardContent>
            <span className="text-3xl font-extrabold text-white">0</span>
            <Badge variant="outline" className="ml-3">Ready</Badge>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle>Session Risk Score</CardTitle>
              <Activity className="h-5 w-5 text-amber-400" />
            </div>
            <CardDescription>Behavioral Anomaly Engine</CardDescription>
          </CardHeader>
          <CardContent>
            <span className="text-3xl font-extrabold text-emerald-400">0.02</span>
            <Badge variant="emerald" className="ml-3">Trusted</Badge>
          </CardContent>
        </Card>
      </Grid>

      <EmptyState
        title="No Recent Vault Activity"
        description="Your stored assets will appear here once items are uploaded and processed by the Content Pipeline."
        actionLabel="Go to Vault"
        onAction={() => window.location.href = "/vault"}
      />
    </PageShell>
  );
}
