<script lang="ts">
	import { isErrorResponse, getWork } from '$lib/requests';
	import type { Work, ErrorResponse } from '../../../types';
	import { _ } from 'svelte-i18n';
	import type { PageData } from './$types';
	import { Heading, Label } from 'flowbite-svelte';
	import { getFileLink } from '$lib/util';
	import Error from '$lib/components/Error.svelte';

	export let data: PageData;

	let loading = true;
	let workData: Work;
	let error: ErrorResponse;

	const fetchWorkData = async () => {
		loading = true;
		const fetchedData = await getWork(Number(data.id));
		loading = false;
		if (isErrorResponse(fetchedData)) {
			error = fetchedData;
		} else {
			workData = fetchedData as Work;
			loading = false;
		}
	};

	fetchWorkData();
</script>

<div id="work-wrapper">
	{#if error}
		<Error error={error?.error || ''} />
	{:else if loading}
		<p></p>
	{:else if workData}
		<div id="work-structure">
			<div id="work-image-wrapper">
				<a href={getFileLink(workData)} target="_blank">
					<img
						id="work-image"
						class="image-shadow"
						src={getFileLink(workData)}
						alt={$_('work.altMessage', {
							values: { name: workData.name, artist: workData.artist.name }
						})}
					/>
				</a>
			</div>
			<div id="work-details">
				<Heading tag="h4">
					{workData.name || $_('work.name.fallback')}
				</Heading>
				<span class="details-entry">
					<Label>{$_('work.artist.label')}</Label>
					<a href={`/artists/${workData.artistID}`}>
						{workData.artist.name || $_('work.artist.fallback')}
					</a>
				</span>
				{#if workData.creationInfo}
					<span class="details-entry">
						<Label>{$_('work.creationInfo.label')}</Label>
						<p class="text-gray-900 dark:text-gray-300">{workData.creationInfo}</p>
					</span>
				{/if}
				{#if workData.material}
					<span class="details-entry">
						<Label>{$_('work.material.label')}</Label>
						<p class="text-gray-900 dark:text-gray-300">{workData.material}</p>
					</span>
				{/if}
				{#if workData.size}
					<span class="details-entry">
						<Label>{$_('work.size.label')}</Label>
						<p class="text-gray-900 dark:text-gray-300">{workData.size}</p>
					</span>
				{/if}
				{#if workData.owner}
					<span class="details-entry">
						<Label>{$_('work.owner.label')}</Label>
						<p class="text-gray-900 dark:text-gray-300">{workData.owner}</p>
					</span>
				{/if}
				{#if workData.source}
					<span class="details-entry">
						<Label>{$_('work.source.label')}</Label>
						<p class="text-gray-900 dark:text-gray-300">{workData.source}</p>
					</span>
				{/if}
			</div>
		</div>
	{/if}
</div>

<style>
	#work-wrapper {
		height: 100%;
		display: flex;
		flex-direction: row;
		align-items: center;
	}

	#work-structure {
		display: flex;
		width: 100%;
	}
	#work-image-wrapper {
		flex: 1.5;
		display: flex;
		justify-content: center;
		align-items: center;
	}
	#work-image {
		max-height: 80vh;
	}
	.image-shadow {
		box-shadow:
			rgba(82, 82, 118, 0.25) 0px 50px 100px -20px,
			rgba(125, 125, 125, 0.3) 0px 30px 60px -30px;
	}

	#work-details {
		flex: 1;
		padding-left: 40px;
		padding-right: 40px;
		display: flex;
		flex-direction: column;
		justify-content: center;
	}

	@media (max-width: 1000px) {
		#work-wrapper {
			display: block;
		}
		#work-structure {
			display: block;
		}

		#work-details {
			padding-left: 0;
			padding-right: 0;
			padding-top: 20px;
		}
	}
</style>
