<script lang="ts">
  import { flip } from "svelte/animate";

  export let enabled: boolean;
</script>

<button on:click>
  {#each [true] as state (state)}
    <span
      class={state === enabled ? "thumb thumb-on" : "thumb thumb-off"}
      animate:flip={{ duration: 200 }}
    />
  {/each}
  <slot name="icon-off" />
  <slot name="icon-on" />
</button>

<style>
  button {
    position: relative;
    display: flex;
    align-items: center;
    justify-content: space-between;
    width: var(--toggle-width);
    height: var(--toggle-height);
    padding: var(--toggle-padding);
    margin: 0;
    border: solid 1px var(--text-on-primary);
    border-radius: calc(var(--toggle-width) / 2);
    color: var(--color-accent);
    background-color: var(--color-surface);
  }

  .thumb {
    position: absolute;
    top: var(--toggle-padding);
    width: calc(var(--toggle-height) - (var(--toggle-padding) * 2));
    height: calc(var(--toggle-height) - (var(--toggle-padding) * 2));
    border-radius: 50%;
    background: var(--text-on-primary);
  }

  .thumb-on {
    right: var(--toggle-padding);
  }

  .thumb-off {
    left: var(--toggle-padding);
  }
</style>
