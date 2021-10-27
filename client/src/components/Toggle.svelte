<script lang="ts">
  import { quintOut } from "svelte/easing";
  import { crossfade } from "svelte/transition";

  export let active: boolean;

  const [send, receive] = crossfade({
    duration: (d) => Math.sqrt(d * 2000),
    easing: quintOut,
  });
</script>

<button on:click>
  {#if active}
    <span
      class="thumb thumb-left"
      in:receive={{ key: 0 }}
      out:send={{ key: 0 }}
    />
  {:else}
    <span
      class="thumb thumb-right"
      in:receive={{ key: 0 }}
      out:send={{ key: 0 }}
    />
  {/if}
  <slot />
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
    cursor: pointer;
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

  .thumb-left {
    left: var(--toggle-padding);
  }

  .thumb-right {
    right: var(--toggle-padding);
  }
</style>
