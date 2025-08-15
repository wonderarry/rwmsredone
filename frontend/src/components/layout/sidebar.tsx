'use client';

import * as React from 'react';
import { Disclosure, Transition } from '@headlessui/react';
import { ChevronRight, Home, Plus } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { cn } from '@/lib/utils/cn';

export type ProcessItem = { id: string; name: string };
export type ProjectItem = { id: string; name: string; processes: ProcessItem[] };

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
  return (
    <aside className="flex h-full w-64 flex-col border-r border-base bg-[hsl(var(--surface-2))]">
      <div className="border-b border-base p-4">
        <h1 className="text-xl font-bold text-[hsl(var(--fg))]">RWMS</h1>
      </div>

      <nav className="flex-1 space-y-6 p-4">
        <button
          onClick={onGoHome}
          className={cn(
            'flex w-full items-center rounded-lg px-3 py-2 text-left text-sm font-medium',
            currentView === 'home'
              ? 'bg-[hsl(var(--brand-600)/.12)] text-[hsl(var(--brand-600))]'
              : 'text-[hsl(var(--fg))] hover:bg-[hsl(var(--surface))]'
          )}
        >
          <Home size={16} className="mr-3" />
          Home
        </button>

        <div>
          <div className="mb-2 flex items-center justify-between">
            <h3 className="text-xs font-semibold uppercase tracking-wide text-[hsl(var(--muted))]">
              Projects
            </h3>
            <Button variant="ghost" size="sm" onClick={onNewProject} aria-label="Create project">
              <Plus size={14} />
            </Button>
          </div>

          <div className="space-y-1">
            {projects.map((p) => {
              const isProjectActive = currentView === 'project' && selectedProjectId === p.id;
              const isProcessActive = currentView === 'process' && selectedProjectId === p.id;
              const active = isProjectActive || isProcessActive;

              return (
                <Disclosure key={p.id} defaultOpen={active}>
                  {({ open }) => (
                    <div className="rounded">
                      <Disclosure.Button
                        className={cn(
                          'flex w-full items-center rounded-lg px-2 py-2 text-left text-sm',
                          open ? '' : '',
                          active
                            ? 'bg-[hsl(var(--brand-600)/.12)] text-[hsl(var(--brand-600))]'
                            : 'text-[hsl(var(--fg))] hover:bg-[hsl(var(--surface))]'
                        )}
                        onClick={() => onSelectProject(p.id)}
                      >
                        <ChevronRight
                          size={14}
                          className={cn('mr-1 transition-transform', open && 'rotate-90')}
                          aria-hidden
                        />
                        <span className="flex-1">{p.name}</span>
                      </Disclosure.Button>

                      <Transition
                        show={open && p.processes.length > 0}
                        enter="transition duration-150 ease-out"
                        enterFrom="transform scale-y-95 opacity-0"
                        enterTo="transform scale-y-100 opacity-100"
                        leave="transition duration-100 ease-in"
                        leaveFrom="transform scale-y-100 opacity-100"
                        leaveTo="transform scale-y-95 opacity-0"
                      >
                        <Disclosure.Panel static>
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
                                      ? 'bg-[hsl(var(--brand-600)/.12)] text-[hsl(var(--brand-600))]'
                                      : 'text-[hsl(var(--muted))] hover:bg-[hsl(var(--surface))]'
                                  )}
                                >
                                  {proc.name}
                                </button>
                              );
                            })}
                          </div>
                        </Disclosure.Panel>
                      </Transition>
                    </div>
                  )}
                </Disclosure>
              );
            })}
          </div>
        </div>
      </nav>
    </aside>
  );
}
