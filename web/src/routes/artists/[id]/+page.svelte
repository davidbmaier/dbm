<script lang="ts">
	import { isErrorResponse, getArtist } from '$lib/requests';
	import type { ErrorResponse, Artist } from '../../../types';
	import { _ } from 'svelte-i18n';
	import type { PageData } from './$types';
	import { Heading, Label, TabItem, Tabs } from 'flowbite-svelte';
	import WorkTileGrid from '$lib/components/WorkTileGrid.svelte';
	import Error from '$lib/components/Error.svelte';

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
					<p>{artistData.birthYear}</p>
				</span>
			{/if}
			{#if artistData.deathYear}
				<span class="details-entry">
					<Label>{$_('artist.deathYear.label')}</Label>
					<p>{artistData.deathYear}</p>
				</span>
			{/if}
		</div>
		<Tabs contentClass="artist-tab">
			<TabItem open title={$_('layout.works.label')}>
				<WorkTileGrid
					artistID={data.id}
					pageStorageID={`artist${data.id}WorksPage`}
					searchStorageID={`artist${data.id}WorksSearch`}
				/>
			</TabItem>
			<TabItem title={$_('artist.biography.label')}>{$_('comingSoon.label')}</TabItem>
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
</style>
