<script lang="ts">
	import '../app.postcss';

	// Floating UI for Popups
	import { computePosition, autoUpdate, flip, shift, offset, arrow } from '@floating-ui/dom';
	import { storePopup } from '@skeletonlabs/skeleton';
	storePopup.set({ computePosition, autoUpdate, flip, shift, offset, arrow });

	// Auth logic
	import { onMount } from 'svelte';
	import auth from '$lib/utils/authService';
	import { isAuthenticated, user } from '$lib/stores/auth';
	import type { Auth0Client } from '@auth0/auth0-spa-js';

	let auth0Client: Auth0Client;

	onMount(async () => {
		auth0Client = await auth.createClient();
		isAuthenticated.set(await auth0Client.isAuthenticated());
		const authUser = await auth0Client.getUser();
		if (authUser) {
			user.set(authUser);
			console.log(authUser);
		} else {
			console.log('Not logged in');
		}
	});

	function login() {
		auth.loginWithPopup(auth0Client);
	}

	function logout() {
		auth.logout(auth0Client);
	}
	// Import components
	import Nav from '$lib/components/Nav.svelte';
</script>

<Nav {login} {logout} />
<slot />
