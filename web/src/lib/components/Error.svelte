<script>
	import { Alert } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';

	export let error = '';
</script>

<Alert color="red">
	<span class="font-medium">{$_('error.label')}</span>
	{error}
</Alert>
