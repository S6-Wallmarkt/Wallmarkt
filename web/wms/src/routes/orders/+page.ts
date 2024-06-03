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
		const response = await fetch(import.meta.env.VITE_GATEWAY_BASE + '/order/getall', {
			headers: {
				Authorization: `Bearer ${token}`
			}
		});

		if (!response.ok) {
			throw new Error('Network response was not ok');
		}

		const orders = await response.json();
		return { props: { orders } };
	} catch (error) {
		console.error('Failed to fetch orders:', error);
		redirect(302, '/');
	}
}
