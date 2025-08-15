'use client';

import * as React from 'react';
import { cn } from '@/lib/utils/cn';

export type AvatarSize = 'sm' | 'md' | 'lg';

const sizes: Record<AvatarSize, string> = {
  sm: 'h-6 w-6 text-xs',
  md: 'h-8 w-8 text-sm',
  lg: 'h-12 w-12 text-base',
};

export interface AvatarProps extends React.HTMLAttributes<HTMLDivElement> {
  name: string;
  size?: AvatarSize;
}

function getInitials(name: string) {
  return name
    .split(' ')
    .filter(Boolean)
    .map((n) => n[0]!)
    .join('')
    .toUpperCase();
}

export const Avatar = ({ name, size = 'md', className, ...props }: AvatarProps) => {
  const initials = getInitials(name);
  return (
    <div
      className={cn(
        sizes[size],
        'bg-blue-600 rounded-full flex items-center justify-center text-white font-medium',
        className
      )}
      aria-label={name}
      {...props}
    >
      {initials}
    </div>
  );
};
