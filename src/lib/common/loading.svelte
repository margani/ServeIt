<script lang="ts">
	import { loading } from '$lib';

	$: if ($loading.status === 'NAVIGATING') {
		setTimeout(() => {
			if ($loading.status === 'NAVIGATING') {
				$loading.status = 'LOADING';
			}
		}, 400);
	}
</script>

{#if $loading.status === 'LOADING'}
	<div>
		<svg width="24" height="24" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
			<style>
			</style>
			<circle fill="#14b8a6" class="circle" cx="4" cy="12" r="3" />
			<circle fill="#14b8a6" class="circle circle1" cx="4" cy="12" r="3" />
			<circle fill="#14b8a6" class="circle circle2" cx="4" cy="12" r="3" />
			<circle fill="#14b8a6" class="circle circle3" cx="4" cy="12" r="3" />
		</svg>
	</div>
{/if}

<style>
	div {
		--size: 24px;
		position: absolute;
		top: -25px;
		right: 15px;
		height: var(--size);
		width: var(--size);
	}

	.circle {
		animation: spinner 2s cubic-bezier(0.36, 0.6, 0.31, 1) infinite;
	}
	.circle1 {
		animation-delay: -0.5s;
	}
	.circle2 {
		animation-delay: -1s;
	}
	.circle3 {
		animation-delay: -1.5s;
	}

	@keyframes spinner {
		0% {
			r: 0;
		}
		25% {
			r: 3px;
			cx: 4px;
		}
		50% {
			r: 3px;
			cx: 12px;
		}
		75% {
			r: 3px;
			cx: 20px;
		}
		100% {
			r: 0;
			cx: 20px;
		}
	}
</style>
