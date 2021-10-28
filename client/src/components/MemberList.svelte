<script>
  import { users } from "../stores/messages";

  import List from "./List.svelte";
  import ListItem from "./ListItem.svelte";
</script>

<List>
  {#each $users as user}
    <ListItem>
      {user}
    </ListItem>
  {/each}
</List>
