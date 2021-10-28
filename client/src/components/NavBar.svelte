<script lang="ts">
  import { isLoggedIn, logout, user } from "../stores/auth";
  import ThemeToggle from "./ThemeToggle.svelte";
</script>

<nav>
  <div class="brand">RealChat</div>
  <div class="tools">
    {#if $isLoggedIn}
      {$user}
      <span class="sep">|</span>
      <a href="/" on:click|preventDefault={logout}> Logout </a>
    {/if}
    <ThemeToggle />
  </div>
</nav>

<style>
  nav {
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 1em;
    background-color: var(--color-primary);
    color: var(--text-on-primary);
  }

  .brand {
    font-weight: bold;
  }

  .tools {
    display: flex;
    align-items: center;
  }

  .tools a {
    color: var(--link-on-primary);
    margin-right: 2em;
  }

  .sep {
    margin: 0 0.5em;
  }
</style>
