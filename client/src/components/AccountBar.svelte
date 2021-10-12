<script lang="ts">
  const user = "John Doe"; // TODO: store

  const logout = (event: Event) => {
    console.log(`logout: ${event}`);
  };
</script>

<p>
  Logged in as <strong>{user}</strong> (<a
    href="/"
    on:click|preventDefault={logout}>Not you?</a
  >)
</p>
