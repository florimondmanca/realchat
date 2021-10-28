<script lang="ts">
  import { signIn } from "../stores/auth";

  let user = "";
  let isValid = false;
  $: isValid = !!user;
</script>

<main>
  <h1>Log in</h1>

  <hr />

  <form on:submit|preventDefault={() => signIn(user)}>
    <div class="form-field">
      <label for="user">Username</label>
      <input
        type="text"
        name="user"
        bind:value={user}
        placeholder="Enter your username..."
        required
      />
    </div>
    <button type="submit" disabled={!isValid}>Enter chat</button>
  </form>
</main>

<style>
  main {
    max-width: 20em;
    margin: 4em auto;
  }
</style>
