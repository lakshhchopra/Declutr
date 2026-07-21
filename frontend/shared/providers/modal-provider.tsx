"use client";

import React, { createContext, useContext, useState, useCallback } from "react";

interface ModalContextType {
  activeModal: string | null;
  modalData: any;
  openModal: (modalId: string, data?: any) => void;
  closeModal: () => void;
}

const ModalContext = createContext<ModalContextType | undefined>(undefined);

export function ModalProvider({ children }: { children: React.ReactNode }) {
  const [activeModal, setActiveModal] = useState<string | null>(null);
  const [modalData, setModalData] = useState<any>(null);

  const openModal = useCallback((modalId: string, data?: any) => {
    setActiveModal(modalId);
    setModalData(data || null);
  }, []);

  const closeModal = useCallback(() => {
    setActiveModal(null);
    setModalData(null);
  }, []);

  return (
    <ModalContext.Provider value={{ activeModal, modalData, openModal, closeModal }}>
      {children}
    </ModalContext.Provider>
  );
}

export function useModal() {
  const context = useContext(ModalContext);
  if (!context) {
    throw new Error("useModal must be used within a ModalProvider");
  }
  return context;
}
