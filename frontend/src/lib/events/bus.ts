'use client';

import type { ApiEvent, DataMutationEvent, UIStateEvent, UserActionEvent } from '@/types/events';

export type RWMSAnyEvent = UserActionEvent | DataMutationEvent | UIStateEvent | ApiEvent;

const listeners = new Set<(e: RWMSAnyEvent) => void>();

export function emit(e: RWMSAnyEvent) {
  for (const l of listeners) l(e);
}

export function on(listener: (e: RWMSAnyEvent) => void) {
  listeners.add(listener);
  return () => listeners.delete(listener);
}