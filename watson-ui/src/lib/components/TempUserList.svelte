<script lang="ts">
	import TempUserDetails from './TempUserDetails.svelte';
	import { activeTempUsers } from '../../stores';

	export let loginUser: User;
	export let token: string;

	let showActiveUsers = $activeTempUsers.length > 0 ? true : false;
</script>

<div class="user-list">
	{#if showActiveUsers}
		{#each $activeTempUsers as user (user?.ID)}
			<TempUserDetails {token} {user} {loginUser} />
		{/each}
	{:else}
		<p>No users currently...</p>
	{/if}
</div>

<style>
	.user-list {
		width: 100%;
		display: flex;
		flex-flow: row wrap;
		justify-content: flex-start;
		align-items: center;
		margin: 0 auto;
	}
</style>
