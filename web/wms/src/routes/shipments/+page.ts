import { accessToken } from '$lib/stores/auth';
import { redirect } from '@sveltejs/kit';

export async function load({ fetch }) {
	let token = '';
	accessToken.subscribe((value) => {
		token = value;
	})();

	try {
		// Check if the token is available before making the request
		if (!token) {
			throw new Error('Access token is not available');
		}

		// Include the access token in the Authorization header
		const response = await fetch(import.meta.env.VITE_GATEWAY_BASE + '/shipping/getall', {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});

		if (!response.ok) {
			throw new Error('Network response was not ok');
		}

		const shipments = await response.json();
		return { props: { shipments } };
	} catch (error) {
		console.error('Failed to fetch shipments:', error);
		redirect(302, '/');
	}
}
