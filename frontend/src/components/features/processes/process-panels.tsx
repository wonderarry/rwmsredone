'use client';

import * as React from 'react';
import { GitBranch, Upload } from 'lucide-react';
import { Transition } from '@headlessui/react';
import { Button } from '@/components/ui/button';
import { Avatar } from '@/components/ui/avatar';
import { StatusBadge, type ProcessStateUI } from '@/components/ui/status-badge';
import { cn } from '@/lib/utils/cn';

export function ProcessHeader({
  name,
  state,
  stage,
  showActions,
  onApprove,
  onReject,
}: {
  name: string;
  state: ProcessStateUI;
  stage: string;
  showActions: boolean;
  onApprove(): void;
  onReject(): void;
}) {
  return (
    <div className="mb-6 flex items-center justify-between">
      <div>
        <h1 className="mb-2 text-3xl font-bold text-[hsl(var(--fg))]">{name}</h1>
        <div className="flex items-center space-x-4">
          <StatusBadge state={state} />
          <span className="text-[hsl(var(--muted))]">Stage: {stage}</span>
        </div>
      </div>

      {showActions && (
        <div className="flex space-x-2">
          <Button variant="destructive" onClick={onReject}>Reject</Button>
          <Button onClick={onApprove}>Approve</Button>
        </div>
      )}
    </div>
  );
}

export function ProcessGraphCard() {
  return (
    <div className="radii-lg border border-base bg-[hsl(var(--surface))] shadow-soft">
      <div className="border-b border-base p-6">
        <h2 className="text-xl font-semibold text-[hsl(var(--fg))]">Process Graph</h2>
      </div>
      <div className="p-6">
        <div className="flex aspect-video items-center justify-center rounded-lg bg-[hsl(var(--surface-2))] text-[hsl(var(--muted))]">
          <GitBranch size={48} />
          <span className="ml-3">Process visualization will appear here</span>
        </div>
      </div>
    </div>
  );
}

export function ArtifactsCard({ onFilesSelected }: { onFilesSelected?: (files: File[]) => void }) {
  const [dragActive, setDragActive] = React.useState(false);
  const inputRef = React.useRef<HTMLInputElement | null>(null);

  const handleFiles = (files: FileList | null) => {
    if (!files || files.length === 0) return;
    onFilesSelected?.(Array.from(files));
  };

  return (
    <div className="radii-lg border border-base bg-[hsl(var(--surface))] shadow-soft">
      <div className="border-b border-base p-6">
        <h2 className="text-xl font-semibold text-[hsl(var(--fg))]">Artifacts</h2>
      </div>

      <div className="p-6">
        <div
          className={cn(
            'relative radii-lg border-2 border-dashed border-base p-8 text-center',
            'bg-[hsl(var(--surface))]'
          )}
          onDragOver={(e) => { e.preventDefault(); setDragActive(true); }}
          onDragEnter={(e) => { e.preventDefault(); setDragActive(true); }}
          onDragLeave={(e) => { e.preventDefault(); setDragActive(false); }}
          onDrop={(e) => {
            e.preventDefault();
            setDragActive(false);
            handleFiles(e.dataTransfer?.files ?? null);
          }}
          role="button"
          tabIndex={0}
          aria-label="Upload files by drag and drop or browse"
        >
          <Upload size={48} className="mx-auto mb-4 text-[hsl(var(--muted))]" />
          <p className="mb-2 text-[hsl(var(--muted))]">
            Drag and drop files here, or click to browse
          </p>
          <Button
            variant="ghost"
            size="sm"
            onClick={() => inputRef.current?.click()}
          >
            Choose Files
          </Button>

          <input
            ref={inputRef}
            type="file"
            multiple
            className="hidden"
            onChange={(e) => handleFiles(e.target.files)}
          />

          <Transition
            show={dragActive}
            enter="transition-opacity duration-150"
            enterFrom="opacity-0"
            enterTo="opacity-100"
            leave="transition-opacity duration-150"
            leaveFrom="opacity-100"
            leaveTo="opacity-0"
          >
            <div className="pointer-events-none absolute inset-0 radii-lg ring-2 ring-[hsl(var(--ring))] ring-offset-2" />
          </Transition>
        </div>
      </div>
    </div>
  );
}

export function MessagesCard() {
  return (
    <div className="radii-lg border border-base bg-[hsl(var(--surface))] shadow-soft">
      <div className="border-b border-base p-6">
        <h2 className="text-xl font-semibold text-[hsl(var(--fg))]">Messages</h2>
      </div>
      <div className="p-6">
        <div className="mb-4 space-y-4">
          <div className="flex space-x-3">
            <Avatar name="John Smith" size="sm" />
            <div className="flex-1">
              <div className="radii-md bg-[hsl(var(--surface-2))] p-3">
                <p className="text-sm text-[hsl(var(--fg))]">
                  Initial review completed. Ready for next stage.
                </p>
              </div>
              <p className="mt-1 text-xs text-[hsl(var(--muted))]">2 hours ago</p>
            </div>
          </div>
        </div>

        <div className="flex space-x-2">
          <input
            type="text"
            placeholder="Type a message..."
            className={cn(
              'flex-1 radii-md border px-3 py-2',
              'border-base bg-[hsl(var(--surface))] text-[hsl(var(--fg))]',
              'focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]'
            )}
          />
          <Button size="sm">Send</Button>
        </div>
      </div>
    </div>
  );
}
