import { writable } from 'svelte/store';
export const isAuthenticated = writable(false);
export const user = writable({} as any);
export const popupOpen = writable(false);
export const error = writable();
export const accessToken = writable('');
