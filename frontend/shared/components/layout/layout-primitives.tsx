import * as React from "react";
import { cn } from "../../utils/cn";

export interface ContainerProps extends React.HTMLAttributes<HTMLDivElement> {
  size?: "sm" | "md" | "lg" | "xl" | "full";
}

export function Container({ size = "lg", className = "", ...props }: ContainerProps) {
  const sizes = {
    sm: "max-w-3xl",
    md: "max-w-5xl",
    lg: "max-w-7xl",
    xl: "max-w-[90rem]",
    full: "max-w-full",
  };

  return <div className={cn("w-full mx-auto px-4 sm:px-6 lg:px-8", sizes[size], className)} {...props} />;
}

export interface GridProps extends React.HTMLAttributes<HTMLDivElement> {
  cols?: 1 | 2 | 3 | 4 | 6 | 12;
}

export function Grid({ cols = 3, className = "", ...props }: GridProps) {
  const gridCols = {
    1: "grid-cols-1",
    2: "grid-cols-1 sm:grid-cols-2",
    3: "grid-cols-1 sm:grid-cols-2 lg:grid-cols-3",
    4: "grid-cols-1 sm:grid-cols-2 lg:grid-cols-4",
    6: "grid-cols-2 sm:grid-cols-3 lg:grid-cols-6",
    12: "grid-cols-12",
  };

  return <div className={cn("grid gap-4 md:gap-6", gridCols[cols], className)} {...props} />;
}

export interface StackProps extends React.HTMLAttributes<HTMLDivElement> {
  direction?: "row" | "col";
  gap?: 1 | 2 | 3 | 4 | 6 | 8;
  align?: "start" | "center" | "end" | "stretch";
  justify?: "start" | "center" | "end" | "between";
}

export function Stack({
  direction = "col",
  gap = 4,
  align = "stretch",
  justify = "start",
  className = "",
  ...props
}: StackProps) {
  const flexDirections = {
    row: "flex-row",
    col: "flex-col",
  };

  const gaps = {
    1: "gap-1",
    2: "gap-2",
    3: "gap-3",
    4: "gap-4",
    6: "gap-6",
    8: "gap-8",
  };

  const aligns = {
    start: "items-start",
    center: "items-center",
    end: "items-end",
    stretch: "items-stretch",
  };

  const justifies = {
    start: "justify-start",
    center: "justify-center",
    end: "justify-end",
    between: "justify-between",
  };

  return (
    <div
      className={cn("flex", flexDirections[direction], gaps[gap], aligns[align], justifies[justify], className)}
      {...props}
    />
  );
}

export function Section({ className = "", ...props }: React.HTMLAttributes<HTMLElement>) {
  return <section className={cn("py-6 md:py-8 border-b border-slate-200 dark:border-slate-800 last:border-0", className)} {...props} />;
}
