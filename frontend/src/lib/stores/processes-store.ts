'use client';

import { create } from 'zustand';
import type { ProcessBrief, ProcessLifecycleState } from '@/types/domain';
import * as Processes from '@/lib/api/processes';
import { emit } from '@/lib/events/bus';

interface ProcessesState {
  byId: Record<string, ProcessBrief>;
  selectedId: string | null;
  loading: boolean;
  error?: string;

  // actions
  select(id: string | null): void;
  create(input: { name: string; projectId: string; templateKey: string }): Promise<string>;
  approve(processId: string, actorRole: string, comment?: string): Promise<void>;
  reject(processId: string, actorRole: string, comment?: string): Promise<void>;
}

export const useProcessesStore = create<ProcessesState>((set) => ({
  byId: {},
  selectedId: null,
  loading: false,

  select(id) {
    set({ selectedId: id });
  },

  async create({ name, projectId, templateKey }) {
    emit({ type: 'api:start', key: 'processes:create' });
    try {
      const { id } = await Processes.createProcess({ name, project_id: projectId, template_key: templateKey });
      const proc: ProcessBrief = {
        id,
        name,
        currentStage: 'init', // placeholder until fetched from server
        state: 'active' as ProcessLifecycleState,
        projectId,
      };
      set((s) => ({ byId: { ...s.byId, [id]: proc } }));
      emit({ type: 'process:create', payload: { projectId } } as any);
      emit({ type: 'api:success', key: 'processes:create' });
      return id;
    } catch (e: any) {
      emit({ type: 'api:error', key: 'processes:create', error: String(e?.message || e) });
      throw e;
    }
  },

  async approve(processId, actorRole, comment) {
    emit({ type: 'api:start', key: 'processes:approve' });
    try {
      await Processes.submitApproval(processId, { decision: 'approve', actor_role: actorRole, comment });
      set((s) => ({ byId: { ...s.byId, [processId]: { ...s.byId[processId], lastDecision: 'approve' } } }));
      emit({ type: 'approval:submit', payload: { processId, decision: 'approve' } } as any);
      emit({ type: 'api:success', key: 'processes:approve' });
    } catch (e: any) {
      emit({ type: 'api:error', key: 'processes:approve', error: String(e?.message || e) });
      throw e;
    }
  },

  async reject(processId, actorRole, comment) {
    emit({ type: 'api:start', key: 'processes:reject' });
    try {
      await Processes.submitApproval(processId, { decision: 'reject', actor_role: actorRole, comment });
      set((s) => ({ byId: { ...s.byId, [processId]: { ...s.byId[processId], lastDecision: 'reject' } } }));
      emit({ type: 'approval:submit', payload: { processId, decision: 'reject' } } as any);
      emit({ type: 'api:success', key: 'processes:reject' });
    } catch (e: any) {
      emit({ type: 'api:error', key: 'processes:reject', error: String(e?.message || e) });
      throw e;
    }
  },
}));