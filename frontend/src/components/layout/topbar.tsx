'use client';

import * as React from 'react';
import { Bell, Settings, LogOut } from 'lucide-react';
import { Menu, Transition, Popover } from '@headlessui/react';
import { Avatar } from '@/components/ui/avatar';
import { cn } from '@/lib/utils/cn';
import { Sun, Moon } from 'lucide-react';
import { PopoverButton } from '@headlessui/react';
import { applyTheme, Theme, themeGetter, ThemeProvider } from '@/lib/providers/theme-provider';

interface TopbarProps {
    userName: string;
    onOpenProfile(): void;
    onOpenSettings(): void;
    onSignOut(): void;
}

export function Topbar({ userName, onOpenProfile, onOpenSettings, onSignOut }: TopbarProps) {
    const [mounted, setMounted] = React.useState(false);
    React.useEffect(() => setMounted(true), []);
    const [themeMode, setThemeMode] = React.useState<Theme>(themeGetter())

    React.useEffect(() => {
        applyTheme(themeMode);
    }, [themeMode])
    return (
        <header className="border-b border-base bg-[hsl(var(--surface))] px-6 py-4">
            <div className="flex items-center justify-between">
                <div className="flex-1" />

                <div className="flex items-center space-x-2">
                    {/* Notifications popover (placeholder content) */}
                    <Popover className="relative align-center">
                        <button
                            onClick={() => setThemeMode((tm) => {
                                return tm === 'dark' ? 'light' : 'dark'
                            })}
                            className={cn(
              'relative overflow-hidden rounded-full p-2 text-[hsl(var(--muted))] focus:outline-none',
              'hover:bg-[hsl(var(--surface-2))]',
              // ripple/fill
              "before:content-[''] before:absolute before:inset-0 before:rounded-full",
              'before:bg-[hsl(var(--fg)/0.18)] before:scale-0 before:origin-center',
              'before:transition-transform before:duration-300 before:ease-out',
              'hover:before:scale-100 active:before:scale-125'
            )}
                        >
                            {mounted ? (themeMode === 'dark' ? <Moon size={20}/> : <Sun size={20}/>) : null}
                        </button>
                        <Popover.Button
                            className="rounded p-2 text-[hsl(var(--muted))] hover:bg-[hsl(var(--surface-2))]"
                            aria-label="Notifications"
                        >
                            <Bell size={20} />
                        </Popover.Button>
                        <Transition
                            enter="transition duration-100 ease-out"
                            enterFrom="transform opacity-0 scale-95"
                            enterTo="transform opacity-100 scale-100"
                            leave="transition duration-75 ease-in"
                            leaveFrom="transform opacity-100 scale-100"
                            leaveTo="transform opacity-0 scale-95"
                        >
                            <Popover.Panel className="absolute right-0 z-50 mt-2 w-64 radii-md border border-base bg-[hsl(var(--surface))] p-3 shadow-soft">
                                <div className="text-sm text-[hsl(var(--muted))]">No notifications (placeholder)</div>
                            </Popover.Panel>
                        </Transition>
                    </Popover>

                    {/* Profile menu */}
                    <Menu as="div" className="relative">
                        <Menu.Button className="flex items-center space-x-3 rounded-lg p-2 hover:bg-[hsl(var(--surface-2))]">
                            <Avatar name={userName} />
                            <span className="text-sm font-medium text-[hsl(var(--fg))]">{userName}</span>
                        </Menu.Button>

                        <Transition
                            enter="transition duration-100 ease-out"
                            enterFrom="transform opacity-0 scale-95"
                            enterTo="transform opacity-100 scale-100"
                            leave="transition duration-75 ease-in"
                            leaveFrom="transform opacity-100 scale-100"
                            leaveTo="transform opacity-0 scale-95"
                        >
                            <Menu.Items className="absolute right-0 z-50 mt-2 w-48 overflow-hidden radii-md border border-base bg-[hsl(var(--surface))] shadow-soft">
                                <Menu.Item>
                                    {({ active }) => (
                                        <button
                                            className={cn(
                                                'block w-full px-4 py-2 text-left text-sm',
                                                active ? 'bg-[hsl(var(--surface-2))]' : ''
                                            )}
                                            onClick={onOpenProfile}
                                        >
                                            Profile
                                        </button>
                                    )}
                                </Menu.Item>
                                <Menu.Item>
                                    {({ active }) => (
                                        <button
                                            className={cn(
                                                'block w-full px-4 py-2 text-left text-sm',
                                                active ? 'bg-[hsl(var(--surface-2))]' : ''
                                            )}
                                            onClick={onOpenSettings}
                                        >
                                            <span className="inline-flex items-center gap-2">
                                                <Settings size={16} aria-hidden />
                                                Settings
                                            </span>
                                        </button>
                                    )}
                                </Menu.Item>
                                <div className="my-1 h-px bg-[hsl(var(--border))]" />
                                <Menu.Item>
                                    {({ active }) => (
                                        <button
                                            className={cn(
                                                'block w-full px-4 py-2 text-left text-sm text-[hsl(var(--danger))]',
                                                active ? 'bg-[hsl(var(--surface-2))]' : ''
                                            )}
                                            onClick={onSignOut}
                                        >
                                            <span className="inline-flex items-center gap-2">
                                                <LogOut size={16} aria-hidden />
                                                Sign out
                                            </span>
                                        </button>
                                    )}
                                </Menu.Item>
                            </Menu.Items>
                        </Transition>
                    </Menu>
                </div>
            </div>
        </header>
    );
}
