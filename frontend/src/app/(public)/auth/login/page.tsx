
'use client';
import * as React from 'react';
import { useRouter, useSearchParams } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { useToast } from '@/hooks/use-toast';
import { loginLocal } from '@/lib/api/auth'; // calls /api/auth/login-local

export default function LoginPage() {
  const [login, setLogin] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [loading, setLoading] = React.useState(false);
  const router = useRouter();
  const params = useSearchParams();
  const next = params.get('next') || '/dashboard';
  const toast = useToast();

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);
    try {
      const { token } = await loginLocal({ login, password });
      // Set HttpOnly cookie via helper API route (demo)
      await fetch('/api/session', { method: 'POST', body: JSON.stringify({ token }) });
      toast.success('Welcome back!');
      router.replace(next);
    } catch {
      toast.error('Invalid credentials');
    } finally {
      setLoading(false);
    }
  }

  return (
    <>
      <h1 className="mb-6 text-center text-2xl font-semibold">Sign in</h1>
      <form className="space-y-4" onSubmit={onSubmit}>
        <div>
          <label className="mb-2 block text-sm text-[hsl(var(--muted))]">Username</label>
          <input
            autoFocus
            value={login}
            onChange={(e) => setLogin(e.target.value)}
            className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 focus:outline-none focus:ring-2 focus:ring-[hsl(var(--brand-600))]"
          />
        </div>
        <div>
          <label className="mb-2 block text-sm text-[hsl(var(--muted))]">Password</label>
          <input
            type="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 focus:outline-none focus:ring-2 focus:ring-[hsl(var(--brand-600))]"
          />
        </div>
        <Button type="submit" disabled={loading} className="w-full">
          {loading ? 'Signing in…' : 'Sign in'}
        </Button>
      </form>

      <p className="mt-4 text-center text-sm text-[hsl(var(--muted))]">
        No account? <a href="/auth/register" className="text-[hsl(var(--brand-700))] hover:underline">Register</a>
      </p>
    </>
  );
}
