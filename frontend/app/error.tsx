"use client";

import React, { useEffect } from "react";
import { ErrorState } from "../shared/components/feedback/error-state";
import { Container } from "../shared/components/layout/layout-primitives";

export default function Error({
  error,
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  useEffect(() => {
    console.error("Application shell error boundary caught:", error);
  }, [error]);

  return (
    <Container size="md" className="py-16">
      <ErrorState
        title="Application Error Occurred"
        message={error.message || "An unexpected error occurred in the application shell."}
        onRetry={reset}
      />
    </Container>
  );
}
