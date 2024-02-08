<script lang="ts">
	import { Pagination, P, Button, Dropdown, DropdownItem } from 'flowbite-svelte';
	import { ChevronDownSolid, ChevronLeftOutline, ChevronRightOutline } from 'flowbite-svelte-icons';
	import type { LinkType } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';

	export let pages: LinkType[] | undefined;
	export let switchToPreviousPage: () => void;
	export let switchToNextPage: () => void;
	export let handleClick: (e: any) => Promise<void>;
	export let startEntry = 1;
	export let endEntry = 20;
	export let totalEntries = 20;

	export let pageSize = 20;
	export let handlePageSizeChange: (newSize: number) => void;
	export let bottom = false;

	const pageSizeOptions = [10, 20, 30, 40, 50];
</script>

<div class="pagination" style={bottom ? 'margin-bottom: 20px' : ''}>
	<div class="pagination-rows">
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
		<div class="pagination-extra">
			<P size="sm" color="text-gray-700 dark:text-gray-400" class="pagination-entries">
				{startEntry} - {endEntry}
				{$_('pagination.of.label')}
				{totalEntries}
			</P>
			<span class="pagination-pagesize">
				<Button size="xs" class="pagination-pagesize-button">
					{pageSize}
					<ChevronDownSolid class="ms-2 h-3 w-3 text-white dark:text-white" />
				</Button>
				<Dropdown>
					{#each pageSizeOptions as option}
						<DropdownItem on:click={() => handlePageSizeChange(Number(option))}>
							{option}
						</DropdownItem>
					{/each}
				</Dropdown>
				<P size="sm" color="text-gray-700 dark:text-gray-400">
					{$_('pagination.itemsPerPage.label')}
				</P>
			</span>
		</div>
	</div>
</div>

<style>
	.pagination {
		display: flex;
		justify-content: center;
	}
	.pagination-rows {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	.pagination-extra {
		display: flex;
		margin-top: 10px;
		align-items: center;
	}
	:global(.pagination-entries) {
		margin-right: 20px;
	}
	.pagination-pagesize {
		display: flex;
		justify-content: center;
		align-items: center;
	}
	:global(.pagination-pagesize-button) {
		margin-right: 5px;
	}
</style>
