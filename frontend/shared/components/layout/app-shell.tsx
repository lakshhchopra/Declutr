"use client";

import React, { useState } from "react";
import Link from "next/link";
import { usePathname } from "next/navigation";
import {
  Shield,
  LayoutDashboard,
  FolderKey,
  Search,
  FolderKanban,
  Sparkles,
  UserCheck,
  ShieldAlert,
  Settings,
  Bell,
  Sun,
  Moon,
  Menu,
  X,
  Layers,
} from "lucide-react";
import { useTheme } from "../../providers/theme-provider";
import { useToast } from "../../providers/toast-provider";
import { cn } from "../../utils/cn";
import { Button } from "../ui/button";
import { Avatar, AvatarFallback } from "../ui/avatar";

export function AppShell({ children }: { children: React.ReactNode }) {
  const pathname = usePathname();
  const { theme, toggleTheme } = useTheme();
  const { toast } = useToast();
  const [mobileOpen, setMobileOpen] = useState(false);

  const navItems = [
    { label: "Dashboard", icon: LayoutDashboard, path: "/dashboard" },
    { label: "Vault", icon: FolderKey, path: "/vault" },
    { label: "Search", icon: Search, path: "/search" },
    { label: "Collections", icon: FolderKanban, path: "/collections" },
    { label: "AI Engine", icon: Sparkles, path: "/ai" },
    { label: "Persona", icon: UserCheck, path: "/persona" },
    { label: "Security", icon: ShieldAlert, path: "/security" },
    { label: "Settings", icon: Settings, path: "/settings" },
    { label: "Design System", icon: Layers, path: "/design-system" },
  ];

  const handleNotificationClick = () => {
    toast({
      type: "info",
      title: "Notifications Placeholder",
      message: "No new security or vault alerts.",
    });
  };

  return (
    <div className="flex min-h-screen w-full bg-slate-950 text-slate-50 font-sans">
      {/* Mobile Drawer Backdrop */}
      {mobileOpen && (
        <div
          className="fixed inset-0 z-[1200] bg-black/60 backdrop-blur-sm lg:hidden"
          onClick={() => setMobileOpen(false)}
        />
      )}

      {/* Desktop / Tablet Sidebar */}
      <aside
        className={cn(
          "fixed top-0 bottom-0 left-0 z-[1250] flex w-64 flex-col border-r border-slate-800 bg-slate-900/95 backdrop-blur-md transition-transform duration-200 lg:static lg:translate-x-0",
          mobileOpen ? "translate-x-0" : "-translate-x-full"
        )}
      >
        {/* Brand Header */}
        <div className="flex h-16 items-center justify-between px-6 border-b border-slate-800">
          <Link href="/dashboard" className="flex items-center gap-2.5">
            <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-emerald-500/15 text-emerald-400 border border-emerald-500/30">
              <Shield className="h-5 w-5" />
            </div>
            <span className="font-extrabold tracking-tight text-lg text-slate-100">Declutr</span>
          </Link>
          <button className="lg:hidden text-slate-400 hover:text-slate-100" onClick={() => setMobileOpen(false)}>
            <X className="h-5 w-5" />
          </button>
        </div>

        {/* Sidebar Navigation */}
        <nav className="flex-1 space-y-1 px-3 py-4 overflow-y-auto">
          {navItems.map((item) => {
            const Icon = item.icon;
            const active = pathname === item.path || (item.path !== "/" && pathname?.startsWith(item.path));
            return (
              <Link
                key={item.path}
                href={item.path}
                onClick={() => setMobileOpen(false)}
                className={cn(
                  "flex items-center gap-3 rounded-lg px-3 py-2.5 text-xs font-medium transition-all duration-150",
                  active
                    ? "bg-emerald-500/15 text-emerald-400 border border-emerald-500/25 font-semibold shadow-sm"
                    : "text-slate-400 hover:bg-slate-800/80 hover:text-slate-100"
                )}
              >
                <Icon className={cn("h-4 w-4 shrink-0", active ? "text-emerald-400" : "text-slate-400")} />
                {item.label}
              </Link>
            );
          })}
        </nav>

        {/* Sidebar Footer */}
        <div className="p-4 border-t border-slate-800 text-[11px] text-slate-500 flex items-center justify-between">
          <span>Zero-Knowledge Architecture</span>
          <span className="text-emerald-400/80 font-mono">v1.0</span>
        </div>
      </aside>

      {/* Main Content Area */}
      <div className="flex flex-1 flex-col overflow-hidden">
        {/* Top Header */}
        <header className="flex h-16 items-center justify-between border-b border-slate-800 bg-slate-900/60 px-4 sm:px-6 backdrop-blur-md">
          <div className="flex items-center gap-3">
            <button
              onClick={() => setMobileOpen(true)}
              className="p-1.5 rounded-lg text-slate-400 hover:bg-slate-800 hover:text-slate-100 lg:hidden"
            >
              <Menu className="h-5 w-5" />
            </button>
            <div className="hidden sm:flex items-center gap-2">
              <span className="text-xs font-semibold text-slate-400 uppercase tracking-wider">Workspace</span>
              <span className="text-xs font-mono bg-slate-800 text-emerald-400 px-2 py-0.5 rounded border border-slate-700">
                Personal Vault
              </span>
            </div>
          </div>

          <div className="flex items-center gap-2 sm:gap-3">
            <Button variant="outline" size="sm" className="hidden sm:inline-flex" leftIcon={<Search className="h-3.5 w-3.5" />}>
              Search (CMD+K)
            </Button>
            <Button variant="icon" onClick={handleNotificationClick} aria-label="Notifications">
              <Bell className="h-4 w-4" />
            </Button>
            <Button variant="icon" onClick={toggleTheme} aria-label="Toggle theme">
              {theme === "dark" ? <Sun className="h-4 w-4 text-amber-400" /> : <Moon className="h-4 w-4 text-slate-300" />}
            </Button>
            <Avatar className="h-8 w-8 cursor-pointer">
              <AvatarFallback className="bg-emerald-500/20 text-emerald-400 text-xs font-bold">DC</AvatarFallback>
            </Avatar>
          </div>
        </header>

        {/* Main Content Container */}
        <main className="flex-1 overflow-y-auto pb-16 lg:pb-0">{children}</main>

        {/* Mobile Bottom Navigation */}
        <nav className="fixed bottom-0 left-0 right-0 z-[1100] flex h-16 items-center justify-around border-t border-slate-800 bg-slate-900/95 backdrop-blur-md lg:hidden">
          {[
            { label: "Dashboard", icon: LayoutDashboard, path: "/dashboard" },
            { label: "Vault", icon: FolderKey, path: "/vault" },
            { label: "Search", icon: Search, path: "/search" },
            { label: "Persona", icon: UserCheck, path: "/persona" },
            { label: "Settings", icon: Settings, path: "/settings" },
          ].map((item) => {
            const Icon = item.icon;
            const active = pathname === item.path || (item.path !== "/" && pathname?.startsWith(item.path));
            return (
              <Link
                key={item.path}
                href={item.path}
                className={cn(
                  "flex flex-col items-center gap-1 text-[10px] font-medium transition-colors p-1",
                  active ? "text-emerald-400 font-bold" : "text-slate-400 hover:text-slate-200"
                )}
              >
                <Icon className="h-5 w-5" />
                {item.label}
              </Link>
            );
          })}
        </nav>
      </div>
    </div>
  );
}
