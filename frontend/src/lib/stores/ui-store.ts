'use client';

import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import { STORAGE_KEYS } from '@/lib/utils/constants';
import { emit } from '@/lib/events/bus';

export type ToastLevel = 'success' | 'error' | 'info';
export type Toast = { id: string; level: ToastLevel; message: string; duration?: number };

interface UIState {
  sidebarOpen: boolean;
  modals: Record<string, boolean>; // e.g. { createProject: true }
  toasts: Toast[];

  // actions
  toggleSidebar(): void;
  openModal(key: string): void;
  closeModal(key: string): void;
  showToast(level: ToastLevel, message: string, duration?: number): string; // returns id
  removeToast(id: string): void;
}

export const useUIStore = create<UIState>()(
  persist(
    (set, get) => ({
      sidebarOpen: true,
      modals: {},
      toasts: [],

      toggleSidebar() {
        set((s) => ({ sidebarOpen: !s.sidebarOpen }));
      },
      openModal(key) {
        emit({ type: 'modal:open', key });
        set((s) => ({ modals: { ...s.modals, [key]: true } }));
      },
      closeModal(key) {
        emit({ type: 'modal:close', key });
        set((s) => ({ modals: { ...s.modals, [key]: false } }));
      },
      showToast(level, message, duration = 3000) {
        const id = `t_${Math.random().toString(36).slice(2, 8)}`;
        emit({ type: 'toast:show', level, message });
        set((s) => ({ toasts: [...s.toasts, { id, level, message, duration }] }));
        // Auto-remove
        if (typeof window !== 'undefined') {
          setTimeout(() => get().removeToast(id), duration);
        }
        return id;
      },
      removeToast(id) {
        set((s) => ({ toasts: s.toasts.filter((t) => t.id !== id) }));
      },
    }),
    { name: `${STORAGE_KEYS.authToken}.ui` }
  )
);
