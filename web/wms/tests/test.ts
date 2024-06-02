import { test, expect } from '@playwright/test';

test.beforeEach(async ({ page }) => {
	await page.goto('http://localhost:5173/');
	await page.waitForTimeout(1000);
	const page1Promise = page.waitForEvent('popup');
	await page.getByRole('button', { name: 'Login' }).click();
	const page1 = await page1Promise;
	await page1.getByLabel('Email address*').click();
	await page1.getByLabel('Password*').click();
	await page1.getByLabel('Password*').fill(process.env.PLAYWRIGHT_PASSWORD || '');
	await page1.getByLabel('Email address*').click();
	await page1.getByLabel('Email address*').fill(process.env.PLAYWRIGHT_USERNAME || '');
	await page1.getByRole('button', { name: 'Continue', exact: true }).click();
	await page.waitForTimeout(1000);
});

test('User can log in using Auth0 SPA SDK and see menu', async ({ page }) => {
	await expect(await page.getByRole('img', { name: 'User Avatar' })).toBeEnabled();
	await expect(await page.getByRole('link', { name: 'Home' })).toBeEnabled();
});

test('User can see products', async ({ page }) => {
	await page.reload();
	await page.getByRole('link', { name: 'Products' }).click();
	await expect(await page.getByText('Minecraft painting Minecraft')).toBeEnabled();
	await expect(await page.getByText('Deer head mount Realistic')).toBeEnabled();
});

test('User can see orders', async ({ page }) => {
	await page.reload();
	await page.getByRole('link', { name: 'Orders' }).click();
	await expect(
		await page.getByText('facebook|12345 6651d10a02fb7f8a4414e17b Payed Shipped')
	).toBeEnabled();
	await expect(await page.locator('div').filter({ hasText: 'auth0|' }).nth(3)).toBeEnabled();
});

test('User can see shipments', async ({ page }) => {
	await page.reload();
	await page.getByRole('link', { name: 'Shipments' }).click();
	await expect(await page.getByText('6656f403ce0c27ef71d06983 2024')).toBeEnabled();
	await expect(await page.getByText('6656f34ace0c27ef71d06980 2024')).toBeEnabled();
});

test('User can logout and stay logged out', async ({ page }) => {
	await page.getByRole('button', { name: 'Logout' }).click();
	await page.waitForTimeout(1000);
	await expect(await page.locator('div').filter({ hasText: 'Login' }).nth(1)).toBeEnabled();
	await page.reload;
	await page.waitForTimeout(5000);
	await expect(await page.locator('div').filter({ hasText: 'Login' }).nth(1)).toBeEnabled();
});
