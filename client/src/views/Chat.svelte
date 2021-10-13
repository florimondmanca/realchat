<script lang="ts">
  import { onDestroy, onMount } from "svelte";
  import type { ISendDetail } from "../models/message";
  import { user } from "../stores/auth";
  import { messages } from "../stores/messages";
  import { messageService } from "../services/messages";
  import MessageForm from "../components/MessageForm.svelte";
  import MessageList from "../components/MessageList.svelte";

  onMount(async () => {
    await messageService.init($user);
  });

  onDestroy(() => messageService.tearDown());

  const onSend = (event: CustomEvent<ISendDetail>) => {
    const { body } = event.detail;
    messageService.send({ userName: $user, body });
  };
</script>

<div>
  <MessageList messages={$messages} />
  <MessageForm on:send={onSend} />
</div>
