<script lang="ts">
	import { goto } from '$app/navigation';
	import { Pagination } from '$lib';
	import { FolderArrowRightSolid } from 'flowbite-svelte-icons';
	import type { Folder } from '../../../models';
	import type PaginationData from '../../../models/common/pagination-data';

	let { data }: { data: { records: Folder[]; pagination: PaginationData } } = $props();
</script>

<div class="container mx-auto">
	{#if data && data.records.length > 0}
		<div class="flex flex-wrap justify-center">
			{#each data.records as folder}
				<div class="m-5 flex h-48 w-56 items-center justify-center rounded-md bg-white shadow-md">
					<p
						class="flex flex-col bg-gradient-to-r from-indigo-600 via-pink-600 to-purple-600 bg-clip-text text-2xl font-bold text-transparent"
					>
						<FolderArrowRightSolid class="h-24 w-24 text-orange-300" />
						<span class="text-center">{folder.name}</span>
					</p>
				</div>
			{/each}
		</div>
	{:else}
		No folders found
	{/if}
	{#if data}
		<Pagination
			data={data.pagination}
			showPageNumbers={false}
			update={(page) => goto(`/folders/${page}`)}
		/>
	{/if}
</div>
