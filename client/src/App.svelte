<script lang="ts">
  import { onMount } from "svelte";
  import AccountBar from "./components/AccountBar.svelte";
  import MessageForm from "./components/MessageForm.svelte";
  import MessageList from "./components/MessageList.svelte";
  import { socketService } from "./services/socket";

  onMount(() => {
    socketService.init();

    return () => socketService.close();
  });
</script>

<main>
  <h1>Chat room</h1>

  <AccountBar />

  <div class="chat">
    <MessageList />
    <MessageForm />
  </div>
</main>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>
