<script lang="ts">
	import '../app.pcss';
	import '$lib/i18n/i18n.ts';
	import { _, isLoading } from 'svelte-i18n';

	import { page } from '$app/stores';
	import {
		Navbar,
		NavBrand,
		NavLi,
		NavUl,
		NavHamburger,
		DarkMode,
		Modal,
		Button,
		Input,
		Label
	} from 'flowbite-svelte';
	import { updatePassword } from '$lib/requests';
	import Error from '$lib/components/Error.svelte';
	import { UserSettingsSolid } from 'flowbite-svelte-icons';
	$: activeUrl = $page.url.pathname;

	let modalOpened = false;
	let newPassword = '';
	let passwordError = '';

	const translatePasswordError = (errorMessage: string) => {
		switch (errorMessage) {
			case 'Bad request, password too short':
				return $_('account.passwordError.tooShort');
			case 'Bad request, password too long':
				return $_('account.passwordError.tooLong');
			default:
				return errorMessage;
		}
	};

	const cleanUpPasswordModal = () => {
		newPassword = '';
		passwordError = '';
	};

	const performPasswordUpdate = async () => {
		const updateError = await updatePassword(newPassword);
		if (updateError) {
			passwordError = updateError.error;
		} else {
			modalOpened = false;
			newPassword = '';
		}
	};
</script>

<div id="wrapper">
	{#if !$isLoading}
		<Navbar>
			<NavBrand href="/works">
				<img src="/favicon.png" class="me-3 h-6 sm:h-9" alt="" />
				<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
					{$_('layout.title')}
				</span>
			</NavBrand>
			<div class="flex md:order-2">
				{#if activeUrl != '/'}
					<Button
						on:click={() => (modalOpened = true)}
						title={$_('account.settings.title')}
						class="settings-button rounded-lg p-2.5 text-sm text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-4 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-700"
					>
						<UserSettingsSolid />
					</Button>
				{/if}
				<DarkMode title={$_('darkMode.title')} />
				<NavHamburger />
			</div>
			<NavUl {activeUrl} activeClass="nav-active">
				{#if activeUrl != '/'}
					<NavLi href="/works">{$_('layout.works.label')}</NavLi>
					<NavLi href="/artists">{$_('layout.artists.label')}</NavLi>
				{/if}
			</NavUl>
		</Navbar>
		<div style="height: 1px" class="bg-gray-200 dark:bg-gray-700" />
		<Modal
			title={$_('account.settings.title')}
			bind:open={modalOpened}
			on:close={cleanUpPasswordModal}
			size="sm"
		>
			<Label>{$_('account.newPassword.label')}</Label>
			<Input type="password" bind:value={newPassword} />
			{#if passwordError}
				<div id="login-error">
					<Error error={translatePasswordError(passwordError)} />
				</div>
			{/if}
			<Button on:click={performPasswordUpdate}>{$_('account.updatePassword.label')}</Button>
		</Modal>

		<main id="content-wrapper">
			<slot />
		</main>
	{/if}
</div>

<style>
	#wrapper {
		height: 100vh;
	}

	#content-wrapper {
		height: calc(100% - 73px);
		padding: 20px 50px;
	}
	@media (max-width: 500px) {
		#content-wrapper {
			padding: 20px 20px;
		}
	}

	:global(.nav-active) {
		font-weight: bold;
	}

	:global(.settings-button) {
		background-color: transparent;
		margin-right: 5px;
	}
</style>
