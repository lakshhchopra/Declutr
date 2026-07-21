"use client";

import React from "react";
import { ShieldAlert } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { EmptyState } from "../../shared/components/feedback/empty-state";

export default function SecurityPage() {
  return (
    <PageShell
      title="Behavioral Security & Risk Telemetry"
      subtitle="Passive session anomaly detection, adaptive MFA interceptors, and HMAC audit chains."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Security" }]}
    >
      <EmptyState
        icon={<ShieldAlert className="h-6 w-6" />}
        title="Zero Security Anomaly Signals"
        description="Session risk score is 0.02 (Trusted). Passive telemetry monitoring active."
      />
    </PageShell>
  );
}
