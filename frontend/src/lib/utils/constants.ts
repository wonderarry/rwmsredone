export const ROLES = ['ProjectLeader', 'Advisor', 'Reviewer', 'Viewer'] as const;
export type Role = typeof ROLES[number];

export const PROCESS_STATES = ['draft', 'pending', 'approved', 'rejected'] as const;
export type ProcessState = typeof PROCESS_STATES[number];

export const STORAGE_KEYS = {
  authToken: 'rwms.token',
};
