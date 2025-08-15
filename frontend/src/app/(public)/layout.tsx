export default function PublicLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-screen bg-[hsl(var(--bg))] text-[hsl(var(--fg))] grid place-items-center p-6">
      <div className="w-full max-w-md radii-lg border border-base bg-[hsl(var(--surface))] p-6 shadow-soft">
        {children}
      </div>
    </div>
  );
}