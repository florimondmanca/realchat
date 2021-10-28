<script lang="ts">
  import { isLoggedIn } from "./stores/auth";
  import NavBar from "./components/NavBar.svelte";
  import Footer from "./components/Footer.svelte";
  import Chat from "./views/Chat.svelte";
  import Login from "./views/Login.svelte";
</script>

<NavBar />

{#if $isLoggedIn}
  <Chat />
{:else}
  <Login />
{/if}

<!-- <Footer /> -->
