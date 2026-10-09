import {
  CpModel,
  CpSat,
  CpSolverStatus,
  setCloudNoticeEnabled,
} from "or-tools-wasm/cp-sat";

setCloudNoticeEnabled(false);

const smallVaseClayOz = 1;
const smallVaseGlazeOz = 1;
const smallVaseRevenue = 3;
const largeVaseClayOz = 4;
const largeVaseGlazeOz = 2;
const largeVaseRevenue = 9;

function validateInventory(value, name) {
  if (!Number.isSafeInteger(value) || value < 0) {
    throw new RangeError(`${name} inventory must be a non-negative safe integer`);
  }
}

function toSafeNumber(value, name) {
  const number = Number(value);
  if (!Number.isSafeInteger(number)) {
    throw new RangeError(`${name} exceeds JavaScript's safe integer range`);
  }
  return number;
}

export async function solve(clayInventoryOz, glazeInventoryOz) {
  validateInventory(clayInventoryOz, "clay");
  validateInventory(glazeInventoryOz, "glaze");

  const model = new CpModel();
  const smallVases = model.newIntVar(0, clayInventoryOz, "small_vases");
  const largeVases = model.newIntVar(
    0,
    Math.floor(clayInventoryOz / largeVaseClayOz),
    "large_vases",
  );

  model.addLinearConstraint(
    smallVases.times(smallVaseClayOz).plus(largeVases.times(largeVaseClayOz)),
    0,
    clayInventoryOz,
  );
  model.addLinearConstraint(
    smallVases.times(smallVaseGlazeOz).plus(largeVases.times(largeVaseGlazeOz)),
    0,
    glazeInventoryOz,
  );
  model.maximize(
    smallVases.times(smallVaseRevenue).plus(largeVases.times(largeVaseRevenue)),
  );

  const result = await CpSat.solve(model, { numWorkers: 1 });
  if (
    !result.hasSolution ||
    (result.status !== CpSolverStatus.OPTIMAL && result.status !== "OPTIMAL")
  ) {
    throw new Error(`CP-SAT did not prove an optimum (status ${result.status})`);
  }

  const smallCount = toSafeNumber(result.value(smallVases), "small vase quantity");
  const largeCount = toSafeNumber(result.value(largeVases), "large vase quantity");
  const clayUsedOz = smallCount * smallVaseClayOz + largeCount * largeVaseClayOz;
  const glazeUsedOz = smallCount * smallVaseGlazeOz + largeCount * largeVaseGlazeOz;
  const revenue = smallCount * smallVaseRevenue + largeCount * largeVaseRevenue;

  return {
    smallVases: smallCount,
    largeVases: largeCount,
    clayUsedOz: toSafeNumber(clayUsedOz, "clay used"),
    glazeUsedOz: toSafeNumber(glazeUsedOz, "glaze used"),
    revenue: toSafeNumber(revenue, "revenue"),
  };
}
