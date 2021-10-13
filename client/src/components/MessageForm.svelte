<script lang="ts">
  import { user } from "../services/stores";
  import { webSocketService } from "../services/websocket";

  let body = "";

  const onSubmit = () => {
    if (!body) return;
    webSocketService.send($user, body);
    body = "";
  };
</script>

<form class="message-entry" on:submit|preventDefault={onSubmit}>
  <input
    type="text"
    name="body"
    bind:value={body}
    placeholder="Write a message..."
    autocomplete="off"
  />
  <button type="submit">Send</button>
</form>

<style>
  form {
    display: grid;
    grid-template-columns: 1fr auto;
    column-gap: 0.5em;
  }
  input {
    width: 100%;
  }
</style>
