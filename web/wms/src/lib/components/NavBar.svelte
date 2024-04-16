<script>
	import { page } from '$app/stores';
	import DarkModeToggle from './DarkModeToggle.svelte';

	let menuItems = [
		{ name: 'Catalog', href: '/catalog' },
		{ name: 'Stock', href: '/stock' },
		{ name: 'Orders', href: '/orders' }
	];

	$: currentPath = $page.url.pathname;
</script>

<nav class="bg-primary">
	<div class="container mx-auto px-4">
		<div class="flex justify-between items-center py-4">
			<div class="flex items-center">
				<a href="/" class="text-lg font-bold text-gray-800">WMS</a>
				<ul class="flex ml-8 space-x-8 text-text">
					{#each menuItems as item}
						<li class={currentPath === item.href ? 'bg-background rounded p-1' : 'p-1'}>
							<a href={item.href}>{item.name}</a>
						</li>
					{/each}
				</ul>
			</div>
			<div class="flex items-center space-x-8">
				<a href="/login" class="text-gray-800">Login</a>
				<DarkModeToggle />
			</div>
		</div>
	</div>
</nav>
