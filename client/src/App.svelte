<script lang="ts">
  import { isLoggedIn } from "./stores/auth";
  import AccountBar from "./components/AccountBar.svelte";
  import Chat from "./views/Chat.svelte";
  import Login from "./views/Login.svelte";
</script>

<main>
  <h1>RealChat</h1>

  {#if $isLoggedIn}
    <AccountBar />
    <Chat />
  {:else}
    <Login />
  {/if}
</main>

<style>
  main {
    display: flex;
    flex-flow: column;
    align-items: center;
    padding: 3em;
    margin: 0 auto;
  }
</style>
