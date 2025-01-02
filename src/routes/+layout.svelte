<script lang="ts">
	import '../app.css';
	import '@fontsource/open-sans/400.css';
	import { browser, dev } from '$app/environment';
	import { navigating } from '$app/state';
	import Nav from '$lib/common/nav.svelte';
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
	<aside class="relative hidden h-screen w-64 bg-teal-600 shadow-xl sm:block">
		<div class="p-6">
			<a href="/" class="text-3xl font-semibold uppercase text-white hover:text-gray-300">
				Serve It
			</a>
		</div>
		<Nav mode="vertical-nav" />
	</aside>
	<div class="flex h-screen w-full flex-col overflow-y-hidden">
		<Nav mode="desktop-header" />
		<Nav mode="mobile-header-nav" />
		<div class="flex w-full flex-col overflow-x-hidden border-t">
			<main class="w-full flex-grow p-6">
				<div class="container mx-auto relative">
					<Loading />
					{@render children?.()}
				</div>
			</main>
		</div>
	</div>
</div>

<style>
	:global(body) {
		font-family: 'Open Sans', sans-serif;
	}
</style>
