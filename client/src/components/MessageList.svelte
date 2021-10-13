<script lang="ts">
  import type { Message } from "../models/message";
  import MessageItem from "./MessageItem.svelte";

  export let messages: Message[];
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
