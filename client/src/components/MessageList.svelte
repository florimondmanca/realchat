<script lang="ts">
  import type { Message } from "../models";
  import { socketService } from "../services/socket";
  import MessageItem from "./MessageItem.svelte";

  let messages: Message[] = [];

  socketService.onMessage((message) => {
    messages = [...messages, message];
  });
</script>

{#if messages.length > 0}
  <ul class="message-list">
    {#each messages as message (message.id)}
      <MessageItem {message} />
    {/each}
  </ul>
{:else}
  <div class="message-list empty">
    <small>No messages here yet.</small>
  </div>
{/if}
