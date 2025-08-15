'use client';

import * as React from 'react';
import { ChevronRight, ChevronDown, Home, Plus } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils/cn';

export type ProcessItem = {
  id: string;
  name: string;
};

export type ProjectItem = {
  id: string;
  name: string;
  processes: ProcessItem[];
};

interface SidebarProps {
  projects: ProjectItem[];
  currentView: 'home' | 'project' | 'process' | 'profile' | 'settings';
  selectedProjectId: string | null;
  selectedProcessId: string | null;
  onGoHome(): void;
  onSelectProject(id: string): void;
  onSelectProcess(projectId: string, processId: string): void;
  onNewProject(): void;
}

export function Sidebar({
  projects,
  currentView,
  selectedProjectId,
  selectedProcessId,
  onGoHome,
  onSelectProject,
  onSelectProcess,
  onNewProject,
}: SidebarProps) {
  const [expanded, setExpanded] = React.useState<Record<string, boolean>>({});

  const toggle = (id: string) =>
    setExpanded((s) => ({ ...s, [id]: !s[id] }));

  return (
    <aside className="flex h-full w-64 flex-col border-r border-gray-200 bg-gray-50">
      <div className="border-b border-gray-200 p-4">
        <h1 className="text-xl font-bold text-gray-900">RWMS</h1>
      </div>

      <nav className="flex-1 space-y-6 p-4">
        <button
          onClick={onGoHome}
          className={cn(
            'flex w-full items-center rounded-lg px-3 py-2 text-left text-sm font-medium',
            currentView === 'home'
              ? 'bg-blue-100 text-blue-700'
              : 'text-gray-700 hover:bg-gray-100'
          )}
        >
          <Home size={16} className="mr-3" />
          Home
        </button>

        <div>
          <div className="mb-2 flex items-center justify-between">
            <h3 className="text-xs font-semibold uppercase tracking-wide text-gray-500">Projects</h3>
            <Button variant="ghost" size="sm" onClick={onNewProject} aria-label="Create project">
              <Plus size={14} />
            </Button>
          </div>

          <div className="space-y-1">
            {projects.map((p) => {
              const isProjectActive = currentView === 'project' && selectedProjectId === p.id;
              const isProcessActive = currentView === 'process' && selectedProjectId === p.id;
              const active = isProjectActive || isProcessActive;
              const isOpen = !!expanded[p.id];

              return (
                <div key={p.id} className="rounded">
                  <div className="flex items-center">
                    <button
                      onClick={() => toggle(p.id)}
                      className="rounded p-1 hover:bg-gray-100"
                      aria-label={isOpen ? 'Collapse project' : 'Expand project'}
                    >
                      {isOpen ? <ChevronDown size={14} /> : <ChevronRight size={14} />}
                    </button>

                    <button
                      onClick={() => onSelectProject(p.id)}
                      className={cn(
                        'flex-1 rounded-lg px-2 py-2 text-left text-sm',
                        active ? 'bg-blue-100 text-blue-700' : 'text-gray-700 hover:bg-gray-100'
                      )}
                    >
                      {p.name}
                    </button>
                  </div>

                  {isOpen && p.processes.length > 0 && (
                    <div className="ml-6 mt-1 space-y-1">
                      {p.processes.map((proc) => {
                        const activeProc =
                          currentView === 'process' &&
                          selectedProcessId === proc.id &&
                          selectedProjectId === p.id;

                        return (
                          <button
                            key={proc.id}
                            onClick={() => onSelectProcess(p.id, proc.id)}
                            className={cn(
                              'w-full rounded px-2 py-1 text-left text-xs',
                              activeProc
                                ? 'bg-blue-100 text-blue-700'
                                : 'text-gray-600 hover:bg-gray-100'
                            )}
                          >
                            {proc.name}
                          </button>
                        );
                      })}
                    </div>
                  )}
                </div>
              );
            })}
          </div>
        </div>
      </nav>
    </aside>
  );
}
