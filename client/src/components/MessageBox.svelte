<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import type { ISendDetail } from "../models/message";
  import { user } from "../stores/auth";
  import { messages } from "../stores/chat";
  import { messageService } from "../services/messages";
  import MessageForm from "../components/MessageForm.svelte";
  import MessageList from "../components/MessageList.svelte";

  let historyEl: Element = null;

  const scrollChat = () => {
    if (!historyEl) {
      return;
    }
    historyEl.scrollTo({ top: historyEl.scrollHeight });
  };

  onMount(() => {
    messageService.init($user);
    return messages.subscribe(() => {
      // HACK: Let new message render first.
      setTimeout(() => scrollChat(), 0);
    });
  });

  onDestroy(() => messageService.tearDown());

  const onSend = (event: CustomEvent<ISendDetail>) => {
    const { body } = event.detail;
    messageService.send(body);
  };
</script>

<div class="box">
  <div class="history" bind:this={historyEl}>
    <MessageList messages={$messages} />
  </div>
  <div class="form">
    <MessageForm on:send={onSend} />
  </div>
</div>

<style>
  .box {
    display: flex;
    flex-flow: column;
  }

  .history {
    position: relative;
    height: 70vh;
    overflow-y: scroll;
    padding: 0 1.5em; /* Leave space for handlebar */
  }

  .form {
    position: sticky;
    bottom: 0;
    display: inline-block;
    padding-top: 0.5em;
    border-top: 1px solid var(--color-surface);
  }
</style>
