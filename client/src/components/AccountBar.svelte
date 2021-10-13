<script lang="ts">
  import { user } from "../services/stores";

  const logout = () => {
    user.set(null);
  };
</script>

<p>
  Logged in as <strong>{$user}</strong> (<a
    href="/"
    on:click|preventDefault={logout}>Not you?</a
  >)
</p>
