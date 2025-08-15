// Domain models are the shape the app uses internally (camelCase, friendly enums).

export type UserRole = 'ProjectLeader' | 'Advisor' | 'Reviewer' | 'Viewer';

export interface Account {
  id: string;
  firstName: string;
  middleName?: string;
  lastName: string;
  groupNumber?: string;
  createdAt: Date;
  updatedAt: Date;
}

export interface ProjectBrief {
  id: string;
  name: string;
  description: string;
  theme: string;
}

export type ProcessState = 'draft' | 'pending' | 'approved' | 'rejected';

export interface ProcessBrief {
  id: string;
  name: string;
  stage: string; 
  state: ProcessState;
  projectId: string;
}

export interface IdResult { id: string }
export interface TokenResult { token: string }