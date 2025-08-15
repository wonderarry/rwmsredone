export function ThemeScript() {
  // Runs before hydration; sets html[data-theme] and color-scheme.
  const code = `
  (function() {
    try {
      var key = 'rwms-theme';
      var stored = localStorage.getItem(key);
      var systemDark = window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
      var theme = stored || (systemDark ? 'dark' : 'light');
      var root = document.documentElement;
      root.dataset.theme = theme;
      root.style.colorScheme = theme;
    } catch (e) {}
  })();`;
  return <script dangerouslySetInnerHTML={{ __html: code }} />;
}