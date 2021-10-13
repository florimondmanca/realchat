<script lang="ts">
  import { onMount } from "svelte";
  import MessageForm from "../components/MessageForm.svelte";
  import MessageList from "../components/MessageList.svelte";
  import { webSocketService } from "../services/websocket";

  onMount(() => {
    webSocketService.init();
    return () => webSocketService.close();
  });
</script>

<div>
  <MessageList />
  <MessageForm />
</div>
