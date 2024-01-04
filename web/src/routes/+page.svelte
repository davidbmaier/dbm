<script>
	import { Input, Button } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';
	import { login } from '$lib/requests';
	import Error from '$lib/components/Error.svelte';

	let email = '';
	let password = '';
	let error = '';

	const performLogin = async () => {
		const loginResponse = await login(email, password);
		error = loginResponse?.error || '';
	};
</script>

<section id="login-wrapper">
	<div>
		{$_('login.email')}
		<Input bind:value={email} />
	</div>
	<div>
		{$_('login.password')}
		<Input bind:value={password} />
	</div>
	<div id="login-button">
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
		margin-top: 50px;
	}
	#login-button {
		margin-top: 10px;
	}
	#login-error {
		margin-top: 10px;
	}
</style>
