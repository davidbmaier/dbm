<script lang="ts">
	import { page } from '$app/stores';
	import Error from '$lib/components/Error.svelte';
</script>

<Error error={$page.error?.message} />
