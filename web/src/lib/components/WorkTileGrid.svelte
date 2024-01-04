<script lang="ts">
	import { getWorks, isErrorResponse } from '$lib/requests';
	import { Pagination, Input } from 'flowbite-svelte';
	import { ChevronLeftOutline, ChevronRightOutline } from 'flowbite-svelte-icons';
	import type { WorksResponse, ErrorResponse } from '../../types';
	import WorkTile from '$lib/components/WorkTile.svelte';
	import { getStorage, updateStorage, type storageEntry, debounce } from '$lib/util';
	import { _ } from 'svelte-i18n';
	import Error from './Error.svelte';

	export let pageStorageID = ``;
	export let searchStorageID = ``;

	export let artistID = '';

	let page = 1;
	let search = '';

	// get values from last visit
	const storedValues: storageEntry = getStorage([pageStorageID, searchStorageID]);
	if (storedValues[pageStorageID]) {
		page = Number(storedValues[pageStorageID]);
	}
	if (storedValues[searchStorageID]) {
		search = storedValues[searchStorageID] || '';
	}

	let pages = [{ name: '1' }];
	let maxPage = 1;
	let loading = true;

	let worksData: WorksResponse;
	let error: ErrorResponse;

	const updatePages = () => {
		if (worksData.works.length > 0) {
			maxPage = Math.ceil(worksData.total / worksData.works.length);

			const tempPages = [];
			if (page - 2 > 0) {
				tempPages.push({ name: `${Number(page) - 2}` });
			}
			if (page - 1 > 0) {
				tempPages.push({ name: `${Number(page) - 1}` });
			}
			tempPages.push({ name: `${Number(page)}`, active: true });
			if (page + 1 <= maxPage) {
				tempPages.push({ name: `${Number(page) + 1}` });
			}
			if (page + 2 <= maxPage) {
				tempPages.push({ name: `${Number(page) + 2}` });
			}
			pages = tempPages;
		} else {
			page = 1;
			pages = [{ name: '1' }];
		}
	};

	const fetchWorksData = async () => {
		loading = true;
		const fetchedData = await getWorks(search, page, Number(artistID));
		loading = false;
		if (isErrorResponse(fetchedData)) {
			error = fetchedData;
		} else {
			worksData = fetchedData as WorksResponse;
			updatePages();
			updateStorage([
				{ id: pageStorageID, value: page },
				{ id: searchStorageID, value: search }
			]);
		}
	};

	fetchWorksData();

	const switchToPreviousPage = async () => {
		if (page - 1 > 0) {
			page--;
			await fetchWorksData();
		}
	};

	const switchToNextPage = async () => {
		if (page + 1 <= maxPage) {
			page++;
			await fetchWorksData();
		}
	};

	const handleClick = async (e: any) => {
		const targetPage = e?.currentTarget?.innerText;
		if (page !== Number(targetPage)) {
			page = Number(targetPage);
			await fetchWorksData();
		}
	};

	const handleSearchChange = async (e: any) => {
		const newSearchValue = e.target.value;
		if (search !== newSearchValue) {
			search = newSearchValue;
			page = 1;
			await fetchWorksData();
		}
	};
	const debounceSearchChange = debounce(handleSearchChange);
</script>

<div>
	{#if error}
		<Error error={error?.error || ''} />
	{:else}
		<div class="search">
			<span>
				<Input
					size="lg"
					class="search-input"
					placeholder={$_('works.search.placeholder')}
					value={search}
					on:input={(e) => debounceSearchChange(e)}
				/>
			</span>
		</div>
		{#if loading}
			<p></p>
		{:else if worksData}
			{#if worksData.total !== 0}
				<div class="tile-grid">
					{#each worksData.works as work}
						<WorkTile {work} />
					{/each}
				</div>
				<div class="pagination">
					<Pagination
						{pages}
						on:previous={switchToPreviousPage}
						on:next={switchToNextPage}
						on:click={handleClick}
						icon
					>
						<svelte:fragment slot="prev">
							<ChevronLeftOutline class="h-2.5 w-2.5" />
						</svelte:fragment>
						<svelte:fragment slot="next">
							<ChevronRightOutline class="h-2.5 w-2.5" />
						</svelte:fragment>
					</Pagination>
				</div>
			{:else}
				{$_('works.nodata')}
			{/if}
		{/if}
	{/if}
</div>

<style>
	.tile-grid {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
		margin: 20px 50px;
	}
	.tile-grid :global(.worktile-card) {
		justify-self: center;
	}

	@media (max-width: 1650px) {
		.tile-grid {
			grid-template-columns: 1fr 1fr 1fr 1fr;
		}
	}

	@media (max-width: 1340px) {
		.tile-grid {
			grid-template-columns: 1fr 1fr;
		}
	}

	@media (max-width: 720px) {
		.tile-grid {
			grid-template-columns: 1fr;
		}
	}
</style>
