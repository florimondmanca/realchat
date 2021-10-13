<script lang="ts">
  import { user } from "./services/stores";
  import AccountBar from "./components/AccountBar.svelte";
  import Chat from "./views/Chat.svelte";
  import Login from "./views/Login.svelte";
</script>

<main>
  <h1>RealChat</h1>

  {#if $user == null}
    <Login />
  {:else}
    <AccountBar />
    <Chat />
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
