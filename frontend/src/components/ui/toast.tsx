'use client';
import * as React from 'react';
import { cn } from '@/lib/utils/cn';

export type ToastType = 'success' | 'error' | 'info';
export interface ToastItem { id: string; message: string; type?: ToastType; durationMs?: number; }
export interface ToastProps extends Omit<ToastItem, 'id'> { onClose?: () => void; className?: string; }

const LIGHT: Record<ToastType, string> = {
    success: 'bg-[hsl(var(--success-bg))] text-[hsl(var(--success-fg))] border border-[hsl(var(--success-border))]',
    error: 'bg-[hsl(var(--danger-bg))]  text-[hsl(var(--danger-fg))]  border border-[hsl(var(--danger-border))]',
    info: 'bg-[hsl(var(--info-bg))]    text-[hsl(var(--info-fg))]    border border-[hsl(var(--info-border))]',
};
const DARK: Record<ToastType, string> = {
    success: 'bg-[hsl(var(--success)/.15)] text-[hsl(var(--success))] border border-[hsl(var(--success)/.35)]',
    error: 'bg-[hsl(var(--danger)/.15)]  text-[hsl(var(--danger))]  border border-[hsl(var(--danger)/.35)]',
    info: 'bg-[hsl(var(--info)/.15)]    text-[hsl(var(--info))]    border border-[hsl(var(--info)/.35)]',
};
const GLYPH: Record<ToastType, string> = { success: '✓', error: '✗', info: 'ℹ' };

export function Toast({ message, type = 'success', durationMs = 3000, onClose, className }: ToastProps) {
    const [hovered, setHovered] = React.useState(false);
    const started = React.useRef<number>(0);
    const raf = React.useRef<number | null>(null);

    React.useEffect(() => {
        if (hovered) return; // returning undefined is fine

        started.current = performance.now();

        const tick = (now: number) => {
            if (now - started.current >= durationMs) {
                onClose?.();
            } else {
                raf.current = requestAnimationFrame(tick);
            }
        };

        raf.current = requestAnimationFrame(tick);

        return () => {
            if (raf.current !== null) {
                cancelAnimationFrame(raf.current);
                raf.current = null; // avoid stale id
            }
        };
    }, [hovered, durationMs, onClose]);

    return (
        <div
            role="status"
            aria-live={type === 'error' ? 'assertive' : 'polite'}
            className={cn(
                'pointer-events-auto flex items-center radii-md px-4 py-3 shadow-soft transition-all duration-200',
                LIGHT[type],
                'dark:' + DARK[type],
                className
            )}
            onMouseEnter={() => setHovered(true)}
            onMouseLeave={() => setHovered(false)}
        >
            <span className="mr-2" aria-hidden>{GLYPH[type]}</span>
            <span className="flex-1">{message}</span>
            {onClose && (
                <button
                    className="ml-3 rounded p-1 text-[hsl(var(--muted))] hover:opacity-80 focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]"
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
export function ToastContainer({ toasts, onRemove, position = 'top-right' }: ToastContainerProps) {
    const pos =
        position === 'top-right' ? 'top-4 right-4' :
            position === 'top-left' ? 'top-4 left-4' :
                position === 'bottom-right' ? 'bottom-4 right-4' : 'bottom-4 left-4';
    return (
        <div className={cn('pointer-events-none fixed z-[60] flex max-w-[calc(100%-2rem)] w-[360px] flex-col gap-3', pos)}
            aria-live="polite" aria-relevant="additions">
            {toasts.map(t => (
                <Toast key={t.id} message={t.message} type={t.type} durationMs={t.durationMs} onClose={() => onRemove(t.id)} />
            ))}
        </div>
    );
}
