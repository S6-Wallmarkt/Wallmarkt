<script>
	export let data;
</script>

<svelte:head>
	<title>WMS - Shipments</title>
</svelte:head>

<div class="p-4">
	<div>
		{#if data}
			{#each data.props.shipments as shipment}
				<div class="card p-4 m-4 space-y-2 variant-glass-primary">
					<h2 class="">{shipment.order}</h2>

					<p>{shipment.timestamp}</p>
					<div class="space-x-2">
						{#if shipment.send}
							<span class="chip variant-filled-success">Send</span>
						{:else}
							<span class="chip variant-filled-error">Send</span>
						{/if}
					</div>
					{#if shipment.send}
						<p>{shipment.packer}</p>
					{/if}
				</div>
			{/each}
		{:else}
			<p>No products found</p>
		{/if}
	</div>
</div>
