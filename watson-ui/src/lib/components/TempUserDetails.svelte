<script lang="ts">
	import { fade, scale } from 'svelte/transition';
	import { PUBLIC_WATSON_BACKEND_URL } from '$env/static/public';
	import { error } from '@sveltejs/kit';
	import { activeTempUsers } from '../../stores';
	import { onMount } from 'svelte';
	import parse from 'parse-duration';

	export let token: string;
	export let user: User;
	export let loginUser: User;

	let updatedAt = user?.UpdatedAt;
	let duration = user?.duration;
	let minutesToDeletion = '';
	let showtime = '';
	let showDelete = false;

	onMount(() => {
		if (loginUser?.email == user?.email) showDelete = true;
		const interval = setInterval(() => {
			//@ts-ignore
			minutesToDeletion = formatDate(updatedAt, duration).toString();
		}, 1000);

		return () => {
			clearInterval(interval);
		};
	});

	//@ts-ignore
	$: showtime =
		minutesToDeletion === '0s' ? `<span style='color: red;'>Expired</span>` : minutesToDeletion;

	const determineLetter = (diffReturn: Date) => {
		if (diffReturn.getHours() > 2)
			return (
				(diffReturn.getHours() - 2).toString() +
				'h' +
				diffReturn.getMinutes().toString() +
				'm' +
				diffReturn.getSeconds().toString() +
				's'
			);
		if (diffReturn.getMinutes() > 0)
			return diffReturn.getMinutes().toString() + 'm' + diffReturn.getSeconds().toString() + 's';
		return diffReturn.getSeconds().toString() + 's';
	};

	const formatDate = (t: string, d: string) => {
		var startTime = new Date(t);
		let date_with_dur = startTime.setMilliseconds(startTime.getMilliseconds() + parse(d));
		var endTime = Date.now();
		var difference = new Date(date_with_dur).getTime() - endTime; // This will give difference in milliseconds

		let diffReturn = new Date(difference);
		return difference > 0 ? determineLetter(diffReturn) : 0 + 's';
	};

	const handleDelete = async () => {
		const deleteUser = await fetch(`/db-access/delete`, {
			method: 'DELETE',
			body: JSON.stringify({ username: user?.email, duration: '0s' }),
			headers: {
				'Content-Type': 'application/json',
				Token: token
			}
		});
		if (!deleteUser.ok) {
			const msg = await deleteUser.json();
			throw error(400, msg.message);
		}

		// delete user from store
		activeTempUsers.update((users) => {
			return users.filter((u) => user?.email != u?.email);
		});

		const result = await deleteUser.json();
		return result;
	};
</script>

<div class="user-card" in:fade out:scale|local>
	{#if showDelete}
		<div class="close-wrapper">
			<i on:click={handleDelete} on:keydown class="close-btn fas fa-trash" />
		</div>
	{/if}
	<div class="rside">
		<img class="user-img" src={user?.picture} alt="Avatar" referrerpolicy="no-referrer" />
	</div>
	<div class="lside">
		<p style="font-size: 13px; padding: 5px 0; font-weight: bold; color: #622e83;">
			{user?.firstname + ' ' + user?.lastname}
		</p>
		<p><i class="fas fa-inbox" /> {user?.email}</p>
		<p><i class="fas fa-server" /> {user?.dbname}</p>
		<span> Available: {@html showtime} </span>
	</div>
</div>

<style>
	.rside {
		width: 25%;
		height: auto;
		padding: 1em;
	}

	.user-card {
		width: 350px;
		height: auto;
		color: black;
		margin: 1em 0.8em;
		padding: 0.7em 0.5em;
		box-shadow: -6px -3px 8px -5px rgba(0, 0, 0, 0.2);
		display: flex;
		justify-content: start;
		align-items: center;
		transition: 0.3s;
		border-radius: 0.3em;
		font-size: 12px;
		position: relative;
		border-bottom: 3px solid #622e83;
	}

	.user-card:hover {
		box-shadow: 4px 6px 8px -6px rgba(0, 0, 0, 0.5);
	}

	i {
		padding-right: 5px;
		color: black;
	}

	img.user-img {
		width: 55px;
		height: auto;
		border-radius: 50%;
		box-shadow: 6px 7px 3px -3px rgb(0 0 0 / 10%);
	}

	.user-card p {
		text-align: left;
		display: block;
		padding: 0;
		margin: 0;
		font-size: 13 px;
	}

	.user-card span {
		font-size: 12px;
		padding-top: 0.7em;
	}

	.close-btn {
		border: 2px solid #622e83;
		background: radial-gradient(ellipse at center, #8e29cd, #2d0447);
		border-radius: 50%;
		display: inline-block;
		position: absolute;
		top: -5px;
		right: -15px;
		color: white;
		font-size: 15px;
		text-align: center;
		cursor: pointer;
		visibility: hidden;
		opacity: 0;
		padding: 0.7em;
		transition: all 0.3s;
	}

	.user-card:hover .close-btn {
		visibility: visible;
		opacity: 1;
	}

	.lside {
		width: 75%;
		height: auto;
		padding: 0.5em 2em;
		border-left: 1px solid #e1e1e1;
		display: flex;
		flex-direction: column;
	}
</style>
