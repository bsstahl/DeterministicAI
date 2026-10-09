import assert from "node:assert/strict";
import test from "node:test";
import { solve } from "./production.js";

test("solve returns the revenue-maximizing production plan", async () => {
  assert.deepEqual(await solve(24, 16), {
    smallVases: 8,
    largeVases: 4,
    clayUsedOz: 24,
    glazeUsedOz: 16,
    revenue: 60,
  });
});

test("solve handles inventories and outputs scaled by 10 million", async () => {
  assert.deepEqual(await solve(240_000_000, 160_000_000), {
    smallVases: 80_000_000,
    largeVases: 40_000_000,
    clayUsedOz: 240_000_000,
    glazeUsedOz: 160_000_000,
    revenue: 600_000_000,
  });
});

test("solve uses the provided inventory", async () => {
  assert.deepEqual(await solve(24, 8), {
    smallVases: 0,
    largeVases: 4,
    clayUsedOz: 16,
    glazeUsedOz: 8,
    revenue: 36,
  });
});

test("solve rejects negative inventory", async () => {
  await assert.rejects(solve(-1, 16), {
    name: "RangeError",
    message: "clay inventory must be a non-negative safe integer",
  });
});

test("solve rejects non-integer inventory", async () => {
  await assert.rejects(solve(24.5, 16), {
    name: "RangeError",
    message: "clay inventory must be a non-negative safe integer",
  });
});
