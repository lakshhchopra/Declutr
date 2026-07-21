import React from "react";
import Link from "next/link";
import { FileQuestion } from "lucide-react";
import { Button } from "../shared/components/ui/button";
import { Container } from "../shared/components/layout/layout-primitives";

export default function NotFound() {
  return (
    <Container size="md" className="py-24 text-center">
      <div className="flex h-16 w-16 items-center justify-center rounded-full bg-slate-900 border border-slate-800 text-slate-400 mx-auto mb-6">
        <FileQuestion className="h-8 w-8 text-emerald-400" />
      </div>
      <h1 className="text-3xl font-extrabold text-white mb-2">404 - Page Not Found</h1>
      <p className="text-sm text-slate-400 max-w-md mx-auto mb-8">
        The requested vault location or page route does not exist or has been relocated.
      </p>
      <Link href="/dashboard">
        <Button variant="default">Return to Dashboard</Button>
      </Link>
    </Container>
  );
}
