import React from "react";
import { cn } from "../../utils/cn";

export interface AvatarProps extends React.HTMLAttributes<HTMLDivElement> {
  src?: string;
  alt?: string;
  name?: string;
  size?: "sm" | "md" | "lg";
}

export function Avatar({ src, alt, name, size = "md", className = "", ...props }: AvatarProps) {
  const sizes = {
    sm: "w-8 h-8 text-xs",
    md: "w-10 h-10 text-sm",
    lg: "w-12 h-12 text-base",
  };

  const getInitials = (n?: string) => {
    if (!n) return "D";
    const parts = n.trim().split(" ");
    if (parts.length >= 2) return `${parts[0][0]}${parts[1][0]}`.toUpperCase();
    return n.slice(0, 2).toUpperCase();
  };

  return (
    <div
      className={cn(
        "relative inline-flex items-center justify-center rounded-full bg-[var(--border-subtle)] text-[var(--text-primary)] font-semibold border border-[var(--border-color)] overflow-hidden shrink-0 select-none",
        sizes[size],
        className
      )}
      {...props}
    >
      {src ? (
        <img src={src} alt={alt || name || "Avatar"} className="w-full h-full object-cover" />
      ) : (
        <span>{getInitials(name)}</span>
      )}
    </div>
  );
}

export interface DividerProps extends React.HTMLAttributes<HTMLDivElement> {
  orientation?: "horizontal" | "vertical";
  label?: string;
}

export function Divider({ orientation = "horizontal", label, className = "", ...props }: DividerProps) {
  if (orientation === "vertical") {
    return <div className={cn("w-px h-full bg-[var(--border-color)] mx-2 self-stretch", className)} {...props} />;
  }

  return (
    <div className={cn("w-full flex items-center my-4", className)} {...props}>
      <div className="flex-grow border-t border-[var(--border-color)]" />
      {label && (
        <span className="px-3 text-xs text-[var(--text-muted)] font-medium uppercase tracking-wider">
          {label}
        </span>
      )}
      <div className="flex-grow border-t border-[var(--border-color)]" />
    </div>
  );
}

export interface TooltipProps {
  content: string;
  children: React.ReactNode;
  position?: "top" | "bottom" | "left" | "right";
}

export function Tooltip({ content, children, position = "top" }: TooltipProps) {
  const [visible, setVisible] = React.useState(false);

  const positions = {
    top: "bottom-full mb-2 left-1/2 -translate-x-1/2",
    bottom: "top-full mt-2 left-1/2 -translate-x-1/2",
    left: "right-full mr-2 top-1/2 -translate-y-1/2",
    right: "left-full ml-2 top-1/2 -translate-y-1/2",
  };

  return (
    <div
      className="relative inline-block"
      onMouseEnter={() => setVisible(true)}
      onMouseLeave={() => setVisible(false)}
      onFocus={() => setVisible(true)}
      onBlur={() => setVisible(false)}
    >
      {children}
      {visible && (
        <div
          role="tooltip"
          className={cn(
            "absolute z-[1700] px-2.5 py-1 text-xs font-medium text-white bg-[var(--text-primary)] rounded shadow-md whitespace-nowrap pointer-events-none transition-opacity duration-150",
            positions[position]
          )}
        >
          {content}
        </div>
      )}
    </div>
  );
}
