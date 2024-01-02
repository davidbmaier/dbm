<script lang="ts">
	import { getWorks, isErrorResponse } from '$lib/requests';
	import { Label, Pagination, Input } from 'flowbite-svelte';
	import { ChevronLeftOutline, ChevronRightOutline } from 'flowbite-svelte-icons';
	import type { WorksResponse, ErrorResponse } from '../../types';
	import WorkTile from '$lib/components/WorkTile.svelte';
	import { getStorage, updateStorage, type storageEntry, debounce } from '$lib/util';
	import { _ } from 'svelte-i18n';

	const pageStorageID = `worksPage`;
	const searchStorageID = `worksSearch`;

	export let page = 1;
	export let search = '';

	// get values from last visit
	const storedValues: storageEntry = getStorage([pageStorageID, searchStorageID]);
	if (storedValues[pageStorageID]) {
		page = Number(storedValues[pageStorageID]);
	}
	if (storedValues[searchStorageID]) {
		search = storedValues[searchStorageID];
	}

	let pages = [{ name: '1' }];
	let maxPage = 1;
	let loading = true;

	let worksData: WorksResponse;
	let error: ErrorResponse;

	const updatePages = () => {
		if (worksData.works.length > 0) {
			maxPage = worksData.total / worksData.works.length;

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
		const fetchedData = await getWorks(search, page);
		if (isErrorResponse(fetchedData)) {
			error = fetchedData;
		} else {
			worksData = fetchedData as WorksResponse;
			updatePages();
			updateStorage([
				{ id: pageStorageID, value: page },
				{ id: searchStorageID, value: search }
			]);
			loading = false;
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
		<p>{error.error}</p>
	{:else}
		<div class="grid-search">
			<span>
				<Input
					size="lg"
					class="grid-search-input"
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
				<div class="grid-pagination">
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
