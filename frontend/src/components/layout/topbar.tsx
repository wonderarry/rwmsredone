'use client';

import * as React from 'react';
import { Bell, Settings, LogOut } from 'lucide-react';
import { Avatar } from '@/components/ui/avatar';

interface TopbarProps {
  userName: string;
  onOpenProfile(): void;
  onOpenSettings(): void;
  onSignOut(): void;
}

export function Topbar({ userName, onOpenProfile, onOpenSettings, onSignOut }: TopbarProps) {
  return (
    <header className="border-b border-gray-200 bg-white px-6 py-4">
      <div className="flex items-center justify-between">
        <div className="flex-1" />

        <div className="flex items-center space-x-4">
          <button className="rounded p-2 text-gray-400 hover:text-gray-600" aria-label="Notifications">
            <Bell size={20} />
          </button>

          <button
            onClick={onOpenProfile}
            className="flex items-center space-x-3 rounded-lg p-2 hover:bg-gray-50"
            aria-label="Open Profile"
          >
            <Avatar name={userName} />
            <span className="text-sm font-medium">{userName}</span>
          </button>

          <button
            onClick={onOpenSettings}
            className="rounded p-2 text-gray-400 hover:text-gray-600"
            aria-label="Open Settings"
          >
            <Settings size={20} />
          </button>

          <button
            onClick={onSignOut}
            className="rounded p-2 text-gray-400 hover:text-gray-600"
            aria-label="Sign out"
          >
            <LogOut size={20} />
          </button>
        </div>
      </div>
    </header>
  );
}
