<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { ISendDetail } from "../models/message";

  let body = "";

  const dispatch = createEventDispatcher<{ send: ISendDetail }>();

  const onSubmit = () => {
    if (!body) return;
    dispatch("send", { body });
    body = "";
  };
</script>

<form on:submit|preventDefault={onSubmit}>
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
