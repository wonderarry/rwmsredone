'use client';

import { useUIStore, ToastLevel } from '@/lib/stores/ui-store';

export function useToast() {
  const show = useUIStore((s) => s.showToast);
  return {
    success: (m: string, d?: number) => show('success', m, d),
    error: (m: string, d?: number) => show('error', m, d),
    info: (m: string, d?: number) => show('info', m, d),
  };
}