"use client";

import React, { createContext, useContext, useState } from "react";

export interface SessionUser {
  id: string;
  email: string;
  vaultId?: string;
}

interface SessionContextType {
  user: SessionUser | null;
  isAuthenticated: boolean;
  setUser: (user: SessionUser | null) => void;
}

const SessionContext = createContext<SessionContextType | undefined>(undefined);

export function SessionProvider({ children }: { children: React.ReactNode }) {
  // Placeholder unauthenticated state until Auth module integration
  const [user, setUser] = useState<SessionUser | null>(null);

  return (
    <SessionContext.Provider value={{ user, isAuthenticated: !!user, setUser }}>
      {children}
    </SessionContext.Provider>
  );
}

export function useSession() {
  const context = useContext(SessionContext);
  if (!context) {
    throw new Error("useSession must be used within a SessionProvider");
  }
  return context;
}
