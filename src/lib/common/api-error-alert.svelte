<script lang="ts">
	import { Alert, Li, List } from 'flowbite-svelte';
	import type { ApiError } from '../../models/common/validation';

	let {
		apiErrors,
		message = 'Change a few things up and try submitting again.'
	}: { apiErrors: ApiError; message?: string } = $props();
</script>

{#if apiErrors && apiErrors.errors}
	<Alert color="red">
		<strong>{message}</strong>
		<List list="disc" class="ps-3 pt-3">
			{#each apiErrors.errors as error}
				<Li>{error}</Li>
			{/each}
		</List>
	</Alert>
{/if}
