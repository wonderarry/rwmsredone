'use client';
import * as React from 'react';
import { cn } from '@/lib/utils/cn';

export type ProcessStateUI = 'draft' | 'pending' | 'approved' | 'rejected';

export function StatusBadge({ state, className }: { state: ProcessStateUI; className?: string }) {
  // light: solid tints; dark: use alpha over dark surface
  const base = 'inline-flex items-center px-2.5 py-0.5 radii-sm text-xs font-medium';

  const mapLight: Record<ProcessStateUI, string> = {
    draft:    'bg-[hsl(var(--info-bg))] text-[hsl(var(--info-fg))] border border-[hsl(var(--info-border))]',
    pending:  'bg-[hsl(var(--warning-bg))] text-[hsl(var(--warning-fg))] border border-[hsl(var(--warning-border))]',
    approved: 'bg-[hsl(var(--success-bg))] text-[hsl(var(--success-fg))] border border-[hsl(var(--success-border))]',
    rejected: 'bg-[hsl(var(--danger-bg))] text-[hsl(var(--danger-fg))] border border-[hsl(var(--danger-border))]',
  };

  const mapDark: Record<ProcessStateUI, string> = {
    draft:    'bg-[hsl(var(--info)/.15)] text-[hsl(var(--info))] border border-[hsl(var(--info)/.35)]',
    pending:  'bg-[hsl(var(--warning)/.15)] text-[hsl(var(--warning))] border border-[hsl(var(--warning)/.35)]',
    approved: 'bg-[hsl(var(--success)/.15)] text-[hsl(var(--success))] border border-[hsl(var(--success)/.35)]',
    rejected: 'bg-[hsl(var(--danger)/.15)] text-[hsl(var(--danger))] border border-[hsl(var(--danger)/.35)]',
  };

  return (
    <span
      className={cn(
        base,
        mapLight[state],
        'dark:' + mapDark[state], // Tailwind v4 supports arbitrary values with "dark:"
        className
      )}
    >
      {state}
    </span>
  );
}
