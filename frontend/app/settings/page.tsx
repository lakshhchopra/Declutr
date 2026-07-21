"use client";

import React, { useState, useEffect } from "react";
import { Settings, User, Moon, Sparkles, Lock, Bell, Shield, Save } from "lucide-react";
import { PageShell } from "../../shared/components/layout/page-shell";
import { Card, CardHeader, CardTitle, CardDescription, CardContent } from "../../shared/components/ui/card";
import { Button } from "../../shared/components/ui/button";
import { Input, Textarea } from "../../shared/components/ui/input";
import { Tabs } from "../../shared/components/ui/tabs";
import { useTheme } from "../../shared/providers/theme-provider";
import { useToast } from "../../shared/providers/toast-provider";
import { ProfileService, UserProfile, UserPreferences } from "../../features/user/services/profile-service";

export default function SettingsPage() {
  const { theme, setTheme } = useTheme();
  const { toast } = useToast();
  const [activeTab, setActiveTab] = useState("general");

  const [profile, setProfile] = useState<UserProfile>({
    displayName: "Declutr User",
    timezone: "UTC",
    language: "en",
    bio: "",
    onboardingCompleted: true,
  });

  const [prefs, setPrefs] = useState<UserPreferences>({
    theme: "dark",
    aiBehavior: "balanced",
    privacyMode: "local_first",
    emailNotifications: true,
    pushNotifications: true,
    desktopNotifications: true,
    securityAlerts: true,
  });

  useEffect(() => {
    ProfileService.getProfile().then(setProfile);
    ProfileService.getPreferences().then(setPrefs);
  }, []);

  const handleSaveProfile = async () => {
    await ProfileService.updateProfile(profile);
    toast({
      type: "success",
      title: "Profile Updated",
      message: "Your display name and user profile details have been saved.",
    });
  };

  const handleSavePrefs = async () => {
    await ProfileService.updatePreferences(prefs);
    toast({
      type: "success",
      title: "Preferences Saved",
      message: "AI behavior, privacy mode, and notification preferences updated.",
    });
  };

  return (
    <PageShell
      title="Application Settings"
      subtitle="Manage your profile, theme mode, AI behavior, privacy preferences, and notifications."
      breadcrumbs={[{ label: "Declutr", href: "/" }, { label: "Settings" }]}
      actions={
        <Button variant="default" onClick={handleSaveProfile} leftIcon={<Save className="h-4 w-4" />}>
          Save Settings
        </Button>
      }
    >
      <Tabs
        activeTab={activeTab}
        onChange={setActiveTab}
        tabs={[
          { id: "general", label: "General & Profile", icon: User },
          { id: "appearance", label: "Appearance", icon: Moon },
          { id: "ai", label: "AI Behavior", icon: Sparkles },
          { id: "privacy", label: "Privacy Mode", icon: Lock },
          { id: "notifications", label: "Notifications", icon: Bell },
        ]}
      />

      <div className="mt-6 max-w-3xl space-y-6">
        {/* General Tab */}
        {activeTab === "general" && (
          <Card>
            <CardHeader>
              <CardTitle>User Profile Settings</CardTitle>
              <CardDescription>Update your public display name, timezone, and biography.</CardDescription>
            </CardHeader>
            <CardContent className="space-y-4">
              <Input
                label="Display Name"
                value={profile.displayName}
                onChange={(e) => setProfile({ ...profile, displayName: e.target.value })}
              />
              <div className="grid grid-cols-2 gap-4">
                <Input
                  label="Timezone"
                  value={profile.timezone}
                  onChange={(e) => setProfile({ ...profile, timezone: e.target.value })}
                />
                <Input
                  label="Language"
                  value={profile.language}
                  onChange={(e) => setProfile({ ...profile, language: e.target.value })}
                />
              </div>
              <Textarea
                label="Biography (Optional)"
                value={profile.bio || ""}
                onChange={(e) => setProfile({ ...profile, bio: e.target.value })}
                placeholder="Short bio about your digital vault workflow..."
              />
              <Button variant="default" onClick={handleSaveProfile}>
                Save Profile
              </Button>
            </CardContent>
          </Card>
        )}

        {/* Appearance Tab */}
        {activeTab === "appearance" && (
          <Card>
            <CardHeader>
              <CardTitle>Appearance & Theme Mode</CardTitle>
              <CardDescription>Choose how Declutr looks across your devices.</CardDescription>
            </CardHeader>
            <CardContent className="flex gap-3">
              <Button variant={theme === "dark" ? "default" : "outline"} onClick={() => setTheme("dark")}>
                Dark Mode (Default)
              </Button>
              <Button variant={theme === "light" ? "default" : "outline"} onClick={() => setTheme("light")}>
                Light Mode
              </Button>
              <Button variant={theme === "system" ? "default" : "outline"} onClick={() => setTheme("system")}>
                System Preference
              </Button>
            </CardContent>
          </Card>
        )}

        {/* AI Behavior Tab */}
        {activeTab === "ai" && (
          <Card>
            <CardHeader>
              <CardTitle>AI Content Intelligence Mode</CardTitle>
              <CardDescription>Configure how actively background AI models assist with indexing.</CardDescription>
            </CardHeader>
            <CardContent className="space-y-3">
              {[
                { id: "minimal", title: "Minimal AI Mode", desc: "Run OCR & vector embeddings only on manual trigger." },
                { id: "balanced", title: "Balanced AI Mode (Default)", desc: "Auto-extract metadata, dates, and transaction entities." },
                { id: "proactive", title: "Highly Proactive AI", desc: "Generate temporal context collections & relationship models." },
              ].map((item) => (
                <button
                  key={item.id}
                  onClick={() => {
                    setPrefs({ ...prefs, aiBehavior: item.id as any });
                    handleSavePrefs();
                  }}
                  className={`w-full text-left p-4 rounded-xl border transition-all ${
                    prefs.aiBehavior === item.id
                      ? "border-emerald-500 bg-emerald-500/10 text-white font-semibold"
                      : "border-slate-800 bg-slate-900/40 text-slate-400 hover:border-slate-700"
                  }`}
                >
                  <h4 className="text-xs font-bold text-emerald-400">{item.title}</h4>
                  <p className="text-[11px] text-slate-400 mt-1">{item.desc}</p>
                </button>
              ))}
            </CardContent>
          </Card>
        )}

        {/* Privacy Tab */}
        {activeTab === "privacy" && (
          <Card>
            <CardHeader>
              <CardTitle>Zero-Knowledge Privacy Mode</CardTitle>
              <CardDescription>Manage client-side encryption and cloud synchronization rules.</CardDescription>
            </CardHeader>
            <CardContent className="space-y-3">
              {[
                { id: "local_first", title: "Maximum Privacy (Local First)", desc: "Master Vault Key derives locally. No unencrypted data stored." },
                { id: "encrypted_cloud", title: "Encrypted Cloud Sync", desc: "Client-encrypted AES-256 chunks backed up to S3." },
                { id: "hybrid", title: "Hybrid Indexing", desc: "Local master key with encrypted vector index support." },
              ].map((item) => (
                <button
                  key={item.id}
                  onClick={() => {
                    setPrefs({ ...prefs, privacyMode: item.id as any });
                    handleSavePrefs();
                  }}
                  className={`w-full text-left p-4 rounded-xl border transition-all ${
                    prefs.privacyMode === item.id
                      ? "border-emerald-500 bg-emerald-500/10 text-white font-semibold"
                      : "border-slate-800 bg-slate-900/40 text-slate-400 hover:border-slate-700"
                  }`}
                >
                  <h4 className="text-xs font-bold text-emerald-400">{item.title}</h4>
                  <p className="text-[11px] text-slate-400 mt-1">{item.desc}</p>
                </button>
              ))}
            </CardContent>
          </Card>
        )}

        {/* Notifications Tab */}
        {activeTab === "notifications" && (
          <Card>
            <CardHeader>
              <CardTitle>Notification Rules</CardTitle>
              <CardDescription>Select which security alerts and vault updates you receive.</CardDescription>
            </CardHeader>
            <CardContent className="space-y-3">
              {[
                { key: "emailNotifications", label: "Email Summaries" },
                { key: "pushNotifications", label: "Mobile Push Alerts" },
                { key: "desktopNotifications", label: "Desktop Notifications" },
                { key: "securityAlerts", label: "Critical Security Anomalies" },
              ].map((n) => (
                <label
                  key={n.key}
                  className="flex items-center justify-between p-3.5 rounded-xl border border-slate-800 bg-slate-900/40 cursor-pointer hover:border-slate-700"
                >
                  <span className="text-xs font-medium text-slate-200">{n.label}</span>
                  <input
                    type="checkbox"
                    checked={(prefs as any)[n.key]}
                    onChange={(e) => {
                      const updated = { ...prefs, [n.key]: e.target.checked };
                      setPrefs(updated);
                      ProfileService.updatePreferences(updated);
                    }}
                    className="h-4 w-4 rounded border-slate-700 bg-slate-900 text-emerald-500 focus:ring-emerald-500"
                  />
                </label>
              ))}
            </CardContent>
          </Card>
        )}
      </div>
    </PageShell>
  );
}
