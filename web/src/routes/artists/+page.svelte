<script lang="ts">
	import { getArtists, isErrorResponse } from '$lib/requests';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';
	import type { ArtistsResponse, ErrorResponse } from '../../types';
	import { _ } from 'svelte-i18n';
	import { getStorage, updateStorage, type storageEntry } from '$lib/util';
	import Error from '$lib/components/Error.svelte';
	import Notification from '$lib/components/Notification.svelte';
	import SearchInput from '$lib/components/SearchInput.svelte';
	import Pagination from '$lib/components/Pagination.svelte';

	const pageStorageID = `artistsPage`;
	const pageSizeStorageID = `artistsPageSize`;
	const searchStorageID = `artistsSearch`;

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
		search = storedValues[searchStorageID];
	}

	let pages = [{ name: '1' }];
	let maxPage = 1;
	let loading = true;

	let artistsData: ArtistsResponse;
	let error: ErrorResponse;

	const updatePages = () => {
		if (artistsData.artists.length > 0) {
			maxPage = Math.ceil(artistsData.total / pageSize);

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
			if (maxPage - page >= 2) {
				tempPages.push({ name: `${maxPage}` });
			}
			pages = tempPages;
		} else {
			page = 1;
			pages = [{ name: '1' }];
		}
	};

	const fetchArtistsData = async () => {
		loading = true;
		const fetchedData = await getArtists(search, page, pageSize);
		loading = false;
		if (isErrorResponse(fetchedData)) {
			error = fetchedData;
		} else {
			artistsData = fetchedData as ArtistsResponse;
			updatePages();
			updateStorage([
				{ id: pageStorageID, value: page },
				{ id: pageSizeStorageID, value: pageSize },
				{ id: searchStorageID, value: search }
			]);
		}
	};

	fetchArtistsData();

	const switchToPreviousPage = async () => {
		if (page - 1 > 0) {
			page--;
			await fetchArtistsData();
		}
	};

	const switchToNextPage = async () => {
		if (page + 1 <= maxPage) {
			page++;
			await fetchArtistsData();
		}
	};

	const handleClick = async (e: any) => {
		const targetPage = e?.currentTarget?.innerText;
		if (page !== Number(targetPage)) {
			page = Number(targetPage);
			await fetchArtistsData();
		}
	};

	const handleSearchChange = async (e: any) => {
		const newSearchValue = e.detail;
		if (search !== newSearchValue) {
			search = newSearchValue;
			page = 1;
			await fetchArtistsData();
		}
	};

	const handlePageSizeChange = async (newSize: number) => {
		if (pageSize !== newSize) {
			pageSize = newSize;
			page = 1;
			await fetchArtistsData();
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
			placeholder={$_('artists.search.placeholder')}
		/>
		{#if loading}
			<p></p>
		{:else if artistsData}
			{#if artistsData.total !== 0}
				<Pagination
					{pages}
					{switchToNextPage}
					{switchToPreviousPage}
					{handleClick}
					startEntry={pageSize * (page - 1) + 1}
					endEntry={pageSize * (page - 1) + artistsData.artists.length}
					totalEntries={artistsData.total}
					{pageSize}
					{handlePageSizeChange}
				/>
				<div id="artists-table">
					<Table>
						<TableHead>
							<TableHeadCell>{$_('work.artist.label')}</TableHeadCell>
							<TableHeadCell>{$_('artists.numberOfWorks.label')}</TableHeadCell>
						</TableHead>
						<TableBody>
							{#each artistsData.artists as artist}
								<TableBodyRow>
									<TableBodyCell>
										<a href={`/artists/${artist.id}`}>
											{artist.name || $_('work.artist.fallback')}
										</a>
									</TableBodyCell>
									<TableBodyCell>
										{artist.numberOfWorks}
									</TableBodyCell>
								</TableBodyRow>
							{/each}
						</TableBody>
					</Table>
				</div>
				<Pagination
					{pages}
					{switchToNextPage}
					{switchToPreviousPage}
					{handleClick}
					startEntry={pageSize * (page - 1) + 1}
					endEntry={pageSize * (page - 1) + artistsData.artists.length}
					totalEntries={artistsData.total}
					{pageSize}
					{handlePageSizeChange}
					bottom
				/>
			{:else}
				<div id="artists-notification">
					<Notification message={$_('artists.nodata')} />
				</div>
			{/if}
		{/if}
	{/if}
</div>

<style>
	#artists-table {
		max-width: 1000px;
		margin: auto;
		margin-top: 20px;
	}

	#artists-table > :global(div) {
		overflow-x: hidden !important;
	}

	#artists-table :global(td) {
		text-overflow: ellipsis;
		overflow: hidden;
	}

	@media (max-width: 500px) {
		#artists-table :global(tr td:first-child) {
			max-width: 180px;
		}
		#artists-table :global(tr td:last-child) {
			max-width: 100px;
		}

		#artists-table :global(thead th:last-child) {
			max-width: 100px;
		}
	}

	#artists-notification {
		display: flex;
		justify-content: center;
		align-items: center;
	}
</style>
