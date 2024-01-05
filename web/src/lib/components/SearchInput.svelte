<script lang="ts">
	import { createEventDispatcher } from 'svelte';
	import { debounce } from '$lib/util';
	import { Input, Button } from 'flowbite-svelte';
	import { _ } from 'svelte-i18n';
	import { CloseSolid } from 'flowbite-svelte-icons';

	const dispatch = createEventDispatcher();

	let search = '';

	export let placeholder = '';

	const debounceSearchChange = debounce(() => dispatch('search', search));
</script>

<div class="search">
	<span class="search-input">
		<Input size="lg" {placeholder} bind:value={search} on:input={debounceSearchChange} />
		{#if search}
			<Button
				title={$_('search.reset.title')}
				on:click={() => {
					search = '';
					dispatch('search', search);
				}}
				class="search-reset-button rounded-lg p-2.5 text-sm text-gray-500 hover:bg-gray-100 focus:outline-none dark:text-gray-400 dark:hover:bg-gray-800"
			>
				<CloseSolid size="sm" class="m-1 inline text-red-700 dark:text-green-300" />
			</Button>
		{/if}
	</span>
</div>

<style>
	.search {
		display: flex;
		justify-content: center;
		margin: 20px 0;
	}
	.search-input {
		width: 500px;
		display: flex;
		flex-direction: row;
		align-items: center;
		justify-content: flex-end;
	}

	@media (max-width: 720px) {
		.search-input {
			width: 300px;
		}
	}

	:global(.search-reset-button) {
		position: absolute;
		background-color: transparent !important;
		margin-right: 5px;
	}

	:global(.search-reset-button:focus-within) {
		--tw-ring-opacity: 0 !important;
	}
</style>
