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
    <div>
      <strong class="message-user">
        {message.data.userName}
      </strong>
    </div>
    <span class="message-body">
      {message.data.body}
    </span>
    <div class="message-time">
      <small>
        {new Date(message.timestampSeconds * 1000).toLocaleString()}
      </small>
    </div>
  </li>
{/if}

<style>
  .notification {
    filter: contrast(20%);
    text-align: center;
    font-style: italic;
  }

  .message {
    background-color: var(--color-surface);
    border-radius: 0.4em;
    padding: 0.6em 1em;
    width: fit-content;
    width: -moz-fit-content;
    margin-top: 0.5em;
    margin-bottom: 0.5em;
    margin-right: auto;
  }

  .message-time {
    font-style: italic;
    float: right;
    margin-left: 2em;
    margin-bottom: -2em;
  }

  .message-self {
    background-color: var(--color-primary);
    margin-left: auto;
    margin-right: 0;
  }
</style>
