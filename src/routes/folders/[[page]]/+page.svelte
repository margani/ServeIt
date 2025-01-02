<script lang="ts">
	import { goto } from '$app/navigation';
	import { Pagination } from '$lib';
	import type { Folder } from '../../../models';
	import type PaginationData from '../../../models/common/pagination-data';
	import {
		Table,
		TableBody,
		TableBodyCell,
		TableBodyRow,
		TableHead,
		TableHeadCell
	} from 'flowbite-svelte';

	let { data }: { data: { records: Folder[]; pagination: PaginationData } } = $props();
</script>

<div class="container mx-auto">
	<Table striped={true} hoverable={true} shadow class="w-full">
		<TableHead theadClass="text-xs uppercase">
			<TableHeadCell class="w-full">Name</TableHeadCell>
		</TableHead>
		<TableBody>
			{#if data && data.records.length > 0}
				{#each data.records as folder}
					<TableBodyRow>
						<TableBodyCell>{folder.name}</TableBodyCell>
					</TableBodyRow>
				{/each}
			{:else}
				<TableBodyRow>
					<TableBodyCell colspan={2}>No folders found</TableBodyCell>
				</TableBodyRow>
			{/if}
		</TableBody>
	</Table>
	{#if data}
		<Pagination
			data={data.pagination}
			showPageNumbers={false}
			update={(page) => goto(`/folders/${page}`)}
		/>
	{/if}
</div>
