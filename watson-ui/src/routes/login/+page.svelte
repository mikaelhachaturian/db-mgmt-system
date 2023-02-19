<script lang="ts">
  import { onMount } from "svelte";
  import { initializeGoogleAccounts, renderGoogleButton } from '$lib/google'
  import { loggedInStore } from "../../stores";
	import { goto } from "$app/navigation";
  

  onMount(async () => {
    initializeGoogleAccounts()
    renderGoogleButton()
    //@ts-ignore
    if (!$loggedInStore) google.accounts.id.prompt()
    else goto("/db-access", {invalidateAll: true}) // add more elegant solution to redirect loggedIn user to db-access (1 frame of login page)
  })

</script>

<div class="loginBox">
  {#if !$loggedInStore }
  <div class="googleWrapper">
    <div id="googleButton" class="googleButton"></div>
  </div>
  {/if}
</div>

<style>
  .loginBox {
    text-align: center;
    margin: 0 auto;
    overflow: hidden;
    height: 100%;
  }
  .googleWrapper {
    display: flex;
    justify-content: center;
    margin: 0 auto;
  }
  
</style>
