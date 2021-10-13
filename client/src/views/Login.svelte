<script lang="ts">
  import { user } from "../services/stores";

  let userValue = "";

  const onSubmit = () => {
    user.set(userValue);
  };
</script>

<h3>Welcome!</h3>
<p>How should your chatmates call you?</p>

<form on:submit|preventDefault={onSubmit}>
  <label for="user">Username</label>
  <input
    type="text"
    name="user"
    bind:value={userValue}
    placeholder="Enter your username..."
    required
  />
  <button type="submit">Enter chat</button>
</form>
