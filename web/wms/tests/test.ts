import { test, expect } from '@playwright/test';

test('User can log in using Auth0 SPA SDK and see menu', async ({ page }) => {
	await page.goto('http://localhost:5173/');
	await page.waitForTimeout(1000);
	const page1Promise = page.waitForEvent('popup');
	await page.getByRole('button', { name: 'Login' }).click();
	const page1 = await page1Promise;
	await page1.getByLabel('Email address*').click();
	await page1.getByLabel('Password*').click();
	await page1.getByLabel('Password*').fill('xjUZ_AA4Aqd.KDW');
	// await page1.getByLabel('Password*').fill(process.env.PLAYWRIGHT_PASSWORD);
	await page1.getByLabel('Email address*').click();
	await page1.getByLabel('Email address*').fill('yaseg80297@qiradio.com');
	// await page1.getByLabel('Email address*').fill(process.env.PLAYWRIGHT_EMAIL);
	await page1.getByRole('button', { name: 'Continue', exact: true }).click();
	await page.waitForTimeout(5000);
	await expect(await page.getByRole('img', { name: 'User Avatar' })).toBeEnabled();
	await expect(await page.getByRole('link', { name: 'Home' })).toBeEnabled();
});
