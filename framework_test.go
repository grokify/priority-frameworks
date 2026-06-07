package priorityframeworks

import "testing"

func TestFrameworkIndexOf(t *testing.T) {
	f := Severity()

	tests := []struct {
		input string
		want  int
	}{
		{"critical", 0},
		{"Critical", 0},
		{"CRITICAL", 0},
		{"Crit", 0},
		{"S1", 0},
		{"high", 1},
		{"medium", 2},
		{"low", 3},
		{"informational", 4},
		{"INFO", 4},
		{"notfound", -1},
	}

	for _, tt := range tests {
		got := f.IndexOf(tt.input)
		if got != tt.want {
			t.Errorf("IndexOf(%q) = %d, want %d", tt.input, got, tt.want)
		}
	}
}

func TestFrameworkParse(t *testing.T) {
	f := MoSCoW()

	level := f.Parse("Must")
	if level == nil {
		t.Fatal("Parse(Must) returned nil")
	}
	if level.ID != "must" {
		t.Errorf("Parse(Must).ID = %q, want %q", level.ID, "must")
	}
	if !level.Actionable {
		t.Error("Parse(Must).Actionable = false, want true")
	}

	level = f.Parse("notfound")
	if level != nil {
		t.Error("Parse(notfound) should return nil")
	}
}

func TestFrameworkCompare(t *testing.T) {
	f := Priority()

	tests := []struct {
		a, b string
		want int
	}{
		{"p0", "p1", 1},  // p0 > p1
		{"p1", "p0", -1}, // p1 < p0
		{"p2", "p2", 0},  // equal
		{"p0", "p4", 1},  // p0 > p4
		{"p4", "p0", -1}, // p4 < p0
	}

	for _, tt := range tests {
		got := f.Compare(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("Compare(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}

func TestFrameworkHighestLowest(t *testing.T) {
	// Test IETF requirements framework
	f := IETF()

	highest := f.Highest()
	if highest.ID != "must" {
		t.Errorf("IETF Highest().ID = %q, want %q", highest.ID, "must")
	}

	lowest := f.Lowest()
	if lowest.ID != "may" {
		t.Errorf("IETF Lowest().ID = %q, want %q", lowest.ID, "may")
	}

	// Test IETF prohibitions framework
	fp := IETFProhibitions()

	highest = fp.Highest()
	if highest.ID != "must-not" {
		t.Errorf("IETFProhibitions Highest().ID = %q, want %q", highest.ID, "must-not")
	}

	lowest = fp.Lowest()
	if lowest.ID != "should-not" {
		t.Errorf("IETFProhibitions Lowest().ID = %q, want %q", lowest.ID, "should-not")
	}
}

func TestActionableLevels(t *testing.T) {
	f := Severity()

	actionable := f.ActionableLevels()
	// Critical, High, Medium, Low are actionable; Informational is not
	if len(actionable) != 4 {
		t.Errorf("ActionableLevels() returned %d levels, want 4", len(actionable))
	}
}

func TestGetBuiltin(t *testing.T) {
	for _, id := range AllBuiltinIDs() {
		f := Get(id)
		if f == nil {
			t.Errorf("Get(%q) returned nil", id)
		}
		if f.ID != id {
			t.Errorf("Get(%q).ID = %q, want %q", id, f.ID, id)
		}
		if len(f.Levels) == 0 {
			t.Errorf("Get(%q).Levels is empty", id)
		}
	}

	if Get("nonexistent") != nil {
		t.Error("Get(nonexistent) should return nil")
	}
}

func TestNormalize(t *testing.T) {
	// Test that different frameworks normalize to comparable values
	severity := Severity()
	priority := Priority()
	moscow := MoSCoW()

	// Critical/P0/Must should all be high priority
	if Normalize(severity, "critical") != NormalizedCritical {
		t.Error("Severity critical should normalize to Critical")
	}
	if Normalize(priority, "p0") != NormalizedCritical {
		t.Error("Priority p0 should normalize to Critical")
	}
	if Normalize(moscow, "must") != NormalizedCritical {
		t.Error("MoSCoW must should normalize to Critical")
	}
}

func TestCompareAcross(t *testing.T) {
	severity := Severity()
	moscow := MoSCoW()

	// Critical (severity) vs Must (moscow) - both highest, should be equal
	result := CompareAcross(severity, "critical", moscow, "must")
	if result != 0 {
		t.Errorf("CompareAcross(critical, must) = %d, want 0", result)
	}

	// Critical (severity) vs Won't (moscow) - Critical > Won't
	result = CompareAcross(severity, "critical", moscow, "wont")
	if result != 1 {
		t.Errorf("CompareAcross(critical, wont) = %d, want 1", result)
	}
}

func TestMapTo(t *testing.T) {
	severity := Severity()
	moscow := MoSCoW()

	// Map critical (severity) to moscow - should get "must"
	mapped := MapTo(severity, "critical", moscow)
	if mapped == nil {
		t.Fatal("MapTo returned nil")
	}
	if mapped.ID != "must" {
		t.Errorf("MapTo(critical, moscow).ID = %q, want %q", mapped.ID, "must")
	}
}
