"use client";

import React from "react";
import { Sparkles } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { EmptyState } from "../../shared/components/feedback/empty-state";

export default function AIPage() {
  return (
    <PageShell
      title="Content Intelligence Engine"
      subtitle="Asynchronous OCR, Whisper audio transcription, and intent classification pipelines."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "AI Engine" }]}
    >
      <EmptyState
        icon={<Sparkles className="h-6 w-6" />}
        title="AI Pipeline Status"
        description="All background workers are online and listening for Redis task queue payloads."
      />
    </PageShell>
  );
}
