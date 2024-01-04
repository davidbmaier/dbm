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
				<span class="self-center whitespace-nowrap text-xl font-semibold dark:text-white">
					{$_('layout.title')}
				</span>
			</NavBrand>
			<NavHamburger />
			<NavUl {activeUrl} activeClass="nav-active">
				<NavLi href="/works">{$_('layout.works.label')}</NavLi>
				<NavLi href="/artists">{$_('layout.artists.label')}</NavLi>
			</NavUl>
		</Navbar>

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
