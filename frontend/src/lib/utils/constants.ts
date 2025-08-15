export const GLOBAL_ROLES = ['CanCreateProjects'] as const;
export type GlobalRole = typeof GLOBAL_ROLES[number];

export const PROJECT_ROLES = ['ProjectLeader', 'ProjectMember'] as const;
export type ProjectRole = typeof PROJECT_ROLES[number];

export const PROCESS_ROLES = ['Advisor', 'Student', 'Reviewer'] as const;
export type ProcessRole = typeof PROCESS_ROLES[number];

export const PROCESS_LIFECYCLE_STATES = ['active', 'completed', 'archived'] as const;
export type ProcessLifecycleState = typeof PROCESS_LIFECYCLE_STATES[number];

export const STORAGE_KEYS = {
  authToken: 'rwms.token',
};