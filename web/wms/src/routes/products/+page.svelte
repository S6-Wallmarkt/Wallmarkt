<script lang="ts">
	import { user } from '$lib/stores/auth';
</script>

<div class="p-4">
	<h1>This is the Products page</h1>
	{#if $user}
		<p>Welcome, {$user?.name}</p>
	{:else}
		<p>You are not logged in</p>
	{/if}
</div>
