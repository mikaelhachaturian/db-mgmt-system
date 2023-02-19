<script lang="ts">
  import "../global.css";
	import { goto } from "$app/navigation";
  import { loggedInStore, userInStore } from "../stores";
	import type { LayoutServerData } from "./$types";
	import TopNav from "$lib/components/TopNav.svelte";
	import SideBar from "$lib/components/SideBar.svelte";

  export let data: LayoutServerData;
  let { loggedIn, user } = data.session

  loggedInStore.set(loggedIn)
  userInStore.set(user)

  const logout = async () => {
		const url = '/auth/logout'
		const res = await fetch(url, {
			method: 'POST'
		})
		if (!res.ok) {
      console.error(`Logout not successful: ${res.statusText} (${res.status})`)
		}
    loggedInStore.set(false)
    goto("/login", {invalidateAll: true})
	}

</script>

<header>
  {#if $loggedInStore}
  <TopNav isLoggedIn={$loggedInStore} currentUser={$userInStore} logoutAction={logout} />
  {/if}
</header>
<main>
  {#if $loggedInStore}
    <SideBar currentSelect={data.url}/>
  {/if}
  <slot />
  <div class="tooltip">
    <a target="_blank" rel="noopener noreferrer" href="https://easy-send-group.slack.com/archives/C03G3UT5G6M"><img id="slack" alt="slack" src="https://a.slack-edge.com/80588/marketing/img/meta/slack_hash_256.png" /></a>
    <span class="tooltiptext">Questions to DevOps Slack Channel</span>
  </div>
</main>
<footer>
  <p>Copyright 2023 - DevOps Team</p>
</footer>


<style>
  header {
    display: flex;
    justify-content: center;
  }
  footer {
    position: fixed;
    left: 0;
    bottom: 0;
    width: 100%;
    padding: 1.5em 0;
    text-align: center;
  }

  main {
    display: flex;
    width: 100%;
  }

  #slack {
    width: 64px;
    height: auto;
    border-radius: 50%;
  }

  .tooltip {
    position: fixed;
    left: 9em;
    bottom: 3.5em;
  }

  .tooltip .tooltiptext {
    width: max-content;
    visibility: hidden;
    background-color: rgb(107, 0, 194);
    color: #fff;
    text-align: center;
    padding: .2em .7em;
    border-radius: 6px;
    position: absolute;
    left: 120%;
    z-index: 1;
    top: 25%;
  }

  .tooltip:hover .tooltiptext {
    visibility: visible;
    animation: fadeIn;
    animation-duration: 2s;
  }

  .tooltip:hover #slack {
    animation: spin30;
    animation-duration: 1s;
  }

</style>
