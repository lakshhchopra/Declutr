"use client";

import React, { useState } from "react";
import { Shield, Sparkles, FolderKey, Search, Lock, RefreshCw, CheckCircle2, Info } from "lucide-react";
import { ThemeProvider, useTheme } from "../../shared/providers/theme-provider";
import { Button } from "../../shared/components/ui/button";
import { Input, PasswordInput, SearchInput, Textarea } from "../../shared/components/ui/input";
import { Card, CardHeader, CardTitle, CardDescription, CardContent, CardFooter } from "../../shared/components/ui/card";
import { Badge } from "../../shared/components/ui/badge";
import { Avatar, AvatarFallback } from "../../shared/components/ui/avatar";
import { Tabs, TabsList, TabsTrigger, TabsContent } from "../../shared/components/ui/tabs";
import { Alert } from "../../shared/components/feedback/alert";
import { Spinner } from "../../shared/components/feedback/spinner";
import { Skeleton } from "../../shared/components/feedback/skeleton";
import { EmptyState } from "../../shared/components/feedback/empty-state";
import { ErrorState } from "../../shared/components/feedback/error-state";
import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogDescription,
  DialogFooter,
  DialogClose,
} from "../../shared/components/overlay/dialog";
import { Container, Grid, Stack, Section } from "../../shared/components/layout/layout-primitives";
import { TopNavigation, Breadcrumb } from "../../shared/components/layout/top-navigation";

function DesignSystemContent() {
  const { theme, toggleTheme } = useTheme();
  const [loading, setLoading] = useState(false);

  return (
    <div className="min-h-screen bg-slate-950 text-slate-50 font-sans pb-16">
      <TopNavigation />

      <Container size="lg" className="pt-8">
        <Breadcrumb items={[{ label: "Declutr", href: "/" }, { label: "Design System Showcase" }]} />

        <div className="flex items-center justify-between mb-8 pb-6 border-b border-slate-800">
          <div>
            <div className="flex items-center gap-2 mb-2">
              <Badge variant="emerald">Issue #001</Badge>
              <Badge variant="outline">shadcn/ui + Radix</Badge>
            </div>
            <h1 className="text-3xl font-extrabold tracking-tight text-white">
              Declutr Shared Design System
            </h1>
            <p className="text-sm text-slate-400 mt-1">
              Reusable UI primitives, theme tokens, and components built on shadcn/ui and Radix UI.
            </p>
          </div>
          <Button variant="secondary" onClick={toggleTheme}>
            Current Theme: {theme.toUpperCase()} (Toggle)
          </Button>
        </div>

        {/* Component Showcase Tabs */}
        <Tabs defaultValue="buttons" className="w-full">
          <TabsList className="mb-8">
            <TabsTrigger value="buttons">Buttons & Action</TabsTrigger>
            <TabsTrigger value="inputs">Form Inputs</TabsTrigger>
            <TabsTrigger value="display">Cards & Badges</TabsTrigger>
            <TabsTrigger value="feedback">Feedback & Alerts</TabsTrigger>
            <TabsTrigger value="overlay">Modals & Overlays</TabsTrigger>
          </TabsList>

          {/* BUTTONS TAB */}
          <TabsContent value="buttons">
            <Section>
              <h2 className="text-lg font-semibold mb-4 text-emerald-400">Button Variants</h2>
              <div className="flex flex-wrap items-center gap-4">
                <Button variant="default">Primary Action</Button>
                <Button variant="secondary">Secondary Action</Button>
                <Button variant="outline">Outline</Button>
                <Button variant="ghost">Ghost Button</Button>
                <Button variant="danger">Danger Action</Button>
                <Button variant="default" isLoading>
                  Loading
                </Button>
                <Button variant="default" leftIcon={<Sparkles className="h-4 w-4" />}>
                  With Icon
                </Button>
              </div>
            </Section>
            <Section>
              <h2 className="text-lg font-semibold mb-4 text-emerald-400">Button Sizes</h2>
              <div className="flex flex-wrap items-center gap-4">
                <Button size="sm">Small (sm)</Button>
                <Button size="default">Default (md)</Button>
                <Button size="lg">Large (lg)</Button>
              </div>
            </Section>
          </TabsContent>

          {/* INPUTS TAB */}
          <TabsContent value="inputs">
            <Section>
              <h2 className="text-lg font-semibold mb-4 text-emerald-400">Form Inputs</h2>
              <Grid cols={2}>
                <Input label="Email Address" placeholder="user@declutr.vault" />
                <PasswordInput label="Master Encryption Passphrase" placeholder="••••••••••••" />
                <SearchInput label="Search Engine" />
                <Textarea label="Vault Notes" placeholder="Enter secure item description..." />
              </Grid>
            </Section>
          </TabsContent>

          {/* DISPLAY TAB */}
          <TabsContent value="display">
            <Section>
              <h2 className="text-lg font-semibold mb-4 text-emerald-400">Cards & Badges</h2>
              <Grid cols={3}>
                <Card>
                  <CardHeader>
                    <div className="flex justify-between items-center mb-1">
                      <Badge variant="emerald">Encrypted</Badge>
                      <Avatar>
                        <AvatarFallback className="bg-emerald-500/20 text-emerald-400">VK</AvatarFallback>
                      </Avatar>
                    </div>
                    <CardTitle>Personal Vault</CardTitle>
                    <CardDescription>Zero-knowledge AES-256 encrypted space</CardDescription>
                  </CardHeader>
                  <CardContent>
                    <p className="text-xs text-slate-400">Contains 42 stored digital items.</p>
                  </CardContent>
                  <CardFooter>
                    <Button size="sm" variant="outline" className="w-full">
                      Open Vault
                    </Button>
                  </CardFooter>
                </Card>

                <Card>
                  <CardHeader>
                    <Badge variant="blue" className="w-fit mb-1">AI Pipeline</Badge>
                    <CardTitle>Content Intelligence</CardTitle>
                    <CardDescription>OCR & 512-dim Vector Search</CardDescription>
                  </CardHeader>
                  <CardContent>
                    <p className="text-xs text-slate-400">Automatic topic extraction and intent tagging.</p>
                  </CardContent>
                  <CardFooter>
                    <Button size="sm" variant="secondary" className="w-full">
                      View Status
                    </Button>
                  </CardFooter>
                </Card>

                <Card>
                  <CardHeader>
                    <Badge variant="amber" className="w-fit mb-1">Security</Badge>
                    <CardTitle>Behavioral Risk Score</CardTitle>
                    <CardDescription>Passive Anomaly Detection</CardDescription>
                  </CardHeader>
                  <CardContent>
                    <p className="text-xs font-semibold text-emerald-400">0.05 (Trusted Access)</p>
                  </CardContent>
                  <CardFooter>
                    <Button size="sm" variant="ghost" className="w-full">
                      View Logs
                    </Button>
                  </CardFooter>
                </Card>
              </Grid>
            </Section>
            <Section>
              <h2 className="text-lg font-semibold mb-4 text-emerald-400">Badge Variants</h2>
              <div className="flex flex-wrap gap-2">
                <Badge variant="default">Default</Badge>
                <Badge variant="secondary">Secondary</Badge>
                <Badge variant="emerald">Emerald</Badge>
                <Badge variant="blue">Blue</Badge>
                <Badge variant="amber">Amber</Badge>
                <Badge variant="rose">Rose</Badge>
                <Badge variant="outline">Outline</Badge>
              </div>
            </Section>
          </TabsContent>

          {/* FEEDBACK TAB */}
          <TabsContent value="feedback">
            <Section>
              <h2 className="text-lg font-semibold mb-4 text-emerald-400">Alerts & Spinners</h2>
              <Stack gap={4}>
                <Alert variant="info" title="Zero-Knowledge Protection Active">
                  Encryption keys are derived locally on your device.
                </Alert>
                <Alert variant="success" title="Vault Backup Verified">
                  Master Vault Key (MVK) passphrase backed up successfully.
                </Alert>
                <Alert variant="warning" title="Session Renewal Needed">
                  Access token expires in 2 minutes.
                </Alert>
                <Alert variant="danger" title="Access Denied">
                  Invalid challenge proof M1 response.
                </Alert>
              </Stack>
            </Section>
            <Section>
              <h2 className="text-lg font-semibold mb-4 text-emerald-400">Skeletons & States</h2>
              <Grid cols={2}>
                <EmptyState
                  title="No Vault Items Stored"
                  description="Upload receipts, documents, or photos to get started."
                  actionLabel="Upload Item"
                  onAction={() => alert("Upload clicked")}
                />
                <ErrorState
                  title="Failed to Load Search Index"
                  message="Could not reach vector search service."
                  onRetry={() => alert("Retrying")}
                />
              </Grid>
              <div className="mt-6 flex items-center gap-4">
                <span className="text-xs text-slate-400">Skeleton Loaders:</span>
                <Skeleton className="h-6 w-24" />
                <Skeleton className="h-6 w-36" />
                <Skeleton className="h-10 w-10 rounded-full" />
              </div>
            </Section>
          </TabsContent>

          {/* OVERLAY TAB */}
          <TabsContent value="overlay">
            <Section>
              <h2 className="text-lg font-semibold mb-4 text-emerald-400">Dialog & Modal Overlays</h2>
              <Dialog>
                <DialogTrigger asChild>
                  <Button variant="default">Open Confirmation Dialog</Button>
                </DialogTrigger>
                <DialogContent>
                  <DialogHeader>
                    <DialogTitle>Unlock Cryptographic Vault?</DialogTitle>
                    <DialogDescription>
                      This action will decrypt your Master Vault Key (MVK) using your derived passphrase.
                    </DialogDescription>
                  </DialogHeader>
                  <div className="py-4">
                    <PasswordInput label="Passphrase" placeholder="Enter passphrase..." />
                  </div>
                  <DialogFooter>
                    <DialogClose asChild>
                      <Button variant="outline">Cancel</Button>
                    </DialogClose>
                    <Button variant="default">Unlock Vault</Button>
                  </DialogFooter>
                </DialogContent>
              </Dialog>
            </Section>
          </TabsContent>
        </Tabs>
      </Container>
    </div>
  );
}

export default function DesignSystemPage() {
  return (
    <ThemeProvider>
      <DesignSystemContent />
    </ThemeProvider>
  );
}
