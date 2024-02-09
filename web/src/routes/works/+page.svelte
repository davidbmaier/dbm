<script lang="ts">
	import { _ } from 'svelte-i18n';
	import WorkTileGrid from '$lib/components/WorkTileGrid.svelte';
	import { getTitle } from '$lib/util';
</script>

<svelte:head>
	<title>{getTitle($_('title.works'))}</title>
</svelte:head>
<div>
	<WorkTileGrid
		pageStorageID="worksPage"
		pageSizeStorageID="worksPageSize"
		searchStorageID="worksSearch"
	/>
</div>
