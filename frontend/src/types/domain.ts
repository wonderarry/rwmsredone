export type GlobalRole = 'CanCreateProjects';
export type ProjectRole = 'ProjectLeader' | 'ProjectMember';
export type ProcessRole = 'Advisor' | 'Student' | 'Reviewer';

export interface Account {
  id: string;
  firstName: string;
  middleName?: string;
  lastName: string;
  groupNumber?: string;
  createdAt: Date;
  updatedAt: Date;
  // Optional: global roles for UI gating (if/when the API returns them)
  globalRoles?: GlobalRole[];
}

export interface ProjectBrief {
  id: string;
  name: string;
  description: string;
  theme: string;
}

// Process lifecycle per backend (active/completed/archived)
export type ProcessLifecycleState = 'active' | 'completed' | 'archived';

export interface ProcessBrief {
  id: string;
  name: string;
  currentStage: string; // e.g. "under-review" or a stage key
  state: ProcessLifecycleState;
  projectId: string;
  // Convenience (client-side only): last approval decision we sent
  lastDecision?: 'approve' | 'reject';
}

// Normalized create responses
export interface IdResult { id: string }
export interface TokenResult { token: string }