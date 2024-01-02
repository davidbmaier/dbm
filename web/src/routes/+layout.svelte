<script>
	import '../app.pcss';
	import '$lib/i18n/i18n.ts'
	import { _ } from 'svelte-i18n'
</script>

<nav>
	<h1>{$_("layout.title")}</h1>
</nav>

<main>
	<slot />
</main>
