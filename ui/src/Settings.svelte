<script lang="ts">
  import { onMount } from "svelte";
  import BackIcon from "./BackIcon.svelte";
  import { LocalEnabled, OpenClientLocation, ToggleLocal } from "../wailsjs/go/launcher/App.js";

  export let closeSettings: () => void;

  let localEnabled: boolean | undefined = undefined;

  onMount(async () => {
    localEnabled = await LocalEnabled();
  });

  function openClientLocation() {
    console.log("openClientLocation")
    OpenClientLocation();
  }

  $: if (localEnabled !== undefined) ToggleLocal(localEnabled);
</script>

<button class="close" on:click={closeSettings}>
  <BackIcon />
</button>
<div>
  <h1>Settings</h1>
  <button class="client" on:click={openClientLocation}>Open client location</button>

  <label>
    <input
      type="checkbox"
      bind:checked={localEnabled}
      disabled={localEnabled === undefined}
    />
    <span>Enable local client<br /><em>(advanced, for developer use)</em></span>
  </label>
</div>

<style>
  div {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    gap: 8px;
  }

  label {
    display: flex;
    flex-direction: row;
    align-items: center;
    gap: 16px;
  }

  label span {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    justify-content: center;
  }

  label input {
    scale: 2;
  }

  button {
    background: none;
    border: none;
    cursor: pointer;
    padding: 8px;
    width: 200px;
    height: 56px;
    color: white;
    border-radius: 8px;
    box-shadow: #333333 0px 0px 4px 0px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
  }

  button.close {
    position: absolute;
    top: 0;
    right: 0;
    width: 48px;
    height: 48px;
    margin: 8px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: center;
    box-shadow: none;
  }

  button.client {
    background-color: #4e3bf5;
    height: auto;
  }
</style>
