/**
 * Declutr Design System — Centralized Token Definitions
 */

export const SPACING = {
  xs: "0.25rem",  // 4px
  sm: "0.5rem",   // 8px
  md: "1rem",     // 16px
  lg: "1.5rem",   // 24px
  xl: "2rem",     // 32px
  "2xl": "3rem",  // 48px
  "3xl": "4rem",  // 64px
} as const;

export const RADIUS = {
  none: "0px",
  sm: "0.375rem", // 6px
  md: "0.5rem",   // 8px
  lg: "0.75rem",  // 12px
  xl: "1rem",     // 16px
  full: "9999px",
} as const;

export const Z_INDEX = {
  hide: -1,
  auto: "auto",
  base: 0,
  dock: 10,
  dropdown: 1000,
  sticky: 1100,
  banner: 1200,
  overlay: 1300,
  modal: 1400,
  popover: 1500,
  toast: 1600,
  tooltip: 1700,
} as const;

export const MOTION = {
  duration: {
    fast: "150ms",
    normal: "250ms",
    slow: "350ms",
  },
  easing: {
    easeInOut: "cubic-bezier(0.4, 0, 0.2, 1)",
    easeOut: "cubic-bezier(0, 0, 0.2, 1)",
    easeIn: "cubic-bezier(0.4, 0, 1, 1)",
    spring: "cubic-bezier(0.175, 0.885, 0.32, 1.275)",
  },
} as const;

export const ICON_SIZE = {
  xs: 12,
  sm: 16,
  md: 20,
  lg: 24,
  xl: 32,
} as const;
