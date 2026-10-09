import { solve } from "../src/production.js";

try {
  const plan = await solve(24, 16);

  console.log("Pete's Pottery Paradise production plan");
  console.log(`Small vases: ${plan.smallVases}`);
  console.log(`Large vases: ${plan.largeVases}`);
  console.log(`Clay used: ${plan.clayUsedOz} oz.`);
  console.log(`Glaze used: ${plan.glazeUsedOz} oz.`);
  console.log(`Maximum revenue: $${plan.revenue}.00`);
} catch (error) {
  console.error(error);
  process.exitCode = 1;
}
