import axios, { AxiosError, AxiosInstance } from 'axios';
import type { IdResult, TokenResult } from '@/types/domain';

export class ApiError extends Error {
  status?: number;
  details?: unknown;
  constructor(message: string, status?: number, details?: unknown) {
    super(message);
    this.name = 'ApiError';
    this.status = status;
    this.details = details;
  }
}

/**
 * The single Axios instance used across the app.
 * baseURL is driven by NEXT_PUBLIC_API_BASE_URL.
 */
export const api: AxiosInstance = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_BASE_URL || '/',
  timeout: 15_000,
});

let authToken: string | null = null;
let onUnauthorized: (() => void) | null = null;

export function setAuthToken(token: string | null) {
  authToken = token;
}

export function setUnauthorizedHandler(handler: (() => void) | null) {
  onUnauthorized = handler;
}

// Attach Authorization header when present
api.interceptors.request.use((config) => {
  if (authToken) {
    config.headers = config.headers || {};
    (config.headers as Record<string, string>)['Authorization'] = `Bearer ${authToken}`;
  }
  return config;
});

// Uniform error surface
api.interceptors.response.use(
  (res) => res,
  (error: AxiosError<any>) => {
    const status = error.response?.status;
    if (status === 401 && onUnauthorized) onUnauthorized();
    const message =
      (error.response?.data && (error.response.data.message || error.response.data.error)) ||
      error.message ||
      'Network error';
    return Promise.reject(new ApiError(message, status, error.response?.data));
  }
);

function extractToken(obj: Record<string, string>): TokenResult {
  const entry = Object.entries(obj).find(([k]) => k.toLowerCase().includes('token'));
  if (!entry) throw new ApiError('Token not found in response', 500, obj);
  return { token: entry[1] };
}

function extractId(obj: Record<string, string>): IdResult {
  const entry = Object.entries(obj).find(([k]) => k.endsWith('_id') || k.toLowerCase() === 'id');
  if (!entry) throw new ApiError('ID not found in response', 500, obj);
  return { id: entry[1] };
}

export const Normalizers = { extractToken, extractId };