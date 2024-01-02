<script lang="ts">
	import { Card, Heading } from 'flowbite-svelte';
	import type { Work } from '../../types';
	import { getFileLink } from '$lib/util';
	import { _ } from 'svelte-i18n';

	export let work: Work;
</script>

<span class="worktile-card">
	<Card img={getFileLink(work)} href={`/works/${work.id}`}>
		<Heading tag="h6">
			{work.name || $_('worktile.name.fallback')}
		</Heading>
		<p>
			{work.artist.name || $_('worktile.artist.fallback')}
		</p>
	</Card>
</span>

<style>
	.worktile-card {
		margin: 15px 0px;
		width: 300px;
	}

	.worktile-card :global(img) {
		object-fit: cover;
		height: 300px;
		width: 300px;
		flex-shrink: 0;
	}
	.worktile-card > :global(*) {
		height: 100%;
		max-width: none;
	}

	@media (min-width: 2140px) {
		.worktile-card {
			width: 400px;
		}
		.worktile-card :global(img) {
			height: 400px;
			width: 400px;
		}
	}
</style>
