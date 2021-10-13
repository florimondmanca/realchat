<script lang="ts">
  import { user, logout } from "../stores/auth";
</script>

<p>
  Logged in as <strong>{$user}</strong> (<a
    href="/"
    on:click|preventDefault={logout}>Not you?</a
  >)
</p>
