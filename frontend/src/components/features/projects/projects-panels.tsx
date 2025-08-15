'use client';

import * as React from 'react';
import { Users, Plus } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { StatusBadge, type ProcessStateUI } from '@/components/ui/status-badge';
import { cn } from '@/lib/utils/cn';

export type ProcessItem = { id: string; name: string; stage: string; state: ProcessStateUI };

export function ProcessesList({
  processes,
  onOpenProcess,
}: {
  processes: ProcessItem[];
  onOpenProcess: (processId: string) => void;
}) {
  return (
    <div className="divide-y divide-[hsl(var(--border))]">
      {processes.map((proc) => (
        <button
          key={proc.id}
          onClick={() => onOpenProcess(proc.id)}
          className={cn(
            'flex w-full items-center justify-between p-6 text-left',
            'hover:bg-[hsl(var(--surface-2))]'
          )}
        >
          <div>
            <h3 className="font-medium text-[hsl(var(--fg))]">{proc.name}</h3>
            <p className="mt-1 text-sm text-[hsl(var(--muted))]">{proc.stage}</p>
          </div>
          <StatusBadge state={proc.state} />
        </button>
      ))}
    </div>
  );
}

export function ProjectDetailsCard({
  theme,
  members,
  onManageMembers,
  onStartProcess,
}: {
  theme: string;
  members: number;
  onManageMembers: () => void;
  onStartProcess: () => void;
}) {
  return (
    <div className="radii-lg border border-base bg-[hsl(var(--surface))] p-6 shadow-soft">
      <h3 className="mb-4 font-semibold text-[hsl(var(--fg))]">Project Details</h3>

      <dl className="space-y-3">
        <div>
          <dt className="text-sm font-medium text-[hsl(var(--muted))]">Theme</dt>
          <dd className="text-sm text-[hsl(var(--fg))]">{theme}</dd>
        </div>
        <div>
          <dt className="text-sm font-medium text-[hsl(var(--muted))]">Members</dt>
          <dd className="text-sm text-[hsl(var(--fg))]">{members}</dd>
        </div>
      </dl>

      <div className="mt-4 space-y-2">
        <Button variant="secondary" className="w-full" onClick={onManageMembers}>
          <Users size={16} className="mr-2" />
          Manage Members
        </Button>
        <Button className="w-full" onClick={onStartProcess}>
          <Plus size={16} className="mr-2" />
          Start Process
        </Button>
      </div>
    </div>
  );
}
