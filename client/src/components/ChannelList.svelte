<script>
  import List from "./List.svelte";
  import ListItem from "./ListItem.svelte";

  const channels = ["general", "random"];
  const selected = "general";
</script>

<List>
  {#each channels as channel}
    <ListItem tag="button" selected={channel === selected}>
      #&nbsp;{channel}
    </ListItem>
  {/each}
</List>
