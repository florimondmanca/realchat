import { derived, writable } from "svelte/store";

const storedTheme =
  localStorage.getItem("theme") ||
  (window.matchMedia("(prefers-color-scheme: dark)").matches
    ? "dark"
    : "light");

export const theme = writable<string>(storedTheme);

export const inactiveTheme = derived(
  theme,
  ($theme, set: (value: string) => void) =>
    set($theme === "light" ? "dark" : "light")
);

theme.subscribe((value) => {
  document.documentElement.setAttribute("data-theme", value);
  localStorage.setItem("theme", value);
});
