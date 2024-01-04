<script>
	import { Alert } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';

	export let error = '';
</script>

<Alert color="red" defaultClass="p-4 gap-3 text-sm dark:bg-gray-900">
	<span class="font-medium">{$_('error.label')}</span>
	{error}
</Alert>
