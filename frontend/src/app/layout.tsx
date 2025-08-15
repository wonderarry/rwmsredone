import './globals.css';
import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import { UIToaster } from './ui-toaster';
import { ThemeScript } from './theme-script';
import { ThemeProvider } from '@/lib/providers/theme-provider';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
    title: 'RWMS',
    description: 'Research Workflow Management System',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return (
        <html lang="en" suppressHydrationWarning>
            <head>
                <ThemeScript />
            </head>
            <body>
                <ThemeProvider>
                    {children}
                    <UIToaster />
                </ThemeProvider>

            </body>
        </html>
    );
}