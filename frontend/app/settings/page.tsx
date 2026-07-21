"use client";

import React from "react";
import { Settings } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "../../shared/components/ui/card";
import { Button } from "../../shared/components/ui/button";
import { useTheme } from "../../shared/providers/theme-provider";

export default function SettingsPage() {
  const { theme, setTheme } = useTheme();

  return (
    <PageShell
      title="Application Settings"
      subtitle="Theme mode, vault privacy settings, and cryptographic key management."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Settings" }]}
    >
      <div className="max-w-2xl space-y-6">
        <Card>
          <CardHeader>
            <CardTitle>Appearance & Theme Mode</CardTitle>
            <CardDescription>Choose how Declutr looks on your device.</CardDescription>
          </CardHeader>
          <CardContent className="flex gap-3">
            <Button variant={theme === "dark" ? "default" : "outline"} onClick={() => setTheme("dark")}>
              Dark Mode (Default)
            </Button>
            <Button variant={theme === "light" ? "default" : "outline"} onClick={() => setTheme("light")}>
              Light Mode
            </Button>
            <Button variant={theme === "system" ? "default" : "outline"} onClick={() => setTheme("system")}>
              System
            </Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Vault Privacy Mode</CardTitle>
            <CardDescription>Select client-side Private Mode vs server Enhanced AI Mode.</CardDescription>
          </CardHeader>
          <CardContent className="flex gap-3">
            <Button variant="default">Private Mode (Zero-Knowledge)</Button>
            <Button variant="outline">Enhanced AI Mode (Opt-In)</Button>
          </CardContent>
        </Card>
      </div>
    </PageShell>
  );
}
