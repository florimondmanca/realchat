<script lang="ts">
  import { signIn } from "../stores/auth";

  let user = "";
</script>

<main>
  <h1>RealChat</h1>

  <h3>Welcome!</h3>
  <p>How should your chatmates call you?</p>

  <form on:submit|preventDefault={() => signIn(user)}>
    <label for="user">Username</label>
    <input
      type="text"
      name="user"
      bind:value={user}
      placeholder="Enter your username..."
      required
    />
    <button type="submit">Enter chat</button>
  </form>
</main>

<style>
  main {
    display: flex;
    flex-flow: column;
    align-items: center;
  }
</style>
