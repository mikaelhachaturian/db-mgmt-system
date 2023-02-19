<script lang="ts">
	import Button from '$lib/components/Button.svelte';
	import DBSearch from '$lib/components/DBSearch.svelte';
	import { error, type HttpError } from '@sveltejs/kit';
	import { activeTempUsers } from '../../stores';
	import Loading from '$lib/components/Loading.svelte';

	export let dbs: string[];
	export let user: User;
	export let token: string;
	let clipped = false;

	function copyToClipboard() {
		var copyText = userPassword;
		//@ts-ignore
		navigator.clipboard.writeText(copyText);
		clipped = true;
	}

	let filteredDBs = [];

	$: if (!fields.db) {
		filteredDBs = [];
		hiLiteIndex = null;
	}

	let hiLiteIndex = null;

	const setInputVal = (db: string) => {
		fields.db = db;
		filteredDBs = [];
		hiLiteIndex = null;
	};

	const filterDBs = () => {
		let storageArr = [];
		if (fields.db) {
			dbs.forEach((db) => {
				if (db.toLowerCase().includes(fields.db.toLowerCase())) {
					//@ts-ignore
					storageArr = [...storageArr, db];
				}
			});
		}
		filteredDBs = storageArr;
	};

	let userPassword = '';
	const setPassword = (pass: string) => {
		userPassword = pass;
	};

	let fields = {
		username: user?.email,
		db: '',
		duration: ''
	};

	let errors = {
		db: '',
		duration: ''
	};

	let showInfo = {
		status: false,
		message: '',
		error: false
	};
	let valid = false;
	let loading = false;

	const createUser = async () => {
		const createUser = await fetch(`/db-access/create`, {
			method: 'POST',
			body: JSON.stringify(fields),
			headers: {
				'Content-Type': 'application/json',
				Token: token
			}
		});
		if (!createUser.ok) {
			const msg = await createUser.json();
			throw error(400, {
				message: msg.message
			});
		}

		const { activeTempUser, createdUser } = await createUser.json();

		// add user to store
		activeTempUsers.update((users) => {
			return [activeTempUser.user, ...users];
		});

		return createdUser;
	};

	const handleCreateUser = async () => {
		valid = true;
		loading = true;
		// validate db name
		if (fields.db.trim().length < 1) {
			valid = false;
			errors.db = 'db can not be empty';
			loading = false;
		} else {
			errors.db = '';
		}
		// validate duration
		if (fields.duration.trim().length < 1) {
			valid = false;
			loading = false;
			errors.duration = 'duration can not be empty';
		} else {
			errors.duration = '';
		}

		if (valid) {
			try {
				const result = await createUser();
				setPassword(result.user.Password);
				showInfo.status = true;
				showInfo.message = `
        <span style="font-size: 30px;"> 🎉 </span>
        User: '${fields.username}' has been created!
        Copy Password to Clipboard.
        `;
				loading = false;
			} catch (e) {
				showInfo.status = true;
				showInfo.message = '🙅🏻‍♂️' + (e as HttpError).body.message;
				showInfo.error = true;
				console.log(e);
				loading = false;
			}
		}
	};
</script>

<form on:submit|preventDefault={handleCreateUser}>
	<div class="form-field">
		<input
			placeholder="Search Database"
			autocomplete="false"
			type="text"
			id="db"
			bind:value={fields.db}
			on:input={filterDBs}
		/>
		<!-- FILTERED LIST OF DBs -->
		{#if filteredDBs.length > 0}
			<ul id="autocomplete-items-list">
				{#each filteredDBs as db, i}
					<DBSearch
						itemLabel={db}
						highlighted={i === hiLiteIndex}
						on:click={() => setInputVal(db)}
					/>
				{/each}
			</ul>
		{/if}
		<div class="error">{errors.db}</div>
	</div>
	<div class="form-field">
		<input
			placeholder="Duration (10m, 1h30m, 4h etc..)"
			type="text"
			id="duration"
			bind:value={fields.duration}
		/>
		<div class="error">{errors.duration}</div>
	</div>
	<Button type="secondary">Create User</Button>
</form>
{#if loading}
	<Loading />
{/if}
{#if showInfo.status}
	<p class={!showInfo.error ? 'show-info' : 'show-error'}>
		{#if !showInfo.error}
			<span>Success! db temp user created.<br />Copy password to clipboard.</span>
		{:else}
			{showInfo.message}
		{/if}
		<!-- {@html showInfo.message} -->
		{#if !showInfo.error}
			<i
				class={clipped ? 'copytoclip fas fa-check' : 'copytoclip fas fa-copy'}
				on:click={copyToClipboard}
				on:keydown
			/>
		{/if}
	</p>
{/if}

<style>
	form {
		width: 400px;
		margin: 0 auto;
		text-align: center;
	}
	.form-field {
		margin: 18px auto;
	}

	.copytoclip {
		color: rgb(255, 255, 255);
		font-size: 25px;
		cursor: pointer;
		/* text-shadow: 1px 1px 0px black; */
		margin-left: 1em;
		background: -webkit-linear-gradient(rgb(86, 158, 236), rgb(5, 37, 116));
		-webkit-background-clip: text;
		background-clip: text;
		-webkit-text-fill-color: transparent;
	}

	input {
		width: 90%;
		height: auto;
		padding: 1.2em 1em;
		color: #6a6a6a;
		border-radius: 0.5em;
		border: 1px solid #d3d3d3;
		background: linear-gradient(
			180deg,
			rgba(255, 255, 255, 1) 0%,
			rgba(255, 255, 255, 1) 77%,
			rgba(236, 236, 236, 1) 100%
		);
		text-shadow: 0px 1px 0px #fff;
		outline: none;
	}

	input:focus {
		border: 1px solid #d2d2d2;
	}

	.show-info {
		animation: growDown;
		animation-duration: 1.3s;
		text-align: left;
		margin: 1.3em auto;
		display: flex;
		justify-content: center;
		align-items: center;
		animation: growDown;
		animation-duration: 0.3s;
		padding: 1em;
		background-color: rgb(228, 252, 240);
		border-radius: 0.7em;
		border-bottom: 2px solid #098f38;
	}
	.show-error {
		width: 100%;
		text-align: left;
		margin: 1.3em auto;
		display: flex;
		justify-content: center;
		align-items: center;
		animation: growDown;
		animation-duration: 0.7s;
		padding: 10px 0.7em;
		background-color: rgb(254 254 254);
		border-radius: 0.7em;
		border-bottom: 2px solid #da294d;
	}
	p {
		text-align: center;
		white-space: pre-wrap;
	}
	.error {
		font-weight: bold;
		font-size: 12px;
		color: #d91b42;
		text-align: left;
		padding: 0.3em 1em;
	}
	#autocomplete-items-list {
		position: relative;
		margin: 0 auto;
		padding: 0;
		top: 10px;
		overflow-y: auto;
		max-height: 140px;
		font-size: 13px;
		padding: 7px;
		animation: fadeDown;
		animation-duration: 0.9s;
	}

	::-webkit-scrollbar {
		width: 10px;
	}

	/* Track */
	::-webkit-scrollbar-track {
		border-radius: 1em;
		background: #ffffff;
		border: 1px solid #ededed;
	}

	/* Handle */
	::-webkit-scrollbar-thumb {
		background: rgb(103, 8, 120);
		border-radius: 1em;
	}

	::-webkit-scrollbar-thumb:hover {
		background: rgb(148, 10, 172);
	}
</style>
