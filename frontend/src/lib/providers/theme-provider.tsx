'use client';

import * as React from 'react';

export type Theme = 'light' | 'dark';
type ThemeCtx = { theme: Theme; toggle: () => void; setTheme: (t: Theme) => void };

const ThemeContext = React.createContext<ThemeCtx | null>(null);
export const STORAGE_THEME_KEY = 'rwms-theme';

export function applyTheme(theme: Theme) {
    const root = document.documentElement;
    root.dataset.theme = theme;
    // Helps native UI match (scrollbars, form controls)
    root.style.colorScheme = theme;
    try {
      localStorage.setItem(STORAGE_THEME_KEY, theme);
    } catch { /* ignore */ }

}

export function themeGetter()  {
        // This initial value is unused on first paint because ThemeScript already set DOM,
        // but we still initialize the state sensibly for hydration.
        if (typeof window === 'undefined') return 'light';
        const stored = localStorage.getItem(STORAGE_THEME_KEY) as Theme | null;
        if (stored === 'light' || stored === 'dark') return stored;
        const prefersDark =
            typeof window !== 'undefined' &&
            window.matchMedia &&
            window.matchMedia('(prefers-color-scheme: dark)').matches;
        return prefersDark ? 'dark' : 'light';
    }

export function ThemeProvider({ children }: { children: React.ReactNode }) {
    const [theme, setThemeState] = React.useState<Theme>(themeGetter);

    React.useEffect(() => {
        applyTheme(theme);
        try { localStorage.setItem(STORAGE_THEME_KEY, theme); } catch { }
    }, [theme]);

    // Cross-tab sync
    React.useEffect(() => {
        const onStorage = (e: StorageEvent) => {
            if (e.key === STORAGE_THEME_KEY && (e.newValue === 'light' || e.newValue === 'dark')) {
                setThemeState(e.newValue);
            }
        };
        window.addEventListener('storage', onStorage);
        return () => window.removeEventListener('storage', onStorage);
    }, []);

    const setTheme = React.useCallback((t: Theme) => setThemeState(t), []);
    const toggle = React.useCallback(() => {
        setThemeState((prev) => (prev === 'dark' ? 'light' : 'dark'));
    }, []);

    const value = React.useMemo(() => ({ theme, toggle, setTheme }), [theme, toggle, setTheme]);

    return (<ThemeContext.Provider value= { value } >
        { children }
        </ThemeContext.Provider>);
}

export function useTheme() {
    const ctx = React.useContext(ThemeContext);
    if (!ctx) throw new Error('useTheme must be used within ThemeProvider');
    return ctx;
}
