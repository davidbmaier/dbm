<script>
	import { Alert } from 'flowbite-svelte';

	export let error = '';
</script>

<Alert color="red">
	<span class="font-medium">Something went wrong!</span>
	{error}
</Alert>
