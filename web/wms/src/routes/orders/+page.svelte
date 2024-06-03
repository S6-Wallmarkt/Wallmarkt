<script>
	export let data;
</script>

<svelte:head>
	<title>WMS - Orders</title>
</svelte:head>

<div class="p-4">
	<div>
		{#if data}
			{#each data.props.orders as order}
				<div class="card p-4 m-4 space-y-2 variant-glass-primary">
					<h2 class="">{order.customer_id}</h2>
					{#each order.products as product}
						<p>{product}</p>
					{/each}
					<div class="space-x-2">
						{#if order.payed}
							<span class="chip variant-filled-success">Payed</span>
						{:else}
							<span class="chip variant-filled-error">Payed</span>
						{/if}
						{#if order.shipped}
							<span class="chip variant-filled-success">Shipped</span>
						{:else}
							<span class="chip variant-filled-error">Shipped</span>
						{/if}
					</div>
				</div>
			{/each}
		{:else}
			<p>No products found</p>
		{/if}
	</div>
</div>
