'use client';
import * as React from 'react';
import { cn } from '@/lib/utils/cn';

export type ButtonVariant = 'primary' | 'secondary' | 'ghost' | 'destructive';
export type ButtonSize = 'sm' | 'md' | 'lg';

export interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant; size?: ButtonSize; fullWidth?: boolean;
}

const base =
  'inline-flex items-center justify-center font-medium transition-colors radii-md ' +
  'focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))] focus:ring-offset-2 ' +
  'disabled:opacity-50 disabled:cursor-not-allowed';

const variants: Record<ButtonVariant, string> = {
  primary:
    'text-white bg-[hsl(var(--brand-600))] hover:bg-[hsl(var(--brand-700))] shadow-soft',
  secondary:
    'bg-[hsl(var(--surface-2))] text-[hsl(var(--fg))] hover:bg-[hsl(var(--border))]',
  ghost:
    'bg-transparent text-[hsl(var(--muted))] hover:bg-[hsl(var(--surface-2))]',
  destructive:
    'text-white bg-[hsl(var(--danger))] hover:brightness-95 shadow-soft',
};

const sizes: Record<ButtonSize, string> = {
  sm: 'px-3 py-1.5 text-sm radii-sm',
  md: 'px-4 py-2 text-sm',
  lg: 'px-6 py-3 text-base radii-lg',
};

export const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant = 'primary', size = 'md', fullWidth, ...props }, ref) => {
    return (
      <button
        ref={ref}
        className={cn(base, variants[variant], sizes[size], fullWidth && 'w-full', className)}
        {...props}
      />
    );
  }
);
Button.displayName = 'Button';
