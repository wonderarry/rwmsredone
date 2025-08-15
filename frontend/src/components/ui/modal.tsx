'use client';

import * as React from 'react';
import { Dialog, DialogPanel, DialogTitle, Transition } from '@headlessui/react';
import { X } from 'lucide-react';
import { cn } from '@/lib/utils/cn';

export interface ModalProps {
  isOpen: boolean;
  onClose: () => void;
  title?: React.ReactNode;
  children?: React.ReactNode;
  size?: 'sm' | 'md' | 'lg';
  className?: string;
}

const sizes = { sm: 'max-w-sm', md: 'max-w-md', lg: 'max-w-lg' };

export function Modal({
  isOpen,
  onClose,
  title,
  children,
  size = 'md',
  className,
}: ModalProps) {
  return (
    <Transition show={isOpen} appear>
      <Dialog onClose={onClose} className="fixed inset-0 z-50">
        {/* Overlay is its own element. Use bg alpha, not element opacity. */}
        <Transition.Child
          enter="transition-opacity duration-150"
          enterFrom="opacity-0"
          enterTo="opacity-100"
          leave="transition-opacity duration-150"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
        >
          <div className="fixed inset-0 z-40 bg-black/50" aria-hidden="true" />
        </Transition.Child>

        {/* Panel container sits above overlay */}
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4">
          <Transition.Child
            enter="transition duration-150 ease-out"
            enterFrom="opacity-0 scale-95"
            enterTo="opacity-100 scale-100"
            leave="transition duration-100 ease-in"
            leaveFrom="opacity-100 scale-100"
            leaveTo="opacity-0 scale-95"
          >
            <DialogPanel
              className={cn(
                'w-full radii-lg border border-base bg-[hsl(var(--surface))] text-[hsl(var(--fg))] shadow-soft',
                sizes[size],
                className
              )}
            >
              <div className="flex items-center justify-between border-b border-base px-6 py-4">
                {title ? <DialogTitle className="text-lg font-semibold">{title}</DialogTitle> : <span />}
                <button
                  onClick={onClose}
                  className="rounded p-1 text-[hsl(var(--muted))] hover:opacity-80 focus:outline-none focus:ring-2 focus:ring-[hsl(var(--ring))]"
                  aria-label="Close"
                >
                  <X size={20} />
                </button>
              </div>
              <div className="px-6 py-5">{children}</div>
            </DialogPanel>
          </Transition.Child>
        </div>
      </Dialog>
    </Transition>
  );
}
