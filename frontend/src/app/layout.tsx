import './globals.css';
import type { Metadata } from 'next';
import { Inter } from 'next/font/google';
import { UIToaster } from './ui-toaster';

const inter = Inter({ subsets: ['latin'] });

export const metadata: Metadata = {
  title: 'RWMS',
  description: 'Research Workflow Management System',
};

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body className={inter.className + ' bg-gray-50'}>
        {children}
        {/* Global toasts from ui-store */}
        <UIToaster />
      </body>
    </html>
  );
}