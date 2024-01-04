<script>
	import '../app.pcss';
	import '$lib/i18n/i18n.ts';
	import { _, isLoading } from 'svelte-i18n';

	import { page } from '$app/stores';
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger } from 'flowbite-svelte';
	$: activeUrl = $page.url.pathname;
</script>

<div id="wrapper">
	{#if !$isLoading}
		<Navbar>
			<NavBrand href="/works">
				<img src="/src/images/logo.png" class="me-3 h-6 sm:h-9" alt="Flowbite Logo" />
				<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
					{$_('layout.title')}
				</span>
			</NavBrand>
			<NavHamburger />
			<span style={activeUrl != '/' ? '' : 'visibility:hidden'}>
				<NavUl {activeUrl} activeClass="nav-active">
					<NavLi href="/works">{$_('layout.works.label')}</NavLi>
					<NavLi href="/artists">{$_('layout.artists.label')}</NavLi>
				</NavUl>
			</span>
		</Navbar>
		<hr />

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
		height: calc(100% - 72px);
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
</style>
