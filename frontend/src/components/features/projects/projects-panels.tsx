'use client';

import * as React from 'react';
import { Users, Plus } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { StatusBadge, type ProcessStateUI } from '@/components/ui/status-badge';

export type ProcessItem = { id: string; name: string; stage: string; state: ProcessStateUI };

export function ProcessesList({
  processes,
  onOpenProcess,
}: {
  processes: ProcessItem[];
  onOpenProcess: (processId: string) => void;
}) {
  return (
    <div className="divide-y divide-gray-200">
      {processes.map((p) => (
        <button
          key={p.id}
          onClick={() => onOpenProcess(p.id)}
          className="w-full p-6 text-left hover:bg-gray-50 flex items-center justify-between"
        >
          <div>
            <h3 className="font-medium text-gray-900">{p.name}</h3>
            <p className="text-sm text-gray-500 mt-1">{p.stage}</p>
          </div>
          <StatusBadge state={p.state} />
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
    <div className="bg-white rounded-xl border border-gray-200 shadow-sm p-6">
      <h3 className="font-semibold text-gray-900 mb-4">Project Details</h3>
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

      <Button variant="secondary" className="w-full mt-4" onClick={onManageMembers}>
        <Users size={16} className="mr-2" />
        Manage Members
      </Button>

      <Button className="w-full mt-2" onClick={onStartProcess}>
        <Plus size={16} className="mr-2" />
        Start Process
      </Button>
    </div>
  );
}
