package pottery

import "testing"

func TestSolveReturnsRevenueMaximizingPlan(t *testing.T) {
	got, err := Solve(24, 16)
	if err != nil {
		t.Fatalf("Solve() error = %v", err)
	}

	want := ProductionPlan{
		SmallVases:  8,
		LargeVases:  4,
		ClayUsedOz:  24,
		GlazeUsedOz: 16,
		Revenue:     60,
	}
	if got != want {
		t.Fatalf("Solve() = %+v, want %+v", got, want)
	}
}

func TestSolveReturnsRevenueMaximizingPlanTimes100K(t *testing.T) {
	got, err := Solve(24000000, 16000000)
	if err != nil {
		t.Fatalf("Solve() error = %v", err)
	}

	want := ProductionPlan{
		SmallVases:  8000000,
		LargeVases:  4000000,
		ClayUsedOz:  24000000,
		GlazeUsedOz: 16000000,
		Revenue:     60000000,
	}
	if got != want {
		t.Fatalf("Solve() = %+v, want %+v", got, want)
	}
}

func TestSolveUsesProvidedInventory(t *testing.T) {
	got, err := Solve(24, 8)
	if err != nil {
		t.Fatalf("Solve(24, 8) error = %v", err)
	}

	want := ProductionPlan{
		LargeVases:  4,
		ClayUsedOz:  16,
		GlazeUsedOz: 8,
		Revenue:     36,
	}
	if got != want {
		t.Fatalf("Solve(24, 8) = %+v, want %+v", got, want)
	}
}

func TestSolveRejectsNegativeInventory(t *testing.T) {
	_, err := Solve(-1, 16)
	if err == nil {
		t.Fatal("Solve(-1, 16) error = nil, want an error")
	}
}
