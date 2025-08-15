export type UserActionEvent =
  | { type: 'click'; target: string }
  | { type: 'submit'; form: string }
  | { type: 'navigate'; to: string };

export type DataMutationEvent =
  | { type: 'project:create'; payload: { name: string } }
  | { type: 'project:update'; payload: { id: string } }
  | { type: 'process:create'; payload: { projectId: string } }
  | { type: 'approval:submit'; payload: { processId: string; decision: 'approve' | 'reject' } };

export type UIStateEvent =
  | { type: 'modal:open'; key: string }
  | { type: 'modal:close'; key: string }
  | { type: 'toast:show'; level: 'success' | 'error' | 'info'; message: string };

export type ApiEvent =
  | { type: 'api:start'; key: string }
  | { type: 'api:success'; key: string }
  | { type: 'api:error'; key: string; error: string };