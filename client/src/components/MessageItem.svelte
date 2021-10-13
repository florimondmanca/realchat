<script lang="ts">
  import { Action } from "../models/action";
  import type { Message } from "../models/message";

  export let message: Message;
  const user = "John Doe"; // TODO: store
</script>

{#if message.action === Action.JOIN}
  <li class="notification">
    <strong>{message.userName}</strong> joined.
  </li>
{:else if message.action === Action.LEAVE}
  <li class="notification">
    <strong>{message.userName}</strong> left.
  </li>
{:else}
  <li class="message" class:message-self={message.userName === user}>
    <strong class="message-user">
      {message.userName}
    </strong>
    <span class="message-body">
      {message.body}
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
