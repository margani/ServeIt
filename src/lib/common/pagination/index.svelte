<script lang="ts">
	import { Button, ButtonGroup } from 'flowbite-svelte';
	import type PaginationData from '../../../models/common/pagination-data';
	import { ChevronLeftOutline, ChevronRightOutline } from 'flowbite-svelte-icons';

	interface Props {
		data: PaginationData;
		showPageNumbers: boolean;
		update: (page: number) => void;
	}

	let { data, showPageNumbers = true, update }: Props = $props();

	function gotoPrevious() {
		if (data.current > 1) {
			goto(data.current - 1);
		}
	}

	function gotoNext() {
		if (data.current < data.totalPages) {
			goto(data.current + 1);
		}
	}

	function goto(page: number) {
		data.current = page;
		update(page);
	}
</script>

{#if data && data.totalPages > 1}
	<div class="my-2 flex flex-col items-center gap-3 md:flex-row">
		<div class="md:w-full">
			Page {data.current} of {data.totalPages}, of {data.totalItems} records.
		</div>
		{#if data.isVisible}
			<ButtonGroup>
				<Button on:click={gotoPrevious} disabled={!data.isPreviousEnabled}>
					<ChevronLeftOutline />
				</Button>
				{#if showPageNumbers}
					{#each { length: data.totalPages } as _, page}
						{#if page + 1 === data.current}
							<Button color="primary">{page + 1}</Button>
						{:else}
							<Button on:click={() => goto(page + 1)}>{page + 1}</Button>
						{/if}
					{/each}
				{/if}
				<Button on:click={gotoNext} disabled={!data.isNextEnabled}>
					<ChevronRightOutline />
				</Button>
			</ButtonGroup>
		{/if}
	</div>
{/if}
