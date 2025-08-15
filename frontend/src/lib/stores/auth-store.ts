'use client';

import { create } from 'zustand';
import { persist } from 'zustand/middleware';
import type { Account } from '@/types/domain';
import { STORAGE_KEYS } from '@/lib/utils/constants';
import { setAuthToken, setUnauthorizedHandler } from '@/lib/api/client';
import * as Auth from '@/lib/api/auth';
import { emit } from '@/lib/events/bus';

interface AuthState {
  user: Account | null;
  token: string | null;
  status: 'idle' | 'loading' | 'authenticated' | 'error';
  error?: string;

  // actions
  login(login: string, password: string): Promise<void>;
  logout(): void;
  fetchMe(): Promise<void>;
  setToken(token: string | null): void;
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      user: null,
      token: null,
      status: 'idle',

      async login(login, password) {
        try {
          set({ status: 'loading', error: undefined });
          emit({ type: 'api:start', key: 'auth:login' });
          const { token } = await Auth.loginLocal({ login, password });
          get().setToken(token);
          await get().fetchMe();
          set({ status: 'authenticated' });
          emit({ type: 'api:success', key: 'auth:login' });
        } catch (e: any) {
          set({ status: 'error', error: e?.message || 'Login failed' });
          emit({ type: 'api:error', key: 'auth:login', error: String(e?.message || e) });
          throw e;
        }
      },

      logout() {
        set({ token: null, user: null, status: 'idle' });
        setAuthToken(null);
      },

      async fetchMe() {
        try {
          const u = await Auth.getMe();
          set({ user: u });
        } catch (e) {
          // keep token but surface error to caller
          throw e;
        }
      },

      setToken(token) {
        set({ token });
        setAuthToken(token);
      },
    }),
    { name: STORAGE_KEYS.authToken, partialize: (s) => ({ token: s.token }) }
  )
);

// Ensure 401s clear auth globally
setUnauthorizedHandler(() => {
  useAuthStore.getState().logout();
});
