'use client';

import { create } from 'zustand';
import type { ProjectBrief } from '@/types/domain';
import * as Projects from '@/lib/api/projects';
import { emit } from '@/lib/events/bus';

interface ProjectsState {
  items: ProjectBrief[];
  selectedId: string | null;
  loading: boolean;
  error?: string;

  // actions
  select(id: string | null): void;
  refresh(): Promise<void>;
  create(input: { name: string; description: string; theme: string }): Promise<string>; // returns id
  edit(id: string, patch: Partial<Pick<ProjectBrief, 'name' | 'description' | 'theme'>>): Promise<void>;
}

export const useProjectsStore = create<ProjectsState>((set, get) => ({
  items: [],
  selectedId: null,
  loading: false,

  select(id) {
    set({ selectedId: id });
  },

  async refresh() {
    set({ loading: true, error: undefined });
    emit({ type: 'api:start', key: 'projects:list' });
    try {
      const items = await Projects.listProjects();
      set({ items, loading: false });
      emit({ type: 'api:success', key: 'projects:list' });
    } catch (e: any) {
      set({ loading: false, error: e?.message || 'Failed to load projects' });
      emit({ type: 'api:error', key: 'projects:list', error: String(e?.message || e) });
      throw e;
    }
  },

  async create(input) {
    emit({ type: 'api:start', key: 'projects:create' });
    try {
      const { id } = await Projects.createProject(input);
      // optimistic prepend
      set((s) => ({ items: [{ id, ...input }, ...s.items] }));
      emit({ type: 'data:project:create' as any, payload: { name: input.name } });
      emit({ type: 'api:success', key: 'projects:create' });
      return id;
    } catch (e: any) {
      emit({ type: 'api:error', key: 'projects:create', error: String(e?.message || e) });
      throw e;
    }
  },

  async edit(id, patch) {
    emit({ type: 'api:start', key: 'projects:edit' });
    const prev = get().items;
    // optimistic
    set({ items: prev.map((p) => (p.id === id ? { ...p, ...patch } : p)) });
    try {
      await Projects.editProject(id, patch as any);
      emit({ type: 'api:success', key: 'projects:edit' });
    } catch (e: any) {
      // rollback
      set({ items: prev });
      emit({ type: 'api:error', key: 'projects:edit', error: String(e?.message || e) });
      throw e;
    }
  },
}));