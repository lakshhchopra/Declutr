"use client";

import React from "react";
import { Search } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { SearchInput } from "../../shared/components/ui/input";
import { Badge } from "../../shared/components/ui/badge";
import { EmptyState } from "../../shared/components/feedback/empty-state";

export default function SearchPage() {
  return (
    <PageShell
      title="Semantic Search Engine"
      subtitle="Hybrid search combining PostgreSQL Full-Text Search and pgvector 512-dim cosine distance."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Search" }]}
    >
      <div className="mb-6 space-y-4">
        <SearchInput placeholder="Enter natural language query e.g. 'hotel booking receipt Mumbai'..." />
        <div className="flex items-center gap-2">
          <span className="text-xs text-slate-400">Quick Filters:</span>
          <Badge variant="emerald" className="cursor-pointer">Receipts</Badge>
          <Badge variant="blue" className="cursor-pointer">PDF Documents</Badge>
          <Badge variant="amber" className="cursor-pointer">Travel</Badge>
          <Badge variant="outline" className="cursor-pointer">Last 30 Days</Badge>
        </div>
      </div>

      <EmptyState
        icon={<Search className="h-6 w-6" />}
        title="Natural Language Hybrid Search"
        description="Type a query or select a filter badge above to execute hybrid vector search."
      />
    </PageShell>
  );
}
