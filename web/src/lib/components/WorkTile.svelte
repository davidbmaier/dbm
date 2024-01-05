<script lang="ts">
	import { Card, Heading } from 'flowbite-svelte';
	import type { Work } from '../../types';
	import { getFileLink } from '$lib/util';
	import { _ } from 'svelte-i18n';

	export let work: Work;
</script>

<span class="worktile-card">
	<Card href={`/works/${work.id}`}>
		<div class="worktile-card-image-wrapper rounded-lg">
			<img src={getFileLink(work)} alt="" />
		</div>

		<div class="worktile-card-text">
			<Heading tag="h6">
				{work.name || $_('work.name.fallback')}
			</Heading>
			<p>
				{work.artist.name || $_('work.artist.fallback')}
			</p>
		</div>
	</Card>
</span>

<style>
	.worktile-card {
		margin: 15px 0px;
		width: 300px;
		overflow: hidden;
	}
	.worktile-card:hover :global(img) {
		transform: scale(1.1);
	}
	.worktile-card > :global(*) {
		height: 100%;
		max-width: none;
	}

	.worktile-card-image-wrapper {
		margin-top: -1.5rem;
		width: calc(100% + 3rem);
		margin-left: -1.5rem;
		overflow: hidden;
		border-bottom-left-radius: 0;
		border-bottom-right-radius: 0;
	}

	@media (max-width: 640px) {
		.worktile-card-image-wrapper {
			margin-top: -1rem;
			width: calc(100% + 2rem);
			margin-left: -1rem;
		}
	}

	.worktile-card-image-wrapper img {
		object-fit: cover;
		height: 300px;
		width: 300px;
		flex-shrink: 0;
		overflow: hidden;
		transition: transform 1s ease;
	}

	.worktile-card-text {
		margin-top: 1rem;
	}

	@media (min-width: 2140px) {
		.worktile-card {
			width: 400px;
		}
		.worktile-card-image-wrapper img {
			height: 400px;
			width: 400px;
		}
	}
</style>
