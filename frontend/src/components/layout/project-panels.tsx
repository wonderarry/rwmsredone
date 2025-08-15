'use client';

import * as React from 'react';
import { Users, Plus } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils/cn';

export type ProcessStateUI = 'draft' | 'pending' | 'approved' | 'rejected';
export type ProcessItem = { id: string; name: string; stage: string; state: ProcessStateUI };

export function StatusBadge({ state }: { state: ProcessStateUI }) {
  const map: Record<ProcessStateUI, string> = {
    draft: 'bg-blue-100 text-blue-800',
    pending: 'bg-yellow-100 text-yellow-800',
    approved: 'bg-green-100 text-green-800',
    rejected: 'bg-red-100 text-red-800',
  };
  return (
    <span className={cn('inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium', map[state])}>
      {state}
    </span>
  );
}

export function ProcessesList({
  processes,
  onOpenProcess,
}: {
  processes: ProcessItem[];
  onOpenProcess: (processId: string) => void;
}) {
  return (
    <div className="divide-y divide-gray-200">
      {processes.map((proc) => (
        <button
          key={proc.id}
          onClick={() => onOpenProcess(proc.id)}
          className="flex w-full items-center justify-between p-6 text-left hover:bg-gray-50"
        >
          <div>
            <h3 className="font-medium text-gray-900">{proc.name}</h3>
            <p className="mt-1 text-sm text-gray-500">{proc.stage}</p>
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
    <div className="rounded-xl border border-gray-200 bg-white p-6 shadow-sm">
      <h3 className="mb-4 font-semibold text-gray-900">Project Details</h3>
      <dl className="space-y-3">
        <div>
          <dt className="text-sm font-medium text-gray-500">Theme</dt>
          <dd className="text-sm text-gray-900">{theme}</dd>
        </div>
        <div>
          <dt className="text-sm font-medium text-gray-500">Members</dt>
          <dd className="text-sm text-gray-900">{members}</dd>
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
