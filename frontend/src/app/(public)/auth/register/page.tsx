'use client';
import * as React from 'react';
import { useRouter } from 'next/navigation';
import { Button } from '@/components/ui/button';
import { useToast } from '@/hooks/use-toast';
import { registerLocal } from '@/lib/api/auth'; // calls /api/auth/register-local

export default function RegisterPage() {
  const [firstName, setFirst] = React.useState('');
  const [lastName, setLast] = React.useState('');
  const [login, setLogin] = React.useState('');
  const [password, setPassword] = React.useState('');
  const [loading, setLoading] = React.useState(false);
  const router = useRouter();
  const toast = useToast();

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setLoading(true);
    try {
      await registerLocal({
        first_name: firstName,
        last_name: lastName,
        login,
        password,
        grant_can_create: false,
      });
      toast.success('Account created. Please sign in.');
      router.replace('/auth/login');
    } catch {
      toast.error('Registration failed');
    } finally {
      setLoading(false);
    }
  }

  return (
    <>
      <h1 className="mb-6 text-center text-2xl font-semibold">Create account</h1>
      <form className="space-y-4" onSubmit={onSubmit}>
        <div className="grid grid-cols-2 gap-3">
          <div>
            <label className="mb-2 block text-sm text-[hsl(var(--muted))]">First name</label>
            <input className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 focus:outline-none focus:ring-2 focus:ring-[hsl(var(--brand-600))]"
                   value={firstName} onChange={(e) => setFirst(e.target.value)} />
          </div>
          <div>
            <label className="mb-2 block text-sm text-[hsl(var(--muted))]">Last name</label>
            <input className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 focus:outline-none focus:ring-2 focus:ring-[hsl(var(--brand-600))]"
                   value={lastName} onChange={(e) => setLast(e.target.value)} />
          </div>
        </div>
        <div>
          <label className="mb-2 block text-sm text-[hsl(var(--muted))]">Username</label>
          <input className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 focus:outline-none focus:ring-2 focus:ring-[hsl(var(--brand-600))]"
                 value={login} onChange={(e) => setLogin(e.target.value)} />
        </div>
        <div>
          <label className="mb-2 block text-sm text-[hsl(var(--muted))]">Password</label>
          <input type="password" className="w-full radii-md border border-base bg-[hsl(var(--surface))] px-3 py-2 focus:outline-none focus:ring-2 focus:ring-[hsl(var(--brand-600))]"
                 value={password} onChange={(e) => setPassword(e.target.value)} />
        </div>
        <Button type="submit" disabled={loading} className="w-full">
          {loading ? 'Creating…' : 'Create account'}
        </Button>
      </form>

      <p className="mt-4 text-center text-sm text-[hsl(var(--muted))]">
        Already have an account? <a href="/auth/login" className="text-[hsl(var(--brand-700))] hover:underline">Sign in</a>
      </p>
    </>
  );
}