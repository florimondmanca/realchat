<script lang="ts">
  import type { Message } from "../models/message";
  import { user } from "../stores/auth";

  export let message: Message;
</script>

<li class="message" class:message-self={message.userName === $user}>
  <div class="message-header">
    <strong class="message-user">
      {message.userName}
    </strong>
    <span class="message-time">
      <small>
        {new Date(message.timestampSeconds * 1000).toLocaleString()}
      </small>
    </span>
  </div>
  <div class="message-body">
    {message.body}
  </div>
</li>

<style>
  :root {
    --br: 0.6em;
  }

  .message {
    background-color: var(--color-surface);
    border-radius: var(--br) var(--br) var(--br) 0;
    padding: 0.6em 1em;
    width: fit-content;
    width: -moz-fit-content;
    margin-top: 0.5em;
    margin-bottom: 0.5em;
    margin-right: auto;
  }

  .message-header {
    margin-bottom: 0.5em;
  }

  .message-time {
    font-size: small;
    font-style: italic;
    white-space: nowrap;
    float: right;
    margin-left: 2em;
  }

  .message-self {
    background-color: var(--color-primary);
    color: var(--text-on-primary);
    border-radius: var(--br) var(--br) 0 var(--br);
    margin-left: auto;
    margin-right: 0;
  }
</style>
