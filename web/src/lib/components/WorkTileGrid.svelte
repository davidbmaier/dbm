<script lang="ts">
	import { getWorks, isErrorResponse } from '$lib/requests';
	import type { WorksResponse, ErrorResponse } from '../../types';
	import WorkTile from '$lib/components/WorkTile.svelte';
	import { getStorage, updateStorage, type storageEntry } from '$lib/util';
	import { _ } from 'svelte-i18n';
	import Error from './Error.svelte';
	import Notification from './Notification.svelte';
	import SearchInput from './SearchInput.svelte';
	import Pagination from './Pagination.svelte';

	export let pageStorageID = ``;
	export let pageSizeStorageID = ``;
	export let searchStorageID = ``;

	export let artistID = '';

	let page = 1;
	let search = '';
	let pageSize = 20;

	// get values from last visit
	const storedValues: storageEntry = getStorage([
		pageStorageID,
		pageSizeStorageID,
		searchStorageID
	]);
	if (storedValues[pageStorageID]) {
		page = Number(storedValues[pageStorageID]);
	}
	if (storedValues[pageSizeStorageID]) {
		pageSize = Number(storedValues[pageSizeStorageID]);
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
			maxPage = Math.ceil(worksData.total / pageSize);

			const tempPages = [];
			// first page
			if (page - 1 >= 3) {
				tempPages.push({ name: `1` });
			}
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
			// last page
			if (maxPage - page > 2) {
				tempPages.push({ name: `${maxPage}` });
			}
			pages = tempPages;
		} else {
			page = 1;
			pages = [{ name: '1' }];
		}
	};

	const fetchWorksData = async () => {
		loading = true;
		const fetchedData = await getWorks(search, page, pageSize, Number(artistID));
		loading = false;
		if (isErrorResponse(fetchedData)) {
			error = fetchedData;
		} else {
			worksData = fetchedData as WorksResponse;
			updatePages();
			updateStorage([
				{ id: pageStorageID, value: page },
				{ id: pageSizeStorageID, value: pageSize },
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
		const newSearchValue = e.detail;
		if (search !== newSearchValue) {
			search = newSearchValue;
			page = 1;
			await fetchWorksData();
		}
	};

	const handlePageSizeChange = async (newSize: number) => {
		if (pageSize !== newSize) {
			pageSize = newSize;
			page = 1;
			await fetchWorksData();
		}
	};
</script>

<div>
	{#if error}
		<Error error={error?.error || ''} />
	{:else}
		<SearchInput
			on:search={handleSearchChange}
			{search}
			placeholder={$_('works.search.placeholder')}
		/>
		{#if loading}
			<p></p>
		{:else if worksData}
			{#if worksData.total !== 0}
				<Pagination
					{pages}
					{switchToNextPage}
					{switchToPreviousPage}
					{handleClick}
					startEntry={pageSize * (page - 1) + 1}
					endEntry={pageSize * (page - 1) + worksData.works.length}
					totalEntries={worksData.total}
					{pageSize}
					{handlePageSizeChange}
				/>
				<div class="tile-grid">
					{#each worksData.works as work}
						<WorkTile {work} />
					{/each}
				</div>
				<Pagination
					{pages}
					{switchToNextPage}
					{switchToPreviousPage}
					{handleClick}
					startEntry={pageSize * (page - 1) + 1}
					endEntry={pageSize * (page - 1) + worksData.works.length}
					totalEntries={worksData.total}
					{pageSize}
					{handlePageSizeChange}
					bottom
				/>
			{:else}
				<div id="tile-grid-notification">
					<Notification message={$_('works.nodata')} />
				</div>
			{/if}
		{/if}
	{/if}
</div>

<style>
	.tile-grid {
		display: grid;
		grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
	}
	.tile-grid :global(.worktile-card) {
		justify-self: center;
	}

	@media (max-width: 1650px) {
		.tile-grid :global(.worktile-card) {
			/* reduce tile width to still fit 5 cards for clean page sizes */
			width: 250px;
		}
	}

	@media (max-width: 1400px) {
		.tile-grid {
			grid-template-columns: 1fr 1fr;
		}
	}

	@media (max-width: 720px) {
		.tile-grid {
			grid-template-columns: 1fr;
		}
	}

	#tile-grid-notification {
		display: flex;
		justify-content: center;
		align-items: center;
	}
</style>
