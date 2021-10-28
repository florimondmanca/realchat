<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import type { ISendDetail } from "../models/message";
  import { user } from "../stores/auth";
  import { messages } from "../stores/messages";
  import { messageService } from "../services/messages";
  import MessageForm from "../components/MessageForm.svelte";
  import MessageList from "../components/MessageList.svelte";

  let messageListContainer: Element;

  const scrollChat = () => {
    messageListContainer.scrollTo({
      top: messageListContainer.scrollHeight,
    });
  };

  onMount(async () => {
    await messageService.init($user);
    return messages.subscribe(() => {
      // HACK: Let new message render first.
      setTimeout(() => scrollChat(), 10);
    });
  });

  onDestroy(() => messageService.tearDown());

  const onSend = (event: CustomEvent<ISendDetail>) => {
    const { body } = event.detail;
    messageService.send({ userName: $user, body });
  };
</script>

<div class="box">
  <div class="history" bind:this={messageListContainer}>
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
    padding: 0 3em;
    margin: 0 auto;
  }

  .history {
    position: relative;
    height: 70vh;
    overflow-y: scroll;
  }

  .form {
    position: sticky;
    bottom: 0;
    display: inline-block;
  }
</style>
