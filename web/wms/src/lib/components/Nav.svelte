<script lang="ts">
	import { LightSwitch, Avatar, TabGroup, TabAnchor } from '@skeletonlabs/skeleton';
	import { page } from '$app/stores';
	import { isAuthenticated, user } from '$lib/stores/auth';
	export let login: () => void;
	export let logout: () => void;

	let tabs = [
		{ name: 'Home', href: '/' },
		{ name: 'Products', href: '/products' }
	];
</script>

<div class="bg-surface-100-800-token flex items-center p-4 space-x-8">
	<TabGroup
		active="variant-ghost-primary"
		hover="hover:variant-soft-primary"
		flex="flex-1 lg:flex-none"
		rounded=""
		border=""
		class="w-full"
	>
		{#each tabs as tab}
			{#if $isAuthenticated && $user?.userroles.includes('Warehouse-Operator')}
				<TabAnchor href={tab.href} selected={$page.url.pathname === tab.href}>{tab.name}</TabAnchor>
			{/if}
		{/each}
	</TabGroup>
	<div>
		<LightSwitch />
	</div>
	{#if $isAuthenticated && $user?.picture}
		<Avatar src={$user.picture} alt="User Avatar" class="border-2 border-primary-600 select-none" />
	{/if}
	<div>
		{#if $isAuthenticated}
			<button on:click={logout} class="btn variant-filled-primary">Logout</button>
		{:else}
			<button on:click={login} class="btn variant-filled-primary">Login</button>
		{/if}
	</div>
</div>
