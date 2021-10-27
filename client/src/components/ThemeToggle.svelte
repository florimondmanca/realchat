<script lang="ts">
  import { inactiveTheme, theme } from "../stores/theme";

  import DarkThemeIcon from "../icons/DarkThemeIcon.svelte";
  import LightThemeIcon from "../icons/LightThemeIcon.svelte";
  import Toggle from "./Toggle.svelte";
</script>

<Toggle
  active={$theme === "light"}
  aria-label="Enable {$inactiveTheme} theme"
  on:click={() => theme.set($inactiveTheme)}
>
  <DarkThemeIcon />
  <LightThemeIcon />
</Toggle>
