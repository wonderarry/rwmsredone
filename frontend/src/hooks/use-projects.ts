'use client';

import { useEffect } from 'react';
import { useProjectsStore } from '@/lib/stores/projects-store';
import { useToast } from './use-toast';

export function useProjects(autoLoad = true) {
  const { success, error } = useToast();
  const items = useProjectsStore((s) => s.items);
  const selectedId = useProjectsStore((s) => s.selectedId);
  const loading = useProjectsStore((s) => s.loading);
  const select = useProjectsStore((s) => s.select);
  const refresh = useProjectsStore((s) => s.refresh);
  const create = useProjectsStore((s) => s.create);
  const edit = useProjectsStore((s) => s.edit);

  useEffect(() => {
    if (autoLoad && !items.length && !loading) {
      refresh().catch((e) => error(e?.message || 'Failed to load projects'));
    }
  }, [autoLoad, items.length, loading, refresh, error]);

  return {
    items,
    selectedId,
    loading,
    select,
    refresh,
    create: async (input: { name: string; description: string; theme: string }) => {
      const id = await create(input);
      success('Project created');
      return id;
    },
    edit: async (id: string, patch: { name?: string; description?: string; theme?: string }) => {
      await edit(id, patch);
      success('Project updated');
    },
  };
}