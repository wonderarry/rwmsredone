import { cookies } from 'next/headers';
import { redirect } from 'next/navigation';
import { Sidebar } from '@/components/layout/sidebar';
import { Topbar } from '@/components/layout/topbar';
import {RWMS_LOGIN_TOKEN_KEY} from "@/app/page"
export default async function AppShellLayout({ children }: { children: React.ReactNode }) {
  const c = await cookies()
  const token = c.get(RWMS_LOGIN_TOKEN_KEY)?.value;
  if (!token) redirect('/auth/login');

  return (
    <div className="flex h-screen bg-[hsl(var(--bg))] text-[hsl(var(--fg))]">
      <Sidebar
        projects={[]}
        currentView="home"
        selectedProjectId={null}
        selectedProcessId={null}
        onGoHome={() => {}}
        onSelectProject={() => {}}
        onSelectProcess={() => {}}
        onNewProject={() => {}}
      />
      <div className="flex flex-1 flex-col overflow-hidden">
        <Topbar
          userName="Demo User"
          onOpenProfile={() => {}}
          onOpenSettings={() => {}}
          onSignOut={() => {}}
        />
        <main className="flex-1 overflow-auto">{children}</main>
      </div>
    </div>
  );
}