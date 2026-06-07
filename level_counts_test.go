package priorityframeworks

import (
	"testing"
)

func TestNewLevelCounts(t *testing.T) {
	lc := NewLevelCounts(Severity())

	// All levels should be initialized to 0
	for _, level := range Severity().Levels {
		if got := lc.Get(level.ID); got != 0 {
			t.Errorf("Get(%s) = %d, want 0", level.ID, got)
		}
	}

	if !lc.IsEmpty() {
		t.Error("IsEmpty() = false, want true")
	}
}

func TestLevelCountsAddAndIncrement(t *testing.T) {
	lc := NewLevelCounts(Severity())

	lc.Increment("critical")
	lc.Increment("critical")
	lc.Add("high", 5)
	lc.Add("medium", 10)

	if got := lc.Get("critical"); got != 2 {
		t.Errorf("Get(critical) = %d, want 2", got)
	}
	if got := lc.Get("high"); got != 5 {
		t.Errorf("Get(high) = %d, want 5", got)
	}
	if got := lc.Get("medium"); got != 10 {
		t.Errorf("Get(medium) = %d, want 10", got)
	}
	if got := lc.Get("low"); got != 0 {
		t.Errorf("Get(low) = %d, want 0", got)
	}
}

func TestLevelCountsTotal(t *testing.T) {
	lc := NewLevelCounts(Severity())

	lc.Add("critical", 1)
	lc.Add("high", 2)
	lc.Add("medium", 3)
	lc.Add("low", 4)
	lc.Add("informational", 5)

	if got := lc.Total(); got != 15 {
		t.Errorf("Total() = %d, want 15", got)
	}
}

func TestLevelCountsActionableTotal(t *testing.T) {
	lc := NewLevelCounts(Severity())

	lc.Add("critical", 1)      // actionable
	lc.Add("high", 2)          // actionable
	lc.Add("medium", 3)        // actionable
	lc.Add("low", 4)           // actionable
	lc.Add("informational", 5) // NOT actionable

	// Critical, High, Medium, Low are actionable; Informational is not
	if got := lc.ActionableTotal(); got != 10 {
		t.Errorf("ActionableTotal() = %d, want 10", got)
	}
}

func TestLevelCountsHigherThan(t *testing.T) {
	lc := NewLevelCounts(Severity())

	lc.Add("critical", 1)
	lc.Add("high", 2)
	lc.Add("medium", 3)
	lc.Add("low", 4)
	lc.Add("informational", 5)

	// Higher than medium = critical (1) + high (2) = 3
	if got := lc.HigherThan("medium"); got != 3 {
		t.Errorf("HigherThan(medium) = %d, want 3", got)
	}

	// Higher than or equal to medium = critical (1) + high (2) + medium (3) = 6
	if got := lc.HigherThanOrEqual("medium"); got != 6 {
		t.Errorf("HigherThanOrEqual(medium) = %d, want 6", got)
	}

	// Lower than medium = low (4) + informational (5) = 9
	if got := lc.LowerThan("medium"); got != 9 {
		t.Errorf("LowerThan(medium) = %d, want 9", got)
	}
}

func TestLevelCountsMerge(t *testing.T) {
	lc1 := NewLevelCounts(Severity())
	lc1.Add("critical", 1)
	lc1.Add("high", 2)

	lc2 := NewLevelCounts(Severity())
	lc2.Add("critical", 3)
	lc2.Add("medium", 4)

	lc1.Merge(lc2)

	if got := lc1.Get("critical"); got != 4 {
		t.Errorf("Get(critical) after merge = %d, want 4", got)
	}
	if got := lc1.Get("high"); got != 2 {
		t.Errorf("Get(high) after merge = %d, want 2", got)
	}
	if got := lc1.Get("medium"); got != 4 {
		t.Errorf("Get(medium) after merge = %d, want 4", got)
	}
}

func TestLevelCountsClone(t *testing.T) {
	lc := NewLevelCounts(Severity())
	lc.Add("critical", 5)

	clone := lc.Clone()

	// Modify original
	lc.Add("critical", 10)

	// Clone should be unchanged
	if got := clone.Get("critical"); got != 5 {
		t.Errorf("Clone Get(critical) = %d, want 5 (original was modified)", got)
	}
}

func TestLevelCountsReset(t *testing.T) {
	lc := NewLevelCounts(Severity())
	lc.Add("critical", 5)
	lc.Add("high", 10)

	lc.Reset()

	if !lc.IsEmpty() {
		t.Error("IsEmpty() after Reset() = false, want true")
	}
}

func TestLevelCountsSlice(t *testing.T) {
	lc := NewLevelCounts(Severity())
	lc.Add("critical", 1)
	lc.Add("high", 2)
	lc.Add("medium", 3)
	lc.Add("low", 4)
	lc.Add("informational", 5)

	slice := lc.Slice()
	expected := []int{1, 2, 3, 4, 5}

	if len(slice) != len(expected) {
		t.Fatalf("Slice() length = %d, want %d", len(slice), len(expected))
	}
	for i, v := range expected {
		if slice[i] != v {
			t.Errorf("Slice()[%d] = %d, want %d", i, slice[i], v)
		}
	}
}

func TestLevelCountsNilFramework(t *testing.T) {
	lc := NewLevelCounts(nil)

	// These should return 0 without panic
	if got := lc.ActionableTotal(); got != 0 {
		t.Errorf("ActionableTotal() with nil framework = %d, want 0", got)
	}
	if got := lc.HigherThan("medium"); got != 0 {
		t.Errorf("HigherThan() with nil framework = %d, want 0", got)
	}
	if slice := lc.Slice(); slice != nil {
		t.Errorf("Slice() with nil framework = %v, want nil", slice)
	}
}
