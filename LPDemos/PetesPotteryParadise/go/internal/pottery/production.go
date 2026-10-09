package pottery

import (
	"fmt"

	"github.com/google/or-tools/ortools/sat/go/cpmodel"
	cmpb "github.com/google/or-tools/ortools/sat/proto/cpmodel"
)

const (
	smallVaseClayOz  = int64(1)
	smallVaseGlazeOz = int64(1)
	smallVaseRevenue = int64(3)

	largeVaseClayOz  = int64(4)
	largeVaseGlazeOz = int64(2)
	largeVaseRevenue = int64(9)
)

// ProductionPlan is the optimal integer production plan and its resource use.
type ProductionPlan struct {
	SmallVases  int64
	LargeVases  int64
	ClayUsedOz  int64
	GlazeUsedOz int64
	Revenue     int64
}

// Solve returns the revenue-maximizing production plan for Pete's Pottery
// Paradise given the available clay and glaze inventory, in ounces.
func Solve(clayInventoryOz, glazeInventoryOz int64) (ProductionPlan, error) {
	if clayInventoryOz < 0 || glazeInventoryOz < 0 {
		return ProductionPlan{}, fmt.Errorf("clay and glaze inventory must be non-negative")
	}

	model := cpmodel.NewCpModelBuilder()
	smallVases := model.NewIntVar(0, clayInventoryOz).WithName("small_vases")
	largeVases := model.NewIntVar(0, clayInventoryOz/largeVaseClayOz).WithName("large_vases")

	clayUsed := cpmodel.NewLinearExpr().
		AddTerm(smallVases, smallVaseClayOz).
		AddTerm(largeVases, largeVaseClayOz)
	glazeUsed := cpmodel.NewLinearExpr().
		AddTerm(smallVases, smallVaseGlazeOz).
		AddTerm(largeVases, largeVaseGlazeOz)

	model.AddLessOrEqual(clayUsed, cpmodel.NewConstant(clayInventoryOz))
	model.AddLessOrEqual(glazeUsed, cpmodel.NewConstant(glazeInventoryOz))

	revenue := cpmodel.NewLinearExpr().
		AddTerm(smallVases, smallVaseRevenue).
		AddTerm(largeVases, largeVaseRevenue)
	model.Maximize(revenue)

	cpModel, err := model.Model()
	if err != nil {
		return ProductionPlan{}, fmt.Errorf("build production model: %w", err)
	}

	response, err := cpmodel.SolveCpModel(cpModel)
	if err != nil {
		return ProductionPlan{}, fmt.Errorf("solve production model: %w", err)
	}
	if response.GetStatus() != cmpb.CpSolverStatus_OPTIMAL {
		return ProductionPlan{}, fmt.Errorf(
			"CP-SAT did not prove an optimum (status %v)",
			response.GetStatus(),
		)
	}

	smallCount := cpmodel.SolutionIntegerValue(response, smallVases)
	largeCount := cpmodel.SolutionIntegerValue(response, largeVases)

	return ProductionPlan{
		SmallVases:  smallCount,
		LargeVases:  largeCount,
		ClayUsedOz:  smallCount*smallVaseClayOz + largeCount*largeVaseClayOz,
		GlazeUsedOz: smallCount*smallVaseGlazeOz + largeCount*largeVaseGlazeOz,
		Revenue:     smallCount*smallVaseRevenue + largeCount*largeVaseRevenue,
	}, nil
}
