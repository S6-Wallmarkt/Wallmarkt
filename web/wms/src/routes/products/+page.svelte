<script>
	export let data;
</script>

<svelte:head>
	<title>WMS - Products</title>
</svelte:head>

<div class="p-4">
	<div>
		{#if data}
			{#each data.props.products as product}
				<div class="card p-4 m-4 space-y-2 variant-glass-primary">
					<h2 class="">{product.name}</h2>
					<p>{product.description}</p>
					<p>€ {product.price}</p>
					<div class="space-x-2">
						{#each product.types as type}
							<span class="chip variant-filled-secondary">{type}</span>
						{/each}
					</div>
				</div>
			{/each}
		{:else}
			<p>No products found</p>
		{/if}
	</div>
</div>
