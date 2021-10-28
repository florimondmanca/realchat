<script lang="ts">
  export let tag = "";
  export let selected = false;
</script>

{#if tag == "button"}
  <button class={selected ? "selected" : ""} role="listitem">
    <slot />
  </button>
{:else}
  <li>
    <slot />
  </li>
{/if}

<style>
  li,
  button {
    border: 1px solid var(--color-surface);
    border-bottom: none;
    padding: 1em;
  }

  li:first-child,
  button:first-child {
    border-radius: 6px 6px 0 0;
  }

  li:last-child,
  button:last-child {
    border-radius: 0 0 6px 6px;
    border-bottom: 1px solid var(--color-surface);
  }

  li:only-child,
  button:only-child {
    border-bottom: 1px solid var(--color-surface);
    border-radius: 6px;
  }

  button {
    /* Reset */
    display: block;
    margin: 0;
    width: 100%;
    height: 100%;
    background-color: transparent;
    color: inherit;
    text-align: inherit;
    /* /Reset */
    cursor: pointer;
  }

  button.selected,
  button:hover {
    background-color: var(--link-on-surface);
    color: var(--text-on-link);
  }
</style>
