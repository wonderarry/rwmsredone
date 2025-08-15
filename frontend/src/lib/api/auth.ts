import { api, Normalizers } from './client';
import type { API } from '@/types/api';
import type { Account } from '@/types/domain';

export async function loginLocal(payload: API.LoginLocalReq) {
  const { data } = await api.post<API.LoginLocalRes>('/api/auth/login-local', payload);
  return Normalizers.extractToken(data);
}

export async function registerLocal(payload: API.RegisterLocalReq) {
  const { data } = await api.post<API.RegisterLocalRes>('/api/auth/register-local', payload);
  return Normalizers.extractId(data); // returns { id: accountId }
}

export async function getMe(): Promise<Account> {
  const { data } = await api.get<API.AccountDTO>('/api/accounts/me');
  return {
    ...data,
    createdAt: new Date(data.createdAt),
    updatedAt: new Date(data.updatedAt),
  };
}