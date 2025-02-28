import { expect, test } from "@playwright/test";
import { license, users } from "../constants";
import { login } from "../helpers";
import { beforeCoderTest } from "../hooks";

test.beforeEach(async ({ page }) => {
	beforeCoderTest(page);
});

test.describe("deployment settings access", () => {
	test("member cannot see deployment settings", async ({ page }) => {
		await login(page, users.member);
		await page.goto("/", { waitUntil: "domcontentloaded" });

		await expect(
			page.getByRole("button", { name: "Admin settings" }),
		).not.toBeVisible();
	});

	test("template admin can see deployment settings", async ({ page }) => {
		await login(page, users.templateAdmin);
		await page.goto("/", { waitUntil: "domcontentloaded" });

		await expect(
			page.getByRole("button", {
				name: "Admin settings",
			}),
		).toBeVisible();
	});

	test("user admin can see deployment settings", async ({ page }) => {
		await login(page, users.userAdmin);
		await page.goto("/", { waitUntil: "domcontentloaded" });

		await expect(
			page.getByRole("button", { name: "Admin settings" }),
		).toBeVisible();
	});

	test("auditor can see deployment settings", async ({ page }) => {
		await login(page, users.auditor);
		await page.goto("/", { waitUntil: "domcontentloaded" });

		await expect(
			page.getByRole("button", { name: "Admin settings" }),
		).toBeVisible();
	});

	test("admin can see deployment settings", async ({ page }) => {
		await login(page, users.admin);
		await page.goto("/", { waitUntil: "domcontentloaded" });

		await expect(
			page.getByRole("button", { name: "Admin settings" }),
		).toBeVisible();
	});
});
