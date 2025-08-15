'use client';

import { useEffect } from 'react';
import { useAuthStore } from '@/lib/stores/auth-store';
import { useToast } from './use-toast';

export function useAuth() {
  const { success, error } = useToast();
  const token = useAuthStore((s) => s.token);
  const user = useAuthStore((s) => s.user);
  const status = useAuthStore((s) => s.status);
  const login = useAuthStore((s) => s.login);
  const logout = useAuthStore((s) => s.logout);
  const fetchMe = useAuthStore((s) => s.fetchMe);

  useEffect(() => {
    // On mount with a token but no user, try to fetch profile
    if (token && !user && status !== 'loading') {
      fetchMe().catch((e) => error(e?.message || 'Failed to load profile'));
    }
  }, [token, user, status, fetchMe, error]);

  return {
    user,
    token,
    status,
    isAuthenticated: !!token,
    login: async (loginStr: string, password: string) => {
      await login(loginStr, password);
      success('Signed in');
    },
    logout: () => {
      logout();
      success('Signed out');
    },
  };
}