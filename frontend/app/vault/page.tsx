"use client";

import React, { useState, useEffect } from "react";
import { FolderKey, ShieldCheck, Database, HardDrive, Upload, Filter, RefreshCw, Layers } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "../../shared/components/ui/card";
import { Badge } from "../../shared/components/ui/badge";
import { Button } from "../../shared/components/ui/button";
import { SearchInput } from "../../shared/components/ui/input";
import { Grid } from "../../shared/components/layout/layout-primitives";
import { EmptyState } from "../../shared/components/feedback/empty-state";
import { VaultService, VaultData } from "../../features/vault/services/vault-service";
import { UploadModal } from "../../features/upload/components/upload-modal";

export default function VaultPage() {
  const [vault, setVault] = useState<VaultData | null>(null);
  const [uploadOpen, setUploadOpen] = useState(false);

  useEffect(() => {
    VaultService.getCurrentVault().then(setVault);
  }, []);

  const formatBytes = (bytes: number) => {
    if (bytes === 0) return "0 B";
    const k = 1024;
    const sizes = ["B", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(k));
    return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + " " + sizes[i];
  };

  return (
    <PageShell
      title={vault?.displayName || "My Life Vault"}
      subtitle="Root Zero-Knowledge Workspace Container. All digital assets, indices, and collections belong here."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Vault Workspace" }]}
      actions={
        <div className="flex items-center gap-2">
          <Badge variant="emerald" className="px-3 py-1 text-xs">
            <ShieldCheck className="h-3.5 w-3.5 mr-1" /> AES-256 ENCRYPTED
          </Badge>
          <Button variant="default" onClick={() => setUploadOpen(true)} leftIcon={<Upload className="h-4 w-4" />}>
            Upload Memory
          </Button>
        </div>
      }
    >
      <UploadModal
        open={uploadOpen}
        onOpenChange={setUploadOpen}
        onUploadComplete={() => {
          VaultService.getCurrentVault().then(setVault);
        }}
      />

      {/* Search & Action Bar */}
      <div className="mb-6 flex flex-col sm:flex-row gap-3 items-center justify-between">
        <div className="w-full sm:max-w-md">
          <SearchInput placeholder="Search within active vault workspace..." />
        </div>
        <div className="flex gap-2 w-full sm:w-auto">
          <Button variant="outline" size="sm" leftIcon={<Filter className="h-3.5 w-3.5" />}>
            Filter
          </Button>
          <Button variant="ghost" size="sm" leftIcon={<RefreshCw className="h-3.5 w-3.5" />}>
            Sync Index
          </Button>
        </div>
      </div>

      {/* Storage Overview & Statistics Cards */}
      <Grid cols={3} className="mb-8">
        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle>Storage Utilization</CardTitle>
              <HardDrive className="h-5 w-5 text-emerald-400" />
            </div>
            <CardDescription>Encrypted S3 Object Storage</CardDescription>
          </CardHeader>
          <CardContent>
            <span className="text-3xl font-extrabold text-white">
              {formatBytes(vault?.storageUsageBytes || 0)}
            </span>
            <span className="text-xs text-slate-400 block mt-1">10 GB Free Tier Quota</span>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle>Digital Assets</CardTitle>
              <Database className="h-5 w-5 text-blue-400" />
            </div>
            <CardDescription>Indexed Documents & Files</CardDescription>
          </CardHeader>
          <CardContent>
            <span className="text-3xl font-extrabold text-white">
              {vault?.itemCount || 0}
            </span>
            <Badge variant="outline" className="ml-3">0 Pending</Badge>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <div className="flex items-center justify-between">
              <CardTitle>Context Collections</CardTitle>
              <Layers className="h-5 w-5 text-amber-400" />
            </div>
            <CardDescription>Temporal Event Groupings</CardDescription>
          </CardHeader>
          <CardContent>
            <span className="text-3xl font-extrabold text-white">
              {vault?.collectionCount || 0}
            </span>
            <Badge variant="emerald" className="ml-3">Auto-Cluster Ready</Badge>
          </CardContent>
        </Card>
      </Grid>

      {/* Premium Empty State */}
      <EmptyState
        icon={<FolderKey className="h-8 w-8 text-emerald-400" />}
        title="Your Vault Workspace is Ready"
        description="Start by uploading your first memory, document, receipt, or audio file. Your Master Vault Key will encrypt items before transmission."
        actionLabel="Upload First File"
        onAction={() => setUploadOpen(true)}
      />
    </PageShell>
  );
}
