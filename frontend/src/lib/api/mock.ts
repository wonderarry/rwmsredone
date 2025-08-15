import type { API } from '@/types/api';
import type { Account, IdResult, ProjectBrief, TokenResult } from '@/types/domain';

const wait = (ms: number) => new Promise((r) => setTimeout(r, ms));

let mockToken = 'mock.jwt.token';
let mockAccount: Account = {
  id: 'acc_001',
  firstName: 'Sarah',
  lastName: 'Chen',
  createdAt: new Date().toISOString() as unknown as Date, // acceptable for mock
  updatedAt: new Date().toISOString() as unknown as Date,
};

let mockProjects: ProjectBrief[] = [
  {
    id: 'prj_001',
    name: 'Climate Change Impact Study',
    description: 'Analyzing the effects of climate change on coastal ecosystems',
    theme: 'Environmental Science',
  },
  {
    id: 'prj_002',
    name: 'Neural Network Architecture',
    description: 'Developing new approaches to deep learning optimization',
    theme: 'Computer Science',
  },
];

export async function loginLocal(_: API.LoginLocalReq): Promise<TokenResult> {
  await wait(400);
  return { token: mockToken };
}

export async function registerLocal(_: API.RegisterLocalReq): Promise<IdResult> {
  await wait(400);
  return { id: 'acc_new' };
}

export async function getMe(): Promise<Account> {
  await wait(300);
  return { ...mockAccount, createdAt: new Date(), updatedAt: new Date() };
}

export async function listProjects(): Promise<ProjectBrief[]> {
  await wait(350);
  return mockProjects;
}

export async function createProject(payload: API.CreateProjectReq): Promise<IdResult> {
  await wait(500);
  const id = `prj_${Math.random().toString(36).slice(2, 8)}`;
  mockProjects = [
    { id, name: payload.name, description: payload.description, theme: payload.theme },
    ...mockProjects,
  ];
  return { id };
}

export async function editProject(id: string, payload: API.EditProjectReq): Promise<void> {
  await wait(300);
  mockProjects = mockProjects.map((p) => (p.id === id ? { ...p, ...payload } : p));
}

export async function addProjectMember(): Promise<void> { await wait(200); }
export async function removeProjectMember(): Promise<void> { await wait(200); }
export async function createProcess(_: API.CreateProcessReq): Promise<IdResult> { await wait(400); return { id: 'proc_001' }; }
export async function addProcessMember(): Promise<void> { await wait(200); }
export async function removeProcessMember(): Promise<void> { await wait(200); }
export async function submitApproval(): Promise<void> { await wait(250); }
