'use client';

import * as React from 'react';
import { GitBranch, Upload } from 'lucide-react';
import { Button } from '@/components/ui/button';
import { Avatar } from '@/components/ui/avatar';
import { StatusBadge, type ProcessStateUI } from '@/components/ui/status-badge';

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
  onApprove: () => void;
  onReject: () => void;
}) {
  return (
    <div className="mb-6 flex items-center justify-between">
      <div>
        <h1 className="mb-2 text-3xl font-bold text-gray-900">{name}</h1>
        <div className="flex items-center space-x-4">
          <StatusBadge state={state} />
          <span className="text-gray-500">Stage: {stage}</span>
        </div>
      </div>

      {showActions && (
        <div className="flex space-x-2">
          <Button variant="destructive" onClick={onReject}>
            Reject
          </Button>
          <Button onClick={onApprove}>Approve</Button>
        </div>
      )}
    </div>
  );
}

export function ProcessGraphCard() {
  return (
    <div className="rounded-xl border border-gray-200 bg-white shadow-sm">
      <div className="border-b border-gray-200 p-6">
        <h2 className="text-xl font-semibold">Process Graph</h2>
      </div>
      <div className="p-6">
        <div className="aspect-video rounded-lg bg-gray-100 text-gray-500 flex items-center justify-center">
          <GitBranch size={48} />
          <span className="ml-3">Process visualization will appear here</span>
        </div>
      </div>
    </div>
  );
}

export function ArtifactsCard() {
  return (
    <div className="rounded-xl border border-gray-200 bg-white shadow-sm">
      <div className="border-b border-gray-200 p-6">
        <h2 className="text-xl font-semibold">Artifacts</h2>
      </div>
      <div className="p-6">
        <div className="border-2 border-dashed border-gray-300 rounded-lg p-8 text-center">
          <Upload size={48} className="mx-auto text-gray-400 mb-4" />
          <p className="text-gray-500 mb-2">Drag and drop files here, or click to browse</p>
          <Button variant="ghost" size="sm">Choose Files</Button>
        </div>
      </div>
    </div>
  );
}

export function MessagesCard() {
  return (
    <div className="rounded-xl border border-gray-200 bg-white shadow-sm">
      <div className="border-b border-gray-200 p-6">
        <h2 className="text-xl font-semibold">Messages</h2>
      </div>
      <div className="p-6">
        <div className="space-y-4 mb-4">
          <div className="flex space-x-3">
            <Avatar name="John Smith" size="sm" />
            <div className="flex-1">
              <div className="bg-gray-100 rounded-lg p-3">
                <p className="text-sm">Initial review completed. Ready for next stage.</p>
              </div>
              <p className="text-xs text-gray-500 mt-1">2 hours ago</p>
            </div>
          </div>
        </div>

        <div className="flex space-x-2">
          <input
            type="text"
            placeholder="Type a message..."
            className="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          />
          <Button size="sm">Send</Button>
        </div>
      </div>
    </div>
  );
}
