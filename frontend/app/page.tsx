"use client";

import React from "react";
import Link from "next/link";
import { Shield, Sparkles, FolderKey, Search, Lock, Layers } from "lucide-react";
import { ThemeProvider, useTheme } from "../shared/providers/theme-provider";
import { Button } from "../shared/components/ui/button";
import { Badge } from "../shared/components/ui/badge";
import { Container, Grid, Section } from "../shared/components/layout/layout-primitives";
import { TopNavigation } from "../shared/components/layout/top-navigation";

function LandingContent() {
  const { toggleTheme } = useTheme();

  return (
    <div className="min-h-screen bg-slate-950 text-slate-50 font-sans">
      <TopNavigation />

      <main>
        <Section className="pt-20 pb-16 text-center border-0">
          <Container size="md">
            <div className="inline-flex items-center gap-2 mb-6">
              <Badge variant="emerald">Issue #001 Completed</Badge>
              <Badge variant="outline">Shared Design System Foundation</Badge>
            </div>

            <h1 className="text-4xl md:text-6xl font-extrabold tracking-tight text-white mb-6 leading-tight">
              AI-Powered Intelligent <br />
              <span className="text-transparent bg-clip-text bg-gradient-to-r from-emerald-400 to-teal-300">
                Digital Life Vault
              </span>
            </h1>

            <p className="text-base md:text-lg text-slate-400 max-w-2xl mx-auto mb-8 leading-relaxed">
              Declutr helps you securely store, organize, contextually connect, and retrieve your digital memory assets using natural human memory associations.
            </p>

            <div className="flex flex-wrap items-center justify-center gap-4">
              <Link href="/design-system">
                <Button size="lg" variant="default" leftIcon={<Layers className="h-5 w-5" />}>
                  Explore Design System Showcase
                </Button>
              </Link>
              <Button size="lg" variant="outline" onClick={toggleTheme}>
                Toggle Theme Mode
              </Button>
            </div>
          </Container>
        </Section>

        <Section className="py-12 border-t border-slate-900">
          <Container size="lg">
            <Grid cols={3}>
              <div className="p-6 rounded-xl border border-slate-800 bg-slate-900/40">
                <div className="h-10 w-10 rounded-lg bg-emerald-500/10 text-emerald-400 flex items-center justify-center mb-4">
                  <Lock className="h-5 w-5" />
                </div>
                <h3 className="font-semibold text-white text-base mb-2">Zero-Knowledge Security</h3>
                <p className="text-xs text-slate-400 leading-relaxed">
                  SRP-6a authentication and client-side AES-256-GCM encryption guarantee server databases never hold plaintext passwords or unencrypted master keys.
                </p>
              </div>

              <div className="p-6 rounded-xl border border-slate-800 bg-slate-900/40">
                <div className="h-10 w-10 rounded-lg bg-blue-500/10 text-blue-400 flex items-center justify-center mb-4">
                  <Sparkles className="h-5 w-5" />
                </div>
                <h3 className="font-semibold text-white text-base mb-2">Intent-Aware Intelligence</h3>
                <p className="text-xs text-slate-400 leading-relaxed">
                  Extracts OCR text, 512-dim vector embeddings, and intent tags to categorize booking references, invoices, and travel plans.
                </p>
              </div>

              <div className="p-6 rounded-xl border border-slate-800 bg-slate-900/40">
                <div className="h-10 w-10 rounded-lg bg-indigo-500/10 text-indigo-400 flex items-center justify-center mb-4">
                  <Layers className="h-5 w-5" />
                </div>
                <h3 className="font-semibold text-white text-base mb-2">shadcn/ui Design System</h3>
                <p className="text-xs text-slate-400 leading-relaxed">
                  Standardized UI component primitives, theme tokens, and accessibility standards built with Radix UI and Tailwind CSS.
                </p>
              </div>
            </Grid>
          </Container>
        </Section>
      </main>
    </div>
  );
}

export default function Home() {
  return (
    <ThemeProvider>
      <LandingContent />
    </ThemeProvider>
  );
}
