"use client";

import React from "react";
import { FolderKey, Upload, Filter } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { Button } from "../../shared/components/ui/button";
import { SearchInput } from "../../shared/components/ui/input";
import { EmptyState } from "../../shared/components/feedback/empty-state";

export default function VaultPage() {
  return (
    <PageShell
      title="Vault Workspaces"
      subtitle="Client-side encrypted files, documents, receipts, and media items."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Vault Workspaces" }]}
      actions={
        <div className="flex gap-2">
          <Button variant="outline" leftIcon={<Filter className="h-4 w-4" />}>
            Filter
          </Button>
          <Button variant="default" leftIcon={<Upload className="h-4 w-4" />}>
            Upload File
          </Button>
        </div>
      }
    >
      <div className="mb-6">
        <SearchInput placeholder="Search within active vault..." />
      </div>

      <EmptyState
        title="Vault Workspace is Empty"
        description="Drag and drop documents or images to encrypt and upload directly to cloud storage."
        actionLabel="Initiate Upload"
        onAction={() => alert("Upload dialog placeholder")}
      />
    </PageShell>
  );
}
