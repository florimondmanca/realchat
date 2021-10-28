import { derived, writable } from "svelte/store";

const storedTheme =
  localStorage.getItem("theme") ||
  (window.matchMedia("(prefers-color-scheme: dark)").matches
    ? "dark"
    : "light");

export const theme = writable<string>(storedTheme);

theme.subscribe((value) => {
  document.documentElement.setAttribute("data-theme", value);
  localStorage.setItem("theme", value);
});

export const inactiveTheme = derived<typeof theme, string>(
  theme,
  ($theme, set) => set($theme === "light" ? "dark" : "light")
);
