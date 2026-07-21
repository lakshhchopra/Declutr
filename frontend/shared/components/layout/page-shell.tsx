"use client";

import React from "react";
import { Breadcrumb, BreadcrumbProps } from "./top-navigation";
import { Container } from "./layout-primitives";

export interface PageShellProps {
  title: string;
  subtitle?: string;
  breadcrumbs?: BreadcrumbProps["items"];
  actions?: React.ReactNode;
  children: React.ReactNode;
}

export function PageShell({
  title,
  subtitle,
  breadcrumbs,
  actions,
  children,
}: PageShellProps) {
  return (
    <Container size="lg" className="py-6">
      {breadcrumbs && <Breadcrumb items={breadcrumbs} />}

      <div className="flex flex-col sm:flex-row sm:items-center justify-between gap-4 mb-8 pb-6 border-b border-slate-200 dark:border-slate-800">
        <div>
          <h1 className="text-2xl sm:text-3xl font-extrabold tracking-tight text-slate-900 dark:text-slate-50">
            {title}
          </h1>
          {subtitle && (
            <p className="text-xs sm:text-sm text-slate-500 dark:text-slate-400 mt-1">
              {subtitle}
            </p>
          )}
        </div>
        {actions && <div className="flex items-center gap-3 shrink-0">{actions}</div>}
      </div>

      <div className="w-full">{children}</div>
    </Container>
  );
}
