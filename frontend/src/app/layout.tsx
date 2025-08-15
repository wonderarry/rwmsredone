import './globals.css';
import { ThemeProvider } from '@/lib/providers/theme-provider';
import ThemeScript from '@/app/theme-script'; // your existing script component
import { UIToaster } from '@/components/ui/ui-toaster';

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head><ThemeScript /></head>
      <body>
        <ThemeProvider>
          {children}
          <UIToaster />
        </ThemeProvider>
      </body>
    </html>
  );
}