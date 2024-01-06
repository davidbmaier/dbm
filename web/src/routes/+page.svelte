<script>
	import { Input, Button, Label } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';
	import { login } from '$lib/requests';
	import Error from '$lib/components/Error.svelte';

	let user = '';
	let password = '';
	let error = '';

	const performLogin = async () => {
		const loginResponse = await login(user, password);
		error = loginResponse?.error || '';
	};
</script>

<section id="login-wrapper">
	<div class="login-element">
		<Label>{$_('login.user')}</Label>
		<Input bind:value={user} />
	</div>
	<div class="login-element">
		<Label>{$_('login.password')}</Label>
		<Input
			type="password"
			bind:value={password}
			on:keydown={(e) => {
				if (e.key === 'Enter') {
					performLogin();
				}
			}}
		/>
	</div>
	<div class="login-element">
		<Button on:click={performLogin}>{$_('login.login')}</Button>
		{#if error}
			<div id="login-error">
				<Error error={error == 'Unauthorized' ? $_('login.unauthorized.error') : error} />
			</div>
		{/if}
	</div>
</section>

<style>
	#login-wrapper {
		max-width: 500px;
		margin: auto;
		margin-top: 150px;
	}
	.login-element {
		margin-top: 10px;
	}
	#login-error {
		margin-top: 10px;
	}
</style>
