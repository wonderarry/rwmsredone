import { api, Normalizers } from './client';
import type { API } from '@/types/api';
import type { ProjectBrief } from '@/types/domain';

export async function listProjects(): Promise<ProjectBrief[]> {
  const { data } = await api.get<API.ProjectBriefDTO[]>('/api/projects/');
  // DTO already matches domain shape here (camelCase); return as-is.
  return data.map((p) => ({ ...p }));
}

export async function createProject(payload: API.CreateProjectReq) {
  const { data } = await api.post<API.CreateProjectRes>('/api/projects/', payload);
  return Normalizers.extractId(data); // { id: project_id }
}

export async function editProject(id: string, payload: API.EditProjectReq): Promise<void> {
  await api.put(`/api/projects/${id}`, payload);
}

export async function addProjectMember(projectId: string, payload: API.ProjectMemberReq): Promise<void> {
  await api.post(`/api/projects/${projectId}/members`, payload);
}

export async function removeProjectMember(projectId: string, payload: API.ProjectMemberReq): Promise<void> {
  await api.delete(`/api/projects/${projectId}/members`, { data: payload });
}