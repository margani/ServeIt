<script lang="ts">
	import '../app.css';
	import '@fontsource/open-sans/400.css';
	import { browser, dev } from '$app/environment';
	import { navigating } from '$app/state';
	import { Loading, loading } from '$lib';
	import { onMount } from 'svelte';

	$effect(() => {
		loading.setNavigate(!!navigating.to);
	});

	interface Props {
		children?: import('svelte').Snippet;
	}

	let { children }: Props = $props();

	onMount(() => {
		loading.finished();
	});
</script>

<svelte:head>
	<title>Serve It App</title>
</svelte:head>

{#if dev && browser && window.navigator.userAgent.includes('Mobile')}
	<script src="//cdn.jsdelivr.net/npm/eruda"></script>
	<script>
		eruda.init();
	</script>
{/if}

<div class="flex bg-gray-100">
	<div class="flex h-screen w-full flex-col overflow-y-hidden">
		<main class="w-full flex-grow p-6">
			Serve It
			<div class="container relative mx-auto">
				<Loading />
				{@render children?.()}
			</div>
		</main>
	</div>
</div>

<style>
	:global(body) {
		font-family: 'Open Sans', sans-serif;
	}
</style>
