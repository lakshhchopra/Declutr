"use client";

import React from "react";
import { UserCheck } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { EmptyState } from "../../shared/components/feedback/empty-state";

export default function PersonaPage() {
  return (
    <PageShell
      title="Reverse Persona Intelligence"
      subtitle="Private probabilistic user modeling with time-based recency decay."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Persona" }]}
    >
      <EmptyState
        icon={<UserCheck className="h-6 w-6" />}
        title="Persona Profile Initialization"
        description="Your private interaction signals will build personalized search weighting over time."
      />
    </PageShell>
  );
}
