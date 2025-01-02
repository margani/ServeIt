<script lang="ts">
	import {
		UserCircleSolid,
		BarsOutline,
		CloseOutline,
		HomeOutline,
		FolderDuplicateOutline,
		DatabaseOutline,
		CogOutline
	} from 'flowbite-svelte-icons';
	import { page } from '$app/stores';

	let { mode }: { mode: 'vertical-nav' | 'desktop-header' | 'mobile-header-nav' } = $props();

	let isOpen = $state(false);
</script>

{#if mode === 'vertical-nav'}
	<nav class="pt-3 text-base font-semibold text-white">
		<a
			href="/"
			class="flex items-center py-4 pl-6 text-white opacity-75 hover:opacity-100"
			class:active-nav-link={$page.url.pathname === '/'}
		>
			<HomeOutline class="me-2" /> Dashboard
		</a>
		<a
			href="/folders"
			class="flex items-center py-4 pl-6 text-white opacity-75 hover:opacity-100"
			class:active-nav-link={$page.url.pathname.split('/')[1] === 'folders'}
		>
			<FolderDuplicateOutline class="me-2" /> Folders
		</a>
	</nav>
	<a
		href={'#'}
		class="active-nav-link absolute bottom-0 flex w-full items-center justify-center bg-teal-700 py-4 text-white opacity-75 hover:opacity-100"
	>
		<CogOutline class="me-2" /> Settings
	</a>
{:else if mode === 'desktop-header'}
	<header class="hidden w-full items-center bg-white px-6 py-2 sm:flex">
		<div class="w-1/2"></div>
		<div class="relative flex w-1/2 justify-end">
			<button
				onclick={() => (isOpen = !isOpen)}
				class="realtive z-10 h-12 w-12 overflow-hidden opacity-75 hover:opacity-100"
			>
				<UserCircleSolid class="h-10 w-10 text-teal-600" />
			</button>
			<div
				class:hidden={!isOpen}
				class="absolute z-10 mt-16 w-32 rounded-lg bg-white py-2 shadow-lg"
			>
				<a href={'#'} class="block px-4 py-2 hover:bg-teal-700 hover:text-white">Account</a>
				<a href={'#'} class="block px-4 py-2 hover:bg-teal-700 hover:text-white">Support</a>
				<a href={'#'} class="block px-4 py-2 hover:bg-teal-700 hover:text-white">Sign Out</a>
			</div>
		</div>
	</header>
{:else if mode === 'mobile-header-nav'}
	<header class="w-full bg-teal-600 px-6 py-5 sm:hidden">
		<div class="flex items-center justify-between">
			<a href="/" class="text-3xl font-semibold uppercase text-white hover:text-gray-300">
				Serve It
			</a>
			<button onclick={() => (isOpen = !isOpen)} class="text-3xl text-white focus:outline-none">
				<span class:hidden={isOpen}><BarsOutline class="text-grey-900 h-10 w-10" /></span>
				<span class:hidden={!isOpen}><CloseOutline class="text-grey-900 h-10 w-10" /></span>
			</button>
		</div>

		<nav class="flex flex-col pt-4 {isOpen ? 'flex' : 'hidden'}">
			<a
				href="/"
				class="flex items-center py-2 pl-4 text-white opacity-75 hover:opacity-100"
				class:active-nav-link={$page.url.pathname === '/'}
			>
				<HomeOutline class="me-2" /> Dashboard
			</a>
			<a
				href="/folders"
				class="flex items-center py-2 pl-4 text-white opacity-75 hover:opacity-100"
				class:active-nav-link={$page.url.pathname.split('/')[1] === 'folders'}
			>
				<FolderDuplicateOutline class="me-2" /> Folders
			</a>
			<a href={'#'} class="flex items-center py-2 pl-4 text-white opacity-75 hover:opacity-100">
				Support
			</a>
			<a href={'#'} class="flex items-center py-2 pl-4 text-white opacity-75 hover:opacity-100">
				My Account
			</a>
			<a href={'#'} class="flex items-center py-2 pl-4 text-white opacity-75 hover:opacity-100">
				Sign Out
			</a>
			<button
				class="mt-3 flex w-full items-center justify-center rounded-lg bg-white py-2 font-semibold shadow-lg hover:bg-gray-300 hover:shadow-xl"
			>
				<CogOutline class="me-2" /> Settings
			</button>
		</nav>
	</header>
{/if}
