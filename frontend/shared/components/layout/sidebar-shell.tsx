"use client";

import * as React from "react";
import { Shield, FolderKey, Search, UserCheck, Activity, Menu, X } from "lucide-react";
import { cn } from "../../utils/cn";
import { Button } from "../ui/button";

export interface SidebarNavProps {
  activePath?: string;
  onNavigate?: (path: string) => void;
  children?: React.ReactNode;
}

export function SidebarShell({ activePath = "/vault", onNavigate, children }: SidebarNavProps) {
  const [mobileOpen, setMobileOpen] = React.useState(false);

  const navItems = [
    { label: "Vault Workspaces", icon: FolderKey, path: "/vault" },
    { label: "Semantic Search", icon: Search, path: "/search" },
    { label: "Reverse Persona", icon: UserCheck, path: "/persona" },
    { label: "Behavioral Risk", icon: Activity, path: "/security" },
  ];

  return (
    <div className="flex min-h-screen w-full bg-slate-950 text-slate-50">
      {/* Mobile Sidebar Overlay */}
      {mobileOpen && (
        <div
          className="fixed inset-0 z-[1200] bg-black/60 backdrop-blur-sm lg:hidden"
          onClick={() => setMobileOpen(false)}
        />
      )}

      {/* Sidebar Content */}
      <aside
        className={cn(
          "fixed top-0 bottom-0 left-0 z-[1250] flex w-64 flex-col border-r border-slate-800 bg-slate-900 transition-transform duration-200 lg:static lg:translate-x-0",
          mobileOpen ? "translate-x-0" : "-translate-x-full"
        )}
      >
        <div className="flex h-16 items-center gap-2.5 px-6 border-b border-slate-800">
          <div className="flex h-8 w-8 items-center justify-center rounded-lg bg-emerald-500/15 text-emerald-400 border border-emerald-500/30">
            <Shield className="h-5 w-5" />
          </div>
          <span className="font-bold tracking-tight text-lg text-slate-100">Declutr</span>
        </div>

        <nav className="flex-1 space-y-1 px-3 py-4">
          {navItems.map((item) => {
            const Icon = item.icon;
            const active = activePath === item.path;
            return (
              <button
                key={item.path}
                onClick={() => {
                  onNavigate?.(item.path);
                  setMobileOpen(false);
                }}
                className={cn(
                  "flex w-full items-center gap-3 rounded-lg px-3 py-2 text-xs font-medium transition-colors",
                  active
                    ? "bg-emerald-500/15 text-emerald-400 border border-emerald-500/20 font-semibold"
                    : "text-slate-400 hover:bg-slate-800 hover:text-slate-100"
                )}
              >
                <Icon className="h-4 w-4 shrink-0" />
                {item.label}
              </button>
            );
          })}
        </nav>

        <div className="p-4 border-t border-slate-800 text-[10px] text-slate-500">
          Zero-Trust Vault Engine v1.0
        </div>
      </aside>

      {/* Main Content Area */}
      <div className="flex flex-1 flex-col overflow-hidden">
        {/* Mobile Header Toggle */}
        <header className="flex h-14 items-center justify-between border-b border-slate-800 bg-slate-900/50 px-4 lg:hidden">
          <button
            onClick={() => setMobileOpen(true)}
            className="text-slate-400 hover:text-slate-100 p-1"
          >
            <Menu className="h-6 w-6" />
          </button>
          <span className="font-bold text-sm">Declutr Vault</span>
          <div className="w-6" />
        </header>

        <main className="flex-1 overflow-y-auto p-6">{children}</main>
      </div>
    </div>
  );
}
