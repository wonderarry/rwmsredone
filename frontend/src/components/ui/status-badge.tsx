'use client';

import * as React from 'react';
import { cn } from '@/lib/utils/cn';

export type ProcessStateUI = 'draft' | 'pending' | 'approved' | 'rejected';

const styles: Record<ProcessStateUI, string> = {
  draft: 'bg-blue-100 text-blue-800',
  pending: 'bg-yellow-100 text-yellow-800',
  approved: 'bg-green-100 text-green-800',
  rejected: 'bg-red-100 text-red-800',
};

export function StatusBadge({ state, className }: { state: ProcessStateUI; className?: string }) {
  return (
    <span
      className={cn(
        'inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium',
        styles[state],
        className
      )}
    >
      {state}
    </span>
  );
}
