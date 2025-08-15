export namespace API {
  // ---------- Accounts
  export interface LoginLocalReq {
    login: string;
    password: string;
  }

  export interface RegisterLocalReq {
    first_name?: string;
    grant_can_create?: boolean;
    group_number?: string;
    last_name?: string;
    login: string;
    middle_name?: string;
    password: string;
  }

  // The spec returns an object with arbitrary string properties for login/register
  // responses. We'll normalize these to strongly typed results in our API layer.
  export type LoginLocalRes = Record<string, string>; // e.g. { token: string }
  export type RegisterLocalRes = Record<string, string>; // e.g. { account_id: string }

  // ---------- Accounts: Me
  // Spec shows camelCase in domain.Account
  export interface AccountDTO {
    id: string;
    firstName: string;
    middleName?: string;
    lastName: string;
    groupNumber?: string;
    createdAt: string; // ISO-8601 UTC
    updatedAt: string; // ISO-8601 UTC
  }

  // ---------- Projects
  export interface ProjectBriefDTO {
    id: string;
    name: string;
    description: string;
    theme: string;
  }

  export interface CreateProjectReq {
    name: string;
    description: string;
    theme: string;
  }

  export interface EditProjectReq {
    name?: string;
    description?: string;
    theme?: string;
  }

  export interface ProjectMemberReq {
    account_id: string;
    role: string;
  }

  // Responses for create endpoints come back as { "project_id": "..." } etc.
  export type CreateProjectRes = Record<string, string>; // expect key project_id

  // ---------- Processes
  export interface CreateProcessReq {
    name: string;
    project_id: string;
    template_key: string;
  }

  export interface ProcessMemberReq {
    account_id: string;
    role: string;
  }

  export type CreateProcessRes = Record<string, string>; // expect key process_id

  export type SubmitApprovalReq = {
    decision: 'approve' | 'reject';
    actor_role: string;
    comment?: string;
  };
}
