<script>
	import { LightSwitch, TabGroup, Tab, TabAnchor } from '@skeletonlabs/skeleton';
	import { page } from '$app/stores';

	let tabs = [
		{ name: 'Home', href: '/' },
		{ name: 'Products', href: '/products' }
	];
</script>

<div class="bg-surface-100-800-token flex items-center p-4 space-x-4">
	<TabGroup
		active="variant-ghost-primary"
		hover="hover:variant-soft-primary"
		flex="flex-1 lg:flex-none"
		rounded=""
		border=""
		class="w-full"
	>
		{#each tabs as tab}
			<TabAnchor href={tab.href} selected={$page.url.pathname === tab.href}>{tab.name}</TabAnchor>
		{/each}
	</TabGroup>
	<LightSwitch />
</div>
