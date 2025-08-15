'use client';

import { ToastContainer } from '@/components/ui/toast';
import { useUIStore } from '@/lib/stores/ui-store';

export function UIToaster() {
  const toasts = useUIStore((s) => s.toasts);
  const removeToast = useUIStore((s) => s.removeToast);
  // map store's level -> toast type
  return (
    <ToastContainer
      toasts={toasts.map(t => ({ id: t.id, message: t.message, type: t.level }))}
      onRemove={removeToast}
    />
  );
}