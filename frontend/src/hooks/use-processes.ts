'use client';

import { useProcessesStore } from '@/lib/stores/processes-store';
import { useToast } from './use-toast';

export function useProcesses() {
  const { success, error } = useToast();
  const byId = useProcessesStore((s) => s.byId);
  const selectedId = useProcessesStore((s) => s.selectedId);
  const select = useProcessesStore((s) => s.select);
  const create = useProcessesStore((s) => s.create);
  const approve = useProcessesStore((s) => s.approve);
  const reject = useProcessesStore((s) => s.reject);

  return {
    processes: byId,
    selectedId,
    select,
    create: async (input: { name: string; projectId: string; templateKey: string }) => {
      const id = await create(input);
      success('Process started');
      return id;
    },
    approve: async (processId: string, actorRole: string, comment?: string) => {
      try {
        await approve(processId, actorRole, comment);
        success('Approved');
      } catch (e: any) {
        error(e?.message || 'Approval failed');
        throw e;
      }
    },
    reject: async (processId: string, actorRole: string, comment?: string) => {
      try {
        await reject(processId, actorRole, comment);
        success('Rejected');
      } catch (e: any) {
        error(e?.message || 'Rejection failed');
        throw e;
      }
    },
  };
}
