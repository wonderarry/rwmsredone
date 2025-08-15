import { api, Normalizers } from './client';
import type { API } from '@/types/api';

export async function createProcess(payload: API.CreateProcessReq) {
  const { data } = await api.post<API.CreateProcessRes>('/api/processes/', payload);
  return Normalizers.extractId(data); // { id: process_id }
}

export async function addProcessMember(processId: string, payload: API.ProcessMemberReq): Promise<void> {
  await api.post(`/api/processes/${processId}/members`, payload);
}

export async function removeProcessMember(processId: string, payload: API.ProcessMemberReq): Promise<void> {
  await api.delete(`/api/processes/${processId}/members`, { data: payload });
}

export async function submitApproval(processId: string, payload: API.SubmitApprovalReq): Promise<void> {
  await api.post(`/api/processes/${processId}/approvals`, payload);
}