import React from "react";
import { cn } from "../../utils/cn";

export interface TypographyProps extends React.HTMLAttributes<HTMLElement> {
  children: React.ReactNode;
  className?: string;
}

export function Display({ children, className = "", ...props }: TypographyProps) {
  return (
    <h1
      className={cn("text-4xl md:text-5xl lg:text-6xl font-extrabold tracking-tight text-[var(--text-primary)]", className)}
      {...props}
    >
      {children}
    </h1>
  );
}

export function Heading1({ children, className = "", ...props }: TypographyProps) {
  return (
    <h1
      className={cn("text-3xl md:text-4xl font-bold tracking-tight text-[var(--text-primary)]", className)}
      {...props}
    >
      {children}
    </h1>
  );
}

export function Heading2({ children, className = "", ...props }: TypographyProps) {
  return (
    <h2
      className={cn("text-2xl md:text-3xl font-semibold tracking-tight text-[var(--text-primary)]", className)}
      {...props}
    >
      {children}
    </h2>
  );
}

export function Heading3({ children, className = "", ...props }: TypographyProps) {
  return (
    <h3
      className={cn("text-xl md:text-2xl font-semibold tracking-tight text-[var(--text-primary)]", className)}
      {...props}
    >
      {children}
    </h3>
  );
}

export function Heading4({ children, className = "", ...props }: TypographyProps) {
  return (
    <h4
      className={cn("text-lg font-semibold text-[var(--text-primary)]", className)}
      {...props}
    >
      {children}
    </h4>
  );
}

export function Heading5({ children, className = "", ...props }: TypographyProps) {
  return (
    <h5
      className={cn("text-base font-semibold text-[var(--text-primary)]", className)}
      {...props}
    >
      {children}
    </h5>
  );
}

export function Heading6({ children, className = "", ...props }: TypographyProps) {
  return (
    <h6
      className={cn("text-sm font-semibold uppercase tracking-wider text-[var(--text-secondary)]", className)}
      {...props}
    >
      {children}
    </h6>
  );
}

export function Subtitle({ children, className = "", ...props }: TypographyProps) {
  return (
    <p
      className={cn("text-lg md:text-xl text-[var(--text-secondary)] leading-relaxed", className)}
      {...props}
    >
      {children}
    </p>
  );
}

export function Body({ children, className = "", ...props }: TypographyProps) {
  return (
    <p
      className={cn("text-base text-[var(--text-primary)] leading-normal", className)}
      {...props}
    >
      {children}
    </p>
  );
}

export function Caption({ children, className = "", ...props }: TypographyProps) {
  return (
    <span
      className={cn("text-xs text-[var(--text-muted)] tracking-normal", className)}
      {...props}
    >
      {children}
    </span>
  );
}

export function Mono({ children, className = "", ...props }: TypographyProps) {
  return (
    <code
      className={cn("font-mono text-sm bg-[var(--border-subtle)] text-[var(--brand-accent)] px-1.5 py-0.5 rounded-md", className)}
      {...props}
    >
      {children}
    </code>
  );
}
