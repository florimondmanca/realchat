<script lang="ts">
  import { socketService } from "../services/socket";

  const user = "John"; // TODO: get from store
  let body = "";

  const onSubmit = () => {
    if (!body) return;
    socketService.send(user, body);
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
</form>
