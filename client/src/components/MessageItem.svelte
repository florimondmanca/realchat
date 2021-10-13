<script lang="ts">
  import type { Message } from "../models/message";
  import { user } from "../stores/auth";

  export let message: Message;
</script>

{#if message.type === "JOIN"}
  <li class="notification">
    <strong>{message.data.userName}</strong> joined.
  </li>
{:else if message.type === "LEAVE"}
  <li class="notification">
    <strong>{message.data.userName}</strong> left.
  </li>
{:else}
  <li class="message" class:message-self={message.data.userName === $user}>
    <strong class="message-user">
      {message.data.userName}
    </strong>
    <span class="message-body">
      {message.data.body}
    </span>
    <em class="message-time">
      <small>
        {new Date(message.timestampSeconds * 1000).toLocaleString()}
      </small>
    </em>
  </li>
{/if}

<style>
  .notification {
    /* TODO */
  }

  .message {
    /* TODO */
  }

  .message-self {
    /* TODO */
  }
</style>
