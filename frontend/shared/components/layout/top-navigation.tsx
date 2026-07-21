"use client";

import * as React from "react";
import { Sun, Moon, Search, ShieldCheck } from "lucide-react";
import { useTheme } from "../../providers/theme-provider";
import { Button } from "../ui/button";
import { Avatar, AvatarFallback } from "../ui/avatar";

export function TopNavigation() {
  const { theme, toggleTheme } = useTheme();

  return (
    <header className="flex h-16 w-full items-center justify-between border-b border-slate-200 dark:border-slate-800 bg-white/50 dark:bg-slate-900/50 px-6 backdrop-blur-md">
      <div className="flex items-center gap-2">
        <div className="flex h-7 w-7 items-center justify-center rounded-md bg-emerald-500/10 text-emerald-500 border border-emerald-500/20">
          <ShieldCheck className="h-4 w-4" />
        </div>
        <span className="text-xs font-semibold uppercase tracking-wider text-slate-500 dark:text-slate-400">
          Zero-Knowledge Vault
        </span>
      </div>

      <div className="flex items-center gap-3">
        <Button variant="outline" size="sm" leftIcon={<Search className="h-3.5 w-3.5" />}>
          Search (CMD+K)
        </Button>
        <Button
          variant="icon"
          onClick={toggleTheme}
          aria-label="Toggle theme"
        >
          {theme === "dark" ? <Sun className="h-4 w-4 text-amber-400" /> : <Moon className="h-4 w-4 text-slate-700" />}
        </Button>
        <Avatar className="h-8 w-8">
          <AvatarFallback className="bg-emerald-500/20 text-emerald-400 text-xs">US</AvatarFallback>
        </Avatar>
      </div>
    </header>
  );
}

export interface BreadcrumbProps {
  items: { label: string; href?: string }[];
}

export function Breadcrumb({ items }: BreadcrumbProps) {
  return (
    <nav aria-label="Breadcrumb" className="flex items-center gap-1.5 text-xs text-slate-500 dark:text-slate-400 mb-4">
      {items.map((item, idx) => (
        <React.Fragment key={idx}>
          {idx > 0 && <span className="text-slate-600 dark:text-slate-500">/</span>}
          {item.href ? (
            <a href={item.href} className="hover:text-emerald-500 transition-colors">
              {item.label}
            </a>
          ) : (
            <span className="font-semibold text-slate-900 dark:text-slate-100">{item.label}</span>
          )}
        </React.Fragment>
      ))}
    </nav>
  );
}
