import React from "react";
import { Skeleton } from "../shared/components/feedback/skeleton";
import { Container, Grid, Section } from "../shared/components/layout/layout-primitives";

export default function Loading() {
  return (
    <Container size="lg" className="py-8">
      <div className="flex flex-col gap-2 mb-8">
        <Skeleton className="h-8 w-64" />
        <Skeleton className="h-4 w-96" />
      </div>

      <Grid cols={3} className="mb-8">
        <Skeleton className="h-36 w-full rounded-xl" />
        <Skeleton className="h-36 w-full rounded-xl" />
        <Skeleton className="h-36 w-full rounded-xl" />
      </Grid>

      <Section>
        <div className="space-y-4">
          <Skeleton className="h-12 w-full rounded-lg" />
          <Skeleton className="h-12 w-full rounded-lg" />
          <Skeleton className="h-12 w-full rounded-lg" />
        </div>
      </Section>
    </Container>
  );
}
