"use client";

import React from "react";
import { FolderKanban } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { EmptyState } from "../../shared/components/feedback/empty-state";

export default function CollectionsPage() {
  return (
    <PageShell
      title="Context Collections"
      subtitle="Grouped digital memory items linked by temporal, spatial, or entity relationships."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Collections" }]}
    >
      <EmptyState
        icon={<FolderKanban className="h-6 w-6" />}
        title="No Collections Formed"
        description="The Relationship Engine will automatically group related items (such as trips, events, or projects) here."
      />
    </PageShell>
  );
}
