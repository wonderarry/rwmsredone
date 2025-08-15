'use client';

import * as React from 'react';
import { cn } from '@/lib/utils/cn';

export type ToastType = 'success' | 'error' | 'info';

export interface ToastItem {
  id: string;
  message: string;
  type?: ToastType;       // default: 'success'
  durationMs?: number;    // default: 3000
}

export interface ToastProps extends Omit<ToastItem, 'id'> {
  onClose?: () => void;
  className?: string;
}

const COLORS: Record<ToastType, string> = {
  success: 'bg-green-50 text-green-800 border-green-200',
  error:   'bg-red-50 text-red-800 border-red-200',
  info:    'bg-blue-50 text-blue-800 border-blue-200',
};

const GLYPH: Record<ToastType, string> = {
  success: '✓',
  error:   '✗',
  info:    'ℹ',
};

export function Toast({
  message,
  type = 'success',
  durationMs = 3000,
  onClose,
  className,
}: ToastProps) {
  // auto-dismiss with pause-on-hover
  const [remaining, setRemaining] = React.useState(durationMs);
  const [hovered, setHovered] = React.useState(false);
  const start = React.useRef<number | null>(null);
  const raf = React.useRef<number | null>(null);

  React.useEffect(() => {
    if (hovered) return;
    start.current = performance.now();
    const tick = (now: number) => {
      const elapsed = now - (start.current ?? now);
      const left = durationMs - elapsed;
      setRemaining(left);
      if (left <= 0) {
        onClose?.();
      } else {
        raf.current = requestAnimationFrame(tick);
      }
    };
    raf.current = requestAnimationFrame(tick);
    return () => {
      if (raf.current) cancelAnimationFrame(raf.current);
    };
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [hovered, durationMs]);

  return (
    <div
      role="status"
      aria-live={type === 'error' ? 'assertive' : 'polite'}
      className={cn(
        'pointer-events-auto flex items-center rounded-xl border px-4 py-3 shadow-lg',
        'transition-all duration-200',
        COLORS[type],
        className
      )}
      onMouseEnter={() => setHovered(true)}
      onMouseLeave={() => setHovered(false)}
    >
      <span className="mr-2" aria-hidden>{GLYPH[type]}</span>
      <span className="flex-1">{message}</span>

      {onClose && (
        <button
          className="ml-3 rounded p-1 text-gray-500 hover:text-gray-700 focus:outline-none focus:ring-2 focus:ring-blue-500"
          onClick={onClose}
          aria-label="Close notification"
        >
          ×
        </button>
      )}
    </div>
  );
}

export interface ToastContainerProps {
  toasts: ToastItem[];
  onRemove: (id: string) => void;
  position?: 'top-right' | 'top-left' | 'bottom-right' | 'bottom-left';
}

export function ToastContainer({
  toasts,
  onRemove,
  position = 'top-right',
}: ToastContainerProps) {
  const posCls =
    position === 'top-right'
      ? 'top-4 right-4'
      : position === 'top-left'
      ? 'top-4 left-4'
      : position === 'bottom-right'
      ? 'bottom-4 right-4'
      : 'bottom-4 left-4';

  return (
    <div
      className={cn(
        'pointer-events-none fixed z-[60] flex w-[360px] max-w-[calc(100%-2rem)] flex-col gap-3',
        posCls
      )}
      aria-live="polite"
      aria-relevant="additions"
    >
      {toasts.map((t) => (
        <Toast
          key={t.id}
          message={t.message}
          type={t.type}
          durationMs={t.durationMs}
          onClose={() => onRemove(t.id)}
        />
      ))}
    </div>
  );
}
