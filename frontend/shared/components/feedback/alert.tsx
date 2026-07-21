import React from "react";
import { cn } from "../../utils/cn";

export interface AlertProps extends React.HTMLAttributes<HTMLDivElement> {
  variant?: "info" | "success" | "warning" | "danger";
  title?: string;
  onClose?: () => void;
}

export function Alert({ children, variant = "info", title, onClose, className = "", ...props }: AlertProps) {
  const variants = {
    info: "bg-[var(--status-info)]/10 border-[var(--status-info)]/30 text-[var(--status-info)]",
    success: "bg-[var(--status-success)]/10 border-[var(--status-success)]/30 text-[var(--status-success)]",
    warning: "bg-[var(--status-warning)]/10 border-[var(--status-warning)]/30 text-[var(--status-warning)]",
    danger: "bg-[var(--status-danger)]/10 border-[var(--status-danger)]/30 text-[var(--status-danger)]",
  };

  return (
    <div
      role="alert"
      className={cn(
        "w-full rounded-xl border p-4 flex items-start justify-between gap-3 text-sm leading-relaxed",
        variants[variant],
        className
      )}
      {...props}
    >
      <div className="flex flex-col gap-1">
        {title && <h4 className="font-semibold text-base">{title}</h4>}
        <div>{children}</div>
      </div>
      {onClose && (
        <button
          onClick={onClose}
          aria-label="Dismiss alert"
          className="text-current opacity-70 hover:opacity-100 transition-opacity p-1"
        >
          ✕
        </button>
      )}
    </div>
  );
}
