<script lang="ts">
	import { isErrorResponse, getArtist } from '$lib/requests';
	import type { ErrorResponse, Artist } from '../../../types';
	import { _ } from 'svelte-i18n';
	import type { PageData } from './$types';
	import { Heading, Label, TabItem, Tabs } from 'flowbite-svelte';
	import WorkTileGrid from '$lib/components/WorkTileGrid.svelte';
	import Error from '$lib/components/Error.svelte';
	import { getTitle } from '$lib/util';

	export let data: PageData;

	let loading = true;
	let artistData: Artist;
	let error: ErrorResponse;

	const fetchArtistData = async () => {
		loading = true;
		const fetchedArtistData = await getArtist(Number(data.id));
		loading = false;
		if (isErrorResponse(fetchedArtistData)) {
			error = fetchedArtistData;
		} else {
			artistData = fetchedArtistData as Artist;
			loading = false;
		}
	};

	fetchArtistData();
</script>

<svelte:head>
	<title>{getTitle(artistData?.name)}</title>
</svelte:head>
<div id="work-wrapper">
	{#if error}
		<Error error={error?.error || ''} />
	{:else if loading}
		<p></p>
	{:else if artistData}
		<div id="artist-details">
			<Heading tag="h4">
				{artistData.name || $_('work.artist.fallback')}
			</Heading>
			{#if artistData.birthYear}
				<span class="details-entry">
					<Label>{$_('artist.birthYear.label')}</Label>
					<p class="text-gray-900 dark:text-gray-300">{artistData.birthYear}</p>
				</span>
			{/if}
			{#if artistData.deathYear}
				<span class="details-entry">
					<Label>{$_('artist.deathYear.label')}</Label>
					<p class="text-gray-900 dark:text-gray-300">{artistData.deathYear}</p>
				</span>
			{/if}
		</div>
		<Tabs contentClass="artist-tab">
			<TabItem open title={$_('layout.works.label')}>
				<WorkTileGrid
					artistID={data.id}
					pageStorageID={`artist${data.id}WorksPage`}
					pageSizeStorageID={`artist${data.id}WorksPageSize`}
					searchStorageID={`artist${data.id}WorksSearch`}
				/>
			</TabItem>
			<TabItem title={$_('artist.biography.label')}>
				<p id="artist-biography" class="text-gray-900 dark:text-gray-300">
					{$_('comingSoon.label')}
				</p>
			</TabItem>
		</Tabs>
	{/if}
</div>

<style>
	#artist-details {
		flex: 1;
		padding-left: 20px;
		padding-right: 20px;
		display: flex;
		flex-direction: column;
		justify-content: center;
		margin-bottom: 20px;
	}
	#artist-biography {
		margin: 20px 50px;
	}
</style>
