<script>
	import '../app.pcss';
	import '$lib/i18n/i18n.ts';
	import { _, isLoading } from 'svelte-i18n';

	import { page } from '$app/stores';
	import { Navbar, NavBrand, NavLi, NavUl, NavHamburger, DarkMode } from 'flowbite-svelte';
	$: activeUrl = $page.url.pathname;
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
				<DarkMode />
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
</style>
