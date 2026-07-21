"use client";

import React, { useState } from "react";
import Link from "next/link";
import {
  Shield,
  User,
  Sun,
  Moon,
  Sparkles,
  Lock,
  Bell,
  CheckCircle2,
  ArrowRight,
  ArrowLeft,
  Check,
} from "lucide-react";
import { Button } from "../../shared/components/ui/button";
import { Input } from "../../shared/components/ui/input";
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "../../shared/components/ui/card";
import { Container } from "../../shared/components/layout/layout-primitives";
import { useTheme } from "../../shared/providers/theme-provider";
import { ProfileService } from "../../features/user/services/profile-service";

export default function OnboardingPage() {
  const [step, setStep] = useState(1);
  const { setTheme, theme } = useTheme();

  // Form states
  const [displayName, setDisplayName] = useState("Declutr User");
  const [selectedAvatar, setSelectedAvatar] = useState("emerald");
  const [aiBehavior, setAiBehavior] = useState<"minimal" | "balanced" | "proactive">("balanced");
  const [privacyMode, setPrivacyMode] = useState<"local_first" | "encrypted_cloud" | "hybrid">("local_first");
  const [notifications, setNotifications] = useState({
    email: true,
    push: true,
    desktop: true,
    security: true,
  });

  const totalSteps = 8;

  const handleNext = () => {
    if (step < totalSteps) {
      setStep((prev) => prev + 1);
    }
  };

  const handleBack = () => {
    if (step > 1) {
      setStep((prev) => prev - 1);
    }
  };

  const handleComplete = async () => {
    await ProfileService.updateProfile({ displayName });
    await ProfileService.updatePreferences({
      theme: theme as any,
      aiBehavior,
      privacyMode,
      emailNotifications: notifications.email,
      pushNotifications: notifications.push,
      desktopNotifications: notifications.desktop,
      securityAlerts: notifications.security,
    });
    await ProfileService.completeOnboarding();
    window.location.href = "/dashboard";
  };

  return (
    <div className="min-h-screen bg-slate-950 text-slate-50 font-sans flex flex-col justify-between p-4 sm:p-6">
      {/* Top Header */}
      <header className="flex items-center justify-between max-w-2xl mx-auto w-full pt-4">
        <div className="flex items-center gap-2">
          <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-emerald-500/15 text-emerald-400 border border-emerald-500/30">
            <Shield className="h-5 w-5" />
          </div>
          <span className="font-extrabold tracking-tight text-lg text-white">Declutr Setup</span>
        </div>
        <span className="text-xs font-mono text-slate-400">
          Step {step} of {totalSteps}
        </span>
      </header>

      {/* Progress Bar */}
      <div className="max-w-2xl mx-auto w-full my-4">
        <div className="h-1.5 w-full bg-slate-900 rounded-full overflow-hidden">
          <div
            className="h-full bg-emerald-500 transition-all duration-300 ease-out"
            style={{ width: `${(step / totalSteps) * 100}%` }}
          />
        </div>
      </div>

      {/* Main Step Container */}
      <Container size="md" className="my-auto py-8">
        <Card className="glass-panel border-slate-800 shadow-2xl p-6 sm:p-8">
          {/* Step 1: Welcome */}
          {step === 1 && (
            <div className="space-y-6 text-center">
              <div className="h-16 w-16 rounded-full bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 flex items-center justify-center mx-auto">
                <Sparkles className="h-8 w-8" />
              </div>
              <div className="space-y-2">
                <h2 className="text-2xl font-extrabold text-white">Welcome to Declutr</h2>
                <p className="text-sm text-slate-400 max-w-md mx-auto">
                  Let's personalize your private digital life vault experience in a few quick steps.
                </p>
              </div>
              <Button size="lg" variant="default" className="w-full sm:w-auto px-8" onClick={handleNext}>
                Get Started <ArrowRight className="h-4 w-4 ml-2" />
              </Button>
            </div>
          )}

          {/* Step 2: Choose Display Name */}
          {step === 2 && (
            <div className="space-y-6">
              <div>
                <h2 className="text-xl font-extrabold text-white mb-1">Choose Display Name</h2>
                <p className="text-xs text-slate-400">How should Declutr address you across your workspace?</p>
              </div>
              <Input
                label="Display Name"
                value={displayName}
                onChange={(e) => setDisplayName(e.target.value)}
                placeholder="Jane Doe"
              />
            </div>
          )}

          {/* Step 3: Choose Avatar */}
          {step === 3 && (
            <div className="space-y-6">
              <div>
                <h2 className="text-xl font-extrabold text-white mb-1">Select Vault Accent Color</h2>
                <p className="text-xs text-slate-400">Choose your avatar accent badge color.</p>
              </div>
              <div className="grid grid-cols-4 gap-4 py-4">
                {[
                  { id: "emerald", bg: "bg-emerald-500/20 text-emerald-400 border-emerald-500" },
                  { id: "blue", bg: "bg-blue-500/20 text-blue-400 border-blue-500" },
                  { id: "amber", bg: "bg-amber-500/20 text-amber-400 border-amber-500" },
                  { id: "indigo", bg: "bg-indigo-500/20 text-indigo-400 border-indigo-500" },
                ].map((av) => (
                  <button
                    key={av.id}
                    onClick={() => setSelectedAvatar(av.id)}
                    className={`h-16 rounded-xl border flex items-center justify-center font-bold text-lg transition-all ${
                      av.bg
                    } ${selectedAvatar === av.id ? "ring-2 ring-emerald-400 scale-105" : "opacity-70"}`}
                  >
                    DC
                  </button>
                ))}
              </div>
            </div>
          )}

          {/* Step 4: Choose Theme */}
          {step === 4 && (
            <div className="space-y-6">
              <div>
                <h2 className="text-xl font-extrabold text-white mb-1">Choose Interface Theme</h2>
                <p className="text-xs text-slate-400">Select your preferred color scheme.</p>
              </div>
              <div className="grid grid-cols-3 gap-3">
                {[
                  { id: "dark", label: "Dark Mode", icon: Moon },
                  { id: "light", label: "Light Mode", icon: Sun },
                  { id: "system", label: "System", icon: Shield },
                ].map((t) => {
                  const Icon = t.icon;
                  const active = theme === t.id;
                  return (
                    <button
                      key={t.id}
                      onClick={() => setTheme(t.id as any)}
                      className={`p-4 rounded-xl border text-center flex flex-col items-center gap-2 transition-all ${
                        active
                          ? "border-emerald-500 bg-emerald-500/10 text-emerald-400 font-semibold"
                          : "border-slate-800 bg-slate-900/40 text-slate-400 hover:border-slate-700"
                      }`}
                    >
                      <Icon className="h-6 w-6" />
                      <span className="text-xs">{t.label}</span>
                    </button>
                  );
                })}
              </div>
            </div>
          )}

          {/* Step 5: Choose AI Behavior */}
          {step === 5 && (
            <div className="space-y-6">
              <div>
                <h2 className="text-xl font-extrabold text-white mb-1">Choose AI Behavior Mode</h2>
                <p className="text-xs text-slate-400">Configure how actively AI assists with document organization.</p>
              </div>
              <div className="space-y-3">
                {[
                  { id: "minimal", title: "Minimal AI", desc: "Only run OCR & vector search when explicitly requested." },
                  { id: "balanced", title: "Balanced AI (Recommended)", desc: "Auto-tag receipts and extract dates/merchants." },
                  { id: "proactive", title: "Highly Proactive AI", desc: "Build intent predictions & suggest temporal collections." },
                ].map((item) => (
                  <button
                    key={item.id}
                    onClick={() => setAiBehavior(item.id as any)}
                    className={`w-full text-left p-4 rounded-xl border transition-all ${
                      aiBehavior === item.id
                        ? "border-emerald-500 bg-emerald-500/10 text-white"
                        : "border-slate-800 bg-slate-900/40 text-slate-400 hover:border-slate-700"
                    }`}
                  >
                    <h4 className="text-xs font-bold text-emerald-400">{item.title}</h4>
                    <p className="text-[11px] text-slate-400 mt-1">{item.desc}</p>
                  </button>
                ))}
              </div>
            </div>
          )}

          {/* Step 6: Privacy Preferences */}
          {step === 6 && (
            <div className="space-y-6">
              <div>
                <h2 className="text-xl font-extrabold text-white mb-1">Privacy Architecture Mode</h2>
                <p className="text-xs text-slate-400">Control client-side zero-knowledge encryption preferences.</p>
              </div>
              <div className="space-y-3">
                {[
                  { id: "local_first", title: "Maximum Privacy (Local-First)", desc: "Zero-Knowledge AES-256 encryption. Server never sees plaintext." },
                  { id: "encrypted_cloud", title: "Encrypted Cloud Storage", desc: "Client-encrypted backups synced to S3/Cloudflare R2." },
                  { id: "hybrid", title: "Hybrid Indexing", desc: "Local keys with server vector search acceleration." },
                ].map((p) => (
                  <button
                    key={p.id}
                    onClick={() => setPrivacyMode(p.id as any)}
                    className={`w-full text-left p-4 rounded-xl border transition-all ${
                      privacyMode === p.id
                        ? "border-emerald-500 bg-emerald-500/10 text-white"
                        : "border-slate-800 bg-slate-900/40 text-slate-400 hover:border-slate-700"
                    }`}
                  >
                    <h4 className="text-xs font-bold text-emerald-400">{p.title}</h4>
                    <p className="text-[11px] text-slate-400 mt-1">{p.desc}</p>
                  </button>
                ))}
              </div>
            </div>
          )}

          {/* Step 7: Notification Preferences */}
          {step === 7 && (
            <div className="space-y-6">
              <div>
                <h2 className="text-xl font-extrabold text-white mb-1">Notification Preferences</h2>
                <p className="text-xs text-slate-400">Select which security and vault updates you want to receive.</p>
              </div>
              <div className="space-y-3">
                {[
                  { key: "email", label: "Email Notifications", desc: "Receive weekly vault telemetry summaries." },
                  { key: "push", label: "Push Notifications", desc: "Mobile alerts for background ingestion completion." },
                  { key: "desktop", label: "Desktop Alerts", desc: "Browser notifications for active uploads." },
                  { key: "security", label: "Critical Security Alerts", desc: "Instant notifications for session anomalies." },
                ].map((n) => {
                  const checked = (notifications as any)[n.key];
                  return (
                    <label
                      key={n.key}
                      className="flex items-center justify-between p-3.5 rounded-xl border border-slate-800 bg-slate-900/40 cursor-pointer hover:border-slate-700"
                    >
                      <div>
                        <h4 className="text-xs font-semibold text-slate-200">{n.label}</h4>
                        <p className="text-[11px] text-slate-400">{n.desc}</p>
                      </div>
                      <input
                        type="checkbox"
                        checked={checked}
                        onChange={(e) =>
                          setNotifications((prev) => ({ ...prev, [n.key]: e.target.checked }))
                        }
                        className="h-4 w-4 rounded border-slate-700 bg-slate-900 text-emerald-500 focus:ring-emerald-500"
                      />
                    </label>
                  );
                })}
              </div>
            </div>
          )}

          {/* Step 8: Completion */}
          {step === 8 && (
            <div className="space-y-6 text-center">
              <div className="h-16 w-16 rounded-full bg-emerald-500/10 text-emerald-400 border border-emerald-500/20 flex items-center justify-center mx-auto">
                <CheckCircle2 className="h-8 w-8" />
              </div>
              <div className="space-y-2">
                <h2 className="text-2xl font-extrabold text-white">Setup Complete!</h2>
                <p className="text-sm text-slate-400 max-w-md mx-auto">
                  Your profile and privacy preferences have been saved. Your encrypted vault workspace is ready.
                </p>
              </div>
              <Button size="lg" variant="default" className="w-full sm:w-auto px-8" onClick={handleComplete}>
                Launch Workspace Dashboard
              </Button>
            </div>
          )}

          {/* Step Navigation Controls */}
          {step > 1 && step < totalSteps && (
            <div className="flex items-center justify-between pt-6 mt-6 border-t border-slate-800">
              <Button variant="outline" size="sm" onClick={handleBack}>
                <ArrowLeft className="h-4 w-4 mr-1" /> Back
              </Button>
              <Button variant="default" size="sm" onClick={handleNext}>
                Continue <ArrowRight className="h-4 w-4 ml-1" />
              </Button>
            </div>
          )}
        </Card>
      </Container>

      {/* Footer */}
      <footer className="text-center text-xs text-slate-500 py-4 max-w-2xl mx-auto w-full">
        Declutr Privacy-First Onboarding • Preferences stored locally and encrypted.
      </footer>
    </div>
  );
}
