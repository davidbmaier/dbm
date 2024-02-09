<script lang="ts">
	import { getArtists, getWorks } from '$lib/requests';
	import { Button, Heading } from 'flowbite-svelte';
	import { ArrowRightOutline, SearchSolid } from 'flowbite-svelte-icons';
	import { _ } from 'svelte-i18n';
	import type { ArtistsResponse, WorksResponse } from '../../types';
	import { getTitle } from '$lib/util';

	let works = 0;
	let artists = 0;
	let randomWork = 0;
	let randomArtist = 0;

	const getData = async () => {
		const worksData = (await getWorks('', 1)) as WorksResponse;
		works = worksData.total;
		const artistsData = (await getArtists('', 1)) as ArtistsResponse;
		artists = artistsData.total;
		randomWork = Math.ceil(Math.random() * works);
		randomArtist = Math.ceil(Math.random() * artists);
	};

	getData();
</script>

<svelte:head>
	<title>{$_('title.home')}</title>
</svelte:head>
{#if works !== 0 && artists !== 0}
	<div id="home-wrapper">
		<Heading tag="h3" class="home-heading home-heading-title">{$_('home.welcome.title')}</Heading>
		<Heading tag="h4" class="home-heading">
			{$_('home.welcome.description', {
				values: { works, artists }
			})}
		</Heading>
		<div id="home-cta">
			<a href="/works" class="home-cta-button">
				<Button size="lg">
					{$_('home.discoverWorks.label')}
					<ArrowRightOutline class="ms-2 h-3.5 w-3.5" />
				</Button>
			</a>
			<a href="/artists" class="home-cta-button">
				<Button size="lg">
					{$_('home.discoverArtists.label')}
					<ArrowRightOutline class="ms-2 h-3.5 w-3.5" />
				</Button>
			</a>
		</div>
		<Heading tag="h6" class="home-heading home-heading-inspiration">
			{$_('home.welcome.subtitle')}
		</Heading>
		<div id="home-random">
			<a href={`/works/${randomWork}`} class="home-cta-button">
				<Button>
					{$_('home.randomWork.label')}
					<SearchSolid class="ms-2 h-3.5 w-3.5" />
				</Button>
			</a>
			<a href={`/artists/${randomArtist}`} class="home-cta-button">
				<Button>
					{$_('home.randomArtist.label')}
					<SearchSolid class="ms-2 h-3.5 w-3.5" />
				</Button>
			</a>
		</div>
		<span id="home-footer" class="bg-white text-gray-700 dark:bg-gray-900 dark:text-gray-200">
			<div style="height: 1px" class="bg-gray-200 dark:bg-gray-700" />
			<div id="home-footer-content">
				<p>{$_('home.footer.title')}</p>
				<p>{$_('home.footer.description')}</p>
			</div>
		</span>
	</div>
{/if}

<style>
	#home-wrapper {
		display: flex;
		align-items: center;
		justify-content: center;
		flex-direction: column;
		height: calc(100% - 45px);
	}

	:global(.home-heading) {
		width: auto !important;
		text-align: center;
	}
	:global(.home-heading-title) {
		margin-bottom: 30px;
	}
	:global(.home-heading-inspiration) {
		margin-top: 45px;
		margin-bottom: 5px;
	}

	@media (max-width: 600px) {
		:global(.home-heading-title) {
			margin-bottom: 15px;
		}
		:global(.home-heading-inspiration) {
			margin-top: 25px;
			margin-bottom: 5px;
		}
	}

	#home-cta {
		display: flex;
		justify-content: center;
		margin-top: 20px;
		flex-wrap: wrap;
	}
	:global(.home-cta-button) {
		margin: 15px;
	}

	#home-random {
		display: flex;
		justify-content: center;
		flex-wrap: wrap;
	}

	#home-footer {
		position: fixed;
		bottom: 0;
		height: 65px;
		margin-bottom: 0;
		width: 100%;
		text-align: center;
	}
	#home-footer-content {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		height: 100%;
	}
</style>
