<script lang="ts">
  import { inactiveTheme, theme } from "../stores/theme";
  import DarkThemeIcon from "../icons/DarkThemeIcon.svelte";
  import LightThemeIcon from "../icons/LightThemeIcon.svelte";
  import Toggle from "./Toggle.svelte";
</script>

<Toggle
  class="theme-toggle"
  enabled={$theme === "dark"}
  aria-label="Enable {$inactiveTheme} theme"
  on:click={() => theme.set($inactiveTheme)}
>
  <DarkThemeIcon slot="icon-off" />
  <LightThemeIcon slot="icon-on" />
</Toggle>

<style>
  :global(.theme-toggle) {
    color: #e99e14;
  }
</style>
