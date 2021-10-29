<script>
  import { users } from "../stores/chat";

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
