import { Auth0Client, createAuth0Client, type PopupLoginOptions } from '@auth0/auth0-spa-js';
import { user, isAuthenticated, popupOpen } from '$lib/stores/auth';

// Create Auth0 client
async function createClient() {
	const auth0Client = await createAuth0Client({
		domain: import.meta.env.VITE_AUTH0_DOMAIN,
		clientId: import.meta.env.VITE_AUTH0_CLIENT_ID,
		authorizationParams: {
			audience: import.meta.env.VITE_AUTH0_AUDIENCE
		}
	});
	return auth0Client;
}

// Login with popup
async function loginWithPopup(client: Auth0Client, options?: PopupLoginOptions | undefined) {
	popupOpen.set(true);
	try {
		await client.loginWithPopup(options);
		user.set((await client.getUser()) || {});
		isAuthenticated.set(true);
	} catch (e) {
		console.error(e);
	} finally {
		popupOpen.set(false);
	}
}

// Logout
function logout(client: Auth0Client) {
	return client.logout();
}

// Export functions
const auth = {
	createClient,
	loginWithPopup,
	logout
};
export default auth;
