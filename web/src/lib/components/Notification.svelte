<script>
	import { Alert } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';

	export let message = '';
</script>

<Alert color="blue" defaultClass="p-4 gap-3 text-sm dark:bg-gray-900">
	<span class="font-medium">{message}</span>
</Alert>
