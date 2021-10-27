<script lang="ts">
  import { isLoggedIn, logout, user } from "../stores/auth";
</script>

<nav>
  <div class="brand">RealChat</div>
  <div class="tools">
    {#if $isLoggedIn}
      <div class="dimmed">
        {$user}
      </div>
      <span class="sep dimmed">|</span>
      <a href="/" on:click|preventDefault={logout}> Logout </a>
    {/if}
  </div>
</nav>

<style>
  nav {
    display: flex;
    justify-content: space-between;
    padding: 1em;
    background-color: var(--color-primary);
    color: var(--text-on-primary);
  }

  .dimmed {
    opacity: 0.7;
  }

  .tools {
    display: flex;
  }

  .tools a {
    color: var(--link-on-primary);
  }

  .sep {
    margin: 0 0.5em;
  }

  .brand {
    font-weight: bold;
  }
</style>
