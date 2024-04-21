<script>
	function toggleDarkMode() {
		document.documentElement.classList.toggle('dark');
	}
</script>

<!-- Write a darkmode toggle button in tailwind, it should contain a -->
<button on:click={toggleDarkMode} class="p-2 bg-gray-800 text-white rounded">
	Toggle Dark Mode
</button>
